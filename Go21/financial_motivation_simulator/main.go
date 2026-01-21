package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==================== СТРУКТУРЫ ДАННЫХ ====================
type DailyResult struct {
	OrdersCompleted int
	TotalEarned     float64
	Expenses        float64
	Balance         float64
	MotivationScore int
	TransportLog    []string
}

type Order struct {
	From      string
	To        string
	BasePrice float64
	Vehicle   string
	Route     string
}

// ==================== ГЛОБАЛЬНЫЕ ПЕРЕМЕННЫЕ ====================
var motivationPhrases = []string{
	"💪 Холодный душ утром — горячий код вечером!",
	"🚀 700 рублей сегодня — 200К завтра как Go-разработчик!",
	"🎯 Финансовый удар — мотивация учить Go ударными темпами!",
	"🔥 Никаких видеоигр! Только Go и коммиты!",
	"💡 2К маме сегодня, 200К себе завтра!",
	"🌟 Курьерка — временно, Go — стабильность!",
	"📈 Каждая поездка на автобусе 1346 — шаг к удаленной работе!",
	"🎮 Фильмы и сериалы подождут, изучай сейчас Go!",
	"🏃‍♂️ Беготня по Москве — тренировка выносливости для марафона кода!",
	"🚀 Тройка в автобусе, Go в карьере — билет к финансовой свободе!",
}

var lifeExpenses = []struct {
	Name     string
	Amount   float64
	Priority int
}{
	{"Перевод маме", 2000.0, 1},
	{"Мясо (вчерашняя покупка)", 800.0, 2},
	{"Коммунальные услуги", 3500.0, 3},
	{"Проезд Тройка", 1500.0, 4},
	{"Еда", 5000.0, 5},
}

// ==================== ФУНКЦИИ СИМУЛЯЦИИ ====================
func simulateMorning() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🌅 УТРО: 10:30")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("⏰ Подъем в пол-одиннадцатого")
	fmt.Println("🚰 Горячей воды нет — умываюсь холодной")
	fmt.Println("🍳 Завтрак, горячий чай")
	time.Sleep(1 * time.Second)
}

func executeOrder(order Order) (float64, bool) {
	fmt.Printf("\n📦 ЗАКАЗ: %s → %s\n", order.From, order.To)
	fmt.Printf("   Транспорт: %s\n", order.Vehicle)
	fmt.Printf("   Маршрут: %s\n", order.Route)

	// Симуляция случайных событий
	events := []string{
		"пробки на Ленинградке",
		"снегопад, движение затруднено",
		"автобус сломался, жду следующий",
		"клиент задержал на 15 минут",
		"погода ясная, еду быстро",
	}

	event := events[rand.Intn(len(events))]
	fmt.Printf("   Событие: %s\n", event)

	// Динамическое ценообразование (от 0.5 до 1.5 от базовой цены)
	multiplier := 0.5 + rand.Float64()
	finalPrice := order.BasePrice * multiplier

	// 25% шанс отмены заказа
	if rand.Intn(100) < 25 {
		fmt.Println("   ❌ ЗАКАЗ ОТМЕНЕН! Компенсация: 0 рублей")
		return 0, false
	}

	fmt.Printf("   ✅ ЗАКАЗ ВЫПОЛНЕН! Заработано: %.2f рублей\n", finalPrice)
	return finalPrice, true
}

func calculateExpenses(earned float64) (float64, float64) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("💰 ФИНАНСОВЫЙ ОТЧЕТ")
	fmt.Println(strings.Repeat("=", 60))

	totalExpenses := 0.0
	criticalExpenses := 0.0

	for _, expense := range lifeExpenses {
		if expense.Priority <= 2 { // Критические расходы (мама, еда)
			totalExpenses += expense.Amount
			criticalExpenses += expense.Amount
			fmt.Printf("   🔴 %s: -%.2f руб (обязательно)\n", expense.Name, expense.Amount)
		} else if earned > totalExpenses+expense.Amount {
			totalExpenses += expense.Amount
			fmt.Printf("   🟡 %s: -%.2f руб\n", expense.Name, expense.Amount)
		} else {
			fmt.Printf("   🟢 %s: -%.2f руб (ОТЛОЖЕНО, не хватает денег)\n", expense.Name, expense.Amount)
		}
	}

	return totalExpenses, criticalExpenses
}

func checkMotivation(balance float64, ordersCompleted int) string {
	if balance < 0 {
		return "💸 КРИТИЧЕСКОЕ ПОЛОЖЕНИЕ! СРОЧНО УЧИТЬ GO!"
	}

	if ordersCompleted == 0 {
		return "😡 ВСЕ ЗАКАЗЫ ОТМЕНЕНЫ! ЕЩЁ ОДНА ПРИЧИНА УЧИТЬ GO!"
	}

	if balance < 1000 {
		return "⚡ МАЛО ДЕНЕГ! ПОРА ПЕРЕХОДИТЬ НА GO-РАЗРАБОТКУ!"
	}

	index := rand.Intn(len(motivationPhrases))
	return motivationPhrases[index]
}

func generateGoals(balance float64) []string {
	goals := []string{
		"🎯 Выучить Go до уровня Junior за 3 месяца",
		"💼 Устроиться на работу с зарплатой от 150К рублей",
		"🚫 Не тратить время на видеоигры и сериалы",
		"📚 Читать документацию Go каждый день",
		"💻 Делать 1 коммит в день в GitHub",
		"🏝️ Смонтировать видео с Филиппин после первой зарплаты",
		"🔥 Пройти 5 собеседований в месяц",
		"🚀 Участвовать в open-source проектах на Go",
	}

	// Добавляем финансовые цели
	if balance < 0 {
		goals = append(goals, "🆘 ВЫЖИТЬ ДО СЛЕДУЮЩЕЙ ДОСТАВКИ")
		goals = append(goals, "💰 ЗАРАБОТАТЬ ХОТЯ БЫ НА ЕДУ")
	} else if balance < 5000 {
		goals = append(goals, fmt.Sprintf("💰 НАКОПИТЬ 5000 руб (сейчас: %.2f)", balance))
	}

	return goals
}

func simulateEvening(earned float64, expenses float64, balance float64, ordersCompleted int) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🌙 ВЕЧЕР: 18:00 - 21:00")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🚿 Горячий душ после холодного дня")
	fmt.Println("🍽️ Ужин, чай с мамой")
	fmt.Println("💻 21:00 - Сажусь за компьютер программировать")

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎯 МОТИВАЦИОННЫЙ АНАЛИЗ")
	fmt.Println(strings.Repeat("=", 60))

	motivation := checkMotivation(balance, ordersCompleted)
	fmt.Printf("   %s\n", motivation)

	fmt.Println("\n🎯 ЦЕЛИ НА БЛИЖАЙШЕЕ БУДУЩЕЕ:")
	goals := generateGoals(balance)
	for i, goal := range goals {
		fmt.Printf("   %d. %s\n", i+1, goal)
	}

	// Геймификация
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎮 ГЕЙМИФИКАЦИЯ ДНЯ")
	fmt.Println(strings.Repeat("=", 60))

	score := 0
	if ordersCompleted > 0 {
		score += ordersCompleted * 10
	}
	if balance > 0 {
		score += 20
	}
	if balance > 1000 {
		score += 30
	}

	fmt.Printf("   🏆 Очки продуктивности: %d/100\n", score)
	fmt.Printf("   💰 Финансовый баланс: %.2f руб\n", balance)
	fmt.Printf("   🎯 Уровень мотивации: %d%%\n", score)
	fmt.Printf("   🚀 До цели 200К: %.3f%%\n", (earned/200000)*100)

	if score > 50 {
		fmt.Println("   ✅ Отличный день! Можно позволить себе 1 час Go!")
	} else {
		fmt.Println("   ⚠️ Тяжелый день. Нужно компенсировать кодом!")
	}
}

// ==================== MAIN ====================
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🚀 ДЕНЬ КУРЬЕРА-РАЗРАБОТЧИКА: БИТВА ЗА ВЫЖИВАНИЕ И GO")
	fmt.Println(strings.Repeat("=", 60))

	// Утро
	simulateMorning()

	// День: работа курьером
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🚚 РАБОЧИЙ ДЕНЬ КУРЬЕРА")
	fmt.Println(strings.Repeat("=", 60))

	// Заказ 1: Совхозная → Центр Москвы
	order1 := Order{
		From:      "Совхозная, Левый Берег Химки",
		To:        "Центр Москвы",
		BasePrice: 700.0,
		Vehicle:   "Автобус 1346 (новый маршрут Москва-Область)",
		Route:     "Химки → Москва",
	}
	earned1, completed1 := executeOrder(order1)

	// Заказ 2: Баррикадная → Новые Химки
	order2 := Order{
		From:      "Баррикадная",
		To:        "Новые Химки (около МЕГА)",
		BasePrice: 500.0,
		Vehicle:   "Автобус 359 → пешком от 'Школы Гольфа'",
		Route:     "м. Планерная → Новые Химки",
	}
	earned2, completed2 := executeOrder(order2)

	// Итоги работы
	totalEarned := earned1 + earned2
	completedOrders := 0
	if completed1 {
		completedOrders++
	}
	if completed2 {
		completedOrders++
	}

	fmt.Printf("\n📊 ИТОГИ РАБОЧЕГО ДНЯ:\n")
	fmt.Printf("   Выполнено заказов: %d/2\n", completedOrders)
	fmt.Printf("   Всего заработано: %.2f рублей\n", totalEarned)

	// Финансовые обязательства
	expenses, criticalExpenses := calculateExpenses(totalEarned)
	balance := totalEarned - criticalExpenses // Считаем только по обязательным расходам

	fmt.Printf("\n   Обязательные расходы: %.2f руб\n", criticalExpenses)
	fmt.Printf("   Общий баланс: %.2f руб\n", balance)

	if balance < 0 {
		fmt.Println("   ❌ ФИНАНСОВЫЙ МИНУС! Нужно больше заказов или... учить Go!")
	}

	// Вечер и мотивация
	simulateEvening(totalEarned, expenses, balance, completedOrders)

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📢 DISCLAIMER")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   Все персонажи и события вымышлены.")
	fmt.Println("   Любые совпадения с реальными людьми и событиями случайны.")
	fmt.Println("   История создана для мотивации изучения Go.")
	fmt.Println("   © Daily Code Life Story - художественный вымысел.")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n🚀 ВЫВОД: Нужно продолжать учить Go!")
	fmt.Println("   Финансовая свобода начинается с первого коммита!")
	fmt.Println("   Завтра - новый день, новый код, новые возможности!")
}
