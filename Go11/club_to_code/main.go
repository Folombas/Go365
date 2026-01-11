package main

import (
	"fmt"
	"strings"
	"time"
)

// ========== СТРУКТУРЫ ДАННЫХ ==========

// PastLife - прошлая жизнь в шоу-бизнесе
type PastLife struct {
	Period       string
	Profession   string
	Achievements []string
	Perks        []string
	Status       string
}

// PresentLife - настоящая жизнь
type PresentLife struct {
	Date        time.Time
	Problems    []string
	Realization string
	Decision    string
}

// FutureGoal - будущая цель
type FutureGoal struct {
	Profession string
	Skills     []string
	Motivation string
	Guarantee  string
}

// Choice - выбор пути
type Choice struct {
	Option      string
	Consequence string
	IsCorrect   bool
}

// ========== КОНСТАНТЫ ==========
const (
	BOX_WIDTH = 50
)

// ========== ОСНОВНАЯ ПРОГРАММА ==========
func main() {
	// Инициализация данных
	past := PastLife{
		Period:     "10-15 лет назад",
		Profession: "Гламурный рэп-артист на клубных фрик-шоу",
		Achievements: []string{
			"Выступления в клубах",
			"Полные залы фанатов",
			"Съёмки в клипах",
			"Гастроли по городам",
			"Гримёрки с закуской и выпивкой",
			"Лимузин с озорной Зиной",
		},
		Perks: []string{
			"Лимузины",
			"Гримёрки с закусками",
			"Фотосессии",
			"Восторженные фанатки",
		},
		Status: "РЭП-ЗВЕЗДА",
	}

	present := PresentLife{
		Date: time.Date(2026, 1, 11, 0, 0, 0, 0, time.UTC),
		Problems: []string{
			"Еле собирает на кварплату",
			"Позабытый продюсерами и коллегами",
			"Не позвали на выступление в субботу",
			"Умеет только читать рэп в микрофон",
		},
		Realization: "Шоу-бизнес - это безумные виражи, головокружительные взлёты и падения",
		Decision:    "Надо учиться чему-то ещё, кроме выступлений",
	}

	future := FutureGoal{
		Profession: "Junior Go Backend Developer",
		Skills: []string{
			"Go программирование",
			"Backend разработка",
			"Базы данных",
			"API проектирование",
			"Docker и DevOps",
		},
		Motivation: "Код нужно писать всегда. Это стабильность. Это постоянство.",
		Guarantee:  "Навыки программирования будут кормить всю жизнь!",
	}

	choices := []Choice{
		{
			Option:      "Вернуться в клубы, упрашивать продюсеров",
			Consequence: "❌ Остаться на улице за дверью клубов",
			IsCorrect:   false,
		},
		{
			Option:      "Пойти таскать коробки по снежным сугробам",
			Consequence: "❌ Заморозить амбиции вместе с руками",
			IsCorrect:   false,
		},
		{
			Option:      "Учиться Go-программированию каждый день",
			Consequence: "✅ Стать востребованным разработчиком",
			IsCorrect:   true,
		},
	}

	// Вывод истории
	printHeader()
	printPastLife(past)
	printPresentLife(present)
	printFutureGoal(future)
	printGameChoice(choices)
	printConclusion()
	printDisclaimer()
}

// ========== ФУНКЦИИ ВЫВОДА ==========

func printHeader() {
	title := "🎤 → 💻 CLUB TO CODE: ИСТОРИЯ ПЕРЕРОЖДЕНИЯ ГОШИ ИЗ РЭП-ЗВЕЗДЫ В ПРОГРАММИСТА"

	fmt.Println()
	fmt.Println("╔" + strings.Repeat("═", BOX_WIDTH-2) + "╗")
	fmt.Printf("║ %-46s ║\n", title)
	fmt.Println("╚" + strings.Repeat("═", BOX_WIDTH-2) + "╝")
	fmt.Println()
	fmt.Println("📅 11 января 2026 года")
	fmt.Println("📍 Перекрёсток судьбы: клубы vs код")
	fmt.Println()
}

func printPastLife(past PastLife) {
	fmt.Println("🎭 ПРОШЛОЕ: КОГДА ГОША БЫЛ ЗВЕЗДОЙ")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))

	fmt.Printf("Период: %s\n", past.Period)
	fmt.Printf("Профессия: %s\n", past.Profession)
	fmt.Printf("Статус: %s\n\n", past.Status)

	fmt.Println("🏆 ДОСТИЖЕНИЯ:")
	for achievementIndex, achievement := range past.Achievements {
		fmt.Printf("  %d. %s\n", achievementIndex+1, achievement)
	}

	fmt.Println("\n✨ ПЛЮШКИ:")
	for _, perk := range past.Perks { // Исправлено: используем _ вместо perkIndex
		fmt.Printf("  • %s\n", perk)
	}

	fmt.Println()
	fmt.Println("💬 Гоша тогда: 'Сегодня ты - рэп-звезда и собираешь полные клубы!'")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))
	fmt.Println()
}

func printPresentLife(present PresentLife) {
	fmt.Println("😔 НАСТОЯЩЕЕ: ХМУРОЕ ПРОБУЖДЕНИЕ ОТ КЛУБНОГО УГАРА")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))

	fmt.Printf("Дата: %s\n\n", present.Date.Format("02.01.2006"))

	fmt.Println("⚠️ ПРОБЛЕМЫ:")
	for problemIndex, problem := range present.Problems {
		fmt.Printf("  %d. %s\n", problemIndex+1, problem)
	}

	fmt.Println()
	fmt.Printf("💡 ОСОЗНАНИЕ: %s\n", present.Realization)
	fmt.Printf("🎯 РЕШЕНИЕ: %s\n", present.Decision)

	fmt.Println()
	fmt.Println("💬 Гоша сейчас: 'Завтра - таскаешь тяжёлые коробки по снежным сугробам'")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))
	fmt.Println()
}

func printFutureGoal(future FutureGoal) {
	fmt.Println("🚀 БУДУЩЕЕ: ЦИФРОВАЯ ТРАНСФОРМАЦИЯ")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))

	fmt.Printf("Целевая профессия: %s\n\n", future.Profession)

	fmt.Println("📚 НАВЫКИ ДЛЯ ОСВОЕНИЯ:")
	for skillIndex, skill := range future.Skills {
		fmt.Printf("  %d. %s\n", skillIndex+1, skill)
	}

	fmt.Println()
	fmt.Printf("💪 МОТИВАЦИЯ: %s\n", future.Motivation)
	fmt.Printf("✅ ГАРАНТИЯ: %s\n", future.Guarantee)

	fmt.Println()
	fmt.Println("💬 Гоша о будущем: 'Выступления в клубах - мимолётный успех. Навыки программирования - на всю жизнь!'")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))
	fmt.Println()
}

func printGameChoice(choices []Choice) {
	fmt.Println("🎮 ГЕЙМИФИКАЦИЯ: ВЫБОР ПУТИ")
	fmt.Println(strings.Repeat("─", BOX_WIDTH))
	fmt.Println()

	fmt.Println("Перед Гошей три пути:")
	fmt.Println()

	for choiceIndex, choice := range choices {
		number := fmt.Sprintf("%d", choiceIndex+1)
		if choice.IsCorrect {
			fmt.Printf("  🟢 %s. %s\n", number, choice.Option)
		} else {
			fmt.Printf("  🔴 %s. %s\n", number, choice.Option)
		}
	}

	fmt.Println()
	fmt.Println("❓ Что выберет Гоша? (Введите 1, 2 или 3)")

	var userChoice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&userChoice)

	fmt.Println()
	fmt.Println("🎯 РЕЗУЛЬТАТ ВЫБОРА:")

	if userChoice >= 1 && userChoice <= len(choices) {
		selected := choices[userChoice-1]

		// Анимация размышления
		fmt.Print("Гоша размышляет")
		for count := 0; count < 3; count++ {
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()

		if selected.IsCorrect {
			fmt.Println("✅ ПРАВИЛЬНЫЙ ВЫБОР!")
			fmt.Println(selected.Consequence)

			// Анимация прогресса
			fmt.Println()
			fmt.Println("📈 ПРОГРЕСС ОБУЧЕНИЯ GO:")
			for day := 0; day <= 10; day++ {
				progress := day * 10
				fmt.Printf("  День %d: [", day)
				for block := 0; block < 10; block++ {
					if block < day {
						fmt.Print("█")
					} else {
						fmt.Print("░")
					}
				}
				fmt.Printf("] %d%%\n", progress)
				time.Sleep(200 * time.Millisecond)
			}

			fmt.Println()
			fmt.Println("🎉 ГОША СТАНОВИТСЯ JUNIOR GO DEVELOPER!")
			fmt.Println("   Зарплата: достойная")
			fmt.Println("   Стабильность: гарантирована")
			fmt.Println("   Будущее: светлое")
		} else {
			fmt.Println("❌ ОШИБКА!")
			fmt.Println(selected.Consequence)

			fmt.Println()
			fmt.Println("💀 ГОША ПРОДОЛЖАЕТ ЖИТЬ В ПРОШЛОМ")
			fmt.Println("   Долги: растут")
			fmt.Println("   Перспективы: туманные")
			fmt.Println("   Будущее: непредсказуемое")
		}
	} else {
		fmt.Println("⚠️ Гоша слишком долго выбирал...")
		fmt.Println("   Система автоматически выбрала путь Go разработчика!")
		fmt.Println("   ✅ Правильный выбор сделан за него!")
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", BOX_WIDTH))
}

func printConclusion() {
	fmt.Println()
	fmt.Println("🎭 МОРАЛЬ ИСТОРИИ")
	fmt.Println(strings.Repeat("═", BOX_WIDTH))

	lessons := []string{
		"1. Шоу-бизнес — это головокружительные виражи, программирование — это стабильность",
		"2. Выступления на сцене мимолётен, навыки программирования — на всю жизнь",
		"3. Забытые продюсерами и конферансье ≠ забытый рынком труда",
		"4. Микрофон в руках временно, Macbook — везде с тобой",
		"5. Клубные выступления кормят сегодня, код будет кормить всегда",
	}

	for _, lesson := range lessons {
		fmt.Println(lesson)
	}

	fmt.Println()
	fmt.Println("💬 ФИНАЛЬНЫЕ СЛОВА ГОШИ:")
	fmt.Println("   'Пусть народ отдыхает и колбасится в клубах под выступления моих дорогих и уважаемых друзей, коллег по сцене...'")
	fmt.Println("   'А я вот выбрал код. Я выбрал стабильность. Я выбрал уверенность в будущем, изучая современную айти-профессию.'")
	fmt.Println("   'Из потешного клубного фрика — в востребованного программиста. Из клубов — в айти.'")

	fmt.Println()
	fmt.Println(strings.Repeat("═", BOX_WIDTH))
}

func printDisclaimer() {
	fmt.Println()
	fmt.Println("⚠️ DISCLAIMER - ВАЖНОЕ УВЕДОМЛЕНИЕ")
	fmt.Println(strings.Repeat("═", BOX_WIDTH))

	disclaimer := []string{
		"ВНИМАНИЕ: Данная программа является художественным произведением.",
		"",
		"📜 ЮРИДИЧЕСКОЕ УВЕДОМЛЕНИЕ:",
		"1. Все персонажи, включая Гошу, его друзей, продюсеров и коллег,",
		"   являются ВЫМЫШЛЕННЫМИ и созданы в творческих целях.",
		"",
		"2. Все истории, сюжеты, события и диалоги ПРИДУМАНЫ автором.",
		"   Любые совпадения с реальными людьми, живыми или умершими,",
		"   организациями, событиями или ситуациями являются СЛУЧАЙНЫМИ",
		"   и НЕПРЕДНАМЕРЕННЫМИ.",
		"",
		"3. Проект представляет собой ХУДОЖЕСТВЕННУЮ ВЫДУМКУ в стиле",
		"   программного кода, созданную для:",
		"   • Развлечения и мотивации",
		"   • Демонстрации возможностей языка Go",
		"   • Практики программирования через сторителлинг",
		"",
		"4. Не является профессиональной консультацией по карьере,",
		"   психологии или образованию.",
		"",
		"🎭 Творческая лицензия: использование элементов драматургии,",
		"   сатиры и гиперболы для создания увлекательного контента о",
		"   программировании.",
	}

	for _, line := range disclaimer {
		fmt.Println(line)
	}

	fmt.Println()
	fmt.Println("© 2026 Go365 | Художественный вымысел в образовательных целях")
	fmt.Println(strings.Repeat("═", BOX_WIDTH))
}
