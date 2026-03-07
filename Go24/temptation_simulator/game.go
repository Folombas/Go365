package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Day           int
	Score         int
	Round         int
	TotalTemptationsResisted int
	TotalTemptations int
	BossDefeated   bool
	StartTime      time.Time
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		Day:   24,
		Score: 0,
		Round: 1,
		StartTime: time.Now(),
	}
}

func (g *Game) StartDay(player *Player) {
	fmt.Println("🌅 НАЧАЛО ДНЯ", g.Day)
	fmt.Println("══════════════════════════════════════")
	fmt.Println("Цель: Не поддаться искушениям и изучить Go")
	fmt.Println("Ваш персонаж:", player.Name)
	fmt.Printf("Уровень: %d | Фокус: %d%%\n", player.Level, player.Focus)
	fmt.Println()
	
	// Генерируем ежедневные квесты
	player.Quests.GenerateDailyQuests()
	fmt.Println()
}

func (g *Game) CheckTemptation() bool {
	// 35% шанс возникновения искушения (увеличили для интереса)
	return rand.Intn(100) < 35
}

func (g *Game) CheckMotivation() bool {
	// 30% шанс получить мотивацию
	return rand.Intn(100) < 30
}

func (g *Game) CheckBossEncounter() bool {
	// 10% шанс встречи с боссом (особо опасное искушение)
	return rand.Intn(100) < 10
}

// HandleTemptationResisted обрабатывает успешное сопротивление
func (g *Game) HandleTemptationResisted(player *Player, t Temptation) {
	g.TotalTemptationsResisted++
	player.Temptations = append(player.Temptations, t.Name)
	
	// Награда за сопротивление
	xpReward := t.Power / 2
	player.AddExperience(xpReward)
	
	// Обновляем прогресс квеста
	player.Quests.UpdateQuestProgress("resist_temptation", 1)
	
	fmt.Printf("   ✨ +%d опыта\n", xpReward)
}

// HandleTemptationFailed обрабатывает поддачу искушению
func (g *Game) HandleTemptationFailed(player *Player, t Temptation) {
	// Потеря опыта
	xpLoss := t.XPLoss
	if player.Experience >= xpLoss {
		player.Experience -= xpLoss
		fmt.Printf("   💀 -%d опыта\n", xpLoss)
	}
	
	// Проверка на понижение уровня
	if player.Experience < 0 {
		player.Level--
		player.Experience = 0
		fmt.Println("   ⚠️  ПОНИЖЕНИЕ УРОВНЯ! Вы потеряли уровень!")
	}
}

// StudyGoSession сессия изучения Go
func (g *Game) StudyGoSession(player *Player, minutes int) {
	fmt.Printf("\n📚 СЕССИЯ ИЗУЧЕНИЯ GO: %d минут\n", minutes)
	
	// Базовый опыт за изучение
	baseXP := minutes / 2
	
	// Бонус от навыков
	knowledgeBonus := player.SkillTree.GetTotalBonus("knowledge")
	totalXP := baseXP + (knowledgeBonus / 5)
	
	player.AddExperience(totalXP)
	player.GoKnowledge += minutes / 5
	if player.GoKnowledge > 100 {
		player.GoKnowledge = 100
	}
	
	// Обновляем квесты
	player.Quests.UpdateQuestProgress("study_go_30min", minutes)
	player.Quests.UpdateQuestProgress("code_practice", minutes/2) // 1 строка кода ~ 2 минуты
	
	// Восстанавливаем дофамин от достижений
	player.Dopamine += minutes / 3
	
	fmt.Printf("   🧠 +%d к знанию Go (всего: %d/100)\n", minutes/5, player.GoKnowledge)
}

// RestSession сессия отдыха
func (g *Game) RestSession(player *Player, minutes int) {
	fmt.Printf("\n💤 ОТДЫХ: %d минут\n", minutes)
	
	player.Focus += minutes / 2
	player.Dopamine += minutes / 3
	
	if player.Focus > 100 {
		player.Focus = 100
	}
	
	fmt.Printf("   😌 Фокус восстановлен: %d%%\n", player.Focus)
}

// UpgradeSkillMenu меню улучшения навыков
func (g *Game) UpgradeSkillMenu(player *Player) {
	fmt.Println("\n🌳 ДЕРЕВО НАВЫКОВ")
	fmt.Println("══════════════════════════════════════")
	
	player.SkillTree.Display()
	
	if player.SkillTree.SkillPoints > 0 {
		fmt.Printf("Доступно очков навыков: %d\n", player.SkillTree.SkillPoints)
		fmt.Println("\nВведите ID навыка для улучшения или 'exit' для выхода:")
		
		var input string
		fmt.Scan(&input)
		
		if input != "exit" {
			if player.SkillTree.UpgradeSkill(input) {
				player.ApplySkillBonuses()
			}
		}
	} else {
		fmt.Println("\n⚠️  Недостаточно очков навыков. Изучайте Go и сопротивляйтесь искушениям!")
	}
}

func (g *Game) EndDay(player *Player, wonBattle bool) {
	// Вычисляем время игры
	playTime := int(time.Since(g.StartTime).Minutes())
	player.AddPlayTime(playTime)
	
	fmt.Println("\n📊 ИТОГИ ДНЯ", g.Day)
	fmt.Println("══════════════════════════════════════")

	g.Score = player.CalculateScore()

	if wonBattle {
		fmt.Println("🎉 ПОБЕДА! Вы успешно сопротивлялись искушениям!")
		g.Score += 1000
		player.Achievements = append(player.Achievements, "Победитель искушений")
		player.AddExperience(200) // Бонус за победу
	} else {
		fmt.Println("💔 Поражение... Искушение оказалось сильнее")
		g.Score -= 500
	}
	
	// Проверяем выполнение квестов
	allCompleted := true
	for _, quest := range player.Quests.Quests {
		if !quest.Completed {
			allCompleted = false
			break
		}
	}
	
	// Забираем награды за квесты
	skillPoints := player.Quests.ClaimRewards()
	if skillPoints > 0 {
		player.SkillTree.EarnSkillPoints(skillPoints)
	}
	
	// Проверяем серию дней
	player.Quests.CheckDayStreak(allCompleted)

	fmt.Printf("Итоговый счет: %d\n", g.Score)
	fmt.Printf("Уровень фокуса: %d%%\n", player.Focus)
	fmt.Printf("Уровень знаний Go: %d\n", player.GoKnowledge)
	fmt.Printf("Сопротивлено искушений: %d/%d\n", g.TotalTemptationsResisted, g.TotalTemptations)
	fmt.Printf("Заработано: %d₽\n", player.Money)
	fmt.Printf("Время игры: %d минут\n", playTime)

	// Уровень игрока
	level := g.Score / 1000
	fmt.Printf("Уровень игрока: %d\n", level)

	if level >= 5 {
		fmt.Println("🏆 ВЫ ДОСТИГЛИ УРОВНЯ GO-MASTER!")
	}
	
	// Показываем расширенную статистику
	player.DisplayStatistics()
}

// ShowQuestsMenu показывает меню квестов
func (g *Game) ShowQuestsMenu(player *Player) {
	player.Quests.DisplayQuests()
}

// SaveGameMenu меню сохранения игры
func (g *Game) SaveGameMenu(player *Player) {
	fmt.Println("\n💾 СОХРАНЕНИЕ ИГРЫ")
	fmt.Println("══════════════════════════════════════")
	fmt.Println("1. Быстрое сохранение")
	fmt.Println("2. Сохранение в файл")
	fmt.Println("3. Загрузить сохранение")
	fmt.Println("4. Показать список сохранений")
	fmt.Println("0. Отмена")
	
	var choice int
	fmt.Print("Выбор: ")
	fmt.Scan(&choice)
	
	switch choice {
	case 1, 2:
		SaveProgress(player, g)
	case 3:
		saves := ListAvailableSaves()
		if len(saves) == 0 {
			fmt.Println("Нет доступных сохранений")
			return
		}
		fmt.Println("Доступные сохранения:")
		for i, save := range saves {
			fmt.Printf("%d. %s\n", i+1, save)
		}
		fmt.Print("Выберите сохранение: ")
		var idx int
		fmt.Scan(&idx)
		if idx > 0 && idx <= len(saves) {
			loadedPlayer := LoadProgress(saves[idx-1])
			if loadedPlayer != nil {
				*player = *loadedPlayer
			}
		}
	case 4:
		listSaves()
	}
}
