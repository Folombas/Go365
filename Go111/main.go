package main

import (
	"fmt"
	"time"
)

func main() {
	// Целевая дата: 21 апреля 2026 года
	targetDate := time.Date(2026, time.April, 21, 0, 0, 0, 0, time.UTC)
	now := time.Now().UTC()

	fmt.Println("🎉 День Go111 — 21 апреля 2026 года! 🎉")
	fmt.Println("======================================")

	// Если сегодня 21 апреля 2026
	if now.Year() == targetDate.Year() && now.YearDay() == targetDate.YearDay() {
		fmt.Println("✅ Сегодня официальный День Go111!")
		fmt.Println("Поздравляем всех гophers и разработчиков!")
	} else if now.Before(targetDate) {
		daysUntil := int(targetDate.Sub(now).Hours() / 24)
		fmt.Printf("📅 До Дня Go111 осталось %d дней.\n", daysUntil)
	} else {
		fmt.Println("🎈 День Go111 уже прошёл, но дух Go живёт вечно!")
	}

	fmt.Println("\n✨ Пусть ваш код компилируется с первого раза, а горутины не утекают! ✨")
}