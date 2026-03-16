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
	screenWidth  = 800
	screenHeight = 600
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
	bombs       []Bomb
	bombTimer   int
	bombDelay   int // время спавна бомб
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

type Bomb struct {
	pos      Point
	timer    int
	maxTime  int // время до взрыва
}

func NewGame() *Game {
	g := &Game{
		snake:     []Point{{gridSizeX / 2, gridSizeY / 2}, {gridSizeX/2 - 1, gridSizeY / 2}, {gridSizeX/2 - 2, gridSizeY / 2}},
		direction: Right,
		score:     0,
		gameOver:  false,
		moveDelay: 8,   // скорость движения змейки
		enemyDelay: 12,  // скорость врагов (медленнее змейки)
		bombDelay: 180,  // спавн бомбы каждые 180 тиков (~3 сек)
		enemies:   []Enemy{},
		bombs:     []Bomb{},
	}
	g.placeFood()
	// Spawn 10 enemies at start
	for i := 0; i < 10; i++ {
		g.spawnEnemy()
	}
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

func (g *Game) spawnBomb() {
	rand.Seed(time.Now().UnixNano())
	for {
		pos := Point{
			X: rand.Intn(gridSizeX),
			Y: rand.Intn(gridSizeY),
		}
		// Don't spawn on snake
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
			if dx < 15 && dy < 10 {
				tooClose = true
				break
			}
		}
		if !tooClose {
			g.bombs = append(g.bombs, Bomb{pos: pos, timer: 0, maxTime: 180}) // 3 секунды до взрыва
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

	// Spawn bombs periodically
	g.bombTimer++
	if g.bombTimer >= g.bombDelay {
		g.bombTimer = 0
		g.spawnBomb()
	}

	// Update bombs
	g.updateBombs()

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

func (g *Game) updateBombs() {
	for i := len(g.bombs) - 1; i >= 0; i-- {
		bomb := &g.bombs[i]
		bomb.timer++

		// Check collision with snake
		for _, segment := range g.snake {
			if segment.X == bomb.pos.X && segment.Y == bomb.pos.Y {
				g.gameOver = true
				return
			}
		}

		// Bomb explodes after maxTime
		if bomb.timer >= bomb.maxTime {
			// Check if snake is near explosion
			for _, segment := range g.snake {
				dx := segment.X - bomb.pos.X
				if dx < 0 {
					dx = -dx
				}
				dy := segment.Y - bomb.pos.Y
				if dy < 0 {
					dy = -dy
				}
				if dx < 3 && dy < 3 {
					g.gameOver = true
					return
				}
			}
			// Remove exploded bomb
			g.bombs = append(g.bombs[:i], g.bombs[i+1:]...)
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
		// Draw eyes on head
		if i == 0 {
			g.drawSnakeEyes(screen, segment, g.direction)
		}
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

	// Draw bombs
	for _, bomb := range g.bombs {
		g.drawBomb(screen, bomb)
	}

	// Draw score
	ebitenutil.DebugPrintAt(screen, "Score: "+string(rune('0'+g.score)), 10, 10)

	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER! Press ENTER to restart", screenWidth/2-150, screenHeight/2)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Final Score: %d - Enemies: %d - Bombs: %d", g.score, len(g.enemies), len(g.bombs)), screenWidth/2-150, screenHeight/2+30)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) drawEnemy(screen *ebiten.Image, enemy Enemy) {
	x := float32(enemy.pos.X * tileSize)
	y := float32(enemy.pos.Y * tileSize)
	size := float32(tileSize) * 1.5 // Increase bug size by 1.5x

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

	// Scary fangs/teeth (white sharp triangles on sides of head)
	fangSize := size / 5
	// Left fang
	leftFangX := headX - size/4
	leftFangY := headY + size/6
	vector.StrokeLine(screen, leftFangX-fangSize/2, leftFangY-fangSize/2, leftFangX, leftFangY+fangSize/2, 3, color.RGBA{255, 255, 255, 255}, false)
	vector.StrokeLine(screen, leftFangX+fangSize/2, leftFangY-fangSize/2, leftFangX, leftFangY+fangSize/2, 3, color.RGBA{255, 255, 255, 255}, false)
	// Right fang
	rightFangX := headX + size/4
	rightFangY := headY + size/6
	vector.StrokeLine(screen, rightFangX-fangSize/2, rightFangY-fangSize/2, rightFangX, rightFangY+fangSize/2, 3, color.RGBA{255, 255, 255, 255}, false)
	vector.StrokeLine(screen, rightFangX+fangSize/2, rightFangY-fangSize/2, rightFangX, rightFangY+fangSize/2, 3, color.RGBA{255, 255, 255, 255}, false)

	// Additional smaller fangs (inner teeth)
	// Left small fang
	vector.StrokeLine(screen, leftFangX-fangSize/2, leftFangY-fangSize/4, leftFangX-fangSize/4, leftFangY+fangSize/4, 2, color.RGBA{230, 230, 230, 255}, false)
	vector.StrokeLine(screen, leftFangX, leftFangY-fangSize/4, leftFangX-fangSize/4, leftFangY+fangSize/4, 2, color.RGBA{230, 230, 230, 255}, false)
	// Right small fang
	vector.StrokeLine(screen, rightFangX, rightFangY-fangSize/4, rightFangX+fangSize/4, rightFangY+fangSize/4, 2, color.RGBA{230, 230, 230, 255}, false)
	vector.StrokeLine(screen, rightFangX+fangSize/2, rightFangY-fangSize/4, rightFangX+fangSize/4, rightFangY+fangSize/4, 2, color.RGBA{230, 230, 230, 255}, false)

	// Eyes (yellow dots)
	vector.DrawFilledCircle(screen, headX-size/8, headY-size/8, 2, color.RGBA{255, 255, 0, 255}, false)
	vector.DrawFilledCircle(screen, headX+size/8, headY-size/8, 2, color.RGBA{255, 255, 0, 255}, false)
}

func (g *Game) drawSnakeEyes(screen *ebiten.Image, head Point, direction Direction) {
	x := float32(head.X * tileSize)
	y := float32(head.Y * tileSize)
	size := float32(tileSize)
	eyeSize := size / 6
	pupilSize := eyeSize / 2

	// Eye positions based on direction
	var leftEyeX, leftEyeY, rightEyeX, rightEyeY float32

	switch direction {
	case Up:
		leftEyeX = x + size/3
		leftEyeY = y + size/3
		rightEyeX = x + 2*size/3
		rightEyeY = y + size/3
	case Down:
		leftEyeX = x + size/3
		leftEyeY = y + 2*size/3
		rightEyeX = x + 2*size/3
		rightEyeY = y + 2*size/3
	case Left:
		leftEyeX = x + size/3
		leftEyeY = y + size/3
		rightEyeX = x + size/3
		rightEyeY = y + 2*size/3
	case Right:
		leftEyeX = x + 2*size/3
		leftEyeY = y + size/3
		rightEyeX = x + 2*size/3
		rightEyeY = y + 2*size/3
	}

	// Draw whites of eyes
	vector.DrawFilledCircle(screen, leftEyeX, leftEyeY, eyeSize, color.RGBA{255, 255, 255, 255}, false)
	vector.DrawFilledCircle(screen, rightEyeX, rightEyeY, eyeSize, color.RGBA{255, 255, 255, 255}, false)

	// Draw pupils (black)
	vector.DrawFilledCircle(screen, leftEyeX, leftEyeY, pupilSize, color.RGBA{0, 0, 0, 255}, false)
	vector.DrawFilledCircle(screen, rightEyeX, rightEyeY, pupilSize, color.RGBA{0, 0, 0, 255}, false)
}

func (g *Game) drawBomb(screen *ebiten.Image, bomb Bomb) {
	x := float32(bomb.pos.X * tileSize)
	y := float32(bomb.pos.Y * tileSize)
	size := float32(tileSize)

	// Bomb body (black circle)
	vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/2-2, color.RGBA{0, 0, 0, 255}, false)

	// Shine on bomb
	vector.DrawFilledCircle(screen, x+size/3, y+size/3, size/6, color.RGBA{50, 50, 50, 255}, false)

	// Fuse (brown stick)
	fuseX := x + size/2
	fuseY := y + size/4
	vector.StrokeLine(screen, fuseX, fuseY, fuseX, fuseY-size/3, 2, color.RGBA{139, 69, 19, 255}, false)

	// Spark at end of fuse (animated - blinking)
	sparkPhase := (bomb.timer % 10) / 5.0
	sparkSize := size/6 + float32(sparkPhase)*size/8

	// Yellow/orange spark glow
	vector.DrawFilledCircle(screen, fuseX, fuseY-size/3, sparkSize, color.RGBA{255, 200, 0, 200}, false)

	// White hot center
	vector.DrawFilledCircle(screen, fuseX, fuseY-size/3, sparkSize/2, color.RGBA{255, 255, 255, 255}, false)

	// Spark particles (random sparks around)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		particleX := fuseX + float32(rand.Intn(8)-4)
		particleY := fuseY - size/3 + float32(rand.Intn(8)-4)
		vector.DrawFilledCircle(screen, particleX, particleY, 1, color.RGBA{255, 100, 0, 255}, false)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Snake - Go365 Go75")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
