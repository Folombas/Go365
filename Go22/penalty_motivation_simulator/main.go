package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==================== СТРУКТУРЫ И ТИПЫ ====================
type Order struct {
	Description string
	Route       string
	Transport   []string
	BasePrice   float64
	Penalty     float64
	Difficulty  int    // 1-10
	Size        string // "small", "medium", "large"
}

type PassengerReaction struct {
	Emotion   string
	Intensity int // 1-10
	Comment   string
}

type DailyResult struct {
	TotalOrders    int
	Completed      int
	Cancelled      int
	TotalEarned    float64
	TotalPenalties float64
	Motivation     int // 1-100
	TraumaScore    int // 0-100 (психологическая нагрузка)
}

// ==================== КОНСТАНТЫ И ГЛОБАЛЬНЫЕ ДАННЫЕ ====================
var motivationalPhrases = []string{
	"🚀 Штраф 200 рублей сегодня — зарплата 200К завтра!",
	"💪 Каждый презрительный взгляд — кирпичик в фундаменте твоей карьеры в Go!",
	"🎯 Объёмный пакет в автобусе — метафора твоего огромного потенциала!",
	"🔥 Пассажиры смотрят с пренебрежением? Пусть смотрят на твой успешный код!",
	"💡 300 рублей за случайный заказ — случайность, которая ведёт к цели!",
	"🌟 Фестивальная → Митино → Сокол → Левобережный — твой путь к Senior Go Developer!",
	"📈 Каждая поездка на автобусе — поездка к удалённой работе!",
	"🎮 Видеомонтаж подождёт — компилятор Go не ждёт!",
	"🏃‍♂️ Физическая усталость сегодня — умственная энергия завтра!",
	"🚀 Холодные взгляды пассажиров — тёплый привет от будущих коллег-разработчиков!",
}

var orders = []Order{
	{
		Description: "Короткий районный заказ (случайный)",
		Route:       "По району на автобусе",
		Transport:   []string{"автобус"},
		BasePrice:   300.0,
		Penalty:     200.0,
		Difficulty:  3,
		Size:        "small",
	},
	{
		Description: "Собачий корм Фестивальная → Митино",
		Route:       "Фестивальная → Митино",
		Transport:   []string{"метро", "автобус"},
		BasePrice:   450.0,
		Penalty:     150.0,
		Difficulty:  6,
		Size:        "medium",
	},
	{
		Description: "Бескаркасное кресло-груша Сокол → Левобережный",
		Route:       "Сокол → Левобережный",
		Transport:   []string{"автобус"},
		BasePrice:   600.0,
		Penalty:     300.0,
		Difficulty:  9,
		Size:        "large",
	},
}

var passengerReactions = []PassengerReaction{
	{"Презрение", 8, "Брезгливо отворачиваются"},
	{"Хладнокровие", 6, "Смотрят сквозь тебя"},
	{"Раздражение", 7, "Неодобрительно качают головой"},
	{"Равнодушие", 4, "Просто игнорируют"},
	{"Сочувствие", 3, "Смотрят с пониманием"},
	{"Любопытство", 5, "Разглядывают груз"},
	{"Неодобрение", 9, "Шепчутся между собой"},
}

// ==================== ФУНКЦИИ СИМУЛЯЦИИ ====================
func simulateMorning() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🌅 УТРО 11:00")
	fmt.Println(strings.Repeat("=", 60))

	actions := []string{
		"⏰ Подъем в 11 утра",
		"📝 Переписываю показания счётчиков",
		"👩 Мама пришла домой",
		"🍵 Пью чай с мамой на кухне",
		"👕 Одеваюсь в уличное",
		"📱 Запускаю приложение для заказов",
	}

	for _, action := range actions {
		fmt.Println(action)
		time.Sleep(400 * time.Millisecond)
	}
}

func acceptOrder(order Order) (bool, float64) {
	fmt.Printf("\n📦 ПРЕДЛОЖЕНИЕ: %s\n", order.Description)
	fmt.Printf("   Маршрут: %s\n", order.Route)
	fmt.Printf("   Транспорт: %v\n", strings.Join(order.Transport, " → "))
	fmt.Printf("   Оплата: %.0f руб | Штраф за отмену: %.0f руб\n", order.BasePrice, order.Penalty)

	// Симуляция раздумий
	fmt.Print("   Думаю... ")
	time.Sleep(1 * time.Second)

	// Случайное принятие (как в легенде)
	if order.Description == "Короткий районный заказ (случайный)" {
		fmt.Println("🤔 Принимаю случайно, хотя не планировал!")
		return true, order.BasePrice
	}

	// Для других заказов - решение на основе сложности
	if order.Difficulty > 7 && rand.Intn(100) < 40 {
		fmt.Printf("😨 Слишком сложно! Отказываюсь. Штраф: -%.0f руб\n", order.Penalty)
		return false, -order.Penalty
	}

	fmt.Println("✅ Принимаю заказ!")
	return true, order.BasePrice
}

func executeOrder(order Order, orderNum int) (float64, int) {
	fmt.Printf("\n🚚 ВЫПОЛНЕНИЕ ЗАКАЗА %d: %s\n", orderNum, order.Description)

	traumaScore := 0

	// Этапы выполнения
	stages := []string{
		"Выхожу на точку забора",
		"Забираю груз",
		"Иду к транспорту",
		"Еду по маршруту",
		"Ищу адрес доставки",
		"Передаю груз получателю",
	}

	for i, stage := range stages {
		fmt.Printf("   %d. %s", i+1, stage)
		time.Sleep(time.Duration(order.Difficulty*50) * time.Millisecond)

		// Вероятность проблемы на этапе
		if rand.Intn(100) < (order.Difficulty * 3) {
			fmt.Printf(" ⚠️\n")
			// Небольшой штраф за проблему
			penalty := order.BasePrice * 0.1
			fmt.Printf("      Проблема! Штраф: -%.0f руб\n", penalty)
			return -penalty, 5
		}
		fmt.Printf(" ✅\n")
	}

	// Особые события для разных заказов
	switch order.Size {
	case "large":
		fmt.Println("\n   🛋️ КРУПНОГАБАРИТНЫЙ ГРУЗ:")
		fmt.Println("      Пакет занимает пол-автобуса")
		fmt.Println("      Все места заняты объёмной посылкой")

		// Реакции пассажиров
		reaction := passengerReactions[rand.Intn(len(passengerReactions))]
		fmt.Printf("   👥 Реакция пассажиров: %s (интенсивность: %d/10)\n",
			reaction.Emotion, reaction.Intensity)
		fmt.Printf("      %s\n", reaction.Comment)

		// Психологическая травма от реакции
		traumaScore = reaction.Intensity * 5
		if reaction.Intensity >= 7 {
			fmt.Println("   😔 Обидно и неприятно...")
		}

	case "medium":
		fmt.Println("\n   🐕 ДОСТАВКА КОРМА:")
		fmt.Println("      Пересадка с метро на автобус")
		fmt.Println("      Тяжелый, но компактный груз")

	case "small":
		fmt.Println("\n   🚌 ПРОГУЛКА НА АВТОБУСЕ:")
		fmt.Println("      Просто качусь по району")
		fmt.Println("      Без метро, только автобусы")
	}

	fmt.Printf("   🎉 ЗАКАЗ УСПЕШНО ДОСТАВЛЕН! +%.0f руб\n", order.BasePrice)
	return order.BasePrice, traumaScore
}

func calculateMotivation(earned float64, penalties float64, trauma int) (int, string) {
	// Базовая мотивация
	baseMotivation := 50

	// Влияние заработка
	if earned > 1000 {
		baseMotivation += 20
	} else if earned > 500 {
		baseMotivation += 10
	}

	// Влияние штрафов
	if penalties > 300 {
		baseMotivation -= 15
	} else if penalties > 0 {
		baseMotivation -= 5
	}

	// Влияние психологической травмы
	baseMotivation -= trauma / 10

	// Ограничение диапазона
	if baseMotivation < 10 {
		baseMotivation = 10
	} else if baseMotivation > 95 {
		baseMotivation = 95
	}

	// Выбор мотивационной фразы
	var phrase string
	if baseMotivation < 30 {
		phrase = "🔥 КРИТИЧЕСКИ НИЗКАЯ МОТИВАЦИЯ! СРОЧНО УЧИТЬ GO!"
	} else if baseMotivation < 50 {
		phrase = "💪 ТЯЖЕЛО, НО НУЖНО ПРОДОЛЖАТЬ! GO ЖДЁТ!"
	} else {
		phrase = motivationalPhrases[rand.Intn(len(motivationalPhrases))]
	}

	return baseMotivation, phrase
}

func simulateEvening(result DailyResult) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🌙 ВЕЧЕР: ПОДВЕДЕНИЕ ИТОГОВ")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Printf("📊 СТАТИСТИКА ДНЯ:\n")
	fmt.Printf("   Всего заказов: %d\n", result.TotalOrders)
	fmt.Printf("   Выполнено: %d | Отменено: %d\n", result.Completed, result.Cancelled)
	fmt.Printf("   Заработано: %.0f руб\n", result.TotalEarned)
	fmt.Printf("   Штрафы: %.0f руб\n", result.TotalPenalties)
	fmt.Printf("   Чистая прибыль: %.0f руб\n", result.TotalEarned-result.TotalPenalties)
	fmt.Printf("   Психологическая нагрузка: %d/100\n", result.TraumaScore)

	motivation, phrase := calculateMotivation(
		result.TotalEarned,
		result.TotalPenalties,
		result.TraumaScore,
	)

	fmt.Printf("\n🎯 УРОВЕНЬ МОТИВАЦИИ: %d/100\n", motivation)
	fmt.Printf("   %s\n", phrase)

	// Геймификация
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎮 ГЕЙМИФИКАЦИЯ ДНЯ")
	fmt.Println(strings.Repeat("=", 60))

	score := result.Completed * 100
	score -= result.Cancelled * 50
	score += int(result.TotalEarned / 10)
	score -= int(result.TotalPenalties / 5)
	score -= result.TraumaScore / 2

	// Достижения
	achievements := []string{}
	if result.Completed >= 2 {
		achievements = append(achievements, "🏅 Мастер доставки")
	}
	if result.TotalPenalties == 0 {
		achievements = append(achievements, "💰 Без штрафов")
	}
	if result.TraumaScore > 50 {
		achievements = append(achievements, "💪 Выдержал давление")
	}
	if result.TotalEarned > 1000 {
		achievements = append(achievements, "🚀 Высокий доход")
	}

	fmt.Printf("   🏆 Очки: %d\n", score)
	if len(achievements) > 0 {
		fmt.Printf("   🎖️ Достижения: %s\n", strings.Join(achievements, ", "))
	}

	// Рекомендации
	fmt.Println("\n💡 РЕКОМЕНДАЦИИ НА ЗАВТРА:")
	if result.TraumaScore > 60 {
		fmt.Println("   - Сегодня было тяжело психологически")
		fmt.Println("   - Завтра больше времени удели Go")
		fmt.Println("   - Помни: программирование избавит от таких ситуаций")
	}
	if result.TotalPenalties > 200 {
		fmt.Println("   - Много штрафов сегодня")
		fmt.Println("   - Каждый штраф — повод учить Go усерднее")
	}
	if result.Completed == result.TotalOrders {
		fmt.Println("   - Отличный день! Все заказы выполнены!")
		fmt.Println("   - Теперь можно с чистой совестью учить Go")
	}
}

// ==================== MAIN ====================
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(strings.Repeat("⭐", 60))
	fmt.Println("🚀 ДЕНЬ КУРЬЕРА: ШТРАФНЫЕ МОТИВАЦИИ И ПУТЬ К GO")
	fmt.Println(strings.Repeat("⭐", 60))

	// Утро
	simulateMorning()

	// Рабочий день
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🛵 РАБОЧИЙ ДЕНЬ")
	fmt.Println(strings.Repeat("=", 60))

	result := DailyResult{}
	var totalTrauma int

	for i, order := range orders {
		fmt.Printf("\n📅 ЗАКАЗ %d/%d\n", i+1, len(orders))

		// Принятие заказа
		accepted, price := acceptOrder(order)
		if !accepted {
			result.Cancelled++
			result.TotalPenalties += -price // price отрицательный при штрафе
			result.TotalOrders++
			continue
		}

		result.TotalOrders++

		// Выполнение заказа
		earned, trauma := executeOrder(order, i+1)
		if earned > 0 {
			result.Completed++
			result.TotalEarned += earned
		} else {
			result.Cancelled++
			result.TotalPenalties += -earned // earned отрицательный при штрафе
		}

		totalTrauma += trauma

		// Пауза между заказами
		time.Sleep(500 * time.Millisecond)
	}

	result.TraumaScore = totalTrauma

	// Вечер: итоги
	simulateEvening(result)

	// Философские выводы
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🤔 ФИЛОСОФСКИЕ ВЫВОДЫ")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   1. Штрафы — не наказание, а плата за опыт")
	fmt.Println("   2. Взгляды пассажиров — отражение их проблем, не твоих")
	fmt.Println("   3. Каждая поездка — километр на пути к разработчику")
	fmt.Println("   4. Физическая усталость сегодня — умственная сила завтра")
	fmt.Println("   5. Go — твой билет из автобуса в офис")

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📢 DISCLAIMER")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("   Все персонажи и события вымышлены.")
	fmt.Println("   Любые совпадения с реальными людьми и событиями случайны.")
	fmt.Println("   История создана для мотивации изучения Go.")
	fmt.Println("   © Daily Code Life Story - художественный вымысел.")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n🎯 ВЫВОД: Продолжай учить Go!")
	fmt.Println("   Завтра — новый день, новый код, новая мотивация!")
	fmt.Println(strings.Repeat("⭐", 60))
}
