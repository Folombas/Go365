package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ================== СТРУКТУРЫ ДАННЫХ ==================

type Thought struct {
	Type        string
	Intensity   int // 1-100, сила мысли
	Description string
	Category    string
	Reward      int // XP за преодоление
}

type FocusState struct {
	GoXP                 int
	FocusLevel           int
	Willpower            int
	Money                int
	DistractionsOvercome int
	Day                  int
}

// ================== КОНСТАНТЫ ==================

var (
	// 10 мотивационных фраз
	motivations = []string{
		"🔥 Каждый час за Go = 625 рублей будущей зарплаты (150К/мес ÷ 240 часов)",
		"🚀 Сегодняшний отказ от CapCut = завтрашний MacBook Pro для реального монтажа",
		"💪 Преодоленная мысль о море = +1% к вероятности устроиться в бигтех",
		"🎯 Go-разработчик в 2026 = как нефтяник в 2006: деньги качают пачками",
		"📈 300 рублей доната за видео vs 150К+ зарплаты = выбор очевиден!",
		"🧠 Мозг, заточенный под горутины, ценнее мозга, заточенного под монтаж",
		"⚡ Каждая пройденная тема по Go = минус один день в статусе 'курьер'",
		"🏆 Через 6 месяцев упорного Go ты будешь смеяться над сегодняшними сомнениями",
		"💼 Финансовая импотенция лечится не CapCut, а строчками кода на Go",
		"🚫 Пальмы подождут. Сначала - офис с кондиционером, потом - отпуск на Бали",
	}

	// Категории мыслей
	thoughtCategories = map[string][]string{
		"ностальгия": {
			"Воспоминание о теплом море",
			"Запах кокосового масла на коже",
			"Шум прибоя в заливе Эль-Нидо",
			"Закат над Палаваном",
			"Девушки в бикини на пляже",
		},
		"творчество": {
			"Хочется установить CapCut",
			"Идея для нового влога",
			"Можно добавить крутые переходы",
			"ИИ-эффекты в новой версии",
			"Собрать все отпускные видео",
		},
		"финансы": {
			"Мама просит 2500 рублей на продукты",
			"Пустой кошелек",
			"Курьерство = 400 руб/час, Go = 1000 руб/час",
			"Ипотека, которую не потянуть",
			"Старость без накоплений",
		},
		"учеба": {
			"Устал смотреть курс",
			"Хочется сделать перерыв",
			"Эта тема слишком сложная",
			"Может, завтра начну",
			"Все равно не устроюсь",
		},
		"альтернативы": {
			"Можно заняться n8n",
			"Blender и 3D-моделирование",
			"Геймдев на Unreal Engine",
			"Python для ML",
			"Frontend на React",
		},
	}
)

// ================== FAN-IN РЕАЛИЗАЦИЯ ==================

// generateThoughts создает мысли в своей категории
func generateThoughts(category string, thoughts chan<- Thought, wg *sync.WaitGroup) {
	defer wg.Done()

	types, exists := thoughtCategories[category]
	if !exists {
		return
	}

	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

		thoughtType := types[rand.Intn(len(types))]
		intensity := rand.Intn(50) + 50 // 50-100
		reward := intensity * 2

		thought := Thought{
			Type:        thoughtType,
			Intensity:   intensity,
			Description: fmt.Sprintf("[%s] %s", category, thoughtType),
			Category:    category,
			Reward:      reward,
		}

		thoughts <- thought
	}
}

// fanIn объединяет несколько каналов в один
func fanIn(sources []chan Thought) <-chan Thought {
	combined := make(chan Thought, 20)
	var wg sync.WaitGroup

	output := func(source <-chan Thought) {
		defer wg.Done()
		for thought := range source {
			combined <- thought
		}
	}

	wg.Add(len(sources))
	for _, source := range sources {
		go output(source)
	}

	go func() {
		wg.Wait()
		close(combined)
	}()

	return combined
}

// ================== ИГРОВАЯ ЛОГИКА ==================

func processDay(day int, state *FocusState) {
	fmt.Printf("\n📅 ДЕНЬ %d: УТРО 8 ФЕВРАЛЯ 2026\n", day)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("🌅 Гоша просыпается. За окном -20°C и сугробы.")
	fmt.Println("💭 В голове: 'Установить CapCut или учить Go?'")
	fmt.Printf("💰 В кармане: %d рублей", state.Money)
	fmt.Println(strings.Repeat("-", 50))

	// Создаем каналы для разных категорий мыслей
	var channels []chan Thought
	var wg sync.WaitGroup

	categories := []string{"ностальгия", "творчество", "финансы", "удеба", "альтернативы"}
	for _, category := range categories {
		ch := make(chan Thought, 2)
		channels = append(channels, ch)
		wg.Add(1)
		go generateThoughts(category, ch, &wg)
	}

	// Fan-In: объединяем все каналы
	combined := fanIn(channels)

	// Обработка мыслей
	fmt.Println("\n🧠 НАЧИНАЮ FAN-IN ОБРАБОТКУ МЫСЛЕЙ...")
	thoughtsProcessed := 0

	for thought := range combined {
		thoughtsProcessed++
		fmt.Printf("\n💭 Мысль %d: %s\n", thoughtsProcessed, thought.Description)
		fmt.Printf("   📊 Интенсивность: %d/100\n", thought.Intensity)
		fmt.Printf("   💪 Воля: %d/100 | 🧠 Фокус: %d/100\n", state.Willpower, state.FocusLevel)

		// Проверяем, преодолеет ли Гоша мысль
		resistance := state.Willpower + state.FocusLevel
		if resistance >= thought.Intensity {
			// Успех!
			state.DistractionsOvercome++
			state.GoXP += thought.Reward
			state.Willpower = min(100, state.Willpower+5)
			state.FocusLevel = min(100, state.FocusLevel+3)

			fmt.Printf("   ✅ ПРЕОДОЛЕНО! +%d Go XP\n", thought.Reward)

			// Периодическая мотивация
			if thoughtsProcessed%2 == 0 {
				motivation := motivations[rand.Intn(len(motivations))]
				fmt.Printf("   💬 МОТИВАЦИЯ: %s\n", motivation)
			}

			// Достижения
			if state.DistractionsOvercome == 5 {
				fmt.Println("   🏆 ДОСТИЖЕНИЕ: 'Железная воля' разблокировано!")
			}
		} else {
			// Неудача
			state.Willpower = max(0, state.Willpower-10)
			state.FocusLevel = max(0, state.FocusLevel-5)
			state.Money -= 200 // Тратит деньги на ненужное

			fmt.Printf("   ❌ ПОДДАЛСЯ! -200 рублей, воля падает\n")

			// Критическое состояние
			if state.Willpower < 30 {
				fmt.Println("   ⚠️  ТРЕВОГА: Уровень воли критический!")
				fmt.Println("   💊 Принял решение: сфокусироваться на Go")
				state.Willpower += 30
			}
		}

		fmt.Printf("   📈 Текущий Go XP: %d\n", state.GoXP)
		time.Sleep(600 * time.Millisecond)

		if thoughtsProcessed >= 12 {
			break
		}
	}

	// Закрываем каналы и ждем завершения
	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("📊", 25))
	fmt.Println("ИТОГИ ДНЯ:")
	fmt.Printf("   ✅ Преодолено мыслей: %d\n", state.DistractionsOvercome)
	fmt.Printf("   🧠 Накопленный Go XP: %d\n", state.GoXP)
	fmt.Printf("   💰 Состояние кошелька: %d рублей\n", state.Money)
	fmt.Printf("   💪 Уровень воли: %d/100\n", state.Willpower)
	fmt.Printf("   🎯 Уровень фокуса: %d/100\n", state.FocusLevel)

	// Прогресс до цели
	targetXP := 10000
	progress := float64(state.GoXP) / float64(targetXP) * 100
	if progress > 100 {
		progress = 100
	}
	fmt.Printf("   🎯 Прогресс до 150К+: %.1f%%\n", progress)

	// Повышение уровня
	if state.GoXP >= 500 && state.Day == 1 {
		fmt.Println("\n⭐ УРОВЕНЬ ПОВЫШЕН! Теперь Гоша 'Начинающий Gopher'")
		fmt.Println("   🎁 Награда: +50 к уверенности, доступ к новым курсам")
	}
}

// ================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ==================

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ================== ОСНОВНАЯ ПРОГРАММА ==================

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(strings.Repeat("🧠", 60))
	fmt.Println("FAN-IN FOCUS: СОБИРАЕМ ВОЛЮ В КУЛАК")
	fmt.Println(strings.Repeat("🧠", 60))
	fmt.Println("Паттерн Fan-In в борьбе с внутренними демонами")
	fmt.Println(strings.Repeat("=", 60))

	// Инициализация состояния
	state := &FocusState{
		GoXP:       350,
		FocusLevel: 65,
		Willpower:  70,
		Money:      3000,
		Day:        1,
	}

	// Симуляция 3 дней
	for day := 1; day <= 3; day++ {
		state.Day = day
		processDay(day, state)

		// Между днями - восстановление
		if day < 3 {
			fmt.Println("\n🌙 Ночь... Гоша анализирует день и готовится к завтрашней битве")
			state.Willpower = min(100, state.Willpower+20)
			state.FocusLevel = min(100, state.FocusLevel+15)
			state.Money += 1200 // Заработок за день
			time.Sleep(1 * time.Second)
		}
	}

	// Финальные результаты
	fmt.Println("\n" + strings.Repeat("🏆", 60))
	fmt.Println("ТРИ ДНЯ БОРЬБЫ ЗАВЕРШЕНЫ!")
	fmt.Println(strings.Repeat("🏆", 60))
	fmt.Printf("\n🎯 ФИНАЛЬНАЯ СТАТИСТИКА:\n")
	fmt.Printf("   🧠 Всего Go XP: %d\n", state.GoXP)
	fmt.Printf("   ✅ Преодолено отвлечений: %d\n", state.DistractionsOvercome)
	fmt.Printf("   💰 Финансы: %d рублей\n", state.Money)
	fmt.Printf("   💪 Воля: %d/100 (начало: 70)\n", state.Willpower)
	fmt.Printf("   🎯 Фокус: %d/100 (начало: 65)\n", state.FocusLevel)

	// Оценка успеха
	successRate := float64(state.DistractionsOvercome) / 36.0 * 100
	fmt.Printf("   📈 Успешность: %.1f%%\n", successRate)

	if successRate > 70 {
		fmt.Println("\n🏅 РЕЗУЛЬТАТ: ОТЛИЧНО! Гоша на верном пути к 150К+")
		fmt.Println("   Он не установил CapCut. Не потратил время на монтаж.")
		fmt.Println("   Каждый день приближает его к цели.")
	} else if successRate > 40 {
		fmt.Println("\n🥈 РЕЗУЛЬТАТ: ХОРОШО! Есть прогресс, но можно лучше")
		fmt.Println("   Гоша иногда сдавался, но в целом держал курс на Go.")
		fmt.Println("   Нужно работать над волей.")
	} else {
		fmt.Println("\n⚠️  РЕЗУЛЬТАТ: НУЖНО УСИЛИТЬ ФОКУС!")
		fmt.Println("   Слишком много мыслей победили Гошу.")
		fmt.Println("   Пора принимать серьезные решения.")
	}

	// Техническая демонстрация Fan-In
	fmt.Println("\n" + strings.Repeat("🔧", 60))
	fmt.Println("ТЕХНИЧЕСКАЯ ЧАСТЬ: PATTERN FAN-IN В GO")
	fmt.Println(strings.Repeat("🔧", 60))
	fmt.Println("Что делает Fan-In:")
	fmt.Println("   • Берет N каналов-источников (категории мыслей)")
	fmt.Println("   • Объединяет их в один канал")
	fmt.Println("   • Позволяет обрабатывать все данные в одном месте")
	fmt.Println("\nКлючевые строки кода:")
	fmt.Println("   combined := fanIn(channels) // Объединение")
	fmt.Println("   for thought := range combined { // Обработка }")
	fmt.Println("\nПреимущества для Гоши:")
	fmt.Println("   ✅ Все мысли в одном месте - легче анализировать")
	fmt.Println("   ✅ Централизованная обработка - последовательные решения")
	fmt.Println("   ✅ Масштабируемость - легко добавить новые категории мыслей")

	// Все мотивационные фразы
	fmt.Println("\n" + strings.Repeat("💪", 60))
	fmt.Println("10 МОТИВАЦИОННЫХ ФРАЗ ДЛЯ БУДУЩЕГО GO-РАЗРАБОТЧИКА:")
	fmt.Println(strings.Repeat("💪", 60))
	for i, phrase := range motivations {
		fmt.Printf("%2d. %s\n", i+1, phrase)
	}

	// Финальная мысль
	fmt.Println("\n" + strings.Repeat("🌟", 60))
	fmt.Println("ФИНАЛЬНАЯ МЫСЛЬ ГОШИ:")
	fmt.Println("'CapCut может подождать. Пальмы никуда не денутся.")
	fmt.Println("А вот возможность изменить свою жизнь - может уйти.")
	fmt.Println("Сегодня я выбираю Go. Завтра я буду благодарен себе.'")
	fmt.Println(strings.Repeat("🌟", 60))

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("⚠️", 60))
	fmt.Println("DISCLAIMER: Все персонажи вымышлены.")
	fmt.Println("История создана для мотивации и изучения паттернов Go.")
	fmt.Println("Совпадения с реальностью случайны.")
	fmt.Println(strings.Repeat("⚠️", 60))
}
