package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Temptation представляет искушение (CapCut, Blender, etc.)
type Temptation struct {
	Name        string
	Type        string  // "video", "3d", "music", "game"
	Power       int     // сила искушения (1-100)
	Distraction int     // отвлекающие минуты
	Relief      float64 // временное облегчение
}

// GoStudy представляет сессию изучения Go
type GoStudy struct {
	Topic    string
	Minutes  int
	Focus    int     // концентрация (1-100)
	Progress float64 // прогресс (0.0-1.0)
	Dopamine float64 // долгосрочный дофамин
}

// CourierDay представляет день курьера
type CourierDay struct {
	Deliveries []string
	Money      int
	Energy     int
	Mood       string
}

func main() {
	fmt.Println("💊 МОДУЛЬ: CAPCUT-WITHDRAWAL-BATTLE - БИТВА С ЛОМКОЙ")
	fmt.Println("═══════════════════════════════════════════════════")

	fmt.Println("📅 19 января 2026. День курьера-разработчика.")
	fmt.Println("   Будильник на шкафу. Метро. Доставки. И... ЛОМКА.")

	// Симуляция дня курьера
	courierDay := simulateCourierDay()

	fmt.Println("🚚 ДЕНЬ КУРЬЕРА ГОШИ:")
	for i, delivery := range courierDay.Deliveries {
		fmt.Printf("   %d. %s\n", i+1, delivery)
		time.Sleep(400 * time.Millisecond)
	}

	fmt.Printf("\n   💰 Заработано: %d руб.\n", courierDay.Money)
	fmt.Printf("   ⚡ Энергии осталось: %d/100\n", courierDay.Energy)
	fmt.Printf("   😔 Настроение: %s\n", courierDay.Mood)

	// Искушения вечером
	fmt.Println("\n🌙 ВЕЧЕР. ДОМ. КОМПЬЮТЕР.")
	fmt.Println("   Рука тянется к CapCut... но нужно учить Go!")

	temptations := []Temptation{
		{"CapCut", "video", 95, 180, 0.8},
		{"Blender 3D", "3d", 70, 120, 0.6},
		{"FL Studio", "music", 60, 90, 0.5},
		{"Unity Engine", "game", 80, 150, 0.7},
		{"Tilda Web", "web", 50, 60, 0.4},
	}

	goStudies := []GoStudy{
		{"Горутины и каналы", 60, 85, 0.3, 2.5},
		{"Интерфейсы в Go", 45, 75, 0.2, 2.0},
		{"Тестирование (testing)", 30, 90, 0.15, 1.8},
		{"Работа с JSON", 40, 80, 0.25, 2.2},
		{"Создание HTTP сервера", 90, 70, 0.4, 3.0},
	}

	// Битва с искушениями
	willpower := 100
	temptationPower := 0
	studyProgress := 0.0
	totalDopamine := 0.0

	fmt.Println("\n⚔️  БИТВА ВОЛИ vs ИСКУШЕНИЙ:")

	for i := 0; i < 5; i++ {
		fmt.Printf("\n──── ХОД %d ────\n", i+1)

		// Искушение атакует
		temptation := temptations[rand.Intn(len(temptations))]
		temptationPower = temptation.Power - (willpower / 2)

		if temptationPower > 0 {
			fmt.Printf("🎭 ИСКУШЕНИЕ: %s (сила: %d)\n", temptation.Name, temptation.Power)
			fmt.Printf("   'Можно на %d минут... всего один раз!'", temptation.Distraction)

			// Проверка силы воли
			resistRoll := rand.Intn(100)
			if resistRoll < willpower {
				fmt.Printf("\n   ✅ СОПРОТИВЛЕНИЕ! Сила воли: %d\n", willpower)
				willpower -= 20
			} else {
				fmt.Printf("\n   ❌ СДАЧА! Установил %s на %d минут\n",
					temptation.Name, temptation.Distraction)
				studyProgress -= 0.1
				willpower -= 40
				break
			}
		}

		// Учеба Go
		study := goStudies[rand.Intn(len(goStudies))]
		fmt.Printf("\n📚 ИЗУЧЕНИЕ GO: %s\n", study.Topic)
		fmt.Printf("   Фокус: %d%%, Время: %d мин\n", study.Focus, study.Minutes)

		if willpower > 30 {
			progress := float64(study.Minutes) * (float64(study.Focus) / 100.0) * 0.01
			studyProgress += progress
			totalDopamine += study.Dopamine
			willpower -= 15

			fmt.Printf("   ✅ ПРОГРЕСС: +%.2f%% | Дофамин: +%.1f\n",
				progress*100, study.Dopamine)
		} else {
			fmt.Printf("   ⚠  СЛИШКОМ УСТАЛ! Сила воли: %d\n", willpower)
			studyProgress += 0.05
			willpower += 10 // небольшой отдых
		}

		time.Sleep(800 * time.Millisecond)
	}

	// Итоги битвы
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("📊 ИТОГИ БИТВЫ С ЛОМКОЙ:")
	fmt.Printf("   Прогресс в Go: %.1f%%\n", studyProgress*100)
	fmt.Printf("   Сила воли: %d/100\n", max(0, willpower))
	fmt.Printf("   Долгосрочный дофамин: %.1f\n", totalDopamine)
	fmt.Printf("   Искушение CapCut: %s\n", getWithdrawalStatus(willpower))
	fmt.Printf("   Мама довольна: %s\n", getMomReaction(studyProgress))

	// Уровни
	level := calculateLevel(studyProgress, willpower)
	fmt.Printf("\n🏆 УРОВЕНЬ: %s\n", level)

	// Мотивация
	fmt.Println("\n💪 МОТИВАЦИЯ НА ЗАВТРА:")
	motivations := []string{
		"Завтра пойду не на доставку, а на собеседование по Go!",
		"Каждая строчка кода — шаг от бесконечных холодных улиц к тёплому офису.",
		"Макарошки с котлеткой вкусны, но зарплата джуна — вкуснее.",
		"CapCut подождёт. Сначала — Работа, потом — Хобби.",
		"Горутины не ждут. JSON не монтируется. Интерфейсы не рендерятся.",
	}

	rand.Shuffle(len(motivations), func(i, j int) {
		motivations[i], motivations[j] = motivations[j], motivations[i]
	})

	fmt.Printf("   '%s'\n", motivations[0])

	// Скрытая команда для победы над ломкой
	if studyProgress > 0.5 && willpower > 50 {
		fmt.Println("\n🎉 СЕКРЕТНАЯ КОМАНДА РАЗБЛОКИРОВАНА:")
		fmt.Println("   $ go run life.go --career=junior-golang")
		fmt.Println("   Результат: ✅ УСТРОЕН НА РАБОТУ")
	}
}

func simulateCourierDay() CourierDay {
	deliveries := []string{
		"Документы → Центр, Бульварное кольцо",
		"Медицинская доставка → Клиника на Ленинградском шоссе",
		"Косметика из Авиапарка → Флотская ул (Речной Вокзал)",
		"Аптечные товары → Флотская ул",
	}

	return CourierDay{
		Deliveries: deliveries,
		Money:      1850,
		Energy:     35,
		Mood:       "Устал, но в плюсе",
	}
}

func getWithdrawalStatus(willpower int) string {
	switch {
	case willpower >= 80:
		return "✅ ПОБЕЖДЕНО (не думаю о CapCut)"
	case willpower >= 50:
		return "⚠  КОНТРОЛЬ (изредка вспоминаю)"
	case willpower >= 20:
		return "😫 БОРЬБА (сильная ломка)"
	default:
		return "💀 СДАЛСЯ (установил обратно)"
	}
}

func getMomReaction(progress float64) string {
	if progress > 0.4 {
		return "✅ 'Сынок, я горжусь тобой!'"
	} else if progress > 0.2 {
		return "👍 'Учись, я поддержу котлетками!'"
	} else {
		return "😐 'Опять в свой компьютер уткнулся?'"
	}
}

func calculateLevel(progress float64, willpower int) string {
	score := int(progress*100) + willpower

	switch {
	case score >= 150:
		return "🚀 БУДУЩИЙ JUNIOR GOLANG"
	case score >= 100:
		return "🎯 УПОРНЫЙ УЧЕНИК"
	case score >= 50:
		return "🔄 КУРЬЕР НА ПЕРЕПУТЬЕ"
	default:
		return "💀 ЗАЛОЖНИК CAPCUT"
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
