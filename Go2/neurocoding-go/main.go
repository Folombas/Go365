package main

import (
	"fmt"
	"strings"
	"time"
)

// 🧠 Нейротрансмиттерные структуры
type (
	NeuroTransmitter struct {
		Name        string
		Level       float64
		Description string
		Emoji       string
	}

	CodeSession struct {
		StartTime    time.Time
		EndTime      time.Time
		LinesWritten int
		TestsPassed  int
		BugsFixed    int
		FocusLevel   float64 // 0.0 - 1.0
	}

	DopamineReward struct {
		Name        string
		Description string
		DopamineHit float64
		Unlocked    bool
	}
)

// 🎮 Глобальные переменные - наша "нейрохимия"
var (
	// 🧪 Нейротрансмиттеры программиста
	dopamine = NeuroTransmitter{
		Name:        "Дофамин",
		Level:       0.3,
		Description: "Ожидание награды, мотивация продолжать",
		Emoji:       "💥",
	}

	serotonin = NeuroTransmitter{
		Name:        "Серотонин",
		Level:       0.4,
		Description: "Удовлетворение от хорошо выполненной работы",
		Emoji:       "😌",
	}

	endorphins = NeuroTransmitter{
		Name:        "Эндорфины",
		Level:       0.2,
		Description: "Естественное обезболивающее, помогает терпеть трудности",
		Emoji:       "🛡️",
	}

	oxytocin = NeuroTransmitter{
		Name:        "Окситоцин",
		Level:       0.3,
		Description: "Социальная связь с коммьюнити Go",
		Emoji:       "🤝",
	}

	// 🏆 Дофаминовые награды за код
	dopamineRewards = []DopamineReward{
		{"FirstCompile", "Первая успешная компиляция", 0.1, true},
		{"TenLines", "Написано 10 строк чистого кода", 0.15, false},
		{"TestGreen", "Все тесты зелёные", 0.2, false},
		{"FeatureComplete", "Завершена новая фича", 0.25, false},
		{"PRMerged", "Pull Request принят в репозиторий", 0.3, false},
		{"BugSlayer", "Исправлен критический баг", 0.35, false},
		{"RefactorKing", "Рефакторинг улучшил код на 50%", 0.4, false},
		{"OpenSourceContrib", "Вклад в open-source проект", 0.5, false},
		{"GoJobOffer", "Получен оффер Go-разработчика", 1.0, false},
	}

	// 📊 Статистика нейрокодинга
	totalSessions    = 1
	totalLines       = 42
	totalTests       = 8
	totalBugs        = 3
	streakDays       = 1
	longestStreak    = 1
	productivity     = 0.68
)

func main() {
	// 🎯 Инициализация сессии кодинга
	session := startCodingSession()

	// 🧠 Нейро-дэшборд
	printNeuroDashboard()

	// 💡 Принципы дофаминового программирования
	printDopaminePrinciples()

	// 🎮 Сессия кодинга
	simulateCodingSession(&session)

	// 🏆 Проверка наград
	checkDopamineRewards()

	// 📈 Обновление нейрохимии
	updateNeurochemistry(session)

	// 🎯 Итоги сессии
	printSessionSummary(session)

	// 🔮 Мотивация на завтра
	printTomorrowMotivation()
}

// 🚀 Начало сессии кодинга
func startCodingSession() CodeSession {
	now := time.Now()
	fmt.Println(strings.Repeat("🧠", 70))
	fmt.Println("                    NEUROCODING-GO: ДОФАМИНОВОЕ ПРОГРАММИРОВАНИЕ")
	fmt.Println("                       GO365 | День 1 | Сессия 1")
	fmt.Println(strings.Repeat("🧠", 70))
	fmt.Printf("\n🕒 Начало сессии: %s\n", now.Format("15:04:05"))

	return CodeSession{
		StartTime:    now,
		FocusLevel:   0.7,
	}
}

// 📊 Нейро-дэшборд
func printNeuroDashboard() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🧪 НЕЙРОХИМИЧЕСКИЙ СТАТУС ПРОГРАММИСТА:")
	fmt.Println(strings.Repeat("─", 70))

	neuroTransmitters := []NeuroTransmitter{dopamine, serotonin, endorphins, oxytocin}

	for _, nt := range neuroTransmitters {
		bar := createNeuroBar(nt.Level, 20)
		fmt.Printf("%s %s: %.1f/1.0\n", nt.Emoji, nt.Name, nt.Level)
		fmt.Printf("   %s\n", bar)
		fmt.Printf("   📝 %s\n\n", nt.Description)
	}

	fmt.Println(strings.Repeat("─", 70))
	fmt.Printf("📊 Продуктивность: %.0f%% | 🔥 Серия дней: %d | 📅 Всего сессий: %d\n",
		productivity*100, streakDays, totalSessions)
}

// 💡 Принципы дофаминового программирования
func printDopaminePrinciples() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("💡 ПРИНЦИПЫ ДОФАМИНОВОГО ПРОГРАММИРОВАНИЯ:")
	fmt.Println(strings.Repeat("─", 70))

	principles := []string{
		"1. 🎯 МАЛЕНЬКИЕ ПОБЕДЫ: Разбивай задачи на микро-шаги (каждая компиляция = дофамин)",
		"2. ⏱️ ПОМОДОРО ТЕХНИКА: 25 минут кода → 5 минут отдыха (циклы дофаминовых пиков)",
		"3. 🏆 МГНОВЕННАЯ НАГРАДА: Каждый green test = микро-выброс дофамина",
		"4. 📈 ВИЗУАЛЬНЫЙ ПРОГРЕСС: Коммиты на GitHub как игровые достижения",
		"5. 🔄 ОБРАТНАЯ СВЯЗЬ: Чем быстрее тесты/компиляция, тем чаще дофаминовые удары",
		"6. 🎮 ИГРОВАЯ МЕХАНИКА: Уровни (junior/middle/senior), опыт (XP), ачивки",
		"7. 👥 СОЦИАЛЬНОЕ ПОДКРЕПЛЕНИЕ: Code review = социальное одобрение (окситоцин)",
		"8. 🧠 НЕЙРОПЛАСТИЧНОСТЬ: Каждый день Go перестраивает мозг под Go-паттерны",
	}

	for _, p := range principles {
		fmt.Println("   " + p)
	}
}

// 🎮 Симуляция сессии кодинга
func simulateCodingSession(session *CodeSession) {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🎮 СИМУЛЯЦИЯ СЕССИИ КОДИНГА:")
	fmt.Println(strings.Repeat("─", 70))

	actions := []struct{
		Action string
		DopamineChange float64
		Result string
	}{
		{"$ go run main.go", 0.05, "✅ Компиляция успешна! Программа запущена."},
		{"Добавление новой функции", 0.08, "✨ Функция calculate() реализована."},
		{"$ go test ./...", 0.12, "🎉 Все 5 тестов зелёные! 100% покрытие."},
		{"Рефакторинг кода", 0.07, "🔧 Улучшена читаемость, удалён дублирующий код."},
		{"Фикс бага с race condition", 0.15, "🐛 Критический баг исправлен с помощью sync.Mutex."},
		{"Документация и комменты", 0.03, "📚 Добавлены комментарии и примеры использования."},
		{"$ go fmt", 0.02, "🎨 Код отформатирован в соответствии с gofmt."},
		{"$ git commit -m 'feat: add concurrency safety'", 0.1, "💾 Коммит создан! История сохранилась."},
		{"$ git push origin main", 0.08, "🚀 Код отправлен в удалённый репозиторий."},
	}

	session.LinesWritten = 42
	session.TestsPassed = 8
	session.BugsFixed = 3

	for i, action := range actions {
		fmt.Printf("\n🎯 Шаг %d: %s\n", i+1, action.Action)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("   %s\n", action.Result)
		fmt.Printf("   💥 Дофамин +%.2f\n", action.DopamineChange)
		dopamine.Level = min(dopamine.Level + action.DopamineChange, 1.0)
	}

	session.EndTime = time.Now()
}

// 🏆 Проверка дофаминовых наград
func checkDopamineRewards() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🏆 ДОФАМИНОВЫЕ НАГРАДЫ ЗА ЭТУ СЕССИЮ:")
	fmt.Println(strings.Repeat("─", 70))

	newRewards := 0
	for i := range dopamineRewards {
		// Симулируем разблокировку некоторых наград
		if !dopamineRewards[i].Unlocked && dopamine.Level > dopamineRewards[i].DopamineHit*0.8 {
			dopamineRewards[i].Unlocked = true
			newRewards++
			fmt.Printf("   🎉 РАЗБЛОКИРОВАНО: %s\n", dopamineRewards[i].Name)
			fmt.Printf("      %s (+%.1f дофамина)\n", dopamineRewards[i].Description, dopamineRewards[i].DopamineHit)
			dopamine.Level = min(dopamine.Level + dopamineRewards[i].DopamineHit/10, 1.0)
		}
	}

	if newRewards == 0 {
		fmt.Println("   📭 Новых наград нет. Продолжай кодить!")
	}

	fmt.Printf("\n📊 Статус наград: %d/%d разблокировано\n",
		countUnlockedRewards(), len(dopamineRewards))
}

// 📈 Обновление нейрохимии после сессии
func updateNeurochemistry(session CodeSession) {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🧪 ОБНОВЛЕНИЕ НЕЙРОХИМИИ ПОСЛЕ СЕССИИ:")
	fmt.Println(strings.Repeat("─", 70))

	// Дофамин: за завершение сессии
	serotonin.Level = min(serotonin.Level + 0.1, 1.0) // Удовлетворение
	endorphins.Level = min(endorphins.Level + 0.05, 1.0) // Преодоление трудностей
	oxytocin.Level = min(oxytocin.Level + 0.03, 1.0) // Чувство принадлежности к коммьюнити

	fmt.Println("   💥 Дофамин:  Мотивация для следующей сессии (+0.15)")
	fmt.Println("   😌 Серотонин: Удовлетворение от выполненной работы (+0.10)")
	fmt.Println("   🛡️ Эндорфины: Устойчивость к сложностям (+0.05)")
	fmt.Println("   🤝 Окситоцин: Связь с Go-коммьюнити (+0.03)")

	productivity = 0.72
	streakDays++
	if streakDays > longestStreak {
		longestStreak = streakDays
	}
}

// 📊 Итоги сессии
func printSessionSummary(session CodeSession) {
	duration := session.EndTime.Sub(session.StartTime)

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("📊 ИТОГИ СЕССИИ NEUROCODING-GO:")
	fmt.Println(strings.Repeat("─", 70))

	fmt.Printf("   🕒 Длительность: %v\n", duration.Round(time.Minute))
	fmt.Printf("   📝 Написано строк: %d\n", session.LinesWritten)
	fmt.Printf("   ✅ Пройдено тестов: %d\n", session.TestsPassed)
	fmt.Printf("   🐛 Исправлено багов: %d\n", session.BugsFixed)
	fmt.Printf("   🎯 Уровень фокуса: %.0f%%\n", session.FocusLevel*100)
	fmt.Printf("   💥 Итоговый уровень дофамина: %.1f/1.0\n", dopamine.Level)

	// Продуктивность сессии
	efficiency := float64(session.LinesWritten) / duration.Minutes()
	fmt.Printf("   ⚡ Эффективность: %.1f строк/минуту\n", efficiency)

	if efficiency > 2.0 {
		fmt.Println("\n   🚀 ОТЛИЧНЫЙ РЕЗУЛЬТАТ! Ты в потоке!")
	} else {
		fmt.Println("\n   👍 ХОРОШАЯ РАБОТА! Каждая строка кода приближает к цели.")
	}
}

// 🔮 Мотивация на завтра
func printTomorrowMotivation() {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🔮 НЕЙРОПРОГНОЗ НА ЗАВТРАШНЮЮ СЕССИЮ:")
	fmt.Println(strings.Repeat("─", 70))

	motivations := []string{
		"🧠 Твой мозг уже начал формировать нейронные пути для Go!",
		"💪 Завтра будет легче: нейропластичность работает на тебя.",
		"🎯 Микро-цель на завтра: написать 50 строк кода или пройти 10 тестов.",
		"🏆 Следующая дофаминовая награда: 'TenLines' (напиши ещё 8 строк).",
		"📈 Если продолжишь стрик 7 дней: +0.3 к базовому уровню дофамина!",
		"",
		"💡 ПОМНИ: Каждый раз когда ты пишешь `go run`, а не открываешь игру,",
		"         ты перепрограммируешь свою систему вознаграждения.",
		"",
		"🎮 Раньше: Игры → Дофамин → Зависимость от игр",
		"🚀 Теперь: Go-код → Дофамин → Зависимость от роста",
		"",
		"🐹 Гофер гордится тобой! Каждая горутина в твоём коде —",
		"   это новый нейрон в твоём мозге.",
	}

	for _, m := range motivations {
		fmt.Println("   " + m)
	}

	fmt.Println("\n" + strings.Repeat("🧠", 70))
	fmt.Println("                 ДО ЗАВТРА! ПУСТЬ ДОФАМИН РАБОТАЕТ НА ТЕБЯ!")
	fmt.Println(strings.Repeat("🧠", 70))
}

// 🛠️ Вспомогательные функции
func createNeuroBar(level float64, width int) string {
	bar := strings.Builder{}
	bar.Grow(width)
	filled := int(level * float64(width))

	for i := 0; i < width; i++ {
		if i < filled {
			switch {
			case i < width/3:
				bar.WriteString("█") // Красный
			case i < 2*width/3:
				bar.WriteString("▓") // Жёлтый
			default:
				bar.WriteString("░") // Зелёный
			}
		} else {
			bar.WriteString(" ")
		}
	}

	return bar.String()
}

func countUnlockedRewards() int {
	count := 0
	for _, r := range dopamineRewards {
		if r.Unlocked {
			count++
		}
	}
	return count
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
