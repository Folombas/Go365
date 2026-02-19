package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed legend.txt
var legend string

//go:embed symbiosis_motivation.txt
var motivation string

const (
	startDate     = "2026-01-18" // начало Go365
	capcutInstall = "2026-02-18" // Гоша установил CapCut вчера
)

func main() {
	fmt.Println("🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬")
	fmt.Println("         ДЕНЬ 50: СИМБИОЗ КОДА И МОНТАЖА        ")
	fmt.Println("🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬🎬")
	fmt.Println()
	fmt.Println("📚 Тема дня: Симбиоз программирования и видеомонтажа")
	fmt.Println()

	// Выводим легенду
	fmt.Println(legend)
	fmt.Println()

	// Подсчёт дней
	start, _ := time.Parse("2006-01-02", startDate)
	install, _ := time.Parse("2006-01-02", capcutInstall)
	today := time.Now()
	goDays := int(today.Sub(start).Hours() / 24)
	editDays := int(today.Sub(install).Hours()/24) + 1 // включая сегодня

	fmt.Printf("🔥 Дней программирования (Go365): %d\n", goDays)
	fmt.Printf("✂️  Дней с CapCut (монтаж): %d\n", editDays)
	fmt.Println()

	// XP: за каждый день программирования +10, за каждый день монтажа +5 (как отдых)
	xp := goDays*10 + editDays*5
	level := xp/100 + 1
	fmt.Printf("⭐ Накоплено XP: %d\n", xp)
	fmt.Printf("📈 Уровень персонажа: %d\n", level)
	fmt.Println()

	// Выводим мотивационные фразы
	fmt.Println(motivation)
	fmt.Println()

	// Похвала
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты прошёл %d дней программирования и %d дней монтажа.\n", goDays, editDays)
	fmt.Println("   Ты не просто учишь Go — ты осваиваешь творчество, которое делает софт лучше.")
	fmt.Println("   Помни: программисты создают инструменты для творцов, и ты можешь стать одним из них!")
	fmt.Println()

	// Дисклеймер
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("CapCut установлен — это не измена, а баланс.")
	fmt.Println("===================")
}
