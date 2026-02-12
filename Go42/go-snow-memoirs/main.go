package main

import (
	"encoding/json"
	"fmt"
	"os"

	//"path/filepath"
	"time"
)

const (
	saveFile   = "gosha_snow_memoirs.json"
	startDate  = "2026-01-18"
	currentDay = 42
)

// Memoir хранит данные о сопротивлении CapCut
type Memoir struct {
	StartDate    string   `json:"start_date"`
	LastCheck    string   `json:"last_check"`
	DaysCount    int      `json:"days_count"`
	Achievements []string `json:"achievements"`
	XP           int      `json:"xp"`
	Slogan       string   `json:"slogan"`
}

var achievementsMap = map[int]string{
	1:   "❄️ Снежинка стойкости (1 день)",
	7:   "☕ Неделя с чаем и вафлями",
	14:  "🧹 Дворник кода (2 недели)",
	30:  "📚 Месяц мозгового штурма",
	42:  "🌀 Ответ на всё (42 дня)",
	60:  "⚔️ Ветеран Go",
	100: "🏆 Легенда 100 дней",
}

func main() {
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println("            ДЕНЬ 42: СНЕЖНЫЙ МОЗГОВОЙ GOLANG-ШТУРМ          ")
	fmt.Println("❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
	fmt.Println()

	memoir := loadOrCreateMemoir()
	today := time.Now().Format("2006-01-02")

	start, _ := time.Parse("2006-01-02", memoir.StartDate)
	current, _ := time.Parse("2006-01-02", today)
	days := int(current.Sub(start).Hours() / 24)

	memoir.DaysCount = days
	memoir.LastCheck = today

	// Проверка достижений
	newAchievements := checkAchievements(memoir)
	memoir.Achievements = append(memoir.Achievements, newAchievements...)
	memoir.XP += len(newAchievements) * 42

	// Генерируем новый слоган дня
	memoir.Slogan = generateSlogan(days)

	saveMemoir(memoir)

	// Вывод
	printSnowLegend(memoir)
	printStats(memoir)
	printMotivation()
	printDisclaimer()

	if len(newAchievements) > 0 {
		fmt.Println("\n✨ РАЗБЛОКИРОВАНО НОВЫХ ДОСТИЖЕНИЙ:", len(newAchievements))
		for _, a := range newAchievements {
			fmt.Println("   🏅", a)
		}
		fmt.Println("   ➕ +", len(newAchievements)*42, "XP")
	}
}

func loadOrCreateMemoir() *Memoir {
	file, err := os.Open(saveFile)
	if os.IsNotExist(err) {
		// Первый запуск
		m := &Memoir{
			StartDate:    startDate,
			LastCheck:    time.Now().Format("2006-01-02"),
			DaysCount:    0,
			Achievements: []string{},
			XP:           0,
			Slogan:       "Вафли и бисквитный рулетик — топливо кода.",
		}
		saveMemoir(m)
		return m
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения файла:", err)
		return &Memoir{StartDate: startDate, LastCheck: time.Now().Format("2006-01-02"), DaysCount: 0}
	}
	defer file.Close()

	var m Memoir
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&m)
	if err != nil {
		fmt.Println("⚠️ Ошибка парсинга JSON. Создаём новую запись.")
		return &Memoir{StartDate: startDate, LastCheck: time.Now().Format("2006-01-02"), DaysCount: 0}
	}
	return &m
}

func saveMemoir(m *Memoir) {
	file, err := os.Create(saveFile)
	if err != nil {
		fmt.Println("⚠️ Не могу сохранить файл:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(m)
}

func checkAchievements(m *Memoir) []string {
	var unlocked []string
	for days, name := range achievementsMap {
		if m.DaysCount >= days {
			found := false
			for _, a := range m.Achievements {
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
		"Дворники чистят снег, а Гоша — код.",
		"Чай остывает, но модули не стареют.",
		"Вафля сегодня, вафля завтра — а прогресс навсегда.",
		"Бисквитный рулетик знаний.",
		"Депрессия уходит в sleep(), а код работает forever.",
		"Снег растает, goroutines останутся.",
		"Курсы, подкасты, статьи про Go — мой daily driver.",
		"CapCut удалён, а каналы открыты.",
		"100 дней позади, 365 — впереди.",
	}
	return slogans[days%len(slogans)]
}

func printSnowLegend(m *Memoir) {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                 🧠 СНЕЖНАЯ ЛЕГЕНДА ДНЯ 42 🧠            ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  12 февраля 2026. Четверг. За окном — снег.              ║\n")
	fmt.Printf("║  Гоша на кухне. Горячий чай, вафли, бисквитный рулетик.  ║\n")
	fmt.Printf("║  Планшет в руках — курс по Go. В голове — каша из слайсов ║\n")
	fmt.Printf("║  и горутин. Дворники расчищают снег, а Гоша — свой мозг. ║\n")
	fmt.Printf("║  Стратегия: «Мозговой штурм». Вебинары, подкасты, статьи.║\n")
	fmt.Printf("║  Всё, что связано с Go — впитываю как губка.             ║\n")
	fmt.Printf("║                                                          ║\n")
	fmt.Printf("║  CapCut удалён 18 января. Сегодня — %3d-й день.          ║\n", m.DaysCount)
	fmt.Printf("║  Слоган дня: «%s»  ║\n", m.Slogan)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printStats(m *Memoir) {
	fmt.Println("📊 СТАТИСТИКА СОПРОТИВЛЕНИЯ")
	fmt.Println("───────────────────────────")
	fmt.Printf("🗓️  Старт: %s\n", m.StartDate)
	fmt.Printf("📅 Сегодня: %s\n", m.LastCheck)
	fmt.Printf("🔥 Дней без CapCut: %d\n", m.DaysCount)
	fmt.Printf("⭐ Опыт (XP): %d\n", m.XP)
	fmt.Printf("🏅 Достижений: %d\n", len(m.Achievements))
	if len(m.Achievements) > 0 {
		fmt.Println("🏆 РАЗБЛОКИРОВАНО:")
		for _, a := range m.Achievements {
			fmt.Println("   •", a)
		}
	}
	fmt.Println()
}

func printMotivation() {
	fmt.Println("💬 10 СНЕЖНЫХ МОТИВАТОРОВ:")
	motivation := []string{
		"1. 42 дня — это 1/8 года. Ты прошёл этот путь, не сломавшись.",
		"2. Каждый день без CapCut — +100 к карме разработчика.",
		"3. Вафля с утра, код вечером — режим настоящего IT-инженера.",
		"4. Дворники чистят снег, а ты чистишь свой код-стайл.",
		"5. Депрессия — не повод ставить CapCut. Повод написать пару горутин.",
		"6. Смотреть курсы по Go в снегопад — эстетика программиста.",
		"7. Сериалы подождут, подкасты про Go — слушай и усваивай.",
		"8. Бары закрыты, а твой мозг открыт для новых знаний.",
		"9. 100 дней уже позади. 365 — станут твоим годом силы.",
		"10. Гоша, ты не просто учишь Go — ты переплавляешь депрессию в код.",
	}
	for _, m := range motivation {
		fmt.Println(m)
	}
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== ДИСКЛЕЙМЕР ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("CapCut — отличный редактор, но сейчас не время его устанавливать.")
	fmt.Println("Вафли «К чаю» одобрены Министерством Go-разработки.")
	fmt.Println("===================")
}
