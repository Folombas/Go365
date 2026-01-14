package main

import (
	"fmt"
	"strings"
	"time"
)

// ==================== ДИСКЛЕЙМЕР ====================
/*
ВНИМАНИЕ: Daily Code Life Story - ХУДОЖЕСТВЕННЫЙ ВЫМЫСЕЛ

Все персонажи, события, истории и сюжеты в модулях Daily Code Life Story
являются ПЛОДОМ ТВОРЧЕСКОГО ВООБРАЖЕНИЯ и созданы исключительно для:

1. ОБРАЗОВАТЕЛЬНЫХ ЦЕЛЕЙ изучения программирования
2. МОТИВАЦИОННОГО КОНТЕКСТА для обучения
3. ДЕМОНСТРАЦИИ технических концепций через нарратив

ВАЖНЫЕ ПОЛОЖЕНИЯ:
- Все персонажи ВЫДУМАНЫ (Гоша, его окружение, тролли)
- Все события ПРИДУМАНЫ для повествования
- Все диалоги и ситуации ЯВЛЯЮТСЯ ФИКЦИЕЙ
- Любые совпадения с реальными людьми, событиями или организациями
  СЛУЧАЙНЫ и НЕПРЕДНАМЕРЕННЫ

Это ХУДОЖЕСТВЕННОЕ ПРОИЗВЕДЕНИЕ в формате программного кода,
где техническое содержание РЕАЛЬНО, а сюжетная обёртка - ВЫМЫСЕЛ.

Цель: сделать обучение Go увлекательным через storytelling.
*/

// ==================== СТРУКТУРЫ ДАННЫХ ====================

type DayStats struct {
	Date           string
	WakeUpTime     string
	Breakfast      string
	Deliveries     int
	Weather        string
	BootsCondition string
	PagesRead      int
	Motivation     int
	DockerInsights []string
	GoLinesWritten int
}

type DockerBook struct {
	Title       string
	Author      string
	PagesTotal  int
	PagesRead   int
	Language    string // Язык, на котором написан Docker
	IsPhysical  bool
	GiftSource  string
}

type WetBoots struct {
	IsWet       bool
	WetnessLevel int // 0-100%
	MotivationBoost int
	Discomfort  int
}

// ==================== КОНСТАНТЫ И МОТИВАЦИЯ ====================

const (
	DOCKER_WRITTEN_IN_GO = "Go"
	MAX_MOTIVATION       = 1000
	BASE_MOTIVATION      = 100
)

// ==================== ОСНОВНАЯ ПРОГРАММА ====================

func main() {
	fmt.Println("🚀 ДЕНЬ 14.01.2026: МОКРЫЕ БОТИНКИ И ПУТЬ К DOCKER")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("📖 СЮЖЕТ: Снег, промокшие ноги, бумажная книга и осознание связи Docker & Go")

	// Инициализация дня
	stats := initializeDay()
	book := initializeDockerBook()
	boots := initializeWetBoots()

	// Симуляция дня
	simulateMorning(stats, boots)
	readInTrain(stats, book)
	realizeDockerGoConnection(stats, book)
	calculateMotivation(stats, boots, book)
	writeGoCode(stats)
	printDaySummary(stats, book, boots)
	printPhilosophicalInsights()
}

func initializeDay() *DayStats {
	return &DayStats{
		Date:       "14.01.2026",
		WakeUpTime: "06:30",
		Breakfast:  "кашка пшёнка",
		Deliveries: 3,
		Weather:    "снегопад",
		BootsCondition: "промокшие насквозь",
		PagesRead:  0,
		Motivation: BASE_MOTIVATION,
		DockerInsights: []string{},
		GoLinesWritten: 0,
	}
}

func initializeDockerBook() *DockerBook {
	return &DockerBook{
		Title:      "Docker: Полное руководство",
		Author:     "Эксперт по контейнеризации",
		PagesTotal: 450,
		PagesRead:  0,
		Language:   DOCKER_WRITTEN_IN_GO,
		IsPhysical: true,
		GiftSource: "подарок на Новый Год",
	}
}

func initializeWetBoots() *WetBoots {
	return &WetBoots{
		IsWet:       true,
		WetnessLevel: 85,
		MotivationBoost: 0,
		Discomfort:  70,
	}
}

func simulateMorning(stats *DayStats, boots *WetBoots) {
	fmt.Println("🌅 УТРЕННИЙ СТАРТ:")
	fmt.Println(strings.Repeat("─", 40))

	events := []struct{
		time string
		action string
		emoji string
	}{
		{"06:30", "Будильник! Подъём в полседьмого", "⏰"},
		{"06:45", "Умывание и душ", "🚿"},
		{"07:00", "Завтрак: " + stats.Breakfast, "🍲"},
		{"07:30", "Выход на улицу - начался снегопад", "❄️"},
		{"08:00", "Доставка 1-го заказа", "📦"},
		{"09:30", "Доставка 2-го заказа", "📦📦"},
		{"10:45", "Доставка 3-го заказа", "📦📦📦"},
		{"11:00", "Констатация: ботинки промокли", "👢💦"},
	}

	for _, event := range events {
		fmt.Printf("   %s %s\n", event.emoji, event.action)
		time.Sleep(200 * time.Millisecond)

		// Особый эффект для мокрых ботинок
		if strings.Contains(event.action, "промокли") {
			fmt.Println("      💭 Мысль: 'Ноги мокрые, но... может, это знак?'")
			boots.MotivationBoost = 25
			stats.Motivation += boots.MotivationBoost
		}
	}

	fmt.Printf("\n   📊 Итог утра: %d заказа, %s погода\n", stats.Deliveries, stats.Weather)
	fmt.Printf("   👢 Состояние обуви: %s (дискомфорт: %d%%)\n", stats.BootsCondition, boots.Discomfort)
}

func readInTrain(stats *DayStats, book *DockerBook) {
	fmt.Println("\n🚂 ЧТЕНИЕ В ЭЛЕКТРИЧКЕ:")
	fmt.Println(strings.Repeat("─", 40))

	fmt.Printf("   📚 Открываю книгу: «%s»\n", book.Title)
	fmt.Printf("   🎁 %s, %d страниц\n", book.GiftSource, book.PagesTotal)
	fmt.Println("   👁️  Преимущество бумажной книги: глаза не устают от маленького экрана")

	// Чтение страниц
	pagesRead := 42
	book.PagesRead = pagesRead
	stats.PagesRead = pagesRead

	fmt.Printf("   📖 Прочитано страниц: %d\n", pagesRead)

	insights := []string{
		"Контейнеризация - это изоляция процессов",
		"Docker использует возможности ядра Linux",
		"Образы Docker - неизменяемые слои",
		"Dockerfile - инструкция по сборке",
	}

	stats.DockerInsights = insights
	fmt.Println("\n   💡 Узнал о Docker:")
	for i, insight := range insights {
		fmt.Printf("      %d. %s\n", i+1, insight)
		time.Sleep(150 * time.Millisecond)
	}

	// Мотивация от чтения
	motivationFromReading := pagesRead * 2
	stats.Motivation += motivationFromReading
	fmt.Printf("   🎯 Мотивация от чтения: +%d\n", motivationFromReading)
}

func realizeDockerGoConnection(stats *DayStats, book *DockerBook) {
	fmt.Println("\n💡 ОСОЗНАНИЕ СВЯЗИ DOCKER И GO:")
	fmt.Println(strings.Repeat("─", 40))

	fmt.Println("   🔍 Листаю книгу дальше...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("   📖 Нахожу раздел «Техническая реализация Docker»")
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n   ⚡ ОТКРОВЕНИЕ: Docker написан на языке %s!\n", book.Language)
	time.Sleep(300 * time.Millisecond)

	realizations := []struct{
		text string
		impact int
	}{
		{"Тот самый Go, который я учу!", 50},
		{"Значит, изучая Go, я приближаюсь к пониманию Docker", 40},
		{"Go → Docker → Контейнеризация → Микросервисы → Успех", 60},
		{"Мой выбор языка был предопределён этой связью", 35},
	}

	fmt.Println("\n   🧠 Цепочка мыслей:")
	for _, r := range realizations {
		fmt.Printf("      • %s\n", r.text)
		stats.Motivation += r.impact
		time.Sleep(200 * time.Millisecond)
	}

	// Особый бонус за осознание
	epiphanyBonus := 100
	stats.Motivation += epiphanyBonus
	fmt.Printf("\n   🌟 БОНУС ОСОЗНАНИЯ: +%d к мотивации\n", epiphanyBonus)
}

func calculateMotivation(stats *DayStats, boots *WetBoots, book *DockerBook) {
	fmt.Println("\n📈 РАСЧЁТ МОТИВАЦИОННОГО БАЛАНСА:")
	fmt.Println(strings.Repeat("─", 40))

	factors := []struct{
		factor string
		value int
		positive bool
	}{
		{"Ранний подъём (06:30)", 20, true},
		{"Полезный завтрак", 15, true},
		{fmt.Sprintf("Выполнено заказов: %d", stats.Deliveries), stats.Deliveries * 10, true},
		{fmt.Sprintf("Промокшие ботинки (%d%%)", boots.WetnessLevel), boots.MotivationBoost, true},
		{"Дискомфорт от мокрых ног", -boots.Discomfort/2, false},
		{fmt.Sprintf("Прочитано страниц: %d", stats.PagesRead), stats.PagesRead * 2, true},
		{"Бумажная книга (не устают глаза)", 30, true},
		{"Осознание связи Docker & Go", 150, true},
		{"Go как язык Docker", 80, true},
	}

	fmt.Println("   Факторы влияния:")
	totalChange := 0
	for _, f := range factors {
		sign := "+"
		if !f.positive {
			sign = ""
		}
		fmt.Printf("      %s %s%d\n", f.factor, sign, f.value)
		totalChange += f.value
		time.Sleep(100 * time.Millisecond)
	}

	stats.Motivation = BASE_MOTIVATION + totalChange
	if stats.Motivation > MAX_MOTIVATION {
		stats.Motivation = MAX_MOTIVATION
	}

	fmt.Printf("\n   📊 Итоговый уровень мотивации: %d/%d\n", stats.Motivation, MAX_MOTIVATION)
}

func writeGoCode(stats *DayStats) {
	fmt.Println("\n💻 НАПИСАНИЕ КОДА НА GO:")
	fmt.Println(strings.Repeat("─", 40))

	// Мокрые ботинки как топливо для кода
	wetnessToCodeRatio := 3 // 1% влажности = 3 строки кода
	linesFromWetBoots := 85 * wetnessToCodeRatio

	// Мотивация к коду
	linesFromMotivation := stats.Motivation / 10

	// Docker инсайты дают идеи для кода
	linesFromInsights := len(stats.DockerInsights) * 15

	totalLines := linesFromWetBoots + linesFromMotivation + linesFromInsights
	stats.GoLinesWritten = totalLines

	fmt.Printf("   👢 Мокрые ботинки (85%% влажности): %d строк кода\n", linesFromWetBoots)
	fmt.Printf("   🎯 Мотивация (%d единиц): %d строк кода\n", stats.Motivation, linesFromMotivation)
	fmt.Printf("   📚 Docker инсайты (%d штук): %d строк кода\n", len(stats.DockerInsights), linesFromInsights)
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("   🏆 ВСЕГО НАПИСАНО СТРОК КОДА НА GO: %d\n", totalLines)

	// Пример кода
	fmt.Println("\n   📝 Пример написанного кода (тема: Docker-like контейнер):")
	fmt.Println(`   type Container struct {
       ID        string
       Image     string
       Status    string
       CreatedAt time.Time
   }

   func NewContainer(image string) *Container {
       return &Container{
           ID:        generateID(),
           Image:     image,
           Status:    "created",
           CreatedAt: time.Now(),
       }
   }`)
}

func printDaySummary(stats *DayStats, book *DockerBook, boots *WetBoots) {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("📊 ИТОГИ ДНЯ 14.01.2026:")
	fmt.Println(strings.Repeat("─", 60))

	summary := []struct{
		label string
		value string
		emoji string
	}{
		{"Дата", stats.Date, "📅"},
		{"Подъём", stats.WakeUpTime, "⏰"},
		{"Завтрак", stats.Breakfast, "🍲"},
		{"Доставки", fmt.Sprintf("%d заказа", stats.Deliveries), "📦"},
		{"Погода", stats.Weather, "❄️"},
		{"Состояние обуви", stats.BootsCondition, "👢"},
		{"Книга", book.Title, "📚"},
		{"Прочитано", fmt.Sprintf("%d/%d страниц", stats.PagesRead, book.PagesTotal), "📖"},
		{"Язык Docker", book.Language, "🐳"},
		{"Мотивация", fmt.Sprintf("%d/%d", stats.Motivation, MAX_MOTIVATION), "🎯"},
		{"Написано кода Go", fmt.Sprintf("%d строк", stats.GoLinesWritten), "💻"},
	}

	for _, item := range summary {
		fmt.Printf("   %s %-25s: %s\n", item.emoji, item.label, item.value)
	}

	fmt.Println(strings.Repeat("─", 60))

	// Уровень дня
	var dayLevel string
	var levelEmoji string

	switch {
	case stats.Motivation >= 800:
		dayLevel = "🚀 ЭПИФАНИЯ DOCKER & GO"
		levelEmoji = "🌟"
	case stats.Motivation >= 600:
		dayLevel = "🔥 МОЩНАЯ МОТИВАЦИЯ"
		levelEmoji = "🔥"
	case stats.Motivation >= 400:
		dayLevel = "💪 ПРОДУКТИВНЫЙ ДЕНЬ"
		levelEmoji = "💪"
	default:
		dayLevel = "📝 ОБЫЧНЫЙ ДЕНЬ"
		levelEmoji = "📝"
	}

	fmt.Printf("   %s УРОВЕНЬ ДНЯ: %s\n", levelEmoji, dayLevel)
}

func printPhilosophicalInsights() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("🧠 ФИЛОСОФСКИЕ ИНСАЙТЫ ДНЯ:")
	fmt.Println(strings.Repeat("─", 60))

	insights := []string{
		"1. Мокрые ботинки ≠ проблема, а ТОПЛИВО для мотивации",
		"2. Бумажная книга — цифровой детокс для программиста",
		"3. Docker написан на Go — это не совпадение, а ЗАКОНОМЕРНОСТЬ",
		"4. Изучая Go, ты изучаешь фундамент современных технологий",
		"5. Доставка заказов и доставка кода — оба создают ценность",
		"6. Снег за окном очищает город, код очищает архитектуру",
		"7. Промокшие ноги напоминают: дискомфорт ведёт к росту",
		"8. Go в Docker как ДНК в живом организме — невидим, но определяет всё",
	}

	for _, insight := range insights {
		fmt.Println("   " + insight)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("💫 ВЫВОД ДНЯ:")
	fmt.Println("   Мокрые ботинки привели к сухим фактам:")
	fmt.Println("   Docker → Go → Твой выбор → Твой путь → Твой успех")
	fmt.Println("\n   Каждая промокшая нить в ботинке — это новая строка кода на Go.")
	fmt.Println("   Каждая страница книги по Docker — это шаг к пониманию мира.")
	fmt.Println("\n🚀 GO IS DOCKER'S DNA. DOCKER IS GO'S LEGACY.")
	fmt.Println(strings.Repeat("═", 60))
}
