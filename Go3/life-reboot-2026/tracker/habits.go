// Go3/life-reboot-2026/tracker/habits.go
package tracker

import (
	"fmt"
	"time"
)

// Habit представляет собой привычку
type Habit struct {
	Name        string
	IsGood      bool
	StartDate   time.Time
	CurrentStreak int // дней подряд
	BestStreak   int // лучшая серия
}

// NewHabit создаёт новую привычку
func NewHabit(name string, isGood bool) *Habit {
	return &Habit{
		Name:      name,
		IsGood:    isGood,
		StartDate: time.Now(),
	}
}

// UpdateStreak обновляет серию выполнения привычки
func (h *Habit) UpdateStreak(completedToday bool) {
	if completedToday {
		h.CurrentStreak++
		if h.CurrentStreak > h.BestStreak {
			h.BestStreak = h.CurrentStreak
		}
	} else {
		h.CurrentStreak = 0
	}
}

// Progress возвращает прогресс в процентах
func (h *Habit) Progress() float64 {
	daysSinceStart := int(time.Since(h.StartDate).Hours()/24) + 1
	return float64(h.CurrentStreak) / float64(daysSinceStart) * 100
}

// Display показывает информацию о привычке
func (h *Habit) Display() {
	icon := "❌"
	if h.IsGood {
		icon = "✅"
	}

	fmt.Printf("%s %s:\n", icon, h.Name)
	fmt.Printf("  Текущая серия: %d дней\n", h.CurrentStreak)
	fmt.Printf("  Лучшая серия: %d дней\n", h.BestStreak)
	fmt.Printf("  Прогресс: %.1f%%\n", h.Progress())
}
