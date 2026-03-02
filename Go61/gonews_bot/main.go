package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
)

// RSS-ленты с новостями о Go
var rssFeeds = []string{
	"https://go.dev/blog/feed.atom",        // официальный блог Go
	"https://habr.com/ru/hub/go/rss/",      // Хабр (Go)
	"https://www.reddit.com/r/golang/.rss", // Reddit r/golang
	"https://golangweekly.com/rss",         // Golang Weekly
	"https://www.calhoun.io/feed",          // Джон Кэлхаун (уроки)
}

// Конфигурация
const (
	MaxNewsHours = 4 // новости за последние 4 часа
)

// UserData хранит статистику пользователя
type UserData struct {
	TotalRequests int `json:"total_requests"`
	TotalEXP      int `json:"total_exp"`
	Level         int `json:"level"`
}

var (
	users         = make(map[int64]*UserData)
	bot           *tgbotapi.BotAPI
	dataFile      = "users.json"
	expPerRequest = 5
)

func main() {
	// Загружаем .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Не задан TELEGRAM_BOT_TOKEN в .env")
	}

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}
	bot.Debug = false
	log.Printf("Бот авторизован: %s", bot.Self.UserName)

	loadUserData()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	saveTicker := time.NewTicker(5 * time.Minute)
	defer saveTicker.Stop()

	for {
		select {
		case <-stop:
			log.Println("Остановка бота, сохраняем данные...")
			saveUserData()
			return
		case <-saveTicker.C:
			saveUserData()
		case update := <-updates:
			if update.Message != nil {
				handleMessage(update.Message)
			} else if update.CallbackQuery != nil {
				handleCallback(update.CallbackQuery)
			}
		}
	}
}

// --- Работа с данными пользователей ---

func getUser(chatID int64) *UserData {
	if _, ok := users[chatID]; !ok {
		users[chatID] = &UserData{
			TotalRequests: 0,
			TotalEXP:      0,
			Level:         1,
		}
	}
	return users[chatID]
}

func addEXP(chatID int64, amount int) {
	user := getUser(chatID)
	user.TotalEXP += amount
	user.TotalRequests++
	updateLevel(user)
}

func updateLevel(user *UserData) {
	newLevel := user.TotalEXP/50 + 1
	if newLevel > user.Level {
		user.Level = newLevel
		// здесь можно отправить сообщение о повышении уровня, но пока оставим
	}
}

func loadUserData() {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Println("Файл пользователей не найден, начинаем с пустой статистикой")
		return
	}
	json.Unmarshal(data, &users)
}

func saveUserData() {
	data, _ := json.MarshalIndent(users, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)
	log.Println("Данные пользователей сохранены")
}

// --- Обработчики ---

func handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	user := getUser(chatID)

	if !msg.IsCommand() {
		return
	}

	switch msg.Command() {
	case "start":
		sendWelcome(chatID)
	case "stats":
		sendStats(chatID, user)
	case "help":
		sendHelp(chatID)
	default:
		bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда. Напиши /help"))
	}
}

func handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	user := getUser(chatID)

	switch callback.Data {
	case "fresh_news":
		// Отвечаем, что ищем новости
		bot.Request(tgbotapi.NewCallback(callback.ID, "Ищу свежие новости..."))
		// Получаем новости
		news := fetchFreshNews()
		if len(news) == 0 {
			bot.Send(tgbotapi.NewMessage(chatID, "😕 За последние 4 часа новостей не найдено. Попробуй позже."))
		} else {
			// Формируем сообщение
			text := "📰 *Свежие новости Go за последние 4 часа:*\n\n"
			for i, item := range news {
				text += fmt.Sprintf("%d. [%s](%s)\n", i+1, item.Title, item.Link)
			}
			msg := tgbotapi.NewMessage(chatID, text)
			msg.ParseMode = "Markdown"
			msg.DisableWebPagePreview = true
			bot.Send(msg)
		}
		// Начисляем EXP
		addEXP(chatID, expPerRequest)
		// Обновляем сообщение с кнопкой (убираем часики)
		bot.Request(tgbotapi.NewCallback(callback.ID, ""))
	default:
		bot.Request(tgbotapi.NewCallback(callback.ID, "Неизвестная команда"))
	}
}

func sendWelcome(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🚀 Fresh IT News", "fresh_news"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот, который собирает свежие IT-новости о языке Go и Go-стеке.\n\nНажми кнопку ниже, чтобы получить новости за последние 4 часа.")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func sendHelp(chatID int64) {
	text := "📋 *Доступные команды:*\n" +
		"/start – приветствие и кнопка новостей\n" +
		"/stats – твоя статистика EXP\n" +
		"/help – эта справка"
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

func sendStats(chatID int64, user *UserData) {
	text := fmt.Sprintf("📊 *Твоя статистика*\n\n"+
		"📈 Уровень: %d\n"+
		"🎮 Всего EXP: %d\n"+
		"📰 Запросов новостей: %d",
		user.Level, user.TotalEXP, user.TotalRequests)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

// --- Получение новостей из RSS ---

type NewsItem struct {
	Title string
	Link  string
	Date  time.Time
}

func fetchFreshNews() []NewsItem {
	fp := gofeed.NewParser()
	now := time.Now()
	threshold := now.Add(-time.Duration(MaxNewsHours) * time.Hour)

	var allNews []NewsItem

	for _, feedURL := range rssFeeds {
		feed, err := fp.ParseURL(feedURL)
		if err != nil {
			log.Printf("Ошибка парсинга %s: %v", feedURL, err)
			continue
		}
		for _, item := range feed.Items {
			var published time.Time
			if item.PublishedParsed != nil {
				published = *item.PublishedParsed
			} else if item.UpdatedParsed != nil {
				published = *item.UpdatedParsed
			} else {
				continue
			}
			if published.After(threshold) {
				allNews = append(allNews, NewsItem{
					Title: item.Title,
					Link:  item.Link,
					Date:  published,
				})
			}
		}
	}

	// Сортируем по дате (свежие сверху)
	sort.Slice(allNews, func(i, j int) bool {
		return allNews[i].Date.After(allNews[j].Date)
	})

	return allNews
}
