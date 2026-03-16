package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gridSize     = 20
	tileSize     = 20
	screenWidth  = gridSize * tileSize
	screenHeight = gridSize * tileSize
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Game struct {
	snake      []Point
	direction  Direction
	food       Point
	score      int
	gameOver   bool
}

type Point struct {
	X int
	Y int
}

func NewGame() *Game {
	g := &Game{
		snake:     []Point{{5, 5}, {4, 5}, {3, 5}},
		direction: Right,
		score:     0,
		gameOver:  false,
	}
	g.placeFood()
	return g
}

func (g *Game) placeFood() {
	rand.Seed(time.Now().UnixNano())
	for {
		g.food = Point{
			X: rand.Intn(gridSize),
			Y: rand.Intn(gridSize),
		}
		// Check if food is not on snake
		onSnake := false
		for _, segment := range g.snake {
			if segment.X == g.food.X && segment.Y == g.food.Y {
				onSnake = true
				break
			}
		}
		if !onSnake {
			break
		}
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		// Press Enter to restart
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			*g = *NewGame()
		}
		return nil
	}

	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.direction != Down {
		g.direction = Up
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.direction != Up {
		g.direction = Down
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.direction != Right {
		g.direction = Left
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && g.direction != Left {
		g.direction = Right
	}

	// Move snake
	head := g.snake[0]
	newHead := head

	switch g.direction {
	case Up:
		newHead.Y--
	case Down:
		newHead.Y++
	case Left:
		newHead.X--
	case Right:
		newHead.X++
	}

	// Check wall collision
	if newHead.X < 0 || newHead.X >= gridSize || newHead.Y < 0 || newHead.Y >= gridSize {
		g.gameOver = true
		return nil
	}

	// Check self collision
	for _, segment := range g.snake {
		if segment.X == newHead.X && segment.Y == newHead.Y {
			g.gameOver = true
			return nil
		}
	}

	// Add new head
	g.snake = append([]Point{newHead}, g.snake...)

	// Check food collision
	if newHead.X == g.food.X && newHead.Y == g.food.Y {
		g.score++
		g.placeFood()
	} else {
		// Remove tail
		g.snake = g.snake[:len(g.snake)-1]
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw snake
	for i, segment := range g.snake {
		green := color.RGBA{0, 255, 0, 255}
		if i == 0 {
			// Head is brighter
			green = color.RGBA{100, 255, 100, 255}
		}
		vector.DrawFilledRect(
			screen,
			float32(segment.X*tileSize),
			float32(segment.Y*tileSize),
			tileSize,
			tileSize,
			green,
			false,
		)
	}

	// Draw food
	vector.DrawFilledRect(
		screen,
		float32(g.food.X*tileSize),
		float32(g.food.Y*tileSize),
		tileSize,
		tileSize,
		color.RGBA{255, 0, 0, 255},
		false,
	)

	// Draw score
	ebitenutil.DebugPrintAt(screen, "Score: "+string(rune('0'+g.score)), 10, 10)

	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER! Press ENTER to restart", screenWidth/2-100, screenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Snake - Go365 Go75")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
