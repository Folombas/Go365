package main

import (
	"fmt"
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
	tileSize     = 20
	screenWidth  = 2560
	screenHeight = 1440
	gridSizeX    = screenWidth / tileSize
	gridSizeY    = screenHeight / tileSize
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Game struct {
	snake       []Point
	direction   Direction
	food        Point
	score       int
	gameOver    bool
	moveTimer   int
	moveDelay   int // тиков между движениями (60 тиков = 1 сек)
	enemies     []Enemy
	enemyTimer  int
	enemyDelay  int // скорость врагов
}

type Point struct {
	X int
	Y int
}

type Enemy struct {
	pos       Point
	direction Direction
	animFrame int
}

func NewGame() *Game {
	g := &Game{
		snake:     []Point{{gridSizeX / 2, gridSizeY / 2}, {gridSizeX/2 - 1, gridSizeY / 2}, {gridSizeX/2 - 2, gridSizeY / 2}},
		direction: Right,
		score:     0,
		gameOver:  false,
		moveDelay: 8,  // скорость движения змейки
		enemyDelay: 12, // скорость врагов (медленнее змейки)
		enemies:   []Enemy{},
	}
	g.placeFood()
	g.spawnEnemy()
	return g
}

func (g *Game) placeFood() {
	rand.Seed(time.Now().UnixNano())
	for {
		g.food = Point{
			X: rand.Intn(gridSizeX),
			Y: rand.Intn(gridSizeY),
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

func (g *Game) spawnEnemy() {
	rand.Seed(time.Now().UnixNano())
	for {
		pos := Point{
			X: rand.Intn(gridSizeX),
			Y: rand.Intn(gridSizeY),
		}
		// Don't spawn on snake or too close to player
		tooClose := false
		for _, segment := range g.snake {
			dx := segment.X - pos.X
			if dx < 0 {
				dx = -dx
			}
			dy := segment.Y - pos.Y
			if dy < 0 {
				dy = -dy
			}
			if dx < 20 && dy < 15 {
				tooClose = true
				break
			}
		}
		if !tooClose {
			dir := Direction(rand.Intn(4))
			g.enemies = append(g.enemies, Enemy{pos: pos, direction: dir, animFrame: 0})
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

	// Update move timer
	g.moveTimer++
	if g.moveTimer < g.moveDelay {
		return nil
	}
	g.moveTimer = 0

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
	if newHead.X < 0 || newHead.X >= gridSizeX || newHead.Y < 0 || newHead.Y >= gridSizeY {
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
		// Spawn new enemy every 2 points
		if g.score%2 == 0 {
			g.spawnEnemy()
		}
	} else {
		// Remove tail
		g.snake = g.snake[:len(g.snake)-1]
	}

	// Update enemies
	g.enemyTimer++
	if g.enemyTimer >= g.enemyDelay {
		g.enemyTimer = 0
		g.updateEnemies()
	}

	return nil
}

func (g *Game) updateEnemies() {
	for i := range g.enemies {
		enemy := &g.enemies[i]
		enemy.animFrame++

		// Move enemy
		newPos := enemy.pos
		switch enemy.direction {
		case Up:
			newPos.Y--
		case Down:
			newPos.Y++
		case Left:
			newPos.X--
		case Right:
			newPos.X++
		}

		// Check bounds - reverse direction if hitting wall
		if newPos.X < 0 || newPos.X >= gridSizeX || newPos.Y < 0 || newPos.Y >= gridSizeY {
			enemy.direction = Direction(rand.Intn(4))
			continue
		}

		enemy.pos = newPos

		// Random direction change
		if rand.Intn(10) < 2 {
			enemy.direction = Direction(rand.Intn(4))
		}

		// Check collision with snake
		for _, segment := range g.snake {
			if segment.X == enemy.pos.X && segment.Y == enemy.pos.Y {
				g.gameOver = true
				return
			}
		}
	}
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

	// Draw enemies (bugs with legs and antennae)
	for _, enemy := range g.enemies {
		g.drawEnemy(screen, enemy)
	}

	// Draw score
	ebitenutil.DebugPrintAt(screen, "Score: "+string(rune('0'+g.score)), 10, 10)

	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER! Press ENTER to restart", screenWidth/2-150, screenHeight/2)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Final Score: %d - Enemies: %d", g.score, len(g.enemies)), screenWidth/2-120, screenHeight/2+30)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) drawEnemy(screen *ebiten.Image, enemy Enemy) {
	x := float32(enemy.pos.X * tileSize)
	y := float32(enemy.pos.Y * tileSize)
	size := float32(tileSize)

	// Body (dark purple oval)
	vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2-2, color.RGBA{128, 0, 128, 255}, false)

	// Head
	headX := x + size/2
	headY := y + size/2
	vector.DrawFilledCircle(screen, headX, headY, size/3, color.RGBA{100, 0, 100, 255}, false)

	// Animated legs (6 legs - 3 on each side)
	legOffset := float32((enemy.animFrame % 20) / 10.0 * 3)
	if enemy.animFrame%40 < 20 {
		legOffset = -legOffset
	}

	// Left legs
	for i := 0; i < 3; i++ {
		legY := y + size/4 + float32(i)*size/4
		vector.StrokeLine(screen, x+size/3, legY, x-size/4, legY+legOffset+float32(i)*2, 2, color.RGBA{100, 0, 100, 255}, false)
	}

	// Right legs
	for i := 0; i < 3; i++ {
		legY := y + size/4 + float32(i)*size/4
		vector.StrokeLine(screen, x+2*size/3, legY, x+5*size/4, legY-legOffset+float32(i)*2, 2, color.RGBA{100, 0, 100, 255}, false)
	}

	// Antennae (animated)
	antennaAngle := float32((enemy.animFrame % 30) / 30.0 * 1.0)
	if enemy.animFrame%60 < 30 {
		antennaAngle = -antennaAngle
	}

	// Left antenna
	vector.StrokeLine(screen, headX-size/6, headY-size/3, headX-size/2-antennaAngle*size, headY-size/2-antennaAngle*size, 1, color.RGBA{150, 50, 50, 255}, false)
	// Right antenna
	vector.StrokeLine(screen, headX+size/6, headY-size/3, headX+size/2+antennaAngle*size, headY-size/2-antennaAngle*size, 1, color.RGBA{150, 50, 50, 255}, false)

	// Eyes (yellow dots)
	vector.DrawFilledCircle(screen, headX-size/8, headY-size/8, 2, color.RGBA{255, 255, 0, 255}, false)
	vector.DrawFilledCircle(screen, headX+size/8, headY-size/8, 2, color.RGBA{255, 255, 0, 255}, false)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Snake - Go365 Go75")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
