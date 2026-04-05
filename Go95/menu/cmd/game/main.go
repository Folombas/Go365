// Menu - Главное меню для выбора игр
// Go365 Challenge - День 95 (5 апреля 2026)
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"menu/internal/game"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("🎮 Go365 Game Collection - Day 95")
}

func main() {
	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
