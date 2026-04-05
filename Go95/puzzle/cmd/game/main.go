// Puzzle - Классическая головоломка "Пятнашки"
// Go365 Challenge - День 95 (5 апреля 2026)
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"puzzle/internal/game"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("🧩 Sliding Puzzle (Пятнашки) - Go365 Day 95")
}

func main() {
	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
