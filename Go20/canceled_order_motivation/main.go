package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Тип для симуляции дохода дня
type DailyResult struct {
	OrdersCompleted int
	TotalEarned     float64
	TimeWasted      time.Duration
	MotivationLevel int // 1-10
}

// Симуляция принятия заказа с вероятностью отмены
func acceptOrder(orderName string) (bool, float64) {
	// 40% шанс отмены заказа
	if rand.Intn(100) < 40 {
		fmt.Printf("❌ Заказ '%s' отменен! Компенсации: 0 руб\n", orderName)
		return false, 0.0
	}

	earnings := 200.0 + rand.Float64()*300.0
	fmt.Printf("✅ Заказ '%s' выполнен! Заработано: %.2f руб\n", orderName, earnings)
	return true, earnings
}

// Расчет мотивации на основе дня
func calculateMotivation(result DailyResult) string {
	motivationPhrases := []string{
		"💪 Каждая отмена заказа - повод открыть Go Playground!",
		"🚀 Электричка трясет, но твой код будет стабильным!",
		"🎯 Сегодня потерял 200 руб, завтра заработаешь 200к как Go-разработчик!",
		"🔥 Не ходи по магазинам - цены высокие! Иди учи Go!",
		"🏝️ Смонтируешь Филиппины, когда устроишься на работу!",
		"💡 Видеомонтаж для души, Go - для карьеры!",
		"🌟 Твои подписчики ждут, но твой код ждёт больше!",
		"🎮 Баланс важен: 70% Go, 30% монтаж!",
		"📈 Курьерка - временно, Go - стабильно!",
		"🚀 Москва-Химки на электричке - 30 минут. Москва-Go разработчик - 100 дней!",
	}

	if result.OrdersCompleted == 0 {
		return motivationPhrases[0] + " 🔥 СРОЧНО ОТКРЫВАЙ Go!"
	}

	index := result.MotivationLevel % len(motivationPhrases)
	return motivationPhrases[index]
}

// Балансировщик времени между программированием и монтажом
func balanceTime() (time.Duration, time.Duration) {
	// Рациональный баланс: 2/3 на Go, 1/3 на монтаж
	goTime := time.Duration(120+rand.Intn(60)) * time.Minute  // 2-3 часа
	editTime := time.Duration(60+rand.Intn(30)) * time.Minute // 1-1.5 часа

	return goTime, editTime
}

func main() {
	rand.Seed(time.Now().UnixNano())

	separator := strings.Repeat("=", 60)

	fmt.Println(separator)
	fmt.Println("🌅 ДЕНЬ КУРЬЕРА-РАЗРАБОТЧИКА: БАЛАНС ХАОСА И КОДА")
	fmt.Println(separator)

	// Часть 1: Утро и первые заказы
	fmt.Println("\n🕘 9:30 - Подъем после второго будильника")
	fmt.Println("🍵 Завтрак с чаем, пришла мама")
	fmt.Println("🚪 Выхожу на работу в уличной одежде")

	fmt.Println("\n📍 Ловлю заказ из дома...")
	success1, earned1 := acceptOrder("Москва → Химки")

	fmt.Println("\n🏞️ В парке Льва Толстого проверяю заказ...")
	if !success1 {
		fmt.Println("😡 Заказ отменен! 30 минут на электричке впустую!")
		fmt.Println("💭 Мысль: 'Надо учить Go и видеомонтаж!'")
	}

	// Часть 2: Рабочий день
	fmt.Println("\n📦 РАБОЧИЙ ДЕНЬ:")
	fmt.Println("📍 Станция Химки → БЦ Останкино (документы)")
	success2, earned2 := acceptOrder("Химки → Останкино")

	fmt.Println("\n🏙️ Останкино → Кутузовская (небоскреб Сбера)")
	success3, earned3 := acceptOrder("Останкино → Сбер")

	fmt.Println("\n🌆 Москва-Сити → Академика Королёва")
	fmt.Println("📍 Башня Меркурий → ул. Академика Королёва")
	success4, earned4 := acceptOrder("Меркурий → Королёва")

	fmt.Println("\n🔄 Обратная поездка: Останкино → Химки")
	fmt.Println("🏞️ Прогулка у Останкинского пруда")

	// Часть 3: Подсчет результатов
	fmt.Println("\n💰 ПОДВЕДЕНИЕ ИТОГОВ ДНЯ:")
	totalEarned := earned1 + earned2 + earned3 + earned4
	completedOrders := 0
	if success1 {
		completedOrders++
	}
	if success2 {
		completedOrders++
	}
	if success3 {
		completedOrders++
	}
	if success4 {
		completedOrders++
	}

	// Симуляция потраченного времени
	wastedTime := time.Duration(rand.Intn(180)) * time.Minute // до 3 часов

	result := DailyResult{
		OrdersCompleted: completedOrders,
		TotalEarned:     totalEarned,
		TimeWasted:      wastedTime,
		MotivationLevel: rand.Intn(10) + 1,
	}

	fmt.Printf("✅ Выполнено заказов: %d/4\n", result.OrdersCompleted)
	fmt.Printf("💰 Заработано: %.2f руб\n", result.TotalEarned)
	fmt.Printf("⏰ Потеряно времени: %v\n", result.TimeWasted.Round(time.Minute))

	// Часть 4: Вечерние размышления
	fmt.Println("\n🌙 ВЕЧЕРНИЕ РАЗМЫШЛЕНИЯ:")
	fmt.Println("💭 'Программирование - для работы'")
	fmt.Println("💭 'Видеомонтаж - для души'")
	fmt.Println("💭 'Как найти баланс?'")

	// Балансировка времени
	goTime, editTime := balanceTime()
	fmt.Printf("\n⚖️ БАЛАНС НА ЗАВТРА:\n")
	fmt.Printf("   Go программирование: %v\n", goTime)
	fmt.Printf("   Видеомонтаж: %v\n", editTime)
	fmt.Printf("   Отношение: %.1f:1\n", float64(goTime)/float64(editTime))

	// Часть 5: Мечты о Филиппинах
	fmt.Println("\n🏝️ МЕЧТЫ О ФИЛИППИНАХ 2019:")
	fmt.Println("📹 30 дней соло-путешествия")
	fmt.Println("🎥 Sony X300 Action Cam - единственный собеседник")
	fmt.Println("💾 1ТБ необработанных видео")
	fmt.Println("🎬 Монтаж travel-влогов - незакрытый гештальт")

	// Часть 6: Решение по CapCut
	fmt.Println("\n🎬 РЕШЕНИЕ ПО CAPCUT:")
	if result.TotalEarned > 500 && result.MotivationLevel > 5 {
		fmt.Println("✅ Устанавливаю CapCut! Начинаю монтаж сегодня!")
	} else {
		fmt.Println("⏸️ Откладываю установку CapCut. Слишком устал.")
		fmt.Println("💡 Сначала Go, потом монтаж!")
	}

	// Мотивационная часть
	fmt.Println("\n" + separator)
	fmt.Println("🚀 МОТИВАЦИЯ НА ЗАВТРА:")
	fmt.Println(calculateMotivation(result))
	fmt.Println(separator)

	// Disclaimer
	fmt.Println("\n" + separator)
	fmt.Println("📢 DISCLAIMER:")
	fmt.Println("   Все персонажи и события вымышлены.")
	fmt.Println("   Любые совпадения с реальными людьми и событиями случайны.")
	fmt.Println("   История создана для мотивации изучения Go.")
	fmt.Println("   © Daily Code Life Story - художественный вымысел.")
	fmt.Println(separator)

	fmt.Println("\n🎯 ВЫВОД: Нужно продолжать учить Go!")
	fmt.Println("   Баланс: 70% программирование, 30% монтаж!")
	fmt.Println("   Цель: Go-разработчик + travel-влоггер!")
}
