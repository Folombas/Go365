package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	dataFile  = "wet_boots_resistance.json"
	startDate = "2026-01-18"
)

// ResistanceData хранит данные о сопротивлении CapCut
type ResistanceData struct {
	StartDate     string   `json:"start_date"`
	LastCheck     string   `json:"last_check"`
	DaysCount     int      `json:"days_count"`
	Achievements  []string `json:"achievements"`
	XP            int      `json:"xp"`
	Slogan        string   `json:"slogan"`
	WetBootsLevel int      `json:"wet_boots_level"` // уровень промокания (метафора сложности)
}

var achievementsMap = map[int]string{
	1:   "🌅 Первый рассвет (1 день)",
	7:   "🚇 Неделя в метро (7 дней)",
	13:  "👟 Промокшие ботинки (13 дней)",
	30:  "💳 Новый проездной (30 дней)",
	43:  "📱 Honor 10x Lite (43 дня)",
	60:  "📸 Флагманский опыт (60 дней)",
	100: "🏆 Легенда пятницы 13 (100 дней)",
}

func main() {
	fmt.Println("🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️")
	fmt.Println("         ДЕНЬ 43: ПЯТНИЦА 13, МОКРЫЕ НУБУКОВЫЕ БОТИНКИ         ")
	fmt.Println("🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️🌧️")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling: I/O & File Handling")
	fmt.Println()

	data := loadOrCreateData()
	today := time.Now().Format("2006-01-02")

	start, _ := time.Parse("2006-01-02", data.StartDate)
	current, _ := time.Parse("2006-01-02", today)
	days := int(current.Sub(start).Hours() / 24)

	data.DaysCount = days
	data.LastCheck = today
	data.WetBootsLevel = days / 7 // чем больше дней, тем выше уровень промокания

	// Проверка новых достижений
	newAchievements := checkAchievements(data)
	data.Achievements = append(data.Achievements, newAchievements...)
	data.XP += len(newAchievements) * 43

	// Генерация слогана дня
	data.Slogan = generateSlogan(days)

	saveData(data)

	// Вывод
	printFriday13Legend(data)
	printStats(data)
	printMotivation()
	printPraise(data)
	printDisclaimer()

	if len(newAchievements) > 0 {
		fmt.Println("\n✨ НОВЫЕ ДОСТИЖЕНИЯ РАЗБЛОКИРОВАНЫ ✨")
		for _, a := range newAchievements {
			fmt.Printf("   🏅 %s (+43 XP)\n", a)
		}
	}
}

func loadOrCreateData() *ResistanceData {
	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		// Первый запуск
		data := &ResistanceData{
			StartDate:     startDate,
			LastCheck:     time.Now().Format("2006-01-02"),
			DaysCount:     0,
			Achievements:  []string{},
			XP:            0,
			Slogan:        "Ботинки промокли, но код сух.",
			WetBootsLevel: 0,
		}
		saveData(data)
		return data
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения файла:", err)
		return &ResistanceData{StartDate: startDate, LastCheck: time.Now().Format("2006-01-02")}
	}
	defer file.Close()

	var data ResistanceData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("⚠️ Ошибка парсинга JSON. Создаём новый файл.")
		return &ResistanceData{StartDate: startDate, LastCheck: time.Now().Format("2006-01-02")}
	}
	return &data
}

func saveData(data *ResistanceData) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("⚠️ Не могу сохранить файл:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(data)
}

func checkAchievements(data *ResistanceData) []string {
	var unlocked []string
	for days, name := range achievementsMap {
		if data.DaysCount >= days {
			found := false
			for _, a := range data.Achievements {
				if a == name {
					found = true
					break
				}
			}
			if !found {
				unlocked = append(unlocked, name)
			}
		}
	}
	return unlocked
}

func generateSlogan(days int) string {
	slogans := []string{
		"Проездной есть, а зонта нет.",
		"Honor 10x Lite снимает код лучше флагманов.",
		"Мокрые ботинки — мокрая архитектура.",
		"Пятница 13 — день, когда JSON не врёт.",
		"Снег растаял, но горутины не тают.",
		"Подписчики подарили ботинки, а Go дарит карьеру.",
		"Метро везёт, а каналы синхронизируют.",
		"Флагманы для селфи, Go для бэкенда.",
		"30 дней безлимита — 30 дней без CapCut.",
		"43 дня — новый уровень промокания.",
	}
	return slogans[days%len(slogans)]
}

func printFriday13Legend(data *ResistanceData) {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              🧟 ЛЕГЕНДА ПЯТНИЦЫ 13 🧟                   ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  13 февраля 2026. Полседьмого утра. Будильник Honor.     ║\n")
	fmt.Printf("║  Сверстники щеголяют в realme GT8 Pro и Xiaomi 15 Ultra, ║\n")
	fmt.Printf("║  а Гоша экономит на проездной — 4 тысячи с копейками.    ║\n")
	fmt.Printf("║  Метро, автобусы, резкое потепление. Нубуковые ботинки   ║\n")
	fmt.Printf("║  промокают в снежной жиже. Другой обуви нет.             ║\n")
	fmt.Printf("║  Но Гоша продолжает обучение Go. CapCut не установлен.   ║\n")
	fmt.Printf("║                                                          ║\n")
	fmt.Printf("║  Слоган дня: «%s»  ║\n", data.Slogan)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printStats(data *ResistanceData) {
	fmt.Println("📊 СТАТИСТИКА ОБУЧЕНИЯ")
	fmt.Println("───────────────────────────")
	fmt.Printf("🗓️  Старт: %s\n", data.StartDate)
	fmt.Printf("📅 Сегодня: %s\n", data.LastCheck)
	fmt.Printf("🔥 Дней без CapCut: %d\n", data.DaysCount)
	fmt.Printf("💧 Уровень промокания: %d\n", data.WetBootsLevel)
	fmt.Printf("⭐ Опыт (XP): %d\n", data.XP)
	fmt.Printf("🏅 Достижений: %d\n", len(data.Achievements))
	if len(data.Achievements) > 0 {
		fmt.Println("🏆 РАЗБЛОКИРОВАНО:")
		for _, a := range data.Achievements {
			fmt.Printf("   • %s\n", a)
		}
	}
	fmt.Println()
}

func printMotivation() {
	fmt.Println("💬 10 МОТИВАТОРОВ В МОКРЫХ БОТИНКАХ:")
	motivation := []string{
		"1. Пока сверстники фоткают еду, ты кодишь на Go.",
		"2. Флагманский смартфон не добавит ума, а Go — добавит.",
		"3. Проездной на 30 дней — инвестиция в карьеру, а не в такси.",
		"4. Мокрые ботинки — временно. Go-навыки — навсегда.",
		"5. Пятница 13 — день, когда баги боятся тебя.",
		"6. Honor 10x Lite — это не стыдно, это экономия на курсы.",
		"7. Нет денег на Xiaomi 15 Ultra? Есть деньги на образование.",
		"8. Блоги подождут, бэкенд на Go не ждёт.",
		"9. Снег растаял — начался сезон кодинга.",
		"10. Ты уже прошёл 42 дня, 43-й будет твоим.",
	}
	for _, m := range motivation {
		fmt.Println(m)
	}
	fmt.Println()
}

func printPraise(data *ResistanceData) {
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	if data.DaysCount >= 43 {
		fmt.Printf("   ГОША! Ты прошёл %d дней! Промок, замёрз, но не сдался. CapCut повержен!\n", data.DaysCount)
	} else if data.DaysCount >= 30 {
		fmt.Printf("   %d дней! Новый проездной — новый уровень. Так держать!\n", data.DaysCount)
	} else {
		fmt.Printf("   Сегодня %d-й день. Каждый день приближает к мечте. Молодец!\n", data.DaysCount)
	}
	fmt.Println("   Ты не установил CapCut, а значит, инвестировал в себя.")
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("Honor 10x Lite — отличный телефон, даже без флагманской камеры.")
	fmt.Println("Ботинки можно просушить, а знания — никогда.")
	fmt.Println("===================")
}
