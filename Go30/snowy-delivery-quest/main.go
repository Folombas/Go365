package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type GoshaRPG struct {
	Level          int     `json:"level"`
	XP             int     `json:"xp"`
	Money          int     `json:"money"`
	Willpower      int     `json:"willpower"`
	DeliveryCount  int     `json:"delivery_count"`
	StudySessions  int     `json:"study_sessions"`
	CapCutTempted  bool    `json:"capcut_tempted"`
	CurrentDay     int     `json:"current_day"`
}

var motivationQuotes = []string{
	"Go — это не язык, это оружие будущего!",
	"Каждый commit — шаг от курьера к разработчику!",
	"Метель снаружи, код внутри — выбирай теплое будущее!",
	"CapCut подождёт. Карьера — нет!",
	"Golang dev зарабатывает больше, чем 10 доставок в метель!",
	"Сугробы сегодня, собеседования завтра!",
	"Прокачай GoshaRPG до уровня Senior!",
	"Кодинг > котлетки > курьерка!",
	"Установи приоритет: Go > всё остальное!",
	"365 дней Go = новая жизнь!",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("🎮 SNOWY DELIVERY QUEST v1.0")
	fmt.Println("📅 Go365 День 30 | 30 января 2026")
	fmt.Println("❄️ Метель снаружи, битва внутри...")
	fmt.Println()

	g := &GoshaRPG{
		Level:         1,
		XP:            0,
		Money:         0,
		Willpower:     100,
		DeliveryCount: 0,
		StudySessions: 0,
		CapCutTempted: false,
		CurrentDay:    30,
	}

	fmt.Printf("👤 Персонаж: Курьер-студент Go (Уровень %d)\n", g.Level)
	fmt.Printf("💪 Сила воли: %d/100\n\n", g.Willpower)

	// Симуляция дня
	simulateDay(g)

	// Финальная статистика
	showFinalStats(g)
}

func simulateDay(g *GoshaRPG) {
	events := []string{
		"📦 Доставка: Третьяковская → Аэропорт (Косметика, 450₽)",
		"📦 Доставка: Авиапарк → Речной Вокзал (Одежда, 380₽)",
		"📦 Доставка: Флотская (Лекарства, 320₽)",
		"💻 Учеба Go: Горутины (60 мин, +85 XP)",
		"💻 Учеба Go: Select statement (+72 XP)",
		"🎬 ИСКУШЕНИЕ: CapCut зовёт! (Travel-Видео 2018)",
		"🍝 Мамины котлетки (+20 воли)",
		"❄️ Метель усилилась! (-15 воли)",
	}

	fmt.Println("=== СИМУЛЯЦИЯ ДНЯ ===")
	for i, event := range events {
		fmt.Printf("\n%2d: %s\n", i+1, event)
		processEvent(g, event)
		time.Sleep(1 * time.Second)
	}
}

func processEvent(g *GoshaRPG, event string) {
	switch {
	case strings.Contains(event, "Доставка"):
		g.Money += rand.Intn(400) + 300
		g.DeliveryCount++
		fmt.Printf("💰 +%d₽ | Всего: %d₽\n", rand.Intn(400)+300, g.Money)

	case strings.Contains(event, "Учеба Go"):
		xp := rand.Intn(75) + 50
		g.XP += xp
		g.StudySessions++
		if g.XP >= 500 && g.Level < 5 {
			g.Level++
			fmt.Printf("🎉 LEVEL UP! Уровень %d! +%d XP\n", g.Level, xp)
		} else {
			fmt.Printf("📚 +%d XP | Всего: %d XP\n", xp, g.XP)
		}

	case strings.Contains(event, "ИСКУШЕНИЕ"):
		fmt.Print("🎬 CapCut: 'Смонтируй тропический отпуск 2018...' ")
		if g.Willpower > 30 {
			fmt.Println("❌ ОТВЕРГНУТО! 'Сначала карьера!' (+15 воли)")
			g.Willpower += 15
		} else {
			fmt.Println("💔 СДАЛСЯ! -120 мин на монтаж (-40 воли)")
			g.CapCutTempted = true
			g.Willpower -= 40
		}

	case strings.Contains(event, "котлетки"):
		g.Willpower += 20
		fmt.Println("🍝 +20 воли | Мамины котлетки ❤️")

	case strings.Contains(event, "Метель"):
		g.Willpower -= 15
		fmt.Println("❄️ -15 воли | Сугробы достали!")

	default:
		fmt.Println("⏳ Обычный момент...")
	}

	// Ограничение воли
	if g.Willpower > 100 {
		g.Willpower = 100
	}
	if g.Willpower < 0 {
		g.Willpower = 0
	}
}

func showFinalStats(g *GoshaRPG) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🏆 ФИНАЛЬНАЯ СТАТИСТИКА ДНЯ")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf("💰 Деньги: %d₽ (%d доставок)\n", g.Money, g.DeliveryCount)
	fmt.Printf("🧠 XP Go: %d (Уровень %d)\n", g.XP, g.Level)
	fmt.Printf("💪 Воля: %d/100\n", g.Willpower)
	fmt.Printf("📚 Сессий: %d\n", g.StudySessions)

	if !g.CapCutTempted {
		fmt.Println("🎉 ДОСТИЖЕНИЕ: 'Железная воля' ✓")
	} else {
		fmt.Println("⚠️ Срыв: CapCut победил сегодня")
	}

	fmt.Println("\n🎯 АЛГОРИТМ ЖИЗНИ:")
	fmt.Println("1. Курьер (деньги) ✓")
	fmt.Println("2. Go учеба (будущее) ✓")
	fmt.Println("3. Программист (цель)")
	fmt.Println("4. Видео (выходные)")

	printMotivation()
}

func printMotivation() {
	fmt.Println("\n🔥 10 МОТИВАЦИОННЫХ ФРАЗ:")
	for i, quote := range motivationQuotes {
		fmt.Printf("%2d. %s\n", i+1, quote)
	}
	fmt.Println("\n❄️ Метель снаружи → код внутри!")
	fmt.Println("📅 Go365 День 30/365 → День 31 ждет!")
}
