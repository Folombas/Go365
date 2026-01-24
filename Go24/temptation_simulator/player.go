package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name          string
	Focus         int      // Фокус на учебе (0-100)
	GoKnowledge   int      // Знание Go (0-100)
	Willpower     int      // Сила воли (0-100)
	Money         int      // Заработанные деньги
	Dopamine      int      // Уровень дофамина
	Temptations   []string // Список преодоленных искушений
	Achievements  []string // Достижения
}

func NewPlayer(name string) *Player {
	rand.Seed(time.Now().UnixNano())
	return &Player{
		Name:        name,
		Focus:       70,
		GoKnowledge: 40,
		Willpower:   65,
		Money:       500,
		Dopamine:    200,
	}
}

func (p *Player) DisplayStatus() {
	fmt.Println("👤 СТАТУС ИГРОКА:")
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Фокус: %d%%\n", p.Focus)
	fmt.Printf("Знание Go: %d/100\n", p.GoKnowledge)
	fmt.Printf("Сила воли: %d%%\n", p.Willpower)
	fmt.Printf(  "Деньги: %d₽\n", p.Money)
	fmt.Printf("Дофамин: %d\n", p.Dopamine)
	fmt.Println()
}

func (p *Player) HandleTemptation(t Temptation) {
	fmt.Printf("\n⚠️  ИСКУШЕНИЕ: %s\n", t.Name)
	fmt.Printf("Сила: %d%% | Описание: %s\n", t.Power, t.Description)

	resistChance := p.Willpower - t.Power + 50

	if resistChance > 0 && rand.Intn(100) < resistChance {
		// Успешное сопротивление
		fmt.Println("✅ ВЫ СОПРОТИВИЛИСЬ! Фокус +10, Сила воли +5")
		p.Focus += 10
		p.Willpower += 5
		p.Dopamine += 50
		p.Temptations = append(p.Temptations, t.Name)

		// Изучение Go
		p.StudyGo(15)
	} else {
		// Поддались искушению
		fmt.Println("❌ ВЫ ПОДДАЛИСЬ! Фокус -20, Сила воли -10")
		p.Focus -= 20
		p.Willpower -= 10
		p.Dopamine -= 100

		if p.Focus < 0 {
			p.Focus = 0
		}
	}
}

func (p *Player) StudyGo(points int) {
	oldLevel := p.GoKnowledge / 10
	p.GoKnowledge += points

	if p.GoKnowledge > 100 {
		p.GoKnowledge = 100
	}

	newLevel := p.GoKnowledge / 10

	if newLevel > oldLevel {
		fmt.Printf("🎓 УРОВЕНЬ ПОВЫШЕН! Знание Go: %d/100\n", p.GoKnowledge)
		p.Dopamine += 200
		p.Achievements = append(p.Achievements,
			fmt.Sprintf("Уровень Go %d достигнут", newLevel))
	}
}

func (p *Player) ReceiveMotivation(m Motivation) {
	fmt.Printf("\n💪 МОТИВАЦИЯ: %s\n", m.Text)
	fmt.Printf("Эффект: %s\n", m.Effect)

	switch m.Effect {
	case "focus+":
		p.Focus += 15
		p.Dopamine += 100
	case "willpower+":
		p.Willpower += 20
		p.Dopamine += 80
	case "knowledge+":
		p.StudyGo(25)
		p.Dopamine += 150
	}

	if p.Focus > 100 {
		p.Focus = 100
	}
	if p.Willpower > 100 {
		p.Willpower = 100
	}
}

func (p *Player) FinalBattle(t Temptation) bool {
	fmt.Println("\n⚔️  ФИНАЛЬНАЯ СХВАТКА С ИСКУШЕНИЕМ!")
	fmt.Println("Бес говорит: 'Установи CapCut! Монтируй тропические видосы!'")
	fmt.Println("Гофер говорит: 'Концентрируйся на Go! Сначала работа, потом хобби!'")

	// Расчет шансов
	successChance := (p.Willpower * 2 + p.Focus) / 3

	fmt.Printf("Ваш шанс на победу: %d%%\n", successChance)
	fmt.Println("Нажмите Enter, чтобы бросить кубик судьбы...")
	fmt.Scanln()

	roll := rand.Intn(100)

	if roll < successChance {
		fmt.Println("🎉 ПОБЕДА! Вы остались верны Go!")
		p.Focus = 100
		p.Willpower = 100
		p.Dopamine += 500
		p.Achievements = append(p.Achievements, "Победил финальное искушение")
		return true
	} else {
		fmt.Println("💀 ПОРАЖЕНИЕ... Искушение победило")
		p.Focus = 30
		p.Willpower = 40
		p.Dopamine -= 300
		return false
	}
}

func (p *Player) CalculateScore() int {
	return p.GoKnowledge * 10 + p.Focus * 5 + p.Willpower * 3 + p.Dopamine / 10
}

func (p *Player) ShowAchievements() {
	if len(p.Achievements) > 0 {
		fmt.Println("\n🏆 ВАШИ ДОСТИЖЕНИЯ:")
		for _, achievement := range p.Achievements {
			fmt.Printf("  ✓ %s\n", achievement)
		}
	}
}
