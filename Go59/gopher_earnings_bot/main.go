package main

import (
	"fmt"
	"log"

	//"os"
	//"strconv"
	//"strings"
	//"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// UserData хранит информацию о пользователе (в памяти, для простоты)
type UserData struct {
	TotalEarned int    // всего заработано (рублей)
	EXP         int    // очки опыта
	LastCommand string // последняя команда (для игрового эффекта)
}

var users = make(map[int64]*UserData) // key: chatID

// Конфигурация
const (
	BotToken = "ВАШ_ТОКЕН_ЗДЕСЬ" // замените на реальный токен
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatalf("Ошибка авторизации: %v", err)
	}
	bot.Debug = false
	log.Printf("Бот авторизован: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // обработка обычных сообщений
			handleMessage(bot, update.Message)
		} else if update.CallbackQuery != nil { // обработка нажатий на инлайн-кнопки
			handleCallback(bot, update.CallbackQuery)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	userData := getUserData(chatID)

	if msg.IsCommand() {
		switch msg.Command() {
		case "start":
			sendStart(bot, chatID)
			addEXP(chatID, 5)
		case "help":
			sendHelp(bot, chatID)
			addEXP(chatID, 2)
		case "earn":
			sendEarn(bot, chatID, userData)
		case "weather":
			sendWeather(bot, chatID)
			addEXP(chatID, 3)
		case "profile":
			sendProfile(bot, chatID, userData)
			addEXP(chatID, 1)
		default:
			bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда. Напиши /help"))
		}
	} else {
		// Если не команда, предлагаем список команд
		bot.Send(tgbotapi.NewMessage(chatID, "Я понимаю только команды. Напиши /help"))
	}
}

func handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	userData := getUserData(chatID)

	data := callback.Data
	switch data {
	case "earn_100":
		userData.TotalEarned += 100
		bot.Send(tgbotapi.NewMessage(chatID, "✅ Зачислено 100 рублей! Теперь у тебя: "+fmt.Sprint(userData.TotalEarned)+" руб."))
		addEXP(chatID, 10)
	case "earn_50":
		userData.TotalEarned += 50
		bot.Send(tgbotapi.NewMessage(chatID, "✅ Зачислено 50 рублей! Теперь у тебя: "+fmt.Sprint(userData.TotalEarned)+" руб."))
		addEXP(chatID, 5)
	case "weather_now":
		bot.Send(tgbotapi.NewMessage(chatID, "🌧 За окном снежная каша. Обувь промокла, но ты молодец, что вышел!"))
		addEXP(chatID, 2)
	default:
		bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная кнопка"))
	}

	bot.Request(tgbotapi.NewCallback(callback.ID, "")) // убираем "часики"
}

func getUserData(chatID int64) *UserData {
	if _, ok := users[chatID]; !ok {
		users[chatID] = &UserData{
			TotalEarned: 0,
			EXP:         0,
		}
	}
	return users[chatID]
}

func addEXP(chatID int64, amount int) {
	user := getUserData(chatID)
	user.EXP += amount
	// Здесь можно проверять достижения, но для простоты пропустим
}

func sendStart(bot *tgbotapi.BotAPI, chatID int64) {
	text := `Привет, я GopherEarningsBot – твой помощник в учёте доходов и прокачке!

Сегодня ты заработал 600 рублей, промочив ноги, но не унывай. Я помогу тебе отслеживать финансы и получать EXP за каждое действие.

Что умею:
/earn – добавить сегодняшний заработок
/weather – узнать погоду (и получить EXP)
/profile – посмотреть свой профиль и уровень
/help – список команд

Поехали!`
	bot.Send(tgbotapi.NewMessage(chatID, text))
}

func sendHelp(bot *tgbotapi.BotAPI, chatID int64) {
	text := `📋 Список команд:
/start – начать работу
/help – эта справка
/earn – добавить заработок (с инлайн-кнопками)
/weather – прогноз погоды
/profile – твой профиль (EXP, заработано)

Каждая команда приносит EXP!`
	bot.Send(tgbotapi.NewMessage(chatID, text))
}

func sendEarn(bot *tgbotapi.BotAPI, chatID int64, userData *UserData) {
	// Создаём инлайн-клавиатуру для выбора суммы
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ 100 руб", "earn_100"),
			tgbotapi.NewInlineKeyboardButtonData("➕ 50 руб", "earn_50"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❓ Другая сумма", "earn_other"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Сколько сегодня заработал?")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func sendWeather(bot *tgbotapi.BotAPI, chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🌧 Сейчас", "weather_now"),
			tgbotapi.NewInlineKeyboardButtonData("⏰ На завтра", "weather_tomorrow"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Погода на сегодня:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func sendProfile(bot *tgbotapi.BotAPI, chatID int64, userData *UserData) {
	level := userData.EXP/100 + 1
	nextLevelExp := level * 100
	progress := float64(userData.EXP%100) / 100 * 100

	text := fmt.Sprintf(`📊 Твой профиль:
💰 Всего заработано: %d руб.
🎮 Опыт (EXP): %d
📈 Уровень: %d
🔜 До следующего уровня: %d EXP (прогресс: %.0f%%)`,
		userData.TotalEarned, userData.EXP, level, nextLevelExp-userData.EXP, progress)

	bot.Send(tgbotapi.NewMessage(chatID, text))
}
