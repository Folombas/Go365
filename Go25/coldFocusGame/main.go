
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==================== ОСНОВНЫЕ ТИПЫ ====================

// Player - игровой персонаж
type Player struct {
	Name        string
	Health      int
	Focus       int
	Energy      int
	Money       int
	Warmth      int
	GoKnowledge int
	Willpower   int
}

// GameWorld - игровой мир
type GameWorld struct {
	Time        string
	Temperature int
	Phase       int // 0:Утро, 1:День, 2:Вечер, 3:Ночь
	DayComplete bool
}

// ==================== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ====================

func renderBar(value, max int, icon string) {
	barLength := 20
	filled := (value * barLength) / max
	if filled > barLength {
		filled = barLength
	}
	if filled < 0 {
		filled = 0
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)
	fmt.Printf("[%s] %s %d%%\n", bar, icon, value)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// ==================== ИГРОК ====================

func NewPlayer(name string) *Player {
	return &Player{
		Name:        name,
		Health:      85,
		Focus:       40,
		Energy:      70,
		Money:       1000,
		Warmth:      30,
		GoKnowledge: 45,
		Willpower:   60,
	}
}

func (p *Player) ShowStatus() {
	fmt.Println("\n👤 СТАТУС ИГРОКА:")
	fmt.Printf("Здоровье: %d/100 ", p.Health)
	renderBar(p.Health, 100, "❤️")

	fmt.Printf("Фокус: %d/100 ", p.Focus)
	renderBar(p.Focus, 100, "🎯")

	fmt.Printf("Энергия: %d/100 ", p.Energy)
	renderBar(p.Energy, 100, "⚡")

	fmt.Printf("Теплота: %d/100 ", p.Warmth)
	renderBar(p.Warmth, 100, "🔥")

	fmt.Printf("Знание Go: %d/100 ", p.GoKnowledge)
	renderBar(p.GoKnowledge, 100, "🐹")

	fmt.Printf("Сила воли: %d/100 ", p.Willpower)
	renderBar(p.Willpower, 100, "🛡️")

	fmt.Printf("Деньги: %d₽\n", p.Money)
}

func (p *Player) StudyGo(points int) {
	p.GoKnowledge += points
	if p.GoKnowledge > 100 {
		p.GoKnowledge = 100
	}
}

// ==================== ИГРОВОЙ МИР ====================

func NewGameWorld() *GameWorld {
	rand.Seed(time.Now().UnixNano())
	return &GameWorld{
		Time:        "11:00",
		Temperature: -5,
		Phase:       0,
		DayComplete: false,
	}
}

func (gw *GameWorld) ShowIntro() {
	fmt.Println(`
❄️  COLD FOCUS: The Go Coder's Survival Simulator
================================================

ТЫ ПРОСЫПАЕШСЯ В ХОЛОДНОЙ КОМНАТЕ...
Балкон был открыт всю ночь. На улице -5°C.
Твоя миссия: пережить день и изучить Go!

ВНИМАНИЕ: Отвлекающие факторы повсюду!
Твоя сила воли будет постоянно проверяться.

НАЖМИ ENTER, ЧТОБЫ НАЧАТЬ...
	`)
	fmt.Scanln()
}

func (gw *GameWorld) RenderStatus(p *Player) {
	clearScreen()
	timeColors := []string{"🌅", "🌞", "🌇", "🌙"}
	phaseNames := []string{"УТРО", "ДЕНЬ", "ВЕЧЕР", "НОЧЬ"}

	fmt.Printf("\n🕐 %s %s | 🌡️  Температура: %d°C\n\n",
		timeColors[gw.Phase], phaseNames[gw.Phase], gw.Temperature)

	p.ShowStatus()
}

// ==================== ИГРОВОЙ ЦИКЛ ====================

func (gw *GameWorld) GetPhaseActions() []string {
	switch gw.Phase {
	case 0:
		return []string{
			"Встать с кровати",
			"Умыться ледяной водой",
			"Позавтракать",
			"Проверить телефон",
		}
	case 1:
		return []string{
			"Взять заказ из Химок",
			"Поехать на электричке",
			"Ковыряться в смартфоне",
			"Купить хлеб и картошку",
		}
	case 2:
		return []string{
			"Поужинать макарошками",
			"Выпить шиповник",
			"Посмотреть телевизор",
			"Уйти в комнату",
		}
	default:
		return []string{
			"Изучить горутины",
			"Попрактиковаться с каналами",
			"Решить задачу на LeetCode",
			"Посмотреть документацию",
		}
	}
}

func (gw *GameWorld) ExecuteAction(p *Player, choice int) string {
	switch gw.Phase {
	case 0:
		return gw.executeMorningAction(p, choice)
	case 1:
		return gw.executeDayAction(p, choice)
	case 2:
		return gw.executeEveningAction(p, choice)
	default:
		return gw.executeNightAction(p, choice)
	}
}

// ==================== ДЕЙСТВИЯ ПО ФАЗАМ ====================

func (gw *GameWorld) executeMorningAction(p *Player, choice int) string {
	switch choice {
	case 1:
		p.Energy += 20
		p.Focus += 10
		gw.Temperature = -3
		return "🛏️  Ты встал с кровати! Энергия +20, Фокус +10"
	case 2:
		p.Energy += 15
		p.Warmth -= 10
		p.Focus += 5
		return "🚿 Умылся ледяной водой! Энергия +15, Теплота -10"
	case 3:
		p.Health += 10
		p.Energy += 25
		p.Money -= 150
		return "🍳 Позавтракал! Здоровье +10, Энергия +25, Деньги -150₽"
	default:
		if rand.Intn(100) < 60 {
			p.Willpower += 5
			return "📱 Удержался от утреннего скроллинга! Сила воли +5"
		}
		p.Energy -= 20
		p.Focus -= 10
		return "📱 Провел 30 минут в соцсетях... Энергия -20, Фокус -10"
	}
}

func (gw *GameWorld) executeDayAction(p *Player, choice int) string {
	switch choice {
	case 1:
		p.Money += 1000
		p.Energy -= 30
		p.Focus += 15
		if rand.Intn(100) < 20 {
			tip := rand.Intn(200) + 50
			p.Money += tip
			return fmt.Sprintf("🚗 Выполнил заказ! +1000₽ (+%d₽ чаевые), Энергия -30", tip)
		}
		return "🚗 Выполнил заказ! +1000₽, Энергия -30"
	case 2:
		p.Energy -= 10
		p.Warmth += 5
		p.Focus += 5
		return "🚆 Поездка на электричке. Энергия -10, Теплота +5"
	case 3:
		if rand.Intn(100) < 70 {
			p.Willpower += 10
			return "⏰ Устоял перед прокрастинацией! Сила воли +10"
		}
		p.Energy -= 25
		p.Focus -= 20
		return "⏰ Потратил час на бессмысленный скроллинг... Энергия -25, Фокус -20"
	default:
		p.Money -= 300
		p.Health += 5
		p.Energy -= 15
		return "🛒 Купил хлеб и картошку. Деньги -300₽, Здоровье +5"
	}
}

func (gw *GameWorld) executeEveningAction(p *Player, choice int) string {
	switch choice {
	case 1:
		p.Health += 20
		p.Energy += 30
		p.Warmth += 10
		p.Focus += 5
		return "🍝 Макарошки с сосисками! Здоровье +20, Энергия +30, Теплота +10"
	case 2:
		p.Health += 10
		p.Warmth += 15
		p.Focus += 10
		return "🍵 Горячий шиповник! Здоровье +10, Теплота +15, Фокус +10"
	case 3:
		if rand.Intn(100) < 65 {
			p.Willpower += 15
			return "📺 Ушел от телевизора! Сила воли +15"
		}
		p.Energy -= 40
		p.Focus -= 25
		return "📺 Засмотрелся на телепередачу... Энергия -40, Фокус -25"
	default:
		p.Focus += 20
		p.Warmth += 5
		return "🚪 Ушел в свою комнату. Фокус +20, Теплота +5"
	}
}

func (gw *GameWorld) executeNightAction(p *Player, choice int) string {
	switch choice {
	case 1:
		p.StudyGo(25)
		p.Energy -= 20
		p.Focus -= 10
		return "🤹 Изучал горутины! Знание Go +25, Энергия -20"
	case 2:
		p.StudyGo(20)
		p.Energy -= 15
		p.Focus -= 5
		return "🔀 Практиковался с каналами. Знание Go +20"
	case 3:
		if rand.Intn(100) < (p.GoKnowledge/2 + 30) {
			p.StudyGo(15)
			p.Focus += 10
			p.Money += 50
			return "🧠 Решил задачу на LeetCode! Знание Go +15, Фокус +10, +50₽"
		}
		p.Energy -= 10
		p.Focus -= 5
		return "🧠 Не смог решить задачу... Энергия -10"
	default:
		p.StudyGo(10)
		p.Energy -= 5
		p.Focus += 5
		return "📚 Читал документацию Go. Знание Go +10, Фокус +5"
	}
}

// ==================== MINI GAME ====================

func playGoQuiz(p *Player) {
	fmt.Println("\n🎮 МИНИ-ИГРА: GO QUIZ")
	fmt.Println("Ответь на 3 вопроса по Go!")
	fmt.Println("Нажми ENTER чтобы начать...")
	fmt.Scanln()

	questions := []struct {
		q      string
		options []string
		answer int
		points int
	}{
		{
			q:      "Что делает 'defer' в Go?",
			options: []string{
				"Откладывает выполнение функции",
				"Прерывает программу",
				"Создает горутину",
				"Обрабатывает ошибки",
			},
			answer: 0,
			points: 10,
		},
		{
			q:      "Как создать буферизированный канал?",
			options: []string{
				"make(chan int, 5)",
				"new(chan int, 5)",
				"chan int{5}",
				"create(chan int).buffer(5)",
			},
			answer: 0,
			points: 15,
		},
		{
			q:      "Что такое горутина?",
			options: []string{
				"Легковесный поток выполнения",
				"Функция обратного вызова",
				"Метод структуры",
				"Тип данных",
			},
			answer: 0,
			points: 12,
		},
	}

	correct := 0
	for i, q := range questions {
		fmt.Printf("\nВопрос %d: %s\n", i+1, q.q)
		for j, opt := range q.options {
			fmt.Printf("  %d. %s\n", j+1, opt)
		}

		var answer int
		fmt.Print("Твой ответ (1-4): ")
		fmt.Scanln(&answer)

		if answer-1 == q.answer {
			fmt.Printf("✅ ПРАВИЛЬНО! +%d знаний\n", q.points)
			p.GoKnowledge += q.points / 2
			correct++
		} else {
			fmt.Printf("❌ НЕПРАВИЛЬНО. Правильный ответ: %d\n", q.answer+1)
		}
	}

	accuracy := float64(correct) / 3 * 100
	fmt.Printf("\n🎯 РЕЗУЛЬТАТЫ: %d/3 правильных (%.0f%%)\n", correct, accuracy)

	if accuracy >= 80 {
		p.GoKnowledge += 20
		fmt.Println("🎉 ОТЛИЧНО! +20 к знанию Go!")
	}
}

// ==================== ОСНОВНАЯ ФУНКЦИЯ ====================

func main() {
	fmt.Println("❄️  COLD FOCUS: The Go Coder's Survival Simulator")
	fmt.Println("================================================")
	fmt.Println("Дата: 25 января 2026, воскресенье")
	fmt.Println("Миссия: Выжить в холодной комнате и выучить Go!")
	fmt.Println()

	game := NewGameWorld()
	player := NewPlayer("Гоша")

	game.ShowIntro()

	// Главный игровой цикл по 4 фазам
	for phase := 0; phase < 4; phase++ {
		game.Phase = phase
		actions := 0
		maxActions := 3

		for actions < maxActions {
			game.RenderStatus(player)

			actionsList := game.GetPhaseActions()
			fmt.Println("\n🎮 ДОСТУПНЫЕ ДЕЙСТВИЯ:")
			for i, action := range actionsList {
				fmt.Printf("%d. %s\n", i+1, action)
			}

			var choice int
			fmt.Printf("\n🎯 Выбери действие (1-%d): ", len(actionsList))
			fmt.Scanln(&choice)

			if choice >= 1 && choice <= len(actionsList) {
				result := game.ExecuteAction(player, choice)
				fmt.Printf("\n📢 %s\n", result)
				fmt.Println("---")

				// Случайное событие
				if rand.Intn(100) < 30 {
					events := []string{
						"❄️  Сквозняк с балкона! Теплота -10",
						"📱 Телефон вибрирует...",
						"📺 Звук телевизора доносится с кухни",
						"💡 Пришла гениальная идея для кода!",
					}
					event := events[rand.Intn(len(events))]
					fmt.Printf("🎲 СЛУЧАЙНОЕ СОБЫТИЕ: %s\n", event)

					if event == "❄️  Сквозняк с балкона! Теплота -10" {
						player.Warmth -= 10
					}
				}

				actions++
				time.Sleep(1 * time.Second)
			}
		}

		// Переход между фазами
		if phase < 3 {
			phaseNames := []string{"Утро", "День", "Вечер", "Программирование"}
			fmt.Printf("\n⏰ ПЕРЕХОД К ФАЗЕ: %s\n", phaseNames[phase+1])

			if phase == 2 {
				game.Temperature = 18
				fmt.Println("🔥 ТЕПЛО ВКЛЮЧЕНО! Комната нагревается до 18°C")
			}

			fmt.Println("Нажми ENTER чтобы продолжить...")
			fmt.Scanln()
		}
	}

	// Финальная фаза - программирование
	fmt.Println(`
💻 НАСТУПИЛО ВРЕМЯ ПРОГРАММИРОВАНИЯ!

Ты сидишь в своей комнате. Телевизор доносится с кухни,
но ты должен сосредоточиться. Это твоё время для изучения Go!
	`)

	playGoQuiz(player)

	// Итоги дня
	fmt.Println("\n" + `
╔══════════════════════════╗
║     ИТОГИ ДНЯ 25         ║
╚══════════════════════════╝
	`)

	fmt.Printf("Игрок: %s\n", player.Name)
	fmt.Printf("Знание Go: %d/100\n", player.GoKnowledge)
	fmt.Printf("Деньги заработаны: %d₽\n", player.Money-1000)
	fmt.Printf("Финальный фокус: %d/100\n", player.Focus)

	score := player.GoKnowledge*10 + player.Focus*5 + player.Money/10
	fmt.Printf("\n🏆 ОБЩИЙ СЧЕТ: %d\n", score)

	if score >= 1500 {
		fmt.Println("🎉 ЛЕГЕНДАРНЫЙ ДЕНЬ! Ты почти Senior!")
	} else if score >= 1000 {
		fmt.Println("🎯 ОТЛИЧНО! Хорошо поработал!")
	} else {
		fmt.Println("💪 МОЖНО ЛУЧШЕ! Завтра постарайся!")
	}

	fmt.Println("\n💾 Прогресс сохранен...")
	fmt.Println("Твои знания сохранены для дня 26!")

	fmt.Println("\n" + `
╔══════════════════════════╗
║  10 МОТИВАЦИОННЫХ ФРАЗ   ║
╚══════════════════════════╝
1.  Холодная комната закаляет характер, как Go закаляет код
2.  Каждый отказ от телевизора — шаг к офису с видом на город
3.  Электричка везет тебя домой, а Go везет тебя к карьере
4.  Макарошки насытят тело, а горутины насытят разум
5.  Смартфон отвлекает, клавиатура приближает к цели
6.  Холод пальцев на клавишах — плата за горячий код
7.  Сегодняшний дискомфорт — завтрашний комфорт senior-зарплаты
8.  Телевизор показывает чужие истории, а ты пишешь свою
9.  Go не спрашивает 'холодно ли тебе?' — он просто работает
10. Завтра ты скажешь себе спасибо за сегодняшний код
	`)

	fmt.Println("\n🎮 ДЕНЬ 25 ЗАВЕРШЁН! УДАЧИ В ИЗУЧЕНИИ GO! 🚀")
}
