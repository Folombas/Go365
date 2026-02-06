package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Игровой персонаж - Гоша
type GoshCharacter struct {
	Name         string
	GoXP         int      // Опыт в Go
	Focus        int      // Уровень фокуса (0-100)
	Motivation   int      // Мотивация (0-100)
	Distractions []string // Отвлекающие факторы
	Level        int      // Уровень разработчика
	Wallet       int      // Текущий баланс
}

// Функция обучения Go с элементами геймификации
func learnGo(g *GoshCharacter, hours int) {
	fmt.Printf("\n🎮 УРОВЕНЬ %d: НАЧАЛО ОБУЧЕНИЯ GO\n", g.Level)
	fmt.Println(strings.Repeat("=", 40))

	lessons := []struct {
		topic string
		xp    int
		cost  int
	}{
		{"Основы синтаксиса Go", 20, 1},
		{"Горутины и каналы", 30, 2},
		{"Пакет sync (WaitGroup, Mutex)", 40, 3},
		{"Работа с БД (PostgreSQL)", 50, 4},
		{"Микросервисы на Go", 60, 5},
		{"Docker и контейнеризация", 70, 6},
		{"Kubernetes для Go-разработчика", 80, 7},
		{"gRPC и protobuf", 90, 8},
		{"Оптимизация и профилирование", 100, 9},
	}

	for i := 0; i < hours; i++ {
		if g.Focus <= 10 {
			fmt.Println("⚠️  ВНИМАНИЕ: Уровень фокуса критически низкий!")
			g.rest(1)
			continue
		}

		// Случайный отвлекающий фактор
		if rand.Intn(100) < 30 { // 30% шанс на отвлечение
			distraction := g.Distractions[rand.Intn(len(g.Distractions))]
			fmt.Printf("🎭 ОТВЛЕЧЕНИЕ: %s\n", distraction)
			g.Focus -= 15
			g.Motivation -= 10
			fmt.Printf("   📉 Фокус: %d, Мотивация: %d\n", g.Focus, g.Motivation)

			// Борьба с отвлечением
			if rand.Intn(100) < g.Motivation {
				fmt.Println("   ✅ Гоша поборол отвлечение и вернулся к учебе!")
				g.Motivation += 5
			} else {
				fmt.Println("   ❌ Гоша поддался отвлечению... Потеряно время.")
				continue
			}
		}

		// Изучение урока
		lesson := lessons[rand.Intn(len(lessons))]
		g.GoXP += lesson.xp
		g.Focus -= lesson.cost

		fmt.Printf("📚 Час %d: Изучаю %s\n", i+1, lesson.topic)
		fmt.Printf("   ➕ +%d XP Go | 📉 Фокус: %d\n", lesson.xp, g.Focus)

		// Проверка на уровень
		if g.GoXP >= g.Level*1000 && g.Level < 10 {
			g.Level++
			fmt.Printf("\n🎉 ПОВЫШЕНИЕ УРОВНЯ! Теперь Гоша %d уровня!\n", g.Level)
			fmt.Println("   🎁 Награда: +50 мотивации, +30 фокуса")
			g.Motivation += 50
			g.Focus += 30
		}

		// Случайная мотивационная фраза
		if rand.Intn(100) < 20 {
			printMotivation()
		}

		time.Sleep(500 * time.Millisecond)
	}
}

// Функция отдыха
func (g *GoshCharacter) rest(hours int) {
	fmt.Printf("\n💤 ОТДЫХ: Гоша отдыхает %d часа\n", hours)
	for i := 0; i < hours; i++ {
		g.Focus += 25
		g.Motivation += 15
		fmt.Printf("   Час %d: 📈 Фокус: %d, Мотивация: %d\n", i+1, g.Focus, g.Motivation)
		time.Sleep(300 * time.Millisecond)
	}
}

// Работа курьером (альтернатива)
func workAsCourier(g *GoshCharacter, hours int) {
	fmt.Printf("\n📦 РАБОТА КУРЬЕРОМ: %d часов\n", hours)
	for i := 0; i < hours; i++ {
		money := 300 + rand.Intn(101) // 300-400 рублей
		g.Wallet += money
		g.Motivation -= 20
		g.GoXP += 5 // Минимальный опыт

		fmt.Printf("   Час %d: Доставил заказ | 💰 +%d руб | 📉 Мотивация: %d\n",
			i+1, money, g.Motivation)

		// Шанс на осознание
		if rand.Intn(100) < 15 {
			fmt.Println("   💡 Мысль: 'Каждый час работы курьером = -20 мотивации. Нужно учить Go!'")
		}
		time.Sleep(400 * time.Millisecond)
	}
}

// 10 мотивационных фраз для Гоши
var motivationPhrases = []string{
	"🔥 Каждая строчка кода — это шаг к зарплате 150К+",
	"🚀 Сегодняшние мучения с Go — завтрашние предложения о работе",
	"💪 Помни: каждый, кто сейчас получает 150К+, тоже начинал с нуля",
	"🎯 Фокус на Go сейчас — свобода выбора завтра",
	"📈 Инвестиция в свои навыки — самая выгодная инвестиция",
	"🧠 Твоя голова стоит миллионов, если наполнить её правильными знаниями",
	"⚡ Каждый раз, когда хочется бросить, представь свою первую зарплату Go-разработчика",
	"🏆 Go — это не просто язык, это билет в другой социальный класс",
	"💼 300 рублей за доставку или 150К+ за Go-код? Выбор за тобой!",
	"🚫 CapCut подождёт. Сначала — финансовая независимость!",
}

func printMotivation() {
	phrase := motivationPhrases[rand.Intn(len(motivationPhrases))]
	fmt.Printf("\n💬 МОТИВАЦИЯ: %s\n", phrase)
}

// Симуляция дня Гоши
func simulateGoshDay() {
	fmt.Println("🎮 FOCUS FIGHT: БИТВА ЗА 150К+")
	fmt.Println("══════════════════════════════════════")
	fmt.Println("👤 Персонаж: Гоша | Бюджет: 0 руб")
	fmt.Println("📱 Смартфон: Honor 10x Lite | Процессор: Kirin 710A")
	fmt.Println("🎯 Цель: Go-разработчик с зарплатой 150К+")
	fmt.Println(strings.Repeat("=", 50))

	// Инициализация персонажа
	gosh := &GoshCharacter{
		Name:       "Гоша",
		GoXP:       350,
		Focus:      85,
		Motivation: 70,
		Distractions: []string{
			"Хочется установить CapCut и монтировать видео",
			"Вспомнил про летний отпуск на Филиппинах в 2019",
			"Задумался о своей нищете...",
			"Проверить соцсети и ленту",
			"Сделать ещё фото каши для блога",
			"Поиграть в мобильную игру",
			"Посмотреть сериал вместо учёбы",
			"Позвонить другу и пожаловаться на жизнь",
		},
		Level:  1,
		Wallet: 0,
	}

	// Утренняя рутина
	fmt.Println("\n🌅 УТРО 6 ФЕВРАЛЯ 2026:")
	fmt.Println("─────────────────────")
	fmt.Println("⏰ 07:00 - Подъём по будильнику")
	fmt.Println("🏋️  - Легкая зарядка")
	fmt.Println("🪒 - Бритьё, умывание")
	fmt.Println("🍌 - Геркулесовая каша с бананом и мёдом")
	fmt.Println("📸 - Фото каши для food-блога")
	fmt.Println("🍵 - Чашка чая")
	fmt.Println("😔 - Накатила депрессия...")

	// Критический выбор
	fmt.Println("\n⚡ КРИТИЧЕСКИЙ ВЫБОР:")
	fmt.Println("───────────────────")
	fmt.Println("1. Установить CapCut и монтировать видео (прокрастинация)")
	fmt.Println("2. Пойти на работу курьером (300-400 руб/час)")
	fmt.Println("3. УЧИТЬ GO (инвестиция в будущее)")

	// Гоша делает выбор (случайно, но с уклоном к правильному)
	choice := 3
	if rand.Intn(100) < 30 { // 70% шанс на правильный выбор
		choice = 3
	} else {
		choice = 1 + rand.Intn(2)
	}

	switch choice {
	case 1:
		fmt.Println("\n❌ ВЫБРАН ПУТЬ ПРОКРАСТИНАЦИИ")
		fmt.Println("Гоша устанавливает CapCut...")
		fmt.Println("Монтирует видео 8-летней давности...")
		fmt.Println("Через 5 часов понимает, что это не приносит денег")
		gosh.Motivation = 10
		gosh.Wallet = 0
		fmt.Println("💔 Итог: 0 рублей, -60 мотивации, депрессия усугубилась")

	case 2:
		fmt.Println("\n⚠️  ВЫБРАН ПУТЬ КУРЬЕРА")
		workAsCourier(gosh, 8)
		fmt.Printf("\n💰 ИТОГ ДНЯ: Заработано %d рублей\n", gosh.Wallet)
		fmt.Printf("📊 Go XP: %d (прирост минимальный)\n", gosh.GoXP)
		fmt.Println("😞 Вечером Гоша понимает: это не путь к 150К+")

	case 3:
		fmt.Println("\n✅ ВЫБРАН ПУТЬ GO-РАЗРАБОТЧИКА!")
		fmt.Println("🔥 Гоша включает режим максимального фокуса")
		fmt.Println("🚫 CapCut удалён, соцсети заблокированы")

		// День обучения
		learnGo(gosh, 8)

		// Вечерняя рефлексия
		fmt.Println("\n🌙 ВЕЧЕРНЯЯ РЕФЛЕКСИЯ:")
		fmt.Println("────────────────────")
		fmt.Printf("📊 Итоги дня:\n")
		fmt.Printf("   🎮 Уровень: %d\n", gosh.Level)
		fmt.Printf("   🧠 Go XP: %d\n", gosh.GoXP)
		fmt.Printf("   💰 Кошелек: %d руб\n", gosh.Wallet)
		fmt.Printf("   🎯 Фокус: %d/100\n", gosh.Focus)
		fmt.Printf("   🔥 Мотивация: %d/100\n", gosh.Motivation)

		// Прогресс к цели
		progress := (float64(gosh.GoXP) / 10000.0) * 100
		if progress > 100 {
			progress = 100
		}
		fmt.Printf("\n📈 ПРОГРЕСС ДО 150К+: %.1f%%\n", progress)

		// Дофаминовая награда
		if gosh.GoXP > 500 {
			fmt.Println("\n🎁 ДОФАМИНОВАЯ НАГРАДА:")
			fmt.Println("   + Чувство контроля над жизнью")
			fmt.Println("   + Уверенность в завтрашнем дне")
			fmt.Println("   + Растущая самооценка")
			fmt.Println("   + Уважение к самому себе")
		}
	}

	// Финальное сообщение
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("🎯 ЗАПОМНИ: Каждый день выбора определяет твоё будущее")
	fmt.Println("💡 Go сегодня — 150К+ завтра")
	fmt.Println(strings.Repeat("═", 50))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	simulateGoshDay()

	// Вывод всех мотивационных фраз в конце
	fmt.Println("\n🔊 10 МОТИВАЦИОННЫХ ФРАЗ ДЛЯ ТЕБЯ:")
	fmt.Println(strings.Repeat("─", 40))
	for i, phrase := range motivationPhrases {
		fmt.Printf("%2d. %s\n", i+1, phrase)
	}

	fmt.Println("\n" + strings.Repeat("★", 50))
	fmt.Println("DISCLAIMER: Все персонажи вымышлены. История создана")
	fmt.Println("для мотивации изучения Go. Совпадения случайны.")
	fmt.Println(strings.Repeat("★", 50))
}
