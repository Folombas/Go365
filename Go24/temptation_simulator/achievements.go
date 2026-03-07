package main

import (
	"fmt"
	"os"
	"encoding/json"
	"time"
	"path/filepath"
)

// SaveData представляет все данные для сохранения
type SaveData struct {
	Player struct {
		Name         string   `json:"name"`
		Focus        int      `json:"focus"`
		GoKnowledge  int      `json:"go_knowledge"`
		Willpower    int      `json:"willpower"`
		Money        int      `json:"money"`
		Dopamine     int      `json:"dopamine"`
		Level        int      `json:"level"`
		Experience   int      `json:"experience"`
		PlayTime     int      `json:"play_time"`
		DaysPlayed   int      `json:"days_played"`
		Temptations  []string `json:"temptations_resisted"`
		Achievements []string `json:"achievements"`
	} `json:"player"`
	SkillTree struct {
		SkillPoints int            `json:"skill_points"`
		TotalPoints int            `json:"total_points"`
		Skills      map[string]int `json:"skills"` // skill_id -> level
	} `json:"skill_tree"`
	Quests struct {
		DayStreak    int      `json:"day_streak"`
		TotalCompleted int    `json:"total_completed"`
		CompletedQuests []string `json:"completed_quests"`
	} `json:"quests"`
	Game struct {
		Date       string `json:"date"`
		Day        int    `json:"day"`
		Score      int    `json:"score"`
		TotalGames int    `json:"total_games"`
	} `json:"game"`
}

// SaveProgress сохраняет прогресс в JSON файл
func SaveProgress(player *Player, game *Game) {
	// Создаем структуру для сохранения
	saveData := SaveData{}
	
	// Данные игрока
	saveData.Player.Name = player.Name
	saveData.Player.Focus = player.Focus
	saveData.Player.GoKnowledge = player.GoKnowledge
	saveData.Player.Willpower = player.Willpower
	saveData.Player.Money = player.Money
	saveData.Player.Dopamine = player.Dopamine
	saveData.Player.Level = player.Level
	saveData.Player.Experience = player.Experience
	saveData.Player.PlayTime = player.PlayTime
	saveData.Player.DaysPlayed = player.DaysPlayed
	saveData.Player.Temptations = player.Temptations
	saveData.Player.Achievements = player.Achievements
	
	// Данные дерева навыков
	saveData.SkillTree.SkillPoints = player.SkillTree.SkillPoints
	saveData.SkillTree.TotalPoints = player.SkillTree.TotalPoints
	saveData.SkillTree.Skills = make(map[string]int)
	for id, skill := range player.SkillTree.Skills {
		saveData.SkillTree.Skills[id] = skill.Level
	}
	
	// Данные квестов
	saveData.Quests.DayStreak = player.Quests.DayStreak
	saveData.Quests.TotalCompleted = player.Quests.TotalCompleted
	
	// Данные игры
	saveData.Game.Date = time.Now().Format("2006-01-02 15:04:05")
	saveData.Game.Day = game.Day
	saveData.Game.Score = player.CalculateScore()
	saveData.Game.TotalGames = game.Day
	
	// Создаем директорию для сохранений
	saveDir := "saves"
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		fmt.Println("❌ Ошибка создания директории сохранений:", err)
		return
	}
	
	// Имя файла с датой
	filename := fmt.Sprintf("save_day%d_%s.json", game.Day, time.Now().Format("20060102_150405"))
	filepath := filepath.Join(saveDir, filename)
	
	// Сохраняем в JSON файл
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("❌ Ошибка сохранения:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(saveData); err != nil {
		fmt.Println("❌ Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println("\n💾 Прогресс сохранён в", filepath)
	fmt.Println("📈 Вы можете продолжить с этого места в любой момент!")
	
	// Показываем список всех сохранений
	listSaves()
}

// LoadProgress загружает прогресс из файла сохранения
func LoadProgress(filename string) *Player {
	filepath := filepath.Join("saves", filename)
	
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("❌ Ошибка загрузки сохранения:", err)
		return nil
	}
	
	var saveData SaveData
	if err := json.Unmarshal(data, &saveData); err != nil {
		fmt.Println("❌ Ошибка декодирования JSON:", err)
		return nil
	}
	
	// Восстанавливаем игрока
	player := &Player{
		Name:         saveData.Player.Name,
		Focus:        saveData.Player.Focus,
		GoKnowledge:  saveData.Player.GoKnowledge,
		Willpower:    saveData.Player.Willpower,
		Money:        saveData.Player.Money,
		Dopamine:     saveData.Player.Dopamine,
		Level:        saveData.Player.Level,
		Experience:   saveData.Player.Experience,
		PlayTime:     saveData.Player.PlayTime,
		DaysPlayed:   saveData.Player.DaysPlayed,
		Temptations:  saveData.Player.Temptations,
		Achievements: saveData.Player.Achievements,
	}
	
	// Восстанавливаем дерево навыков
	player.SkillTree = NewSkillTree()
	player.SkillTree.SkillPoints = saveData.SkillTree.SkillPoints
	player.SkillTree.TotalPoints = saveData.SkillTree.TotalPoints
	
	// Восстанавливаем уровни навыков
	for id, level := range saveData.SkillTree.Skills {
		if skill, exists := player.SkillTree.Skills[id]; exists {
			skill.Level = level
			// Разблокируем навык если уровень > 0
			if level > 0 {
				skill.Unlocked = true
			}
		}
	}
	
	// Восстанавливаем квесты
	player.Quests = NewQuestSystem()
	player.Quests.DayStreak = saveData.Quests.DayStreak
	player.Quests.TotalCompleted = saveData.Quests.TotalCompleted
	
	// Применяем бонусы
	player.ApplySkillBonuses()
	
	fmt.Println("\n💾 Загружено сохранение:", filename)
	fmt.Printf("   Игрок: %s | Уровень: %d | День: %d\n", 
		player.Name, player.Level, saveData.Game.Day)
	
	return player
}

// listSaves показывает список всех сохранений
func listSaves() {
	saveDir := "saves"
	files, err := os.ReadDir(saveDir)
	if err != nil {
		return
	}
	
	if len(files) == 0 {
		return
	}
	
	fmt.Println("\n📂 СОХРАНЕНИЯ:")
	fmt.Println("──────────────────────────────────────")
	for i, file := range files {
		if !file.IsDir() && len(file.Name()) > 5 && file.Name()[len(file.Name())-5:] == ".json" {
			fmt.Printf("  %d. %s\n", i+1, file.Name())
		}
	}
}

// ListAvailableSaves возвращает список доступных сохранений
func ListAvailableSaves() []string {
	saveDir := "saves"
	files, err := os.ReadDir(saveDir)
	if err != nil {
		return nil
	}
	
	var saves []string
	for _, file := range files {
		if !file.IsDir() && len(file.Name()) > 5 && file.Name()[len(file.Name())-5:] == ".json" {
			saves = append(saves, file.Name())
		}
	}
	
	return saves
}

// DeleteSave удаляет сохранение
func DeleteSave(filename string) error {
	filepath := filepath.Join("saves", filename)
	return os.Remove(filepath)
}

// GetLatestSave возвращает имя последнего сохранения
func GetLatestSave() string {
	saves := ListAvailableSaves()
	if len(saves) == 0 {
		return ""
	}
	return saves[len(saves)-1] // Последний файл (сортируется по имени)
}

// AutoSave автоматически сохраняет прогресс
func AutoSave(player *Player, game *Game) {
	// Автосохранение каждые 10 минут игрового времени
	if player.PlayTime > 0 && player.PlayTime%10 == 0 {
		fmt.Println("\n💾 Автосохранение...")
		SaveProgress(player, game)
	}
}
