package main

import (
	"fmt"
	"os"
	"encoding/json"
)

func SaveProgress(player *Player) {
	// Создаем структуру для сохранения
	progress := map[string]interface{}{
		"player": map[string]interface{}{
			"name":          player.Name,
			"focus":         player.Focus,
			"go_knowledge":  player.GoKnowledge,
			"willpower":     player.Willpower,
			"money":         player.Money,
			"dopamine":      player.Dopamine,
			"achievements":  player.Achievements,
			"temptations_resisted": len(player.Temptations),
		},
		"date": "2026-01-24",
		"day":  24,
	}

	// Сохраняем в JSON файл
	file, err := os.Create("progress_day24.json")
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(progress)

	fmt.Println("\n💾 Прогресс сохранен в progress_day24.json")
	fmt.Println("📈 Вы можете продолжить завтра с этого уровня!")
}

func LoadProgress() *Player {
	// Загрузка прогресса (упрощенная версия)
	// В реальной игре здесь бы было чтение из файла
	return NewPlayer("Гоша")
}
