package main

import (
	"fmt"
	"time"
)

// Friend представляет "воображаемого друга" в метро/ТЦ
type Friend struct {
	Name       string
	Location   string
	Imaginary  bool
	Dopamine   float64
}

// Book представляет книгу по программированию
type Book struct {
	Title     string
	Format    string // "fb2", "pdf", "epub"
	Readable  bool
	Topics    []string
}

func main() {
	fmt.Println("🚇 МОДУЛЬ: IMAGINARY-FRIENDS-VENDOR")
	fmt.Println("====================================")

	fmt.Println("📖 Легенда дня 18 января 2026:")
	fmt.Println("   Гоша в ТЦ 'Киргизия' воображает друзей на фуд-корте.")
	fmt.Println("   Читает книгу про стеки/очереди в метро.")
	fmt.Println("   Вендорит воображаемых друзей, как зависимости Go")

	// Создаем "воображаемых друзей" (как vendor-зависимости)
	friends := []Friend{
		{"Алексей (бариста)", "Фуд-корт ТЦ 'Киргизия'", true, 0.8},
		{"Мария (с ребенком)", "Столик рядом", true, 0.5},
		{"Старик с книгой", "Диван у эскалатора", true, 0.3},
		{"Курьер-коллега", "Вход в метро", false, 1.2},
	}

	fmt.Println("👥 ВЕНДОРИМ 'ДРУЗЕЙ' В ПАМЯТЬ:")
	dopamineTotal := 0.0

	for i, f := range friends {
		fmt.Printf("Друг #%d: %s\n", i+1, f.Name)
		fmt.Printf("   Место: %s\n", f.Location)
		fmt.Printf("   Воображаемый: %v\n", f.Imaginary)
		fmt.Printf("   Дофамин: +%.1f\n", f.Dopamine)

		if f.Imaginary {
			fmt.Println("   ⚠  ВЕНДОРИНГ ВООБРАЖАЕМОГО: риск зависимости от иллюзий")
		} else {
			fmt.Println("   ✅ РЕАЛЬНЫЙ КОНТАКТ: здоровый vendor социальных связей")
		}

		dopamineTotal += f.Dopamine
		time.Sleep(600 * time.Millisecond)
		fmt.Println()
	}

	// Книга по программированию (аналогия с vendor-зависимостями)
	book := Book{
		Title:    "Алгоритмы: стеки и очереди",
		Format:   "fb2",
		Readable: true,
		Topics:   []string{"Стек", "Очередь", "Бинарный поиск", "Пузырьковая сортировка"},
	}

	fmt.Println("📚 ВЕНДОРИНГ ЗНАНИЙ:")
	fmt.Printf("Книга: %s\n", book.Title)
	fmt.Printf("Формат: %s", book.Format)
	if book.Format == "fb2" {
		fmt.Print(" (удобно читать в телефоне!)\n")
	} else {
		fmt.Print(" (неудобно на маленьком экране)\n")
	}

	fmt.Println("Темы (универсальные для всех языков):")
	for _, topic := range book.Topics {
		fmt.Printf("   - %s\n", topic)
	}

	// Итоги дня
	fmt.Println("\n════════════════════════════════════════")
	fmt.Println("📊 ИТОГИ ДНЯ 18 ЯНВАРЯ 2026:")
	fmt.Printf("   Заработано: +600 руб.\n")
	fmt.Printf("   Воображаемых друзей: %d\n", countImaginaryFriends(friends))
	fmt.Printf("   Реальных контактов: %d\n", countRealFriends(friends))
	fmt.Printf("   Дофамин от социализации: %.1f/5.0\n", dopamineTotal)

	mood := analyzeMood(dopamineTotal, book.Readable)
	fmt.Printf("   Психическое состояние: %s\n", mood)
	fmt.Printf("   Формат знаний: %s\n", book.Format)
	fmt.Println("════════════════════════════════════════")

	// Геймификация
	fmt.Println("\n🎮 УРОВНИ 'ВЕНДОРИНГА СОЦИАЛЬНЫХ СВЯЗЕЙ':")
	fmt.Println("   5.0+  → 🏆 'Социальный вендор-гуру'")
	fmt.Println("   3.0+  → 👍 'Баланс реального и воображаемого'")
	fmt.Println("   1.0+  → 😐 'Преобладание воображаемых друзей'")
	fmt.Println("   ниже  → 🆘 'Одиночество в большом городе'")

	fmt.Printf("\n🎯 Ваш уровень: %.1f - ", dopamineTotal)
	fmt.Println(getLevel(dopamineTotal))

	// Мотивация
	fmt.Println("\n💪 МОТИВАЦИЯ НА ЗАВТРА:")
	fmt.Println("   'Как стеки и очереди универсальны во всех языках,")
	fmt.Println("    так и навыки общения универсальны во всех сферах жизни.'")
}

func countImaginaryFriends(friends []Friend) int {
	count := 0
	for _, f := range friends {
		if f.Imaginary {
			count++
		}
	}
	return count
}

func countRealFriends(friends []Friend) int {
	count := 0
	for _, f := range friends {
		if !f.Imaginary {
			count++
		}
	}
	return count
}

func analyzeMood(dopamine float64, hasBook bool) string {
	switch {
	case dopamine >= 3.0 && hasBook:
		return "ПРОДУКТИВНАЯ ИЗОЛЯЦИЯ 📚"
	case dopamine >= 2.0:
		return "УМЕРЕННАЯ СОЦИАЛИЗАЦИЯ 👍"
	case dopamine >= 0.5:
		return "ШИЗО-ТЕРАПИЯ ОДИНОЧЕСТВА 🎭"
	default:
		return "ЭКЗИСТЕНЦИАЛЬНЫЙ КРИЗИС 🆘"
	}
}

func getLevel(dopamine float64) string {
	switch {
	case dopamine >= 4.0:
		return "Социальный вендор-гуру"
	case dopamine >= 2.5:
		return "Балансирующий реалист"
	case dopamine >= 1.0:
		return "Воображаемый социализатор"
	default:
		return "Одинокий вендор"
	}
}
