// fatorderhunter/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"encoding/json"
	"strings"
)

// Глобальные константы игры
const (
	MAX_LEVEL     = 10
	EXP_PER_ORDER = 10
	FAT_ORDER_MIN = 1000
	DAILY_EXPENSES = 700
)

// Структура игрока
type Player struct {
	Name      string  `json:"name"`
	Level     int     `json:"level"`
	Exp       int     `json:"exp"`
	Balance   int     `json:"balance"`
	Orders    int     `json:"orders_completed"`
	FatOrders int     `json:"fat_orders_found"`
	Streak    int     `json:"current_streak"`
	Inventory []string `json:"inventory"`
}

// Структура заказа
type Order struct {
	ID     int
	Amount int
	Zone   string
	Urgent bool
}

// Глобальный игрок
var gosha Player

// Мотивационные фразы
var motivations = []string{
	"💪 Каждая строчка кода на Go — кирпичик в фундаменте твоей карьеры!",
	"🚀 Пока другие ищут 'жирные заказы', ты создаёшь 'жирные навыки'!",
	"🎯 Deadlock в коде исправить проще, чем deadlock в карьере без Go!",
	"🔥 1000 рублей за доставку — временно. Навыки Go — навсегда!",
	"⚡ IPFS написан на Go — и твоё будущее тоже будет!",
	"🎮 Прокачивай не персонажа в игре, а своего внутреннего гофера!",
	"💼 Завтра ты будешь благодарить себя за каждый сегодняшний коммит!",
	"🏆 Жирный заказ — это хорошо. Жирное портфолио — лучше!",
	"🎪 n8n автоматизирует задачи, а Go автоматизирует карьерный рост!",
	"🚚 Развозить заказы может любой. Развозить production-код — избранные!",
}

func init() {
	rand.Seed(time.Now().UnixNano())
	gosha = Player{
		Name:      "Гоша",
		Level:     1,
		Exp:       0,
		Balance:   1500,
		Orders:    0,
		FatOrders: 0,
		Streak:    0,
		Inventory: []string{"Смартфон", "Наушники", "Куртка"},
	}
}

func main() {
	clearScreen()
	printTitle()

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎮 ДЕНЬ 31: ОХОТА ЗА ЖИРНЫМ ЗАКАЗОМ")
	fmt.Println(strings.Repeat("=", 60))

	// Выводим случайную мотивационную фразу
	fmt.Printf("\n📢 %s\n\n", motivations[rand.Intn(len(motivations))])

	// Симуляция дня
	fmt.Println("🌅 10:00 - Гоша проснулся. Начинается новый день охоты...")
	time.Sleep(1 * time.Second)

	// Мониторинг заказов
	fmt.Println("\n📱 Мониторинг заказов в приложении...")
	time.Sleep(800 * time.Millisecond)

	ordersToday := simulateDay()

	// Подведение итогов
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 ИТОГИ ДНЯ:")
	fmt.Println(strings.Repeat("=", 60))

	printStats(ordersToday)
	levelUpCheck()
	expenses()

	// Драматический финал
	fmt.Println("\n" + strings.Repeat("=", 60))
	if gosha.Balance > 0 {
		fmt.Println("🎉 УДАЧНЫЙ ДЕНЬ! Гоша в плюсе и прокачал навыки!")
	} else {
		fmt.Println("💔 СЛОЖНЫЙ ДЕНЬ... Но зато Гоша покатался по городу!")
		fmt.Println("   Завтра будет лучше! Главное — не сдаваться!")
	}

	// Сохраняем прогресс
	saveProgress()

	// Выводим все мотивационные фразы
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🔥 10 МОТИВАЦИОННЫХ ФРАЗ ДЛЯ GO-РАЗРАБОТЧИКА:")
	fmt.Println(strings.Repeat("=", 60))
	for i, phrase := range motivations {
		fmt.Printf("%d. %s\n", i+1, phrase)
	}

	// Финальное сообщение
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎯 ЗАПОМНИ: Каждый день с Go приближает тебя к цели!")
	fmt.Println("   Не отвлекайся на барахолку жизни — строй свою карьеру!")
	fmt.Println(strings.Repeat("=", 60))

	// Интерактивный элемент
	fmt.Print("\nНажми Enter чтобы продолжить путь к успеху...")
	fmt.Scanln()
	fmt.Println("\n🚀 Продолжай учить Go! Завтра - Day 32!")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printTitle() {
	title := `
╔══════════════════════════════════════════════════════════╗
║  🚚 FAT ORDER HUNTER - ГОНКА ЗА ЖИРНЫМ ЗАКАЗОМ 🚚        ║
║  Симулятор ценовой волатильности и карьерного роста      ║
╚══════════════════════════════════════════════════════════╝
	`
	fmt.Println(title)
}

func simulateDay() []Order {
	var orders []Order
	orderCount := rand.Intn(8) + 3 // От 3 до 10 заказов за день

	fmt.Printf("📦 Сегодня в системе: %d заказов\n", orderCount)

	for i := 1; i <= orderCount; i++ {
		time.Sleep(300 * time.Millisecond)

		// Генерируем заказ со случайной ценой
		amount := 300 + rand.Intn(1200) // От 300 до 1500 руб
		urgent := rand.Float32() < 0.3  // 30% шанс срочного заказа

		order := Order{
			ID:     i,
			Amount: amount,
			Zone:   randomZone(),
			Urgent: urgent,
		}

		orders = append(orders, order)

		// Визуализация заказа
		emoji := "📦"
		if urgent {
			emoji = "🚨"
		}

		if amount >= FAT_ORDER_MIN {
			emoji = "💰"
			fmt.Printf("%s ЗАКАЗ #%d: %d руб (ЖИРНЫЙ!) [%s]\n", emoji, i, amount, order.Zone)
		} else {
			fmt.Printf("%s Заказ #%d: %d руб [%s]\n", emoji, i, amount, order.Zone)
		}

		// Гоша принимает решение
		if shouldTakeOrder(order) {
			gosha.processOrder(order)
		} else {
			fmt.Printf("   ⏩ Гоша пропускает... ждёт жирный заказ!\n")
		}
	}

	return orders
}

func (p *Player) processOrder(o Order) {
	p.Orders++
	p.Exp += EXP_PER_ORDER

	if o.Amount >= FAT_ORDER_MIN {
		p.FatOrders++
		p.Balance += o.Amount
		fmt.Printf("   ✅ ПРИНЯТ! +%d руб. (ЖИРНЫЙ!)\n", o.Amount)
	} else {
		p.Balance += o.Amount
		fmt.Printf("   ✅ Принят. +%d руб.\n", o.Amount)
	}

	// Проверка повышения уровня
	if p.Exp >= p.Level*100 {
		p.Level++
		fmt.Printf("\n🎉 УРОВЕНЬ ПОВЫШЕН! Теперь уровень %d!\n", p.Level)
		p.Streak++
	}
}

func shouldTakeOrder(o Order) bool {
	// Логика принятия решения
	if o.Urgent {
		return true // Срочные всегда берём
	}

	// С ростом уровня становимся разборчивее
	chance := float32(gosha.Level) / 10.0

	if o.Amount >= FAT_ORDER_MIN {
		return true // Жирные заказы всегда берём
	}

	// Иначе берём с вероятностью, зависящей от уровня
	return rand.Float32() < chance
}

func randomZone() string {
	zones := []string{"Центр", "Спальный район", "Бизнес-парк", "Торговый центр", "Университет"}
	return zones[rand.Intn(len(zones))]
}

func printStats(orders []Order) {
	totalOrders := len(orders)
	fatOrders := 0
	totalAmount := 0

	for _, o := range orders {
		totalAmount += o.Amount
		if o.Amount >= FAT_ORDER_MIN {
			fatOrders++
		}
	}

	fmt.Printf("👤 Игрок: %s\n", gosha.Name)
	fmt.Printf("⭐ Уровень: %d\n", gosha.Level)
	fmt.Printf("📊 Опыт: %d/%d\n", gosha.Exp, gosha.Level*100)
	fmt.Printf("💰 Баланс: %d руб\n", gosha.Balance)
	fmt.Printf("📦 Всего заказов сегодня: %d\n", totalOrders)
	fmt.Printf("💰 Жирных заказов: %d\n", fatOrders)
	fmt.Printf("🔥 Серия успешных дней: %d\n", gosha.Streak)

	if gosha.FatOrders > 0 {
		fmt.Printf("🏆 Всего жирных заказов за карьеру: %d\n", gosha.FatOrders)
	}
}

func levelUpCheck() {
	nextLevelExp := gosha.Level * 100
	if gosha.Exp >= nextLevelExp {
		fmt.Printf("\n🎮 ГЕЙМИФИКАЦИЯ: До следующего уровня %d опыта!\n",
			nextLevelExp - gosha.Exp)
	}
}

func expenses() {
	fmt.Println("\n🏪 Посещение магазина...")
	time.Sleep(500 * time.Millisecond)

	products := []string{
		"Майонез 'Печагин'",
		"Хлеб 'Столичный'",
		"Чай",
		"Репчатый лук",
		"Шоколадки",
		"Кефир",
	}

	fmt.Println("🛒 Куплено:")
	for _, product := range products {
		fmt.Printf("   - %s\n", product)
		time.Sleep(200 * time.Millisecond)
	}

	gosha.Balance -= DAILY_EXPENSES
	fmt.Printf("💸 Потрачено: %d руб.\n", DAILY_EXPENSES)
	fmt.Printf("📉 Новый баланс: %d руб.\n", gosha.Balance)

	// Добавляем продукты в инвентарь
	gosha.Inventory = append(gosha.Inventory, products...)
}

func saveProgress() {
	// Сохраняем прогресс в JSON файл
	data, err := json.MarshalIndent(gosha, "", "  ")
	if err == nil {
		os.WriteFile("achievements/progress.json", data, 0644)
		fmt.Println("\n💾 Прогресс сохранён в achievements/progress.json")
	}
}
