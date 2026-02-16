package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dataFile     = "delivery_stats.txt"
	capcutStart  = "2026-01-18"
	separator    = "|"
)

type Stats struct {
	LastRunDate   string
	TotalDays     int
	TotalEarnings int
	TotalXP       int
}

func main() {
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println("         ДЕНЬ 45: ВЫЖИВАНИЕ В МЕТЕЛИ         ")
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling: os and bufio")
	fmt.Println()

	stats := loadStats()
	today := time.Now().Format("2006-01-02")

	// Подсчёт дней без CapCut
	start, _ := time.Parse("2006-01-02", capcutStart)
	current, _ := time.Parse("2006-01-02", today)
	daysWithoutCapCut := int(current.Sub(start).Hours() / 24)

	// Если запуск сегодняшний, не обновляем лишний раз
	if stats.LastRunDate != today {
		stats.TotalDays = daysWithoutCapCut
		stats.LastRunDate = today

		// Добавляем XP за день учёбы
		stats.TotalXP += 10

		// Спрашиваем про доход
		fmt.Print("💰 Сколько заработал сегодня (руб)? ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		earn, _ := strconv.Atoi(input)
		stats.TotalEarnings += earn
		stats.TotalXP += earn / 10 // 1 XP за каждые 10 рублей

		saveStats(stats)
		fmt.Println("✅ Данные сохранены!")
	} else {
		fmt.Println("⚠️ Сегодня уже запускал программу. Показываю статистику.")
	}

	// Вывод легенды
	printLegend(stats)
	printMotivation()
	printPraise(stats)
	printDisclaimer()
}

func loadStats() Stats {
	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		// Файла нет — возвращаем пустую статистику
		return Stats{}
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения файла:", err)
		return Stats{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stats Stats
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, separator)
		if len(parts) != 4 {
			continue
		}
		stats.LastRunDate = parts[0]
		stats.TotalDays, _ = strconv.Atoi(parts[1])
		stats.TotalEarnings, _ = strconv.Atoi(parts[2])
		stats.TotalXP, _ = strconv.Atoi(parts[3])
	}
	return stats
}

func saveStats(stats Stats) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("⚠️ Ошибка сохранения файла:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	line := fmt.Sprintf("%s%s%d%s%d%s%d\n",
		stats.LastRunDate, separator,
		stats.TotalDays, separator,
		stats.TotalEarnings, separator,
		stats.TotalXP)
	writer.WriteString(line)
	writer.Flush()
}

func printLegend(stats Stats) {
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║           🧠 ЛЕГЕНДА ДНЯ: МЕТЕЛЬ И ГРЕЧКА 🧠            ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  16 февраля 2026. Понедельник. Гоша встал в 11 утра.     ║\n")
	fmt.Printf("║  Побрился, умылся, съел гречку с солёным огурцом.        ║\n")
	fmt.Printf("║  Развёз пару заказов, чтобы не киснуть в четырёх стенах. ║\n")
	fmt.Printf("║  На улице метель, но он покатался на автобусах, посмотрел║\n")
	fmt.Printf("║  на красивых женщин и немного заработал.                 ║\n")
	fmt.Printf("║                                                          ║\n")
	fmt.Printf("║  Заработано сегодня: %d руб.                             ║\n", stats.TotalEarnings)
	fmt.Printf("║  Всего дней без CapCut: %d                                ║\n", stats.TotalDays)
	fmt.Printf("║  Накоплено XP: %d                                         ║\n", stats.TotalXP)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printMotivation() {
	fmt.Println("💬 10 МОТИВАТОРОВ В МЕТЕЛЬ:")
	motivation := []string{
		"1. Деньги на хлеб есть — значит, можно кодить дальше.",
		"2. Метель закончится, а Go-навыки останутся.",
		"3. Красивые женщины в автобусах и метро — это вдохновение, а не повод отвлекаться.",
		"4. Гречка с огурцом — топливо для мозга, который учит Go.",
		"5. Оплата «на батон» — временно, оффер — навсегда.",
		"6. CapCut удалён — мозг свободен для горутин.",
		"7. Депрессия лечится не сериалами, а коммитами.",
		"8. Каждая доставка — это XP для будущей карьеры.",
		"9. В метель кодится особенно уютно.",
		"10. Ты не просто доставляешь коробки — ты доставляешь себя к цели.",
	}
	for _, m := range motivation {
		fmt.Println(m)
	}
	fmt.Println()
}

func printPraise(stats Stats) {
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	if stats.TotalDays >= 45 {
		fmt.Printf("   ГОША! Ты держишься %d дней без CapCut! Это олимпийский рекорд!\n", stats.TotalDays)
	} else {
		fmt.Printf("   Сегодня %d-й день без CapCut. Ты красавчик!\n", stats.TotalDays)
	}
	fmt.Printf("   XP за всё время: %d. Так держать!\n", stats.TotalXP)
	fmt.Println("   Помни: каждая строчка кода приближает к офферу.")
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Гречка с огурцом — реальна, остальное — художественный вымысел.")
	fmt.Println("===================")
}

