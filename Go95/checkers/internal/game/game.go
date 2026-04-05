// Package game - основная игровая логика Checkers
// Go365 Day 95 - Checkers
package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"checkers/internal/ai"
	"checkers/internal/board"
	"checkers/internal/sprite"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	boardSize    = 640 // 8 * 80
	tileSize     = 80
	boardOffsetX = (screenWidth - boardSize) / 2
	boardOffsetY = (screenHeight - boardSize) / 2 + 20
)

// GameState состояние игры
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePlayerTurn
	StateAITurn
	StateGameOver
)

// Game основная игра
type Game struct {
	state       GameState
	board       *board.Board
	spriteStore *sprite.SpriteStore
	ai          *ai.AI
	selectedRow int
	selectedCol int
	validMoves  []board.Move
	highlighted []board.Move
	isWhiteTurn bool
	message     string
	messageTimer int
}

// NewGame создаёт новую игру
func NewGame() *Game {
	g := &Game{
		state:         StateMenu,
		selectedRow:   -1,
		selectedCol:   -1,
		isWhiteTurn:   true,
		spriteStore:   sprite.NewSpriteStore(),
	}
	g.ai = ai.NewAI(false, 3) // AI играет за чёрных
	return g
}

// Reset сбрасывает игру
func (g *Game) Reset() {
	g.board = board.NewBoard()
	g.state = StatePlayerTurn
	g.selectedRow = -1
	g.selectedCol = -1
	g.validMoves = nil
	g.highlighted = nil
	g.isWhiteTurn = true
	g.message = ""
	g.messageTimer = 0
}

// Update обновляет игру
func (g *Game) Update() error {
	// ESC
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		switch g.state {
		case StateMenu, StateGameOver:
			return ebiten.Termination
		case StatePlaying, StatePlayerTurn, StateAITurn:
			g.state = StateMenu
		}
	}

	// Меню
	if g.state == StateMenu && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
	}

	// Game Over
	if g.state == StateGameOver && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
	}

	// Ход игрока
	if g.state == StatePlayerTurn {
		g.handleInput()
	}

	// Ход AI
	if g.state == StateAITurn {
		g.aiMove()
	}

	// Таймер сообщения
	if g.messageTimer > 0 {
		g.messageTimer--
	}

	return nil
}

// handleInput обрабатывает ввод
func (g *Game) handleInput() {
	// Клик мышью
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		// Преобразование в координаты доски
		col := (mx - boardOffsetX) / tileSize
		row := (my - boardOffsetY) / tileSize

		if col < 0 || col >= 8 || row < 0 || row >= 8 {
			return
		}

		g.handleClick(row, col)
	}
}

// handleClick обрабатывает клик
func (g *Game) handleClick(row, col int) {
	piece := g.board.Get(row, col)

	// Если выбрана шашка и клик на допустимый ход
	if g.selectedRow >= 0 {
		for _, m := range g.validMoves {
			if m.ToRow == row && m.ToCol == col {
				// Делаем ход
				g.board.MakeMove(m)
				g.selectedRow = -1
				g.selectedCol = -1
				g.validMoves = nil
				g.highlighted = nil

				// Проверка на серию взятий
				if m.CapturedRow >= 0 {
					moreCaptures := g.board.GetPieceCaptures(row, col, g.board.Get(row, col))
					if len(moreCaptures) > 0 {
						// Продолжаем взятие
						g.selectedRow = row
						g.selectedCol = col
						g.validMoves = moreCaptures
						g.highlighted = moreCaptures
						g.showMessage("Продолжай взятие!")
						return
					}
				}

				// Передача хода
				g.isWhiteTurn = false
				g.state = StateAITurn
				return
			}
		}
	}

	// Выбор своей шашки
	if board.IsOwnPiece(piece, g.isWhiteTurn) {
		g.selectedRow = row
		g.selectedCol = col

		// Получаем ходы
		allMoves := g.board.GetValidMoves(g.isWhiteTurn)

		// Фильтруем ходы для этой шашки
		g.validMoves = make([]board.Move, 0)
		for _, m := range allMoves {
			if m.FromRow == row && m.FromCol == col {
				g.validMoves = append(g.validMoves, m)
			}
		}

		g.highlighted = g.validMoves
	}
}

// aiMove ход AI
func (g *Game) aiMove() {
	g.state = StateAITurn // Блокируем повторные вызовы

	// Небольшая задержка для визуального эффекта
	if g.messageTimer > 0 {
		return
	}

	move := g.ai.GetMove(g.board)
	if move.FromRow >= 0 {
		g.board.MakeMove(move)

		// Проверка на серию взятий
		if move.CapturedRow >= 0 {
			moreCaptures := g.board.GetPieceCaptures(move.ToRow, move.ToCol, g.board.Get(move.ToRow, move.ToCol))
			if len(moreCaptures) > 0 {
				// AI продолжает взятие
				for len(moreCaptures) > 0 {
					// Выбираем первое взятие
					nextMove := moreCaptures[0]
					g.board.MakeMove(nextMove)
					moreCaptures = g.board.GetPieceCaptures(nextMove.ToRow, nextMove.ToCol, g.board.Get(nextMove.ToRow, nextMove.ToCol))
				}
			}
		}
	}

	g.isWhiteTurn = true
	g.state = StatePlayerTurn

	// Проверка окончания игры
	if g.board.IsGameOver(g.isWhiteTurn) {
		g.state = StateGameOver
		winner := g.board.GetWinner()
		if winner == "white" {
			g.message = "Белые победили! 🎉"
		} else if winner == "black" {
			g.message = "Чёрные победили! 🤖"
		} else {
			g.message = "Нет ходов - ничья!"
		}
		g.messageTimer = 300
	}
}

// showMessage показывает сообщение
func (g *Game) showMessage(msg string) {
	g.message = msg
	g.messageTimer = 120
}

// Draw рисует игру
func (g *Game) Draw(screen *ebiten.Image) {
	// Фон
	screen.Fill(color.RGBA{30, 30, 40, 255})

	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying, StatePlayerTurn, StateAITurn:
		g.drawBoard(screen)
		g.drawPieces(screen)
		g.drawHighlights(screen)
		g.drawHUD(screen)
		if g.messageTimer > 0 {
			g.drawMessage(screen)
		}
	case StateGameOver:
		g.drawBoard(screen)
		g.drawPieces(screen)
		g.drawGameOver(screen)
	}
}

// drawBoard рисует доску
func (g *Game) drawBoard(screen *ebiten.Image) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			x := boardOffsetX + col*tileSize
			y := boardOffsetY + row*tileSize

			var tile *ebiten.Image
			if (row+col)%2 == 0 {
				tile = g.spriteStore.LightTile
			} else {
				tile = g.spriteStore.DarkTile
			}

			if tile != nil {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(tile, opts)
			}
		}
	}

	// Рамка доски
	vector.DrawFilledRect(screen,
		float32(boardOffsetX-4), float32(boardOffsetY-4),
		float32(boardSize+8), float32(boardSize+8),
		color.RGBA{100, 70, 40, 255}, false)
}

// drawPieces рисует шашки
func (g *Game) drawPieces(screen *ebiten.Image) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := g.board.Get(row, col)
			if piece == board.Empty {
				continue
			}

			var spriteImg *ebiten.Image
			switch piece {
			case board.White:
				spriteImg = g.spriteStore.WhitePiece
			case board.Black:
				spriteImg = g.spriteStore.BlackPiece
			case board.WhiteKing:
				spriteImg = g.spriteStore.WhiteKing
			case board.BlackKing:
				spriteImg = g.spriteStore.BlackKing
			}

			if spriteImg != nil {
				x := boardOffsetX + col*tileSize + 10
				y := boardOffsetY + row*tileSize + 10

				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(spriteImg, opts)
			}
		}
	}
}

// drawHighlights рисует подсветку
func (g *Game) drawHighlights(screen *ebiten.Image) {
	// Подсветка выбранной шашки
	if g.selectedRow >= 0 {
		x := boardOffsetX + g.selectedCol*tileSize
		y := boardOffsetY + g.selectedRow*tileSize

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(g.spriteStore.Highlight, opts)
	}

	// Подсветка допустимых ходов
	for _, m := range g.highlighted {
		x := boardOffsetX + m.ToCol*tileSize + 10
		y := boardOffsetY + m.ToRow*tileSize + 10

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x), float64(y))
		opts.ColorScale.ScaleAlpha(0.7)
		screen.DrawImage(g.spriteStore.ValidMove, opts)
	}
}

// drawHUD рисует интерфейс
func (g *Game) drawHUD(screen *ebiten.Image) {
	whiteCount, blackCount := g.board.CountPieces()

	// Панель сверху
	vector.DrawFilledRect(screen, 0, 0, screenWidth, 50, color.RGBA{0, 0, 0, 180}, true)

	// Чей ход
	turnText := "⚪ Ход белых"
	if !g.isWhiteTurn {
		turnText = "⚫ Ход чёрных (AI думает...)"
	}
	if g.state == StateAITurn {
		turnText = "⚫ AI думает..."
	}
	ebitenutil.DebugPrintAt(screen, turnText, 20, 15)

	// Счёт
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("⚪ Белые: %d | ⚫ Чёрные: %d", whiteCount, blackCount), 400, 15)

	// Подсказка
	ebitenutil.DebugPrintAt(screen, "Кликни на шашку, затем на зелёную точку", 20, screenHeight-30)
}

// drawMessage рисует сообщение
func (g *Game) drawMessage(screen *ebiten.Image) {
	// Полупрозрачный фон
	msgWidth := len(g.message) * 20 + 40
	x := screenWidth/2 - msgWidth/2
	y := screenHeight/2 - 30

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(msgWidth), 50, color.RGBA{0, 0, 0, 200}, true)

	alpha := 1.0
	if g.messageTimer < 30 {
		alpha = float64(g.messageTimer) / 30.0
	}

	ebitenutil.DebugPrintAt(screen, g.message, x+20, y+15)
	_ = alpha
}

// drawMenu рисует меню
func (g *Game) drawMenu(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 200})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	menuText := `
╔═══════════════════════════════════════════╗
║         ♟️ CHECKERS (ШАШКИ) ♟️            ║
║                                           ║
║           [SPACE] - Начать игру           ║
║           [ESC] - Выход                   ║
║                                           ║
║  🎮 Управление:                           ║
║     Клик мышью - Выбор шашки и ход        ║
║                                           ║
║  📜 Правила:                              ║
║     • Белые ходят первыми                 ║
║     • Взятие обязательно                  ║
║     • Дамка ходит на любое расстояние     ║
║     • Обычная шашка бьёт назад            ║
║                                           ║
║  🤖 Ты играешь за белых, AI за чёрных    ║
╚═══════════════════════════════════════════╝
`
	ebitenutil.DebugPrintAt(screen, menuText, screenWidth/2-220, 80)
}

// drawGameOver рисует Game Over
func (g *Game) drawGameOver(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 180})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	winner := g.board.GetWinner()
	resultText := "Ничья!"
	if winner == "white" {
		resultText = "🎉 Ты победил!"
	} else if winner == "black" {
		resultText = "🤖 AI победил!"
	}

	gameOverText := fmt.Sprintf(`
╔═══════════════════════════════════════════╗
║         🏁 ИГРА ОКОНЧЕНА 🏁              ║
╠═══════════════════════════════════════════╣
║                                           ║
║           %s                      ║
║                                           ║
║           [SPACE] - Новая игра            ║
║           [ESC] - Выход                   ║
╚═══════════════════════════════════════════╝
`, resultText)

	ebitenutil.DebugPrintAt(screen, gameOverText, screenWidth/2-220, screenHeight/2-100)
}

// Layout размер экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// GetValidMovesForPiece получает ходы для шашки (для AI)
func (g *Game) GetValidMovesForPiece(row, col int) []board.Move {
	piece := g.board.Get(row, col)
	if !board.IsOwnPiece(piece, g.isWhiteTurn) {
		return nil
	}

	allMoves := g.board.GetValidMoves(g.isWhiteTurn)
	moves := make([]board.Move, 0)
	for _, m := range allMoves {
		if m.FromRow == row && m.FromCol == col {
			moves = append(moves, m)
		}
	}

	return moves
}

// Distance вычисляет расстояние
func Distance(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}
