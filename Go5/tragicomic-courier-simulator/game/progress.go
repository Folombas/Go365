package game

import (
	"fmt"
	"time"
)

type Progress struct {
	XP       int
	NewSkill string
}

func CodingSession(minutes, day int) Progress {
	fmt.Printf("💻 Начинаю %d-минутную сессию Go...\n", minutes)
	
	// Имитация обучения
	totalXP := 0
	skills := []string{
		"Понимание интерфейсов",
		"Embedding структур",
		"Горутины",
		"Каналы",
		"Работа с JSON",
		"Тестирование",
		"Создание модулей",
	}
	
	for i := 0; i < minutes/15; i++ {
		time.Sleep(200 * time.Millisecond)
		xp := 10 + (day * 2) // С каждым днем больше XP
		totalXP += xp
		fmt.Printf("   [%d мин] +%d XP\n", (i+1)*15, xp)
	}
	
	// Новый навык каждый 3-й день
	newSkill := ""
	if day%3 == 0 {
		newSkill = skills[day%len(skills)]
		fmt.Printf("   🎉 Освоил: %s!\n", newSkill)
	}
	
	return Progress{
		XP:       totalXP,
		NewSkill: newSkill,
	}
}

type Level struct {
	Number   int
	Title    string
	XPToNext int
}

func CalculateLevel(xp int) Level {
	levels := []struct {
		minXP int
		title string
	}{
		{0, "Начинающий гофер"},
		{100, "Курьер с условиями"},
		{250, "Понимающий интерфейсы"},
		{500, "Мастер горутин"},
		{1000, "Гуру каналов"},
		{2000, "Архитектор микросервисов"},
	}
	
	currentLevel := 1
	currentTitle := "Новичок"
	nextXP := 100
	
	for i, level := range levels {
		if xp >= level.minXP {
			currentLevel = i + 1
			currentTitle = level.title
			if i+1 < len(levels) {
				nextXP = levels[i+1].minXP - xp
			} else {
				nextXP = 0
			}
		}
	}
	
	return Level{
		Number:   currentLevel,
		Title:    currentTitle,
		XPToNext: nextXP,
	}
}

func DailyChallenge(day int) string {
	challenges := map[int]string{
		1: "Установить Go и написать Hello World",
		2: "Изучить переменные и типы данных",
		3: "Разобраться с функциями",
		4: "Понять структуры и методы",
		5: "Освоить интерфейсы",
		6: "Изучить embedding структур",
		7: "Попрактиковаться с горутинами",
		8: "Работа с каналами",
		9: "Создать свой пакет",
		10: "Написать тесты",
	}
	
	if challenge, exists := challenges[day]; exists {
		return challenge
	}
	return "Повторить пройденное и придумать свой проект!"
}
