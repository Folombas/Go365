package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	dataFile    = "resistance.json"
	startCapCut = "2026-01-18" // дата удаления CapCut
)

type Stats struct {
	LastBarRefusal string `json:"last_bar_refusal"`
	BarFreeDays    int    `json:"bar_free_days"`
	CapCutDays     int    `json:"capcut_days"`
	TotalXP        int    `json:"total_xp"`
}

func main() {
	// Определяем флаги
	daysOnly := flag.Bool("days", false, "показать только количество дней без CapCut")
	barDaysOnly := flag.Bool("bar-days", false, "показать только количество дней без бара")
	name := flag.String("name", "Гоша", "ваше имя")
	flag.Parse()

	// Загружаем или создаём статистику
	stats := loadStats()

	// Вычисляем дни без CapCut
	start, _ := time.Parse("2006-01-02", startCapCut)
	today := time.Now()
	capDays := int(today.Sub(start).Hours() / 24)

	// Обновляем статистику по CapCut
	stats.CapCutDays = capDays

	// Обрабатываем дату последнего отказа от бара
	var barDays int
	if stats.LastBarRefusal == "" {
		// Если в статистике нет, устанавливаем сегодня как первый день отказа
		stats.LastBarRefusal = today.Format("2006-01-02")
		barDays = 1
	} else {
		barStart, _ := time.Parse("2006-01-02", stats.LastBarRefusal)
		barDays = int(today.Sub(barStart).Hours()/24) + 1 // включая сегодня
	}
	stats.BarFreeDays = barDays

	// Начисляем XP (за каждый день без бара +5, за каждый день без CapCut +10)
	stats.TotalXP = stats.CapCutDays*10 + stats.BarFreeDays*5

	// Сохраняем статистику
	saveStats(stats)

	// Вывод в зависимости от флагов
	if *daysOnly {
		fmt.Println(stats.CapCutDays)
		return
	}
	if *barDaysOnly {
		fmt.Println(stats.BarFreeDays)
		return
	}

	// Полный вывод
	printHeader()
	printLegend(*name, stats)
	printMotivation()
	printPraise(stats)
	printDisclaimer()
}

func loadStats() *Stats {
	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		return &Stats{
			LastBarRefusal: "",
			BarFreeDays:    0,
			CapCutDays:     0,
			TotalXP:        0,
		}
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения файла:", err)
		return &Stats{}
	}
	defer file.Close()

	var stats Stats
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&stats)
	if err != nil {
		fmt.Println("⚠️ Ошибка парсинга JSON. Создаём новую статистику.")
		return &Stats{}
	}
	return &stats
}

func saveStats(stats *Stats) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("⚠️ Не могу сохранить файл:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(stats)
}

func printHeader() {
	fmt.Println("🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺")
	fmt.Println("         ДЕНЬ 45: БАР-ВОЗДЕРЖАНИЕ И GO         ")
	fmt.Println("🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺🍺")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling: flag and time")
	fmt.Println()
}

func printLegend(name string, stats *Stats) {
	fmt.Printf("╔══════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║              🧠 ЛЕГЕНДА ДНЯ: ИНТРОВЕРТ СИЛЫ 🧠          ║\n")
	fmt.Printf("╠══════════════════════════════════════════════════════════╣\n")
	fmt.Printf("║  15 февраля 2026. %s встал в 9:30, умылся, попил воды   ║\n", name)
	fmt.Printf("║  и сразу сел за компьютер программировать на Go.        ║\n")
	fmt.Printf("║  Вечером реальные друзья (Дима-велоремонтник и рэппер)  ║\n")
	fmt.Printf("║  звали в бар, но %s отказался — интроверт, аутист,      ║\n", name)
	fmt.Printf("║  нацеленный на первую работу Go-разработчиком.          ║\n")
	fmt.Printf("║                                                          ║\n")
	fmt.Printf("║  CapCut удалён %s. Сегодня %d-й день без монтажа.       ║\n", startCapCut, stats.CapCutDays)
	fmt.Printf("║  Бар-воздержание: %d дней подряд.                       ║\n", stats.BarFreeDays)
	fmt.Printf("║  Текущий XP: %d                                         ║\n", stats.TotalXP)
	fmt.Printf("╚══════════════════════════════════════════════════════════╝\n")
	fmt.Println()
}

func printMotivation() {
	fmt.Println("💬 10 МОТИВАТОРОВ ДЛЯ ИНТРОВЕРТА:")
	motivation := []string{
		"1. Бар — это временно, код — навсегда.",
		"2. Друзья зовут в бар, а Go зовёт коммитить.",
		"3. Одна строка кода сегодня = 100 строк резюме завтра.",
		"4. Да, ты - сфокусированный айти-гений.",
		"5. Пиво тормозит нейроны, а Go их прокачивает.",
		"6. Сериалы подождут, горутины — нет.",
		"7. Пока они тусуются в барах, ты создаёшь микросервисы.",
		"8. CapCut — прошлый век, Go — будущее.",
		"9. Каждый день без бара приближает к офферу.",
		"10. Интроверты обучаются, пока экстраверты ходят по барам.",
	}
	for _, m := range motivation {
		fmt.Println(m)
	}
	fmt.Println()
}

func printPraise(stats *Stats) {
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	if stats.BarFreeDays >= 30 {
		fmt.Printf("   ГОША! Ты не был в баре %d дней! Это олимпийский рекорд!\n", stats.BarFreeDays)
	} else if stats.BarFreeDays >= 7 {
		fmt.Printf("   Уже %d дней без бара! Сила воли растёт!\n", stats.BarFreeDays)
	} else {
		fmt.Printf("   Сегодня %d-й день без бара. Ты держишься, молодец!\n", stats.BarFreeDays)
	}
	fmt.Printf("   Дней без CapCut: %d — ты круче, чем вчера!\n", stats.CapCutDays)
	fmt.Println("   Продолжай в том же духе — работа найдёт тебя сама.")
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Даже если друзья реальны, их имена изменены.")
	fmt.Println("===================")
}
