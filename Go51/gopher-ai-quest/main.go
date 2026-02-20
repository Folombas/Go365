package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed legend.txt
var legend string

//go:embed ai_facts.txt
var aiFacts string

const startDate = "2026-01-18"

func main() {
	fmt.Println("🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖")
	fmt.Println("         ДЕНЬ 51: GO И НЕЙРО-СФЕРА         ")
	fmt.Println("🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖🤖")
	fmt.Println()
	fmt.Println("📚 Тема дня: Тренды AI и роль Go в нейро-сфере")
	fmt.Println()

	// Выводим легенду
	fmt.Println(legend)
	fmt.Println()

	// Подсчёт дней
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)

	fmt.Printf("🔥 Дней обучения Go (Go365): %d\n", days)
	fmt.Println()

	// XP и уровень (как в RPG)
	xp := days * 12 // +12 XP за каждый день (символично: 12 — число завершённости)
	level := xp/100 + 1
	fmt.Printf("⭐ Накоплено XP: %d\n", xp)
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
	if days >= 50 {
		fmt.Println("   ✅ Золотая полусотня (50 дней)")
	}
	if level >= 5 {
		fmt.Println("   ✅ Уровень 5 — начинающий гитхаб-воин")
	}
	if level >= 10 {
		fmt.Println("   ✅ Уровень 10 — AI-энтузиаст")
	}
	fmt.Println()

	// Факты о Go и AI
	fmt.Println(aiFacts)
	fmt.Println()

	// Похвала и мотивация
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты прошёл %d дней, съел пшёнку со сникерсом и узнал про OpenCode.\n", days)
	fmt.Println("   Go — это не просто язык, это ключ к миру AI. Продолжай в том же духе!")
	fmt.Println("   Помни: DeepSeek V4 уже близко, а теперь ещё и OpenCode написан на Go —")
	fmt.Println("   значит, твой стек знаний становится ещё ценнее.")
	fmt.Println()

	// Дисклеймер
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Пшённая каша с маслом — настоящая, Snickers — тоже.")
	fmt.Println("===================")
}
