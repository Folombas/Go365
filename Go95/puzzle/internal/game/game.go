// Package game - основная игровая логика Puzzle
// Go365 Day 95 - Puzzle
package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"puzzle/internal/puzzle"
	"puzzle/internal/sprite"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	tileSize     = 100
	gridSize     = 400 // 4 * 100
	gridOffsetX  = (screenWidth - gridSize) / 2
	gridOffsetY  = (screenHeight - gridSize) / 2 + 20
)

// GameState состояние игры
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateSolved
)

// Game основная игра
type Game struct {
	state       GameState
	puzzle      *puzzle.Puzzle
	spriteStore *sprite.SpriteStore
	rng         *rand.Rand
	difficulty  int // количество перемешиваний
	hoverRow    int
	hoverCol    int
	hovering    bool
}

// NewGame создаёт новую игру
func NewGame() *Game {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := &Game{
		state:       StateMenu,
		difficulty:  100,
		hoverRow:    -1,
		hoverCol:    -1,
		rng:         rng,
		spriteStore: sprite.NewSpriteStore(),
	}
	return g
}

// Reset сбрасывает игру
func (g *Game) Reset() {
	g.puzzle = puzzle.NewPuzzle()
	g.puzzle.Shuffle(g.difficulty, g.rng)
	g.state = StatePlaying
}

// Update обновляет игру
func (g *Game) Update() error {
	// ESC
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		switch g.state {
		case StateMenu, StateSolved:
			return ebiten.Termination
		case StatePlaying:
			g.state = StateMenu
		}
	}

	// Меню
	if g.state == StateMenu && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
	}

	// Решено
	if g.state == StateSolved && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
	}

	// Игра
	if g.state == StatePlaying {
		g.handleInput()
		g.updateHover()
	}

	return nil
}

// handleInput обрабатывает ввод
func (g *Game) handleInput() {
	// Клик мышью
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		// Преобразование в координаты сетки
		col := (mx - gridOffsetX) / tileSize
		row := (my - gridOffsetY) / tileSize

		if col < 0 || col >= 4 || row < 0 || row >= 4 {
			return
		}

		// Попытка хода
		if g.puzzle.Move(row, col) {
			if g.puzzle.Solved {
				g.state = StateSolved
			}
		}
	}

	// R - перемешать заново
	if ebiten.IsKeyPressed(ebiten.KeyR) && g.state == StatePlaying {
		g.puzzle.Shuffle(g.difficulty, g.rng)
	}
}

// updateHover обновляет состояние наведения
func (g *Game) updateHover() {
	mx, my := ebiten.CursorPosition()

	col := (mx - gridOffsetX) / tileSize
	row := (my - gridOffsetY) / tileSize

	if col >= 0 && col < 4 && row >= 0 && row < 4 {
		g.hoverRow = row
		g.hoverCol = col
		g.hovering = true
	} else {
		g.hovering = false
	}
}

// Draw рисует игру
func (g *Game) Draw(screen *ebiten.Image) {
	// Фон
	screen.Fill(color.RGBA{40, 50, 80, 255})

	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying, StateSolved:
		g.drawGrid(screen)
		g.drawTiles(screen)
		g.drawHUD(screen)
		if g.state == StateSolved {
			g.drawSolved(screen)
		}
	}
}

// drawGrid рисует сетку
func (g *Game) drawGrid(screen *ebiten.Image) {
	// Фон сетки
	vector.DrawFilledRect(screen,
		float32(gridOffsetX-5), float32(gridOffsetY-5),
		float32(gridSize+10), float32(gridSize+10),
		color.RGBA{80, 70, 60, 255}, true)

	// Внутренний фон
	vector.DrawFilledRect(screen,
		float32(gridOffsetX), float32(gridOffsetY),
		float32(gridSize), float32(gridSize),
		color.RGBA{50, 50, 60, 255}, true)

	// Линии сетки
	gridLineColor := color.RGBA{100, 100, 120, 255}
	for i := 0; i <= 4; i++ {
		// Вертикальные
		vector.DrawFilledRect(screen,
			float32(gridOffsetX+i*tileSize), float32(gridOffsetY),
			2, float32(gridSize),
			gridLineColor, true)
		// Горизонтальные
		vector.DrawFilledRect(screen,
			float32(gridOffsetX), float32(gridOffsetY+i*tileSize),
			float32(gridSize), 2,
			gridLineColor, true)
	}
}

// drawTiles рисует плитки
func (g *Game) drawTiles(screen *ebiten.Image) {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			tile := g.puzzle.GetTile(row, col)

			x := gridOffsetX + col*tileSize + 2
			y := gridOffsetY + row*tileSize + 2

			if tile == 0 {
				// Пустая клетка
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(g.spriteStore.EmptyTile, opts)
			} else {
				// Плитка с числом
				spriteImg := g.spriteStore.Tiles[tile]
				if spriteImg != nil {
					opts := &ebiten.DrawImageOptions{}

					// Подсветка при наведении
					if g.hovering && g.hoverRow == row && g.hoverCol == col && g.puzzle.CanMove(row, col) {
						opts.ColorScale.Scale(1.1, 1.1, 1.1, 1.0)
					}

					opts.GeoM.Translate(float64(x), float64(y))
					screen.DrawImage(spriteImg, opts)
				}
			}
		}
	}
}

// drawHUD рисует интерфейс
func (g *Game) drawHUD(screen *ebiten.Image) {
	// Панель сверху
	vector.DrawFilledRect(screen, 0, 0, screenWidth, 50, color.RGBA{0, 0, 0, 180}, true)

	// Ходы
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("🔢 Ходы: %d", g.puzzle.Moves), 20, 15)

	// Сложность
	difficultyText := "Легко"
	if g.difficulty > 50 {
		difficultyText = "Средне"
	}
	if g.difficulty > 150 {
		difficultyText = "Сложно"
	}
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("📊 Сложность: %s", difficultyText), 300, 15)

	// Подсказка
	ebitenutil.DebugPrintAt(screen, "Кликни на плитку рядом с пустой клеткой | R - перемешать | ESC - меню", 20, screenHeight-30)
}

// drawMenu рисует меню
func (g *Game) drawMenu(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 200})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	menuText := `
╔═══════════════════════════════════════════╗
║         🧩 SLIDING PUZZLE 🧩              ║
║         (Пятнашки)                        ║
║                                           ║
║           [SPACE] - Начать игру           ║
║           [ESC] - Выход                   ║
║                                           ║
║  🎮 Управление:                           ║
║     Клик мышью - Сдвинуть плитку          ║
║     R - Перемешать заново                 ║
║                                           ║
║  🎯 Цель: Расставь числа по порядку!      ║
║     1  2  3  4                            ║
║     5  6  7  8                            ║
║     9  10 11 12                           ║
║     13 14 15  ⬜                          ║
║                                           ║
║  💡 Пазл всегда решаем!                   ║
╚═══════════════════════════════════════════╝
`
	ebitenutil.DebugPrintAt(screen, menuText, screenWidth/2-220, 80)

	// Выбор сложности
	difficultyText := `
╔═══════════════════════════════════════════╗
║  Выбор сложности (в коде):                ║
║  50 ходов  - Легко                        ║
║  100 ходов - Средне                       ║
║  200 ходов - Сложно                       ║
╚═══════════════════════════════════════════╝
`
	ebitenutil.DebugPrintAt(screen, difficultyText, screenWidth/2-220, 420)
}

// drawSolved рисует экран решения
func (g *Game) drawSolved(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 100, 0, 150})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	solvedText := fmt.Sprintf(`
╔═══════════════════════════════════════════╗
║     🎉 ПАЗЛ СОБРАН! 🎉                    ║
╠═══════════════════════════════════════════╣
║                                           ║
║     Ходов сделано: %6d                     ║
║                                           ║
║     [SPACE] - Новая игра                  ║
║     [ESC] - Выход                         ║
╚═══════════════════════════════════════════╝
`, g.puzzle.Moves)

	ebitenutil.DebugPrintAt(screen, solvedText, screenWidth/2-220, screenHeight/2-100)
}

// Layout размер экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
