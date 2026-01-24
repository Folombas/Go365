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
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		Day:   24,
		Score: 0,
		Round: 1,
	}
}

func (g *Game) StartDay(player *Player) {
	fmt.Println("🌅 НАЧАЛО ДНЯ 24")
	fmt.Println("Цель: Не поддаться искушениям и изучить Go")
	fmt.Println("Ваш персонаж:", player.Name)
	fmt.Println("Начальный уровень фокуса:", player.Focus)
	fmt.Println()
}

func (g *Game) CheckTemptation() bool {
	// 30% шанс возникновения искушения
	return rand.Intn(100) < 30
}

func (g *Game) CheckMotivation() bool {
	// 25% шанс получить мотивацию
	return rand.Intn(100) < 25
}

func (g *Game) EndDay(player *Player, wonBattle bool) {
	fmt.Println("\n📊 ИТОГИ ДНЯ 24:")
	fmt.Println("==================")

	g.Score = player.CalculateScore()

	if wonBattle {
		fmt.Println("🎉 ПОБЕДА! Вы успешно сопротивлялись искушениям!")
		g.Score += 1000
		player.Achievements = append(player.Achievements, "Победитель искушений")
	} else {
		fmt.Println("💔 Поражение... Искушение оказалось сильнее")
		g.Score -= 500
	}

	fmt.Printf("Итоговый счет: %d\n", g.Score)
	fmt.Printf("Уровень фокуса: %d%%\n", player.Focus)
	fmt.Printf("Уровень знаний Go: %d\n", player.GoKnowledge)
	fmt.Printf("Сопротивлено искушений: %d\n", g.TotalTemptationsResisted)
	fmt.Printf("Заработано: %d₽\n", player.Money)

	// Уровень игрока
	level := g.Score / 1000
	fmt.Printf("Уровень игрока: %d\n", level)

	if level >= 5 {
		fmt.Println("🏆 ВЫ ДОСТИГЛИ УРОВНЯ GO-MASTER!")
	}
}
