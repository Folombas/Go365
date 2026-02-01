package main

import (
	"fmt"
	"strings"
	"time"
)

type Distraction struct {
	Name        string
	Description string
	XPChange    int
	DangerLevel int // 1-10
}

type Player struct {
	XP                   int
	Level                string
	FocusHours           int
	DistractionsResisted int
}

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                 FOCUS FIGHTER v1.0                      ║")
	fmt.Println("║       Битва за концентрацию в изучении Go              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")

	fmt.Println("\n📖 ЛОГ ДНЯ: Воскресенье, 1 февраля 2026")
	fmt.Println("   Температура: -20°C, Гоша остался дома")
	fmt.Println("   Статус: Нищий курьер | Цель: Go-разработчик в офисе")
	fmt.Println("   ────────────────────────────────────────")

	player := Player{XP: 0, Level: "Нищий курьер", FocusHours: 0, DistractionsResisted: 0}

	distractions := []Distraction{
		{"🎮 Видеоигры", "Стрелялка или стратегия?", -20, 8},
		{"🎬 Фильм/сериал", "Всего одна серия...", -15, 7},
		{"✂️ Монтаж видео", "Вспомнить отпуск в Сочи", -25, 6},
		{"🍻 Бар/клуб", "Всего пару пив", -30, 9},
		{"🤖 n8n", "Создать ещё одного бота-синоптика", -50, 10},
		{"📱 Соцсети", "Просто проверить уведомления", -10, 5},
		{"😴 Дремота", "Всего 20 минут...", -5, 3},
		{"🍔 Заказ еды", "Хочется пиццы, а не готовить", -8, 4},
	}

	hoursInDay := 16 // С 8 утра до 12 ночи
	currentHour := 8

	fmt.Println("\n⏰ ДЕНЬ НАЧИНАЕТСЯ! У тебя 16 часов. Каждый час - выбор.")
	fmt.Println("   Изучение Go: +10 XP | Отвлечение: потеря XP")
	fmt.Println("   ────────────────────────────────────────")

	for hour := 1; hour <= hoursInDay; hour++ {
		fmt.Printf("\n🕐 ЧАС %d (%02d:00) | Текущий уровень: %s | XP: %d\n",
			hour, currentHour, player.Level, player.XP)

		// Каждые 2 часа появляется особо опасный враг
		if hour%2 == 0 {
			showBossWarning(distractions[4]) // n8n - босс
		}

		fmt.Println("Выбери действие:")
		fmt.Println("1. 🚀 Учить Go (курс, 1 час) +10 XP")
		fmt.Println("2. 📚 Читать документацию +8 XP")
		fmt.Println("3. 💻 Практика (написать код) +12 XP")
		fmt.Println("4. 🎯 Сопротивляться отвлечению (выбрать случайное)")
		fmt.Println("5. ❌ Сдаться и отвлечься (потеря XP)")

		var choice int
		fmt.Print("Твой выбор (1-5): ")
		fmt.Scan(&choice)

		switch choice {
		case 1, 2, 3:
			xpGain := 10
			if choice == 2 {
				xpGain = 8
			} else if choice == 3 {
				xpGain = 12
			}
			player.XP += xpGain
			player.FocusHours++
			fmt.Printf("✅ +%d XP! Фокус усилен. Всего часов фокуса: %d\n", xpGain, player.FocusHours)
			showMotivation()

		case 4:
			// Сопротивление отвлечению
			resistSuccess := resistDistraction()
			if resistSuccess {
				player.DistractionsResisted++
				player.XP += 5
				fmt.Printf("🛡️ Ты устоял! +5 XP за силу воли. Всего устоял: %d раз\n",
					player.DistractionsResisted)
			} else {
				// Случайное отвлечение всё равно происходит
				distraction := distractions[hour%len(distractions)]
				player.XP += distraction.XPChange
				fmt.Printf("💥 Не устоял: %s %d XP\n", distraction.Name, distraction.XPChange)
				showWarning(distraction)
			}

		case 5:
			// Сознательное отвлечение
			distraction := distractions[hour%len(distractions)]
			player.XP += distraction.XPChange
			fmt.Printf("💀 Сознательное отвлечение: %s %d XP\n",
				distraction.Name, distraction.XPChange)
			showWarning(distraction)

		default:
			fmt.Println("⏭️ Пропущен час... XP не изменилась")
		}

		// Проверяем уровень
		player.Level = updateLevel(player.XP)

		// Каждый 4-й час - проверка прогресса
		if hour%4 == 0 {
			showProgressReport(player)
		}

		currentHour++
		time.Sleep(500 * time.Millisecond) // Небольшая пауза для реализма
	}

	// Финальные результаты
	showFinalResults(player)
}

func showBossWarning(boss Distraction) {
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Printf(" ║    🚨 БОСС-ИСКУСИТЕЛЬ: %s 🚨     ║\n", boss.Name)
	fmt.Printf(" ║    Опасность: %d/10 | Потеря XP: %d       ║\n", boss.DangerLevel, boss.XPChange)
	fmt.Println("║    'Создай бота, это же проще, чем учить Go!'  ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
}

func resistDistraction() bool {
	// 70% шанс успешного сопротивления
	return time.Now().Unix()%10 < 7
}

func showMotivation() {
	motivations := []string{
		"Каждая строка кода на Go — это гвоздь в крышку гроба твоей карьеры курьера!",
		"Помни: офис с кондиционером ждёт тебя, но только после 1000 XP!",
		"n8n никуда не денется, а время на обучение Go — ограничено!",
		"Лучше сегодня написать 100 строк кода, чем завтра развозить 100 пицц!",
		"Интерфейсы в Go проще, чем интерфейсы с голодными клиентами доставки!",
	}
	fmt.Printf("💪 %s\n", motivations[time.Now().Unix()%int64(len(motivations))])
}

func showWarning(distraction Distraction) {
	warnings := []string{
		"Сначала стань разработчиком, потом развлекайся с %s!",
		"Каждая минута на %s отдаляет тебя от офисного кресла!",
		"Ты хочешь быть мастером %s или мастером Go? Выбирай!",
		"%s подождут, а возраст для смены карьеры — нет!",
	}
	warning := warnings[time.Now().Unix()%int64(len(warnings))]
	fmt.Printf("⚠️  %s\n", fmt.Sprintf(warning, distraction.Name))
}

func updateLevel(xp int) string {
	switch {
	case xp < 0:
		return "💀 Полный провал"
	case xp < 50:
		return "🚴 Нищий курьер"
	case xp < 150:
		return "👶 Начинающий гофер"
	case xp < 300:
		return "🎓 Ученик Go"
	case xp < 500:
		return "💼 Младший разработчик"
	case xp < 800:
		return "👨‍💻 Мидл-разработчик"
	case xp < 1200:
		return "👨‍🔬 Старший разработчик"
	default:
		return "🏆 Готов к офису!"
	}
}

func showProgressReport(player Player) {
	fmt.Println("\n📊 ОТЧЁТ О ПРОГРЕССЕ:")
	fmt.Printf("   Уровень: %s\n", player.Level)
	fmt.Printf("   Всего XP: %d\n", player.XP)
	fmt.Printf("   Часов фокуса: %d\n", player.FocusHours)
	fmt.Printf("   Отвлечений отбито: %d\n", player.DistractionsResisted)

	if player.XP < 100 {
		fmt.Println("   ❗ Ты всё ещё ближе к доставке пицц, чем к офису!")
	} else if player.XP < 300 {
		fmt.Println("   ✅ Хороший старт! Продолжай в том же духе!")
	} else {
		fmt.Println("   🚀 Отличный прогресс! Офис уже близко!")
	}
	fmt.Println("   ────────────────────────────────────────")
}

func showFinalResults(player Player) {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("🎯 ИТОГИ ДНЯ")
	fmt.Println(strings.Repeat("═", 60))

	fmt.Printf("Финальный уровень: %s\n", player.Level)
	fmt.Printf("Всего заработано XP: %d\n", player.XP)
	fmt.Printf("Часов сфокусированной работы: %d\n", player.FocusHours)
	fmt.Printf("Отвлечений успешно отбито: %d\n", player.DistractionsResisted)

	fmt.Println("\n" + strings.Repeat("─", 60))

	if player.XP >= 500 {
		fmt.Println("🏆 ПОБЕДА! Ты доказал, что можешь фокусироваться!")
		fmt.Println("   Завтра ты сможешь подняться ещё на уровень выше!")
		fmt.Println("   Помни: офис с кондиционером уже не за горами!")
	} else if player.XP >= 200 {
		fmt.Println("👍 НЕПЛОХО! Ты на правильном пути, но можно лучше.")
		fmt.Println("   Завтра постарайся избегать n8n и видеоигр.")
		fmt.Println("   Каждый час фокуса = шаг к нормальной зарплате!")
	} else {
		fmt.Println("💪 ЗАВТРА НОВЫЙ ДЕНЬ! Не сдавайся!")
		fmt.Println("   Помни: ты учишься не для галочки, а для карьеры!")
		fmt.Println("   n8n подождёт, твоё будущее — нет!")
	}

	fmt.Println("\n🧠 10 МОТИВАЦИОННЫХ ФРАЗ НА ЗАВТРА:")
	motivations := []string{
		"1. Go — это не просто язык, это билет в мир нормальных зарплат",
		"2. Каждый день без изучения Go — день в пользу доставки пицц",
		"3. Горутины масштабируются лучше, чем твой доход курьером",
		"4. Интерфейсы в Go проще, чем объяснять клиентам, почему доставка задерживается",
		"5. Пока другие играют в игры, ты строишь свою карьеру",
		"6. n8n — это игрушка для тех, кто уже имеет стабильный доход",
		"7. Видео из отпуска можно смонтировать, когда будешь в отпуске с работы разработчика",
		"8. Каждый коммит в GitHub ценнее, чем каждый доставленный заказ",
		"9. Система типов Go строже, чем начальник службы доставки",
		"10. Завтра тот день, когда твой будущий работодатель посмотрит твой GitHub",
	}

	for _, m := range motivations {
		fmt.Println("   " + m)
	}

	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("DISCLAIMER: Все персонажи вымышлены. Любые совпадения случайны.")
	fmt.Println("Цель проекта — мотивация к изучению Go через игровые механики.")
	fmt.Println(strings.Repeat("═", 60))
}
