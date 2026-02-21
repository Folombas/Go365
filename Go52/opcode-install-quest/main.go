package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed legend.txt
var legend string

//go:embed install_guide.txt
var installGuide string

const startDate = "2026-01-18"

func main() {
	fmt.Println("🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖")
	fmt.Println("      ДЕНЬ 52: OPENCODE — GO-ИНСТРУМЕНТ ДЛЯ AI     ")
	fmt.Println("🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖")
	fmt.Println()
	fmt.Println("📚 Тема дня: Как скачать и установить OpenCode: Все способы")
	fmt.Println()

	// Выводим легенду
	fmt.Println(legend)
	fmt.Println()

	// Подсчёт дней и XP
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)

	// Бонусные XP за два дня метельного анализа
	xp := days*12 + 100 // +100 за открытие OpenCode
	level := xp/100 + 1

	fmt.Printf("🔥 Дней обучения Go (Go365): %d\n", days)
	fmt.Printf("⭐ Накоплено XP: %d (+100 бонус за OpenCode!)\n", xp)
	fmt.Printf("📈 Уровень персонажа: %d\n", level)
	fmt.Println()

	// Достижения
	fmt.Println("🏆 РАЗБЛОКИРОВАННЫЕ ДОСТИЖЕНИЯ:")
	if days >= 7 {
		fmt.Println("   ✅ Неделя с Go (7 дней)")
	}
	if days >= 30 {
		fmt.Println("   ✅ Месяц кода (30 дней)")
	}
	if days >= 52 {
		fmt.Println("   ✅ 52 дня — как карта в колоде, полная возможностей")
	}
	if level >= 6 {
		fmt.Println("   ✅ Уровень 6 — Go-путешественник")
	}
	if level >= 10 {
		fmt.Println("   ✅ Уровень 10 — AI-исследователь")
	}
	fmt.Println("   ✅ Открытие OpenCode — AI-агент на Go")
	fmt.Println()

	// Выводим гайд по установке
	fmt.Println(installGuide)
	fmt.Println()

	// Похвала и мотивация
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты не просто учишь Go, ты следишь за трендами.\n")
	fmt.Printf("   OpenCode написан на Go — это подтверждение: ты выбрал правильный язык.\n")
	fmt.Printf("   Два дня метели не прошли даром: теперь у тебя есть карта сокровищ —\n")
	fmt.Println("   все способы установки OpenCode на одной странице.")
	fmt.Println("   Выбирай любой, ставь и прокачивай кодинг с AI!")
	fmt.Println()

	// Дисклеймер
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("OpenCode реален и написан на Go.")
	fmt.Println("===================")
}
