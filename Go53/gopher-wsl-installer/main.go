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
	fmt.Println("🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟")
	fmt.Println("      ДЕНЬ 53: OPENCODE В WSL ЧЕРЕЗ GO INSTALL     ")
	fmt.Println("🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟🥟")
	fmt.Println()
	fmt.Println("📚 Тема дня: Установка OpenCode на WSL Linux Ubuntu 24 через go install")
	fmt.Println()

	// Выводим легенду
	fmt.Println(legend)
	fmt.Println()

	// Подсчёт дней и XP
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)

	// Бонус за отказ от похода за посылкой и решение кодить
	xp := days*12 + 150
	level := xp/100 + 1

	fmt.Printf("🔥 Дней обучения Go (Go365): %d\n", days)
	fmt.Printf("⭐ Накоплено XP: %d (+150 за правильный выбор!)\n", xp)
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
	if days >= 53 {
		fmt.Println("   ✅ 53 дня — как 53 причина учить Go")
	}
	if level >= 7 {
		fmt.Println("   ✅ Уровень 7 — WSL-воин")
	}
	if level >= 10 {
		fmt.Println("   ✅ Уровень 10 — AI-мастер")
	}
	fmt.Println("   ✅ OpenCode-installer — теперь ты знаешь, как ставить AI-агента")
	fmt.Println()

	// Выводим инструкцию по установке
	fmt.Println(installGuide)
	fmt.Println()

	// Похвала и мотивация
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты не пошёл за посылкой, а остался настраивать окружение.\n")
	fmt.Printf("   Это правильный выбор: WSL Ubuntu 24 и OpenCode ждут тебя.\n")
	fmt.Println("   OpenCode написан на Go — твой язык обволакивает мир AI.")
	fmt.Println("   Установи его по инструкции и прокачивай кодинг с AI-помощником!")
	fmt.Println()

	// Дисклеймер
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Пельмени и чай — настоящие, бакенбарды — отращиваются.")
	fmt.Println("===================")
}
