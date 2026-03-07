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
	
	// Новые поля для расширенной системы
	SkillTree     *SkillTree
	Quests        *QuestSystem
	Level         int      // Уровень игрока
	Experience    int      // Опыт для следующего уровня
	PlayTime      int      // Время игры в минутах
	DaysPlayed    int      // Количество сыгранных дней
}

func NewPlayer(name string) *Player {
	rand.Seed(time.Now().UnixNano())
	player := &Player{
		Name:        name,
		Focus:       70,
		GoKnowledge: 40,
		Willpower:   65,
		Money:       500,
		Dopamine:    200,
		Level:       1,
		Experience:  0,
		PlayTime:    0,
		DaysPlayed:  1,
	}
	
	// Инициализируем систему навыков и квестов
	player.SkillTree = NewSkillTree()
	player.Quests = NewQuestSystem()
	
	// Применяем начальные бонусы от навыков
	player.ApplySkillBonuses()
	
	return player
}

func (p *Player) DisplayStatus() {
	fmt.Println("👤 СТАТУС ИГРОКА:")
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Уровень: %d (Опыт: %d/%d)\n", p.Level, p.Experience, p.Level*100)
	fmt.Printf("Фокус: %d%% (бонус: +%d)\n", p.Focus, p.SkillTree.GetTotalBonus("focus"))
	fmt.Printf("Знание Go: %d/100 (бонус: +%d)\n", p.GoKnowledge, p.SkillTree.GetTotalBonus("knowledge"))
	fmt.Printf("Сила воли: %d%% (бонус: +%d)\n", p.Willpower, p.SkillTree.GetTotalBonus("willpower"))
	fmt.Printf("Деньги: %d₽ (бонус: +%d)\n", p.Money, p.SkillTree.GetTotalBonus("money"))
	fmt.Printf("Дофамин: %d (бонус: +%d)\n", p.Dopamine, p.SkillTree.GetTotalBonus("dopamine"))
	fmt.Printf("Дней сыграно: %d | Время: %d мин\n", p.DaysPlayed, p.PlayTime)
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
	fmt.Printf("Эффект: %s | XP: +%d\n", m.Effect, m.XPBonus)

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

// ApplySkillBonuses применяет все бонусы от навыков
func (p *Player) ApplySkillBonuses() {
	p.Focus += p.SkillTree.GetTotalBonus("focus")
	p.Willpower += p.SkillTree.GetTotalBonus("willpower")
	p.GoKnowledge += p.SkillTree.GetTotalBonus("knowledge")
	p.Money += p.SkillTree.GetTotalBonus("money")
	p.Dopamine += p.SkillTree.GetTotalBonus("dopamine")
	
	// Ограничиваем значения
	if p.Focus > 100 { p.Focus = 100 }
	if p.Willpower > 100 { p.Willpower = 100 }
	if p.GoKnowledge > 100 { p.GoKnowledge = 100 }
}

// AddExperience добавляет опыт и проверяет повышение уровня
func (p *Player) AddExperience(xp int) {
	p.Experience += xp
	fmt.Printf("\n✨ +%d опыта\n", xp)
	
	// Проверяем повышение уровня
	xpNeeded := p.Level * 100
	for p.Experience >= xpNeeded {
		p.LevelUp()
		xpNeeded = p.Level * 100
	}
}

// LevelUp повышает уровень игрока
func (p *Player) LevelUp() {
	p.Experience -= p.Level * 100
	p.Level++
	p.Focus = 100
	p.Willpower = 100
	
	fmt.Printf("\n🎉 УРОВЕНЬ ПОВЫШЕН! Теперь уровень %d\n", p.Level)
	fmt.Println("   Фокус и сила воли восстановлены!")
	
	// Начисляем очки навыков
	skillPoints := 2 + (p.Level / 5) // 2 очка + 1 каждые 5 уровней
	p.SkillTree.EarnSkillPoints(skillPoints)
	
	// Добавляем достижение
	p.Achievements = append(p.Achievements, fmt.Sprintf("Достигнут уровень %d", p.Level))
	
	// Особые достижения
	if p.Level == 5 {
		p.Achievements = append(p.Achievements, "🏆 Go-Новичок: Уровень 5")
	}
	if p.Level == 10 {
		p.Achievements = append(p.Achievements, "🏆 Go-Разработчик: Уровень 10")
	}
	if p.Level == 20 {
		p.Achievements = append(p.Achievements, "🏆 Go-Мастер: Уровень 20")
	}
	if p.Level == 30 {
		p.Achievements = append(p.Achievements, "🏆 Go-Легенда: Уровень 30")
	}
}

// AddPlayTime добавляет время игры
func (p *Player) AddPlayTime(minutes int) {
	p.PlayTime += minutes
}

// DisplayStatistics отображает расширенную статистику
func (p *Player) DisplayStatistics() {
	fmt.Println("\n📊 СТАТИСТИКА ИГРЫ")
	fmt.Println("══════════════════════════════════════")
	fmt.Printf("🎮 Уровень: %d\n", p.Level)
	fmt.Printf("⭐ Опыт: %d/%d\n", p.Experience, p.Level*100)
	fmt.Printf("🕐 Время в игре: %d минут\n", p.PlayTime)
	fmt.Printf("📅 Дней сыграно: %d\n", p.DaysPlayed)
	fmt.Printf("💪 Преодолено искушений: %d\n", len(p.Temptations))
	fmt.Printf("🏆 Достигнуто достижений: %d\n", len(p.Achievements))
	fmt.Printf("🎯 Выполнено квестов: %d\n", p.Quests.TotalCompleted)
	fmt.Printf("🔥 Текущая серия: %d дней\n", p.Quests.DayStreak)
	fmt.Printf("✨ Очков навыков: %d (всего: %d)\n", p.SkillTree.SkillPoints, p.SkillTree.TotalPoints)
	
	// Рассчитываем общий рейтинг
	rating := p.CalculateScore() + (p.Level * 100) + (len(p.Achievements) * 50)
	fmt.Printf("\n🏅 ОБЩИЙ РЕЙТИНГ: %d\n", rating)
	
	if rating < 500 {
		fmt.Println("Ранг: 🌱 Начинающий гофер")
	} else if rating < 1500 {
		fmt.Println("Ранг: 🌿 Ученик разработчика")
	} else if rating < 3000 {
		fmt.Println("Ранг: 🌳 Junior Go Developer")
	} else if rating < 5000 {
		fmt.Println("Ранг: 🏢 Middle Go Developer")
	} else {
		fmt.Println("Ранг: 🚀 Senior Go Master")
	}
}
