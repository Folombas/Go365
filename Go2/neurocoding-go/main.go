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
	// 📅 Челлендж-переменные
	challengeStart   = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDayNumber int
	moduleTopic      = "Нейрокодинг и дофаминовая зависимость"

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
	totalSessions = 1
	totalLines    = 42
	totalTests    = 8
	totalBugs     = 3
	streakDays    = 1
	longestStreak = 1
	productivity  = 0.68
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

	// Вычисляем текущий день челленджа
	currentDayNumber = calculateCurrentDay()

	fmt.Println(strings.Repeat("🧠", 70))
	fmt.Printf("                    NEUROCODING-GO: ДОФАМИНОВОЕ ПРОГРАММИРОВАНИЕ\n")
	fmt.Printf("                       GO365 | Go%d | Тема: %s\n", currentDayNumber, moduleTopic)
	fmt.Println(strings.Repeat("🧠", 70))

	// 📊 Статистика прогресса
	fmt.Printf("\n📅 День челленджа: Go%d (из 365)\n", currentDayNumber)
	fmt.Printf("📊 Прогресс: %.1f%% завершено\n", float64(currentDayNumber)/365*100)
	fmt.Printf("🕒 Начало сессии: %s\n", now.Format("15:04:05"))

	return CodeSession{
		StartTime:  now,
		FocusLevel: 0.7,
	}
}

// 📅 Вычисление текущего дня челленджа
func calculateCurrentDay() int {
	now := time.Now().UTC()
	daysDiff := int(now.Sub(challengeStart).Hours() / 24)

	// Если текущая дата раньше начала челленджа
	if daysDiff < 0 {
		return 1
	}

	// Если прошло больше 365 дней
	if daysDiff > 364 {
		return 365
	}

	return daysDiff + 1 // +1 потому что Go1 - это 1 января
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
	fmt.Printf("🎯 День челленджа: Go%d | 🐹 Гофер-уровень: %.0f%%\n",
		currentDayNumber, float64(currentDayNumber)/365*100)
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
	fmt.Printf("🎮 СИМУЛЯЦИЯ СЕССИИ КОДИНГА (Go%d):\n", currentDayNumber)
	fmt.Println(strings.Repeat("─", 70))

	actions := []struct {
		Action         string
		DopamineChange float64
		Result         string
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

	session.LinesWritten = 42 + currentDayNumber*3 // Больше строк в последующие дни
	session.TestsPassed = 8 + currentDayNumber
	session.BugsFixed = 3

	for i, action := range actions {
		fmt.Printf("\n🎯 Шаг %d: %s\n", i+1, action.Action)
		fmt.Printf("   %s\n", action.Result)
		fmt.Printf("   💥 Дофамин +%.2f\n", action.DopamineChange)
		dopamine.Level = min(dopamine.Level+action.DopamineChange, 1.0)
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
		// Автоматически разблокируем награды в зависимости от дня челленджа
		if !dopamineRewards[i].Unlocked && currentDayNumber >= getDayForReward(dopamineRewards[i].Name) {
			dopamineRewards[i].Unlocked = true
			newRewards++
			fmt.Printf("   🎉 РАЗБЛОКИРОВАНО: %s\n", dopamineRewards[i].Name)
			fmt.Printf("      %s (+%.1f дофамина)\n", dopamineRewards[i].Description, dopamineRewards[i].DopamineHit)
			dopamine.Level = min(dopamine.Level+dopamineRewards[i].DopamineHit/10, 1.0)
		}
	}

	if newRewards == 0 {
		fmt.Println("   📭 Новых наград нет. Продолжай кодить!")
	}

	fmt.Printf("\n📊 Статус наград: %d/%d разблокировано\n",
		countUnlockedRewards(), len(dopamineRewards))
}

// 📅 Получение дня для разблокировки награды
func getDayForReward(rewardName string) int {
	rewardDays := map[string]int{
		"FirstCompile":      1,
		"TenLines":          2,
		"TestGreen":         5,
		"FeatureComplete":   10,
		"PRMerged":          15,
		"BugSlayer":         20,
		"RefactorKing":      30,
		"OpenSourceContrib": 60,
		"GoJobOffer":        365,
	}

	if day, exists := rewardDays[rewardName]; exists {
		return day
	}
	return 999 // Очень далеко
}

// 📈 Обновление нейрохимии после сессии
func updateNeurochemistry(session CodeSession) {
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("🧪 ОБНОВЛЕНИЕ НЕЙРОХИМИИ ПОСЛЕ СЕССИИ:")
	fmt.Println(strings.Repeat("─", 70))

	// Нейрохимия улучшается с каждым днём челленджа
	dailyImprovement := float64(currentDayNumber) * 0.005

	serotonin.Level = min(serotonin.Level+0.1+dailyImprovement, 1.0)
	endorphins.Level = min(endorphins.Level+0.05+dailyImprovement, 1.0)
	oxytocin.Level = min(oxytocin.Level+0.03+dailyImprovement, 1.0)

	fmt.Printf("   💥 Дофамин:  Мотивация для следующей сессии (+0.15)\n")
	fmt.Printf("   😌 Серотонин: Удовлетворение от выполненной работы (+%.2f)\n", 0.1+dailyImprovement)
	fmt.Printf("   🛡️ Эндорфины: Устойчивость к сложностям (+%.2f)\n", 0.05+dailyImprovement)
	fmt.Printf("   🤝 Окситоцин: Связь с Go-коммьюнити (+%.2f)\n", 0.03+dailyImprovement)

	productivity = 0.68 + (float64(currentDayNumber) * 0.001)
	streakDays = currentDayNumber
	if streakDays > longestStreak {
		longestStreak = streakDays
	}
}

// 📊 Итоги сессии
func printSessionSummary(session CodeSession) {
	duration := session.EndTime.Sub(session.StartTime)

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Printf("📊 ИТОГИ СЕССИИ NEUROCODING-GO (Go%d):\n", currentDayNumber)
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

	// Оценка сессии в зависимости от дня
	if currentDayNumber <= 7 {
		fmt.Println("\n   🐣 НАЧАЛО ПУТИ! Первая неделя Go365 - ты уже молодец!")
	} else if currentDayNumber <= 30 {
		fmt.Println("\n   🚀 ОТЛИЧНЫЙ СТАРТ! Месяц обучения - фундамент заложен!")
	} else if currentDayNumber <= 100 {
		fmt.Println("\n   💪 СИЛЬНЫЙ ПРОГРЕСС! 100 дней Go - ты на верном пути!")
	} else {
		fmt.Println("\n   🏆 ВПЕЧАТЛЯЮЩЕ! Продолжаешь движение к цели!")
	}
}

// 🔮 Мотивация на завтра
func printTomorrowMotivation() {
	nextDay := currentDayNumber + 1

	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Printf("🔮 НЕЙРОПРОГНОЗ НА ЗАВТРА (Go%d):\n", nextDay)
	fmt.Println(strings.Repeat("─", 70))

	motivations := []string{
		fmt.Sprintf("🧠 Go%d: Твой мозг формирует нейронные пути для Go!", nextDay),
		"💪 Завтра будет легче: нейропластичность работает на тебя.",
		fmt.Sprintf("🎯 Микро-цель на Go%d: написать %d строк кода", nextDay, 40+nextDay*2),
		fmt.Sprintf("🏆 Следующая дофаминовая награда: 'TenLines' (напиши ещё %d строк).", 8),
		fmt.Sprintf("📈 Если продолжишь стрик %d дней: +%.1f к базовому уровню дофамина!", nextDay+1, float64(nextDay)*0.005),
		"",
		"💡 ПОМНИ: Каждый раз когда ты пишешь `go run`, а не открываешь игру,",
		"         ты перепрограммируешь свою систему вознаграждения.",
		"",
		"🎮 Раньше: Игры → Дофамин → Зависимость от игр",
		"🚀 Теперь: Go-код → Дофамин → Зависимость от роста",
		"",
		fmt.Sprintf("🐹 Гофер гордится тобой! Go%d - это %d новых нейронов для Go!",
			nextDay, nextDay*1000),
	}

	for _, m := range motivations {
		fmt.Println("   " + m)
	}

	fmt.Println("\n" + strings.Repeat("🧠", 70))
	fmt.Printf("                 Go%d ЖДЁТ! ПУСТЬ ДОФАМИН РАБОТАЕТ НА ТЕБЯ!\n", nextDay)
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
