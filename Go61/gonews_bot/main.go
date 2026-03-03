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

// RSS-ленты: Go, вайб-кодинг, ИИ (полная версия)
var rssFeeds = []string{
	// Официальные и стабильные (работают всегда)
	"https://go.dev/blog/feed.atom", // Официальный блог Go

	// Русскоязычные (новые ссылки Хабра, которые работают)
	"https://habr.com/ru/rss/hubs/go/feed.xml",          // Хабр: Go (новая ссылка)
	"https://habr.com/ru/rss/hubs/ai/feed.xml",          // Хабр: ИИ (новая ссылка)
	"https://habr.com/ru/rss/hubs/programming/feed.xml", // Хабр: Программирование

	// Англоязычные (легкодоступные)
	"https://golangweekly.com/rss", // Golang Weekly
	"https://dev.to/feed/tag/go",   // DEV.to (Go)
}

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
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не задан")
	}

	var err error
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
	saveUserData() // сохраняем сразу, чтобы не потерять
}

func updateLevel(user *UserData) {
	newLevel := user.TotalEXP/50 + 1
	if newLevel > user.Level {
		user.Level = newLevel
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

	if !msg.IsCommand() {
		return
	}

	switch msg.Command() {
	case "start":
		sendWelcome(chatID)
	case "stats":
		user := getUser(chatID)
		sendStats(chatID, user)
	case "help":
		sendHelp(chatID)
	default:
		bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда. Напиши /help"))
	}
}

func handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID

	switch callback.Data {
	case "fresh_news":
		// Сразу отвечаем, чтобы убрать "часики"
		bot.Request(tgbotapi.NewCallback(callback.ID, ""))

		news := fetchFreshNews()
		if len(news) == 0 {
			bot.Send(tgbotapi.NewMessage(chatID, "😕 За последние 4 часа новостей не найдено. Попробуй позже."))
			return
		}

		text := "📰 *Свежие новости Go за последние 4 часа:*\n\n"
		for i, item := range news {
			text += fmt.Sprintf("%d. [%s](%s)\n", i+1, item.Title, item.Link)
		}
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ParseMode = "Markdown"
		msg.DisableWebPagePreview = true
		bot.Send(msg)

		addEXP(chatID, expPerRequest)

	default:
		bot.Request(tgbotapi.NewCallback(callback.ID, "Неизвестная команда"))
	}
}

func sendWelcome(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Fresh IT News", "fresh_news"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Привет! Я собираю свежие IT-новости о Go и Go-стеке.\n\nНажми кнопку ниже, чтобы получить новости за последние 4 часа.")
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

	sort.Slice(allNews, func(i, j int) bool {
		return allNews[i].Date.After(allNews[j].Date)
	})

	return allNews
}
