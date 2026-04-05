// Package game - главное меню для выбора игр
// Go365 Day 95 - Menu
package game

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"menu/internal/sprite"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

// Button кнопка меню
type Button struct {
	X, Y       int
	Width      int
	Height     int
	Text       string
	Icon       *ebiten.Image
	Hovered    bool
	GamePath   string
}

// Game главное меню
type Game struct {
	spriteStore *sprite.SpriteStore
	buttons     []*Button
	selected    int
}

// NewGame создаёт главное меню
func NewGame() *Game {
	g := &Game{
		spriteStore: sprite.NewSpriteStore(),
		selected:    -1,
	}
	g.createButtons()
	return g
}

// createButtons создаёт кнопки игр
func (g *Game) createButtons() {
	// Определяем пути к играм относительно текущей директории
	games := []struct {
		name     string
		icon     *ebiten.Image
		gamePath string
	}{
		{
			name:     "💣 Bomberman",
			icon:     g.spriteStore.BombermanIcon,
			gamePath: "../bomberman/cmd/game/main.go",
		},
		{
			name:     "♟️ Checkers (Шашки)",
			icon:     g.spriteStore.CheckersIcon,
			gamePath: "../checkers/cmd/game/main.go",
		},
		{
			name:     "🧩 Puzzle (Пятнашки)",
			icon:     g.spriteStore.PuzzleIcon,
			gamePath: "../puzzle/cmd/game/main.go",
		},
	}

	buttonWidth := 300
	buttonHeight := 100
	spacing := 40
	totalWidth := len(games)*buttonWidth + (len(games)-1)*spacing
	startX := (screenWidth - totalWidth) / 2
	startY := screenHeight/2 - 50

	for i, game := range games {
		btn := &Button{
			X:      startX + i*(buttonWidth+spacing),
			Y:      startY,
			Width:  buttonWidth,
			Height: buttonHeight,
			Text:   game.name,
			Icon:   game.icon,
			GamePath: game.gamePath,
		}
		g.buttons = append(g.buttons, btn)
	}
}

// Update обновляет меню
func (g *Game) Update() error {
	// ESC
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Обновление наведения
	mx, my := ebiten.CursorPosition()

	for _, btn := range g.buttons {
		btn.Hovered = mx >= btn.X && mx < btn.X+btn.Width &&
			my >= btn.Y && my < btn.Y+btn.Height
	}

	// Клик
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for i, btn := range g.buttons {
			if btn.Hovered {
				g.selected = i
				g.launchGame(btn.GamePath)
				break
			}
		}
	}

	// Клавиатурная навигация
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.selected > 0 {
		g.selected--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.selected < len(g.buttons)-1 {
		g.selected++
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && g.selected >= 0 && g.selected < len(g.buttons) {
		g.launchGame(g.buttons[g.selected].GamePath)
	}

	return nil
}

// launchGame запускает выбранную игру
func (g *Game) launchGame(gamePath string) {
	// Получаем абсолютный путь
	absPath, err := filepath.Abs(gamePath)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		return
	}

	// Определяем директорию модуля
	moduleDir := filepath.Dir(filepath.Dir(filepath.Dir(absPath)))

	// Запускаем игру через go run
	cmd := exec.Command("go", "run", absPath)
	cmd.Dir = moduleDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Запускаем в фоне
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error launching game: %v\n", err)
		return
	}

	// Завершаем меню
	fmt.Printf("Launching: %s\n", gamePath)
}

// Draw рисует меню
func (g *Game) Draw(screen *ebiten.Image) {
	// Фон
	if g.spriteStore.Background != nil {
		screen.DrawImage(g.spriteStore.Background, nil)
	} else {
		screen.Fill(color.RGBA{30, 40, 70, 255})
	}

	// Заголовок
	g.drawTitle(screen)

	// Кнопки
	g.drawButtons(screen)

	// Подсказка
	g.drawHint(screen)
}

// drawTitle рисует заголовок
func (g *Game) drawTitle(screen *ebiten.Image) {
	title := `
╔═══════════════════════════════════════════════════╗
║                                                   ║
║        🎮 GO365 GAME COLLECTION 🎮               ║
║        День 95 - 5 апреля 2026                    ║
║                                                   ║
║     Выбери игру и наслаждайся!                    ║
║                                                   ║
╚═══════════════════════════════════════════════════╝
`
	ebitenutil.DebugPrintAt(screen, title, screenWidth/2-260, 40)
}

// drawButtons рисует кнопки
func (g *Game) drawButtons(screen *ebiten.Image) {
	for _, btn := range g.buttons {
		// Фон кнопки
		var bg *ebiten.Image
		if btn.Hovered {
			bg = g.spriteStore.ButtonHover
		} else {
			bg = g.spriteStore.ButtonBg
		}

		if bg != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(btn.X), float64(btn.Y))
			screen.DrawImage(bg, opts)
		}

		// Иконка
		if btn.Icon != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(btn.X+10), float64(btn.Y+10))
			screen.DrawImage(btn.Icon, opts)
		}

		// Текст
		ebitenutil.DebugPrintAt(screen, btn.Text, btn.X+100, btn.Y+35)
	}
}

// drawHint рисует подсказку
func (g *Game) drawHint(screen *ebiten.Image) {
	// Панель снизу
	vector.DrawFilledRect(screen, 0, float32(screenHeight-40), screenWidth, 40, color.RGBA{0, 0, 0, 180}, true)

	hint := "← → - Выбор | Enter - Запуск | ESC - Выход | Клик мышью - Выбор игры"
	ebitenutil.DebugPrintAt(screen, hint, 20, screenHeight-30)
}

// Layout размер экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
