package main

import (
	"fmt"
	"time"
)

// Decision представляет одно решение дня
type Decision struct {
	Description string
	Dopamine    float64 // Изменение дофамина
	Energy      int     // Затраты энергии
	Risk        string  // Уровень риска
}

// simulateDay симулирует драматичный день курьера-разработчика
func simulateDay() (float64, string) {
	fmt.Println("🎭 СИМУЛЯТОР: 'ДРАМА ОТМЕНЕННЫХ БОТИНОК И ДОФАМИНОВЫЕ КАЧЕЛИ'")
	fmt.Println("=============================================================")

	decisions := []Decision{
		{"Позвонить клиенту с ботинками (риск 500 рублей)", -0.5, 15, "СРЕДНИЙ"},
		{"Отменить заказ по телефону (не идти напрасно)", 1.2, -10, "НИЗКИЙ"},
		{"Идти пешком через МКАД в Химки по сугробам", -1.8, 85, "ВЫСОКИЙ"},
		{"Забрать 2 заказа рядом (эффективность!)", 2.3, 25, "НИЗКИЙ"},
		{"Новый автобус 1344 сработал по расписанию", 1.5, -30, "СРЕДНИЙ"},
		{"Вежливый охранник подсказал про ящик", 0.8, -5, "НИЗКИЙ"},
		{"Купить сметану и фетаксу как послушный сын", 0.3, 20, "НИЗКИЙ"},
		{"Лечь спать без уборки (тяжелый день)", 0.7, -40, "НИЗКИЙ"},
	}

	totalDopamine := 0.0
	fmt.Println("⏰ 10:15 - Пробуждение. Уровень дофамина: нейтральный")

	for i, d := range decisions {
		fmt.Printf("[Решение %d] %s\n", i+1, d.Description)
		fmt.Printf("   Дофамин: %+.1f | Энергия: %d | Риск: %s\n",
			d.Dopamine, d.Energy, d.Risk)

		totalDopamine += d.Dopamine
		time.Sleep(700 * time.Millisecond)

		// Драматические моменты
		if d.Risk == "ВЫСОКИЙ" && d.Dopamine < 0 {
			fmt.Println("   ⚠  КРИТИЧЕСКАЯ СИТУАЦИЯ! Мотивация падает...")
		} else if d.Dopamine > 1.0 {
			fmt.Println("   ✨ ВСПЛЕСК МОТИВАЦИИ! Хочется кодить!")
		}
		fmt.Println()
	}

	// Определяем итоговое состояние
	mood := getMentalState(totalDopamine)
	return totalDopamine, mood
}

// getMentalState определяет психическое состояние по дофамину
func getMentalState(dopamine float64) string {
	switch {
	case dopamine >= 5.0:
		return "ЭЙФОРИЯ КОДА 💻✨"
	case dopamine >= 2.0:
		return "УСТОЙЧИВАЯ МОТИВАЦИЯ 👍"
	case dopamine >= 0.0:
		return "ЛЕГКАЯ ФРУСТРАЦИЯ 😐"
	case dopamine >= -2.0:
		return "ЦИНИЧНОЕ ОТЧАЯНИЕ 😒"
	default:
		return "ЭКЗИСТЕНЦИАЛЬНЫЙ КРИЗИС 🆘"
	}
}

// getMotivation возвращает мотивационную фразу
func getMotivation(dopamine float64) string {
	motivations := []string{
		"Твой код компилируется с первого раза! Продолжай!",
		"Горутины танцуют в идеальной синхронизации!",
		"Ещё один коммит, и ты приблизишь финиш марафона!",
		"Помни: каждый net/http хендлер когда-то был новичком.",
		"Уборку можно отложить, а сбор мусора в Go — нет!",
		"Даже Брайан Керниган иногда искал ошибки 3 часа...",
		"Химки-Ховрино-Смоленская — и всё ради этого коммита!",
		"Завтра твой код будет благодарен сегодняшнему тебе.",
	}

	idx := int(time.Now().UnixNano()) % len(motivations)
	return motivations[idx]
}

// getTomorrowDecision возвращает решение на завтра
func getTomorrowDecision(mood string) string {
	switch mood {
	case "ЭЙФОРИЯ КОДА 💻✨":
		return "Писать код 10 часов подряд, забыв про уборку"
	case "УСТОЙЧИВАЯ МОТИВАЦИЯ 👍":
		return "Утром — код, вечером — пылесосить"
	case "ЛЕГКАЯ ФРУСТРАЦИЯ 😐":
		return "Отложить уборку ещё на день"
	default:
		return "Лежать и смотреть в потолок"
	}
}

func main() {
	// Запускаем симуляцию дня
	dopamine, mood := simulateDay()

	// Выводим итоги
	fmt.Println("════════════════════════════════════════════════════")
	fmt.Printf("📊 ИТОГ ДНЯ: 17 января 2026\n")
	fmt.Printf("   Дофаминовый баланс: %.2f пунктов\n", dopamine)
	fmt.Printf("   Психическое состояние: %s\n", mood)
	fmt.Printf("   Мотивация: %s\n", getMotivation(dopamine))
	fmt.Printf("   Решение на завтра: %s\n", getTomorrowDecision(mood))
	fmt.Println("════════════════════════════════════════════════════")

	// Геймификация
	fmt.Println("\n🎮 УРОВНИ ПСИХИЧЕСКОЙ УСТОЙЧИВОСТИ:")
	fmt.Println("   5.0+  → 🏆 ЭЙФОРИЯ КОДА")
	fmt.Println("   2.0+  → 👍 УСТОЙЧИВАЯ МОТИВАЦИЯ")
	fmt.Println("   0.0+  → 😐 ЛЕГКАЯ ФРУСТРАЦИЯ")
	fmt.Println("   -2.0+ → 😒 ЦИНИЧНОЕ ОТЧАЯНИЕ")
	fmt.Println("   ниже  → 🆘 КРИЗИС")

	fmt.Printf("\n🎯 Ваша цель на завтра: достичь 'ЭЙФОРИИ КОДА'!\n")
}
