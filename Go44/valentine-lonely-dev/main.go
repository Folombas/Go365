package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	journalFile = "valentine_journal.json"
	startDate   = "2026-01-18"
)

// JournalEntry хранит запись одного дня
type JournalEntry struct {
	Date       string   `json:"date"`
	Events     []string `json:"events"`
	XP         int      `json:"xp"`
	Mood       string   `json:"mood"`
	CapcutDays int      `json:"capcut_days"`
}

// Journal содержит все записи
type Journal struct {
	Entries []JournalEntry `json:"entries"`
}

func main() {
	fmt.Println("💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️")
	fmt.Println("         ДЕНЬ 44: СУББОТА, ДЕНЬ ВЛЮБЛЁННЫХ, НО Я С GO         ")
	fmt.Println("💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️💇‍♂️")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling")
	fmt.Println()

	// Загружаем или создаём журнал
	journal := loadJournal()

	// Вычисляем дни без CapCut
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)

	// Создаём запись сегодняшнего дня
	entry := JournalEntry{
		Date:       today.Format("2006-01-02"),
		Events:     getTodayEvents(),
		XP:         calculateXP(),
		Mood:       getMood(),
		CapcutDays: days,
	}

	// Добавляем запись в журнал
	journal.Entries = append(journal.Entries, entry)
	saveJournal(journal)

	// Выводим легенду
	printLegend(entry)

	// Выводим статистику
	printStats(journal)

	// Мотивация из отдельного файла
	PrintMotivation()

	// Похвала
	printPraise(days)

	// Дисклеймер
	printDisclaimer()
}

func loadJournal() *Journal {
	file, err := os.Open(journalFile)
	if os.IsNotExist(err) {
		return &Journal{Entries: []JournalEntry{}}
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения журнала:", err)
		return &Journal{}
	}
	defer file.Close()

	var journal Journal
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&journal)
	if err != nil {
		fmt.Println("⚠️ Ошибка парсинга JSON. Создаём новый журнал.")
		return &Journal{}
	}
	return &journal
}

func saveJournal(journal *Journal) {
	file, err := os.Create(journalFile)
	if err != nil {
		fmt.Println("⚠️ Не могу сохранить журнал:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(journal)
}

func getTodayEvents() []string {
	return []string{
		"🌅 Проснулся в 10:00 в гордом одиночестве",
		"💇 Премиальный барбершоп — стильная стрижка за 700₽ (подписчики в восторге!)",
		"📸 Выложил фото стрижки в блог — лайки и похвала",
		"🛒 Закупил продуктов: 2 десятка яиц, кефир, капуста, свёкла, картошка, дрожжи, мука, майонез, сметана, творог",
		"👞 Нубуковые ботинки промокли в снежной жиже",
		"🥧 Мама напекла беляшей и пирожков",
		"📅 Планы на понедельник: встать в 6-7 утра и курьерить до 5-6 вечера",
	}
}

func calculateXP() int {
	// События дня дают XP
	return 10 + // стрижка
		5 + // фото
		5 + // покупки
		10 + // беляши
		5 // планы
}

func getMood() string {
	return "😐 Одиноко, но стильно. Впереди работа и Go."
}

func printLegend(entry JournalEntry) {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              💔 ЛЕГЕНДА ДНЯ СВЯТОГО ВАЛЕНТИНА 💔         ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  %s\n", entry.Date)
	fmt.Printf("║  Настроение: %s\n", entry.Mood)
	fmt.Println("║  События дня:")
	for _, e := range entry.Events {
		fmt.Printf("║    • %s\n", e)
	}
	fmt.Printf("║  XP сегодня: +%d\n", entry.XP)
	fmt.Printf("║  Всего дней без CapCut: %d\n", entry.CapcutDays)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printStats(journal *Journal) {
	totalXP := 0
	for _, e := range journal.Entries {
		totalXP += e.XP
	}
	fmt.Println("📊 СТАТИСТИКА ПРОКАЧКИ")
	fmt.Println("───────────────────────")
	fmt.Printf("📅 Записей в дневнике: %d\n", len(journal.Entries))
	fmt.Printf("⭐ Общий опыт (XP): %d\n", totalXP)
	fmt.Printf("🏆 Уровень: %d\n", totalXP/100+1)
	fmt.Println()
}

func printPraise(days int) {
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	if days >= 44 {
		fmt.Printf("   ГОША! Ты прошёл %d дней без CapCut! Даже в одиночестве ты крут!\n", days)
	} else {
		fmt.Printf("   Сегодня %d-й день. Ты держишься, молодец!\n", days)
	}
	fmt.Println("   Ты не установил CapCut, а значит, инвестировал в себя.")
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Стрижка за 700 рублей — реальна, беляши — съедобны.")
	fmt.Println("===================")
}
