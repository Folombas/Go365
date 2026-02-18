package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed legend.txt
var legend string

//go:embed deepseek_motivation.txt
var motivation string

const startDate = "2026-01-18"

func main() {
	fmt.Println("☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️")
	fmt.Println("         ДЕНЬ 48: УТРЕННИЙ КОД С EMBED        ")
	fmt.Println("☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️☀️")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling: go:embed")
	fmt.Println()

	// Выводим встроенную легенду
	fmt.Println(legend)
	fmt.Println()

	// Подсчёт дней без CapCut
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)
	fmt.Printf("🔥 Дней без CapCut: %d\n", days)

	// XP: за каждый день +10, за ожидание V4 + бонус (если дни > 30)
	xp := days * 10
	if days > 30 {
		xp += 50 // бонус за стойкость
	}
	fmt.Printf("⭐ Накоплено XP: %d\n", xp)
	fmt.Printf("📈 Уровень персонажа: %d\n", xp/100+1)
	fmt.Println()

	// Выводим мотивационные фразы про DeepSeek V4
	fmt.Println("🤖 DEEPSEEK V4 НА ПОДХОДЕ!")
	lines := strings.Split(motivation, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()

	// Похвала
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты встал в 6:10 и сразу сел за код — это уровень настоящего профи!\n")
	fmt.Printf("   %d дней без CapCut — твоя дисциплина зашкаливает. DeepSeek V4 скоро войдёт в твою жизнь и ускорит прокачку в 10 раз. Держись!\n", days)
	fmt.Println()

	// Дисклеймер (выдумка)
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все события вымышлены, но кашка с блинами — настоящая.")
	fmt.Println("===================")
}
