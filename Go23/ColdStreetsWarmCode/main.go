package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player (Гоша) - структура игрока
type Player struct {
	Name           string
	Age            int
	Location       string
	Cash           int
	Energy         int
	Focus          int
	GoSkill        int
	Distractions   []string
	DeliveryPoints []string
	Level          int
	Motivation     string
}

// NewGosha создает нового персонажа Гошу
func NewGosha() *Player {
	return &Player{
		Name:     "Гоша",
		Age:      37,
		Location: "Кухня, панелька в спальном районе промзоны",
		Cash:     2500,
		Energy:   90,
		Focus:    65,
		GoSkill:  23, // Уровень дня!
		Distractions: []string{
			"Видеомонтаж отпуска 2019",
			"3D-модель дракона в Blender",
			"Настройка нового вокодера",
			"Steam: новая игра со скидкой 90%",
		},
		DeliveryPoints: []string{
			"ММДЦ Москва-Сити, башня 'Федерация'",
			"БЦ 'Остров', 18 этаж",
			"Ленинский пр-т, 123, офис 404",
			"ЦАО, Старая площадь, 8",
			"МФК 'Город столиц', холл",
		},
		Level:      23,
		Motivation: "Теплый офис к 40 годам. Backend. Go.",
	}
}

// MorningRoutine - утренний ритуал
func (p *Player) MorningRoutine() {
	fmt.Printf("\n🌅 6:30. %s. %s.\n", p.Location, p.Motivation)
	fmt.Println("Действия:")
	fmt.Println("  1. Умыться [✓]")
	fmt.Println("  2. Позавтракать, чай [✓]")
	fmt.Println("  3. Проводить маму [✓]")
	fmt.Println("  4. Проверить телефон [✓]")

	// Случайное утреннее искушение
	distraction := p.Distractions[rand.Intn(len(p.Distractions))]
	fmt.Printf("⚠️  Искушение: '%s'\n", distraction)

	// Проверка фокуса
	if p.Focus > 50 {
		fmt.Println("💪 Фокус достаточен. Игнорируем. Запускаем приложение курьера.")
		p.Focus -= 10
	} else {
		fmt.Println("😩 Фокус на нуле. 15 минут потеряны...")
		p.Energy -= 15
		p.Cash -= 300 // Упущенный заказ
	}
}

// DeliverPackages - симуляция рабочего дня курьера
func (p *Player) DeliverPackages() {
	fmt.Printf("\n🚚 НАЧАЛО РАБОЧЕГО ДНЯ. Уровень энергии: %d/100\n", p.Energy)
	fmt.Println("================================================")

	completedDeliveries := 0
	for i, point := range p.DeliveryPoints {
		if p.Energy <= 0 {
			fmt.Println("\n💀 ЭНЕРГИЯ НА НУЛЕ. Гоша не может продолжать.")
			break
		}

		fmt.Printf("\nДоставка %d/%d: %s\n", i+1, len(p.DeliveryPoints), point)

		// Случайные события на маршруте
		eventRoll := rand.Intn(100)
		switch {
		case eventRoll < 20: // Гололёд
			fmt.Println("   ❄️  ОПАСНОСТЬ: Гололёд на переходе!")
			p.Energy -= 25
			p.Cash += 180 // Премия за риск
			fmt.Println("   +180 руб. (премия за риск), -25 энергии")

		case eventRoll < 40: // Пробка
			fmt.Println("   🚗 ПРОБКА: ТТК, 12 км...")
			time.Sleep(800 * time.Millisecond)
			p.Energy -= 15
			fmt.Println("   -15 энергии, -20 минут")

		case eventRoll < 60: // Успешная доставка
			fmt.Println("   ✅ ДОСТАВЛЕНО! Клиент доволен.")
			p.Cash += 450
			p.Energy -= 20
			completedDeliveries++
			fmt.Println("   +450 руб., -20 энергии")

		case eventRoll < 80: // Соблазн
			distraction := p.Distractions[rand.Intn(len(p.Distractions))]
			fmt.Printf("   🎮 СОБЛАЗН: '%s'\n", distraction)
			fmt.Print("   Сопротивляться? (фокус проверка > 60): ")
			if p.Focus > 60 {
				fmt.Println("Успех! Фокус сохранен.")
				p.Focus += 5
			} else {
				fmt.Println("Провал... 30 минут потеряны.")
				p.Energy -= 30
			}

		default: // Неожиданный бонус
			fmt.Println("   🎉 БОНУС: Заказчик дал чаевые + наводку на курсы Go!")
			p.Cash += 700
			p.GoSkill += 2
			fmt.Println("   +700 руб., +2 к навыку Go")
		}

		// Мини-восстановление между заказами
		if i < len(p.DeliveryPoints)-1 {
			fmt.Println("   ⏸️  Перекур? Нет. Глоток воды и дальше.")
			p.Energy += 5
		}
	}

	fmt.Printf("\n📦 ИТОГ ДНЯ: Доставок %d/%d, Заработано: %d руб.\n",
		completedDeliveries, len(p.DeliveryPoints), completedDeliveries*450)
}

// EveningCoding - вечернее изучение Go
func (p *Player) EveningCoding() {
	fmt.Println("\n🌙 21:00. Квартира. Чай. Компуктер.")
	fmt.Println("========================================")

	if p.Energy < 30 {
		fmt.Printf("⚠️  Энергия всего %d/100. Шанс уснуть за клавиатурой: высокий.\n", p.Energy)
		successChance := p.Focus * 100 / 70
		fmt.Printf("   Шанс успешного изучения: %d%%\n", successChance)
	}

	// Выбор темы для изучения
	topics := []struct {
		name   string
		energy int
		skill  int
		focus  int
	}{
		{"Горутины и каналы", 35, 5, -10},
		{"Интерфейсы и embed", 25, 3, -5},
		{"Работа с ошибками", 20, 2, 0},
		{"Тестирование (go test)", 30, 4, -15},
	}

	topic := topics[rand.Intn(len(topics))]
	fmt.Printf("🎯 Тема дня: %s\n", topic.name)

	if p.Energy >= topic.energy {
		p.Energy -= topic.energy
		p.GoSkill += topic.skill
		p.Focus += topic.focus

		fmt.Printf("✅ УСПЕХ! Потрачено энергии: %d\n", topic.energy)
		fmt.Printf("   Навык Go: %d (+%d)\n", p.GoSkill, topic.skill)

		// Критическое осознание
		if rand.Intn(100) > 70 {
			fmt.Println("   💡 ОЗАРЕНИЕ: Внезапно понял механизм работы context!")
			p.GoSkill += 3
		}
	} else {
		fmt.Println("💤 ПРОВАЛ... Уснул на третьей строчке кода.")
		fmt.Println("   Навык Go не изменился, энергия на минимуме.")
		p.Energy = 10
	}
}

// DisplayStats - отображение статистики персонажа
func (p *Player) DisplayStats() {
	fmt.Println("\n📊 СТАТИСТИКА ПЕРСОНАЖА:")
	fmt.Println("══════════════════════════════════")
	fmt.Printf("  Имя:      %s\n", p.Name)
	fmt.Printf("  Возраст:  %d лет\n", p.Age)
	fmt.Printf("  Уровень:  %d\n", p.Level)
	fmt.Printf("  Деньги:   %d руб.\n", p.Cash)
	fmt.Printf("  Энергия:  %d/100\n", p.Energy)
	fmt.Printf("  Фокус:    %d/100\n", p.Focus)
	fmt.Printf("  Навык Go: %d/100\n", p.GoSkill)
	fmt.Printf("  Мотивация: %s\n", p.Motivation)

	// Прогресс до цели
	progress := float64(p.GoSkill) / 100.0 * 100.0
	fmt.Printf("\n  📈 Прогресс до офиса: %.1f%%\n", progress)

	// Визуализация прогресса
	bar := ""
	for i := 0; i < 20; i++ {
		if float64(i)/20.0*100.0 <= progress {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	fmt.Printf("  [%s]\n", bar)

	// Расчёт до цели
	if p.GoSkill < 80 {
		daysLeft := (80 - p.GoSkill) * 2
		fmt.Printf("  ⏳ До минимального уровня для джуна: ~%d дней\n", daysLeft)
	} else {
		fmt.Println("  🎯 Минимальный уровень для джуна достигнут!")
	}
}

// TriggerCrisis - кризисный момент (трагикомедия)
func (p *Player) TriggerCrisis() {
	fmt.Println("\n🔴 КРИЗИСНЫЙ МОМЕНТ")
	fmt.Println("══════════════════════════════════")

	crises := []string{
		"Внезапно понял, что потратил 3 часа на чтение споров про 'лучший фреймворк'",
		"Коллега-курьер рассказал, что устроился в Тинькофф после 6 месяцев Python",
		"Мама звонила: 'Сынок, может, пойдешь на сварщика? Там стабильно'",
		"Увидел в телеграме: 'Go-разработчик, удаленка, от 250к, опыт от 3 лет'",
		"Начал пет-проект, сломал git, потерял 2 дня работы",
	}

	crisis := crises[rand.Intn(len(crises))]
	fmt.Printf("💔 %s\n", crisis)

	// Реакция на кризис
	reactions := []struct {
		text   string
		energy int
		focus  int
		skill  int
	}{
		{"Заесть печенькой и продолжать", -10, 5, 1},
		{"Лечь спать, завтра новый день", 40, 20, 0},
		{"Написать гневный пост на Хабр", -30, -20, 0},
		{"Открыть учебник и пройти еще одну тему", -25, 15, 3},
	}

	reaction := reactions[rand.Intn(len(reactions))]
	fmt.Printf("🤔 Реакция: %s\n", reaction.text)

	p.Energy += reaction.energy
	p.Focus += reaction.focus
	p.GoSkill += reaction.skill

	// Ограничиваем статы
	if p.Energy > 100 { p.Energy = 100 }
	if p.Focus > 100 { p.Focus = 100 }
	if p.GoSkill > 100 { p.GoSkill = 100 }
	if p.Energy < 0 { p.Energy = 0 }
	if p.Focus < 0 { p.Focus = 0 }
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎮 ХОЛОДНЫЕ УЛИЦЫ, ТЁПЛЫЙ КОД")
	fmt.Println("========================================")
	fmt.Println("Уровень 23: Баланс между мечтой и реальностью")
	fmt.Println("========================================")

	gosha := NewGosha()

	// Игровой цикл дня
	gosha.MorningRoutine()
	gosha.DeliverPackages()

	// Случайный кризис после работы
	if rand.Intn(100) > 40 {
		gosha.TriggerCrisis()
	}

	gosha.EveningCoding()
	gosha.DisplayStats()

	// Финальная оценка дня
	fmt.Println("\n⭐ ИТОГ ДНЯ 23:")
	if gosha.GoSkill >= 25 {
		fmt.Println("ПРОГРЕСС ЕСТЬ! Навык Go вырос. Мечта об офисе становится ближе.")
		fmt.Println("Гоша засыпает с мыслью о завтрашнем учебном модуле.")
	} else if gosha.Energy < 20 {
		fmt.Println("ПЕРЕУТОМЛЕНИЕ. Нужен отдых. Завтра новый день.")
		fmt.Println("Иногда отступление - часть пути к цели.")
	} else {
		fmt.Println("СТАБИЛЬНО. Ни шагу назад, ни шагу вперед.")
		fmt.Println("Завтра нужно будет сделать больше.")
	}

	fmt.Println("\n💭 'Backend. Go. Теплый офис.' - повторяет Гоша, закрывая глаза.")
	fmt.Println("================================================")
}
