package main

import (
	"fmt"
	"time"
)

// GameManager управляет игровой механикой
type GameManager struct {
	Day         int
	TotalDays   int
	Achievements []string
}

// NewGameManager создает менеджер игры
func NewGameManager() *GameManager {
	return &GameManager{
		Day:       23,
		TotalDays: 365,
		Achievements: []string{
			"Начал путь (День 1)",
			"Поборол первый синтаксический страх (День 7)",
			"Написал первый полезный скрипт (День 15)",
			"Понял горутины (День 42)",
			"Прошел собеседование (День 180)",
			"Устроился на работу (День 270)",
		},
	}
}

// CalculateProgress вычисляет прогресс
func (gm *GameManager) CalculateProgress(goSkill int) float64 {
	// Прогресс = (день + навык) / общее * 100
	return float64(gm.Day+goSkill) / float64(gm.TotalDays) * 100.0
}

// DisplayAchievements показывает достижения
func (gm *GameManager) DisplayAchievements() {
	fmt.Println("\n🏆 ДОСТИЖЕНИЯ:")
	fmt.Println("══════════════════════════════════")
	for i, ach := range gm.Achievements {
		if i < 3 { // Первые 3 уже получены
			fmt.Printf("  ✓ %s\n", ach)
		} else if i == 3 && gm.Day >= 42 {
			fmt.Printf("  ✓ %s\n", ach)
		} else {
			fmt.Printf("  [ ] %s\n", ach)
		}
	}
}

// SimulateInterview симулирует собеседование
func (gm *GameManager) SimulateInterview(goSkill int) {
	fmt.Println("\n🤵 СИМУЛЯЦИЯ СОБЕСЕДОВАНИЯ")
	fmt.Println("══════════════════════════════════")

	if goSkill < 50 {
		fmt.Println("Вопрос: Что такое interface в Go?")
		fmt.Print("Ответ Гоши: ")
		time.Sleep(2 * time.Second)
		fmt.Println("'Эээ... Это как контракт для структур?'")
		fmt.Println("Результат: ❌ Не принят. Нужно учить основы.")
	} else if goSkill < 75 {
		fmt.Println("Вопрос: Как работает garbage collector в Go?")
		fmt.Print("Ответ Гоши: ")
		time.Sleep(1 * time.Second)
		fmt.Println("'Управляет памятью, использует красивую маркировку...'")
		fmt.Println("Результат: ⚠️  На рассмотрении. Есть потенциал.")
	} else {
		fmt.Println("Вопрос: Расскажите про проблемы concurrent map access.")
		fmt.Print("Ответ Гоши: ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("'Нужно использовать sync.Mutex или sync.RWMutex, либо sharding...'")
		fmt.Println("Результат: ✅ Принят! Оффер на руках!")
	}
}

// DailyChallenge ежедневное задание
func (gm *GameManager) DailyChallenge() string {
	challenges := []string{
		"Написать конвейер (pipeline) из 3 горутин",
		"Реализовать кэш с TTL на map + мьютексах",
		"Написать HTTP-сервер с middleware",
		"Протестировать пакет с 90% покрытием",
		"Разобрать open-source проект на GitHub",
	}
	return challenges[gm.Day%len(challenges)]
}
