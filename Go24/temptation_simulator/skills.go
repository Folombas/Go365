package main

import (
	"fmt"
	"time"
)

// Skill представляет навык в дереве развития
type Skill struct {
	ID          string
	Name        string
	Description string
	Level       int      // Уровень навыка (0-5)
	MaxLevel    int      // Максимальный уровень
	CostPerLevel int     // Стоимость улучшения за уровень
	BonusType   string   // Тип бонуса: focus, willpower, knowledge, money, dopamine
	BonusValue  int      // Значение бонуса за уровень
	Unlocked    bool     // Разблокирован ли навык
	Prerequisites []string // Требуемые навыки для разблокировки
}

// SkillTree представляет дерево навыков игрока
type SkillTree struct {
	Skills       map[string]*Skill
	SkillPoints  int  // Доступные очки навыков
	TotalPoints  int  // Всего заработано очков
}

// DailyQuest представляет ежедневный квест
type DailyQuest struct {
	ID          string
	Title       string
	Description string
	Goal        int  // Цель (например, изучить 50 очков Go)
	Progress    int  // Текущий прогресс
	Reward      int  // Награда (очки навыков)
	Completed   bool // Выполнен ли квест
	Deadline    string // Дедлайн (дата)
}

// QuestSystem управляет ежедневными квестами
type QuestSystem struct {
	Quests      []*DailyQuest
	DayStreak   int  // Серия успешных дней
	TotalCompleted int // Всего выполнено квестов
}

// NewSkillTree создает новое дерево навыков
func NewSkillTree() *SkillTree {
	tree := &SkillTree{
		Skills:      make(map[string]*Skill),
		SkillPoints: 0,
		TotalPoints: 0,
	}

	// Инициализируем дерево навыков
	tree.initSkills()

	return tree
}

// initSkills инициализирует навыки для Go-разработчика
func (st *SkillTree) initSkills() {
	// Базовые навыки Go
	st.Skills["go_basics"] = &Skill{
		ID:          "go_basics",
		Name:        "Основы Go",
		Description: "Синтаксис, типы данных, функции",
		Level:       1,
		MaxLevel:    5,
		CostPerLevel: 1,
		BonusType:   "knowledge",
		BonusValue:  5,
		Unlocked:    true,
	}

	st.Skills["concurrency"] = &Skill{
		ID:          "concurrency",
		Name:        "Конкурентность",
		Description: "Горутины, каналы, sync package",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "knowledge",
		BonusValue:  8,
		Unlocked:    false,
		Prerequisites: []string{"go_basics"},
	}

	st.Skills["interfaces"] = &Skill{
		ID:          "interfaces",
		Name:        "Интерфейсы",
		Description: "Интерфейсы, полиморфизм, паттерны",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "knowledge",
		BonusValue:  7,
		Unlocked:    false,
		Prerequisites: []string{"go_basics"},
	}

	// Навыки фокуса
	st.Skills["focus_master"] = &Skill{
		ID:          "focus_master",
		Name:        "Мастер Фокуса",
		Description: "Умение концентрироваться на задачах",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 1,
		BonusType:   "focus",
		BonusValue:  5,
		Unlocked:    true,
	}

	st.Skills["meditation"] = &Skill{
		ID:          "meditation",
		Name:        "Медитация",
		Description: "Восстановление фокуса и энергии",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "dopamine",
		BonusValue:  10,
		Unlocked:    false,
		Prerequisites: []string{"focus_master"},
	}

	// Навыки силы воли
	st.Skills["willpower"] = &Skill{
		ID:          "willpower",
		Name:        "Сила Воли",
		Description: "Сопротивление искушениям",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "willpower",
		BonusValue:  8,
		Unlocked:    true,
	}

	st.Skills["discipline"] = &Skill{
		ID:          "discipline",
		Name:        "Дисциплина",
		Description: "Ежедневное следование плану",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 3,
		BonusType:   "willpower",
		BonusValue:  10,
		Unlocked:    false,
		Prerequisites: []string{"willpower"},
	}

	// Финансовые навыки
	st.Skills["money_management"] = &Skill{
		ID:          "money_management",
		Name:        "Управление деньгами",
		Description: "Экономия и инвестиции в обучение",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "money",
		BonusValue:  50,
		Unlocked:    false,
		Prerequisites: []string{"willpower"},
	}

	// Продвинутые навыки
	st.Skills["web_frameworks"] = &Skill{
		ID:          "web_frameworks",
		Name:        "Web Фреймворки",
		Description: "Gin, Echo, Fiber - создание API",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 3,
		BonusType:   "knowledge",
		BonusValue:  10,
		Unlocked:    false,
		Prerequisites: []string{"concurrency", "interfaces"},
	}

	st.Skills["database"] = &Skill{
		ID:          "database",
		Name:        "Базы данных",
		Description: "PostgreSQL, MongoDB, Redis",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 3,
		BonusType:   "knowledge",
		BonusValue:  10,
		Unlocked:    false,
		Prerequisites: []string{"concurrency"},
	}

	st.Skills["microservices"] = &Skill{
		ID:          "microservices",
		Name:        "Микросервисы",
		Description: "gRPC, Docker, Kubernetes",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 4,
		BonusType:   "knowledge",
		BonusValue:  12,
		Unlocked:    false,
		Prerequisites: []string{"web_frameworks", "database"},
	}

	st.Skills["clean_architecture"] = &Skill{
		ID:          "clean_architecture",
		Name:        "Чистая архитектура",
		Description: "SOLID, DDD, паттерны проектирования",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 4,
		BonusType:   "knowledge",
		BonusValue:  12,
		Unlocked:    false,
		Prerequisites: []string{"interfaces", "web_frameworks"},
	}

	st.Skills["anti_procrastination"] = &Skill{
		ID:          "anti_procrastination",
		Name:        "Борьба с прокрастинацией",
		Description: "Техники Pomodoro, тайм-менеджмент",
		Level:       0,
		MaxLevel:    5,
		CostPerLevel: 2,
		BonusType:   "focus",
		BonusValue:  8,
		Unlocked:    false,
		Prerequisites: []string{"focus_master"},
	}

	st.Skills["cold_shower"] = &Skill{
		ID:          "cold_shower",
		Name:        "Холодный душ",
		Description: "Закаливание и утренние ритуалы",
		Level:       0,
		MaxLevel:    3,
		CostPerLevel: 1,
		BonusType:   "willpower",
		BonusValue:  5,
		Unlocked:    true,
	}
}

// Display отображает дерево навыков
func (st *SkillTree) Display() {
	fmt.Println("\n🌳 ДЕРЕВО НАВЫКОВ")
	fmt.Println("══════════════════════════════════════")
	fmt.Printf("Очки навыков: %d (всего заработано: %d)\n", st.SkillPoints, st.TotalPoints)
	fmt.Println()

	// Группируем навыки по категориям
	categories := map[string][]string{
		"📚 GO-НАВЫКИ":        {"go_basics", "concurrency", "interfaces", "web_frameworks", "database", "microservices", "clean_architecture"},
		"🎯 ФОКУС":            {"focus_master", "meditation", "anti_procrastination"},
		"💪 СИЛА ВОЛИ":        {"willpower", "discipline", "cold_shower"},
		"💰 ФИНАНСЫ":          {"money_management"},
	}

	for categoryName, skillIDs := range categories {
		fmt.Println(categoryName)
		fmt.Println("──────────────────────────────────────")
		for _, skillID := range skillIDs {
			skill := st.Skills[skillID]
			st.displaySkill(skill)
		}
		fmt.Println()
	}
}

// displaySkill отображает один навык
func (st *SkillTree) displaySkill(skill *Skill) {
	status := "🔒"
	if skill.Unlocked {
		status = "✅"
	}

	levelBar := ""
	for i := 0; i < skill.MaxLevel; i++ {
		if i < skill.Level {
			levelBar += "█"
		} else {
			levelBar += "░"
		}
	}

	fmt.Printf("  %s %-20s [%s] Ур.%d/%d\n",
		status, skill.Name, levelBar, skill.Level, skill.MaxLevel)
	fmt.Printf("     %s\n", skill.Description)

	if skill.Unlocked && skill.Level < skill.MaxLevel {
		fmt.Printf("     💰 Улучшение: %d очк. | +%d %s\n",
			skill.CostPerLevel, skill.BonusValue, translateBonusType(skill.BonusType))
	}
	fmt.Println()
}

// UpgradeSkill улучшает навык
func (st *SkillTree) UpgradeSkill(skillID string) bool {
	skill, exists := st.Skills[skillID]
	if !exists {
		fmt.Println("❌ Навык не найден")
		return false
	}

	if !skill.Unlocked {
		fmt.Println("❌ Навык заблокит. Изучите требуемые навыки сначала.")
		return false
	}

	if skill.Level >= skill.MaxLevel {
		fmt.Println("⚠️  Навык уже максимального уровня")
		return false
	}

	cost := skill.CostPerLevel
	if st.SkillPoints < cost {
		fmt.Printf("❌ Недостаточно очков навыков (нужно %d, есть %d)\n", cost, st.SkillPoints)
		return false
	}

	st.SkillPoints -= cost
	skill.Level++

	// Применяем бонус
	applySkillBonus(skill)

	fmt.Printf("🎉 Навык '%s' улучшен до уровня %d!\n", skill.Name, skill.Level)
	fmt.Printf("   +%d к %s\n", skill.BonusValue, translateBonusType(skill.BonusType))

	// Проверяем разблокировку следующих навыков
	st.checkUnlocks(skillID)

	return true
}

// checkUnlocks проверяет и разблокирует новые навыки
func (st *SkillTree) checkUnlocks(skillID string) {
	for _, skill := range st.Skills {
		if skill.Unlocked {
			continue
		}

		// Проверяем все prerequisites
		allMet := true
		for _, prereq := range skill.Prerequisites {
			prereqSkill := st.Skills[prereq]
			if prereqSkill == nil || prereqSkill.Level == 0 {
				allMet = false
				break
			}
		}

		if allMet {
			skill.Unlocked = true
			fmt.Printf("\n🔓 РАЗБЛОКИРОВАН НОВЫЙ НАВЫК: %s\n", skill.Name)
		}
	}
}

// applySkillBonus применяет бонус от навыка
func applySkillBonus(skill *Skill) {
	// Бонусы применяются через Player.ApplySkillBonus()
	// Эта функция будет вызвана из основного цикла игры
}

// EarnSkillPoints начисляет очки навыков
func (st *SkillTree) EarnSkillPoints(points int) {
	st.SkillPoints += points
	st.TotalPoints += points
	fmt.Printf("\n✨ Получено %d очков навыков! (всего: %d)\n", points, st.SkillPoints)
}

// GetSkill возвращает навык по ID
func (st *SkillTree) GetSkill(skillID string) *Skill {
	return st.Skills[skillID]
}

// GetTotalBonus возвращает суммарный бонус по типу
func (st *SkillTree) GetTotalBonus(bonusType string) int {
	total := 0
	for _, skill := range st.Skills {
		if skill.BonusType == bonusType {
			total += skill.Level * skill.BonusValue
		}
	}
	return total
}

// NewQuestSystem создает новую систему квестов
func NewQuestSystem() *QuestSystem {
	return &QuestSystem{
		Quests:      make([]*DailyQuest, 0),
		DayStreak:   0,
		TotalCompleted: 0,
	}
}

// GenerateDailyQuests генерирует ежедневные квесты
func (qs *QuestSystem) GenerateDailyQuests() {
	today := time.Now().Format("2006-01-02")

	qs.Quests = []*DailyQuest{
		{
			ID:          "study_go_30min",
			Title:       "30 минут Go",
			Description: "Изучи новую тему Go в течение 30 минут",
			Goal:        30,
			Progress:    0,
			Reward:      2,
			Completed:   false,
			Deadline:    today,
		},
		{
			ID:          "resist_temptation",
			Title:       "Борец с искушениями",
			Description: "Сопротивляйся 3 искушениям сегодня",
			Goal:        3,
			Progress:    0,
			Reward:      3,
			Completed:   false,
			Deadline:    today,
		},
		{
			ID:          "code_practice",
			Title:       "Практика кода",
			Description: "Напиши 50+ строк кода на Go",
			Goal:        50,
			Progress:    0,
			Reward:      3,
			Completed:   false,
			Deadline:    today,
		},
		{
			ID:          "morning_routine",
			Title:       "Утренний ритуал",
			Description: "Выполни утренний ритуал (зарядка, душ, завтрак)",
			Goal:        1,
			Progress:    0,
			Reward:      1,
			Completed:   false,
			Deadline:    today,
		},
		{
			ID:          "no_social_media",
			Title:       "Цифровой детокс",
			Description: "Не заходи в соцсети 4 часа",
			Goal:        4,
			Progress:    0,
			Reward:      2,
			Completed:   false,
			Deadline:    today,
		},
	}

	fmt.Println("\n📋 ЕЖЕДНЕВНЫЕ КВЕСТЫ:")
	fmt.Println("══════════════════════════════════════")
	for _, quest := range qs.Quests {
		fmt.Printf("🎯 %s: %s\n", quest.Title, quest.Description)
		fmt.Printf("   Награда: %d очков навыков\n", quest.Reward)
	}
}

// UpdateQuestProgress обновляет прогресс квеста
func (qs *QuestSystem) UpdateQuestProgress(questID string, progress int) {
	for _, quest := range qs.Quests {
		if quest.ID == questID && !quest.Completed {
			quest.Progress += progress
			if quest.Progress >= quest.Goal {
				quest.Completed = true
				qs.TotalCompleted++
				fmt.Printf("\n✅ КВЕСТ ВЫПОЛНЕН: %s\n", quest.Title)
				fmt.Printf("   Награда: +%d очков навыков\n", quest.Reward)
			} else {
				fmt.Printf("\n📊 Прогресс: %s %d/%d\n", quest.Title, quest.Progress, quest.Goal)
			}
			return
		}
	}
}

// DisplayQuests отображает текущие квесты
func (qs *QuestSystem) DisplayQuests() {
	fmt.Println("\n📋 ТЕКУЩИЕ КВЕСТЫ:")
	fmt.Println("══════════════════════════════════════")

	if len(qs.Quests) == 0 {
		fmt.Println("Нет активных квестов")
		return
	}

	for _, quest := range qs.Quests {
		status := "⏳"
		if quest.Completed {
			status = "✅"
		}

		fmt.Printf("%s %-25s %d/%d\n", status, quest.Title, quest.Progress, quest.Goal)
	}

	fmt.Printf("\n🔥 Серия успешных дней: %d\n", qs.DayStreak)
	fmt.Printf("🏆 Всего выполнено квестов: %d\n", qs.TotalCompleted)
}

// ClaimRewards забирает награды за выполненные квесты
func (qs *QuestSystem) ClaimRewards() int {
	totalReward := 0
	for _, quest := range qs.Quests {
		if quest.Completed && quest.Reward > 0 {
			totalReward += quest.Reward
			quest.Reward = 0 // Сбрасываем награду после получения
		}
	}

	if totalReward > 0 {
		fmt.Printf("\n💰 Получено %d очков навыков за квесты!\n", totalReward)
	}

	return totalReward
}

// CheckDayStreak проверяет и обновляет серию дней
func (qs *QuestSystem) CheckDayStreak(allCompleted bool) {
	if allCompleted {
		qs.DayStreak++
		fmt.Printf("\n🔥 СЕРИЯ ПРОДОЛЖЕНА! %d дней подряд\n", qs.DayStreak)

		// Бонус за серию
		if qs.DayStreak%7 == 0 {
			fmt.Printf("🎁 БОНУС ЗА НЕДЕЛЮ: +%d очков навыков!\n", qs.DayStreak)
		}
	} else {
		if qs.DayStreak > 0 {
			fmt.Printf("\n💔 Серия прервана... Было %d дней\n", qs.DayStreak)
		}
		qs.DayStreak = 0
	}
}

// translateBonusType переводит тип бонуса на русский
func translateBonusType(bonusType string) string {
	switch bonusType {
	case "focus":
		return "Фокус"
	case "willpower":
		return "Сила воли"
	case "knowledge":
		return "Знание Go"
	case "money":
		return "Деньги"
	case "dopamine":
		return "Дофамин"
	default:
		return bonusType
	}
}
