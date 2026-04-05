// Package game - основная игровая логика Bomberman
// Go365 Day 95 - Bomberman
package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"bomberman/internal/entity"
	"bomberman/internal/level"
	"bomberman/internal/sprite"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	tileSize     = 32
)

// GameState состояние игры
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateLevelComplete
	StateVictory
)

// Game основная игра
type Game struct {
	state       GameState
	player      *entity.Player
	level       *level.Level
	bombs       []*entity.Bomb
	explosions  []*entity.Explosion
	enemies     []*entity.Enemy
	powerUps    []*entity.PowerUp
	score       int
	levelNum    int
	spriteStore *sprite.SpriteStore
	rng         *rand.Rand
	cameraX     float64
	cameraY     float64
	bombPlaced  map[string]bool
}

// NewGame создаёт новую игру
func NewGame() *Game {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := &Game{
		state:      StateMenu,
		rng:        rng,
		bombPlaced: make(map[string]bool),
	}
	g.spriteStore = sprite.NewSpriteStore()
	return g
}

// Reset сбрасывает игру
func (g *Game) Reset() {
	g.levelNum = 1
	g.score = 0
	g.startLevel()
}

// startLevel запускает уровень
func (g *Game) startLevel() {
	g.level = level.GenerateLevel(g.levelNum, g.rng)
	g.player = entity.NewPlayer(g.level.PlayerX, g.level.PlayerY, g.spriteStore)
	g.bombs = make([]*entity.Bomb, 0)
	g.explosions = make([]*entity.Explosion, 0)
	g.enemies = make([]*entity.Enemy, 0)
	g.powerUps = make([]*entity.PowerUp, 0)
	g.bombPlaced = make(map[string]bool)

	// Создание врагов
	for _, es := range g.level.Enemies {
		enemy := entity.NewEnemy(es.X, es.Y, es.Type, g.spriteStore)
		g.enemies = append(g.enemies, enemy)
	}

	// Создание улучшений
	for _, ps := range g.level.PowerUps {
		powerUp := entity.NewPowerUp(ps.X, ps.Y, ps.Type)
		g.powerUps = append(g.powerUps, powerUp)
	}

	g.cameraX = 0
	g.cameraY = 0
}

// Update обновляет игру
func (g *Game) Update() error {
	// ESC
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		switch g.state {
		case StatePlaying:
			g.state = StatePaused
		case StatePaused:
			g.state = StatePlaying
		case StateMenu, StateGameOver, StateVictory:
			return ebiten.Termination
		}
	}

	// Меню
	if g.state == StateMenu && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
		g.state = StatePlaying
	}

	// Game Over
	if g.state == StateGameOver && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
		g.state = StatePlaying
	}

	// Level Complete
	if g.state == StateLevelComplete && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.levelNum++
		if g.levelNum > 5 {
			g.state = StateVictory
		} else {
			g.startLevel()
			g.state = StatePlaying
		}
	}

	// Victory
	if g.state == StateVictory && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
		g.state = StatePlaying
	}

	// Игра
	if g.state == StatePlaying {
		g.updateGame()
	}

	return nil
}

// updateGame обновляет игровую логику
func (g *Game) updateGame() {
	// Управление
	g.handleInput()

	// Обновление игрока
	g.player.Update()

	// Коллизия игрока со стенами
	g.resolvePlayerCollision()

	// Обновление бомб
	g.updateBombs()

	// Обновление взрывов
	g.updateExplosions()

	// Обновление врагов
	g.updateEnemies()

	// Сбор улучшений
	g.collectPowerUps()

	// Камера
	g.updateCamera()

	// Проверка смерти
	if g.player.Lives <= 0 {
		g.state = StateGameOver
	}

	// Проверка победы (все враги убиты)
	if len(g.enemies) == 0 {
		g.state = StateLevelComplete
	}
}

// handleInput обрабатывает ввод
func (g *Game) handleInput() {
	dx, dy := 0.0, 0.0

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = 1
	}

	// Нормализация диагонали
	if dx != 0 && dy != 0 {
		dx *= 0.707
		dy *= 0.707
	}

	if dx != 0 || dy != 0 {
		g.player.Move(dx, dy)
	}

	// Постановка бомбы
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.placeBomb()
	}
}

// placeBomb ставит бомбу
func (g *Game) placeBomb() {
	key := fmt.Sprintf("%d,%d", g.player.GridX, g.player.GridY)
	if g.bombPlaced[key] {
		return
	}

	// Считаем активные бомбы игрока
	activeBombs := 0
	for _, b := range g.bombs {
		if b.Active {
			activeBombs++
		}
	}

	if activeBombs >= g.player.MaxBombs {
		return
	}

	bomb := entity.NewBomb(g.player.GridX, g.player.GridY, g.player.FlameRange)
	g.bombs = append(g.bombs, bomb)
	g.bombPlaced[key] = true
}

// resolvePlayerCollision разрешает коллизии игрока
func (g *Game) resolvePlayerCollision() {
	px := g.player.GridX
	py := g.player.GridY

	// Проверка границ
	if px < 0 || px >= g.level.Width || py < 0 || py >= g.level.Height {
		g.player.X = g.player.X
		g.player.Y = g.player.Y
		return
	}

	// Проверка стен
	if !g.level.IsWalkable(px, py) {
		// Откат позиции
		g.player.X = float64(g.player.GridX) * 32
		g.player.Y = float64(g.player.GridY) * 32
	}
}

// updateBombs обновляет бомбы
func (g *Game) updateBombs() {
	activeBombs := make([]*entity.Bomb, 0)

	for _, b := range g.bombs {
		b.Update()

		if !b.Active {
			// Взрыв!
			g.createExplosion(b.X, b.Y, b.FlameRange)
			key := fmt.Sprintf("%d,%d", b.X, b.Y)
			delete(g.bombPlaced, key)
		} else {
			activeBombs = append(activeBombs, b)
		}
	}

	g.bombs = activeBombs
}

// createExplosion создаёт взрыв
func (g *Game) createExplosion(x, y, flameRange int) {
	// Центр
	exp := entity.NewExplosion(x, y)
	g.explosions = append(g.explosions, exp)

	// Направления: вверх, вниз, влево, вправо
	directions := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	for _, dir := range directions {
		for i := 1; i <= flameRange; i++ {
			nx := x + dir.dx*i
			ny := y + dir.dy*i

			// Проверка твёрдой стены
			if !g.level.IsWalkable(nx, ny) && !g.level.IsBreakable(nx, ny) {
				break
			}

			// Разрушение стены
			if g.level.IsBreakable(nx, ny) {
				g.level.BreakWall(nx, ny)
				g.explosions = append(g.explosions, entity.NewExplosion(nx, ny))
				g.score += 10
				break
			}

			g.explosions = append(g.explosions, entity.NewExplosion(nx, ny))
		}
	}

	// Проверка попадания во врагов
	g.checkExplosionDamage(x, y, flameRange)

	// Проверка попадания в игрока
	g.checkPlayerExplosionDamage(x, y, flameRange)
}

// checkExplosionDamage проверяет урон врагам
func (g *Game) checkExplosionDamage(cx, cy, flameRange int) {
	aliveEnemies := make([]*entity.Enemy, 0)

	for _, e := range g.enemies {
		hit := false

		// Центр
		if e.GridX == cx && e.GridY == cy {
			hit = true
		}

		// Направления
		directions := []struct{ dx, dy int }{
			{0, -1}, {0, 1}, {-1, 0}, {1, 0},
		}

		for _, dir := range directions {
			for i := 1; i <= flameRange; i++ {
				nx := cx + dir.dx*i
				ny := cy + dir.dy*i

				if !g.level.IsWalkable(nx, ny) && !g.level.IsBreakable(nx, ny) {
					break
				}

				if e.GridX == nx && e.GridY == ny {
					hit = true
					break
				}

				if g.level.IsBreakable(nx, ny) {
					break
				}
			}
		}

		if !hit {
			aliveEnemies = append(aliveEnemies, e)
		} else {
			g.score += 100
		}
	}

	g.enemies = aliveEnemies
}

// checkPlayerExplosionDamage проверяет урон игроку
func (g *Game) checkPlayerExplosionDamage(cx, cy, flameRange int) {
	if g.player.Invincible > 0 {
		return
	}

	hit := false

	if g.player.GridX == cx && g.player.GridY == cy {
		hit = true
	}

	directions := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	for _, dir := range directions {
		for i := 1; i <= flameRange; i++ {
			nx := cx + dir.dx*i
			ny := cy + dir.dy*i

			if !g.level.IsWalkable(nx, ny) && !g.level.IsBreakable(nx, ny) {
				break
			}

			if g.player.GridX == nx && g.player.GridY == ny {
				hit = true
				break
			}

			if g.level.IsBreakable(nx, ny) {
				break
			}
		}
	}

	if hit {
		g.player.Lives--
		g.player.Invincible = 120
	}
}

// updateExplosions обновляет взрывы
func (g *Game) updateExplosions() {
	active := make([]*entity.Explosion, 0)

	for _, e := range g.explosions {
		e.Update()
		if e.Active {
			active = append(active, e)
		}
	}

	g.explosions = active
}

// updateEnemies обновляет врагов
func (g *Game) updateEnemies() {
	for _, e := range g.enemies {
		e.Update()

		// Коллизия со стенами
		if !g.level.IsWalkable(e.GridX, e.GridY) {
			e.X = float64(e.GridX) * 32
			e.Y = float64(e.GridY) * 32
			e.RandomDirection()
			e.ChangeTimer = 30
		}

		// Коллизия с игроком
		if e.GridX == g.player.GridX && e.GridY == g.player.GridY {
			if g.player.Invincible <= 0 {
				g.player.Lives--
				g.player.Invincible = 120
			}
		}
	}
}

// collectPowerUps собирает улучшения
func (g *Game) collectPowerUps() {
	active := make([]*entity.PowerUp, 0)

	for _, p := range g.powerUps {
		if p.X == g.player.GridX && p.Y == g.player.GridY {
			// Применить улучшение
			switch p.Type {
			case "bomb_up":
				g.player.MaxBombs++
			case "flame_up":
				g.player.FlameRange++
			case "speed_up":
				g.player.Speed += 0.5
			}
			g.score += 50
		} else {
			active = append(active, p)
		}
	}

	g.powerUps = active
}

// updateCamera обновляет камеру
func (g *Game) updateCamera() {
	targetX := g.player.X - screenWidth/2
	targetY := g.player.Y - screenHeight/2

	g.cameraX += (targetX - g.cameraX) * 0.1
	g.cameraY += (targetY - g.cameraY) * 0.1

	if g.cameraX < 0 {
		g.cameraX = 0
	}
	if g.cameraY < 0 {
		g.cameraY = 0
	}

	maxX := float64(g.level.Width*tileSize) - screenWidth
	maxY := float64(g.level.Height*tileSize) - screenHeight

	if g.cameraX > maxX {
		g.cameraX = maxX
	}
	if g.cameraY > maxY {
		g.cameraY = maxY
	}
}

// Draw рисует игру
func (g *Game) Draw(screen *ebiten.Image) {
	// Фон
	screen.Fill(color.RGBA{40, 60, 40, 255})

	if g.level == nil {
		return
	}

	// Сетка уровня
	g.drawLevel(screen)

	// Улучшения
	for _, p := range g.powerUps {
		g.drawPowerUp(screen, p)
	}

	// Бомбы
	for _, b := range g.bombs {
		g.drawBomb(screen, b)
	}

	// Взрывы
	for _, e := range g.explosions {
		g.drawExplosion(screen, e)
	}

	// Враги
	for _, e := range g.enemies {
		e.Draw(screen, g.cameraX, g.cameraY)
	}

	// Игрок
	g.player.Draw(screen, g.cameraX, g.cameraY)

	// UI
	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		g.drawHUD(screen)
	case StatePaused:
		g.drawHUD(screen)
		g.drawPause(screen)
	case StateGameOver:
		g.drawGameOver(screen)
	case StateLevelComplete:
		g.drawLevelComplete(screen)
	case StateVictory:
		g.drawVictory(screen)
	}
}

// drawLevel рисует уровень
func (g *Game) drawLevel(screen *ebiten.Image) {
	for y := 0; y < g.level.Height; y++ {
		for x := 0; x < g.level.Width; x++ {
			screenX := float64(x*tileSize) - g.cameraX
			screenY := float64(y*tileSize) - g.cameraY

			// Пол
			floorSprite := g.spriteStore.Walls["floor"]
			if floorSprite != nil {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(screenX, screenY)
				screen.DrawImage(floorSprite, opts)
			}

			// Стены
			tile := g.level.Grid[y][x]
			var spriteImg *ebiten.Image

			switch tile {
			case level.TileSolid:
				spriteImg = g.spriteStore.Walls["solid"]
			case level.TileBreakable:
				spriteImg = g.spriteStore.Walls["breakable"]
			}

			if spriteImg != nil {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(screenX, screenY)
				screen.DrawImage(spriteImg, opts)
			}
		}
	}
}

// drawBomb рисует бомбу
func (g *Game) drawBomb(screen *ebiten.Image, b *entity.Bomb) {
	spriteImg := g.spriteStore.Bombs[b.GetFrame()]
	if spriteImg != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(b.X*32)-g.cameraX, float64(b.Y*32)-g.cameraY)
		screen.DrawImage(spriteImg, opts)
	}
}

// drawExplosion рисует взрыв
func (g *Game) drawExplosion(screen *ebiten.Image, e *entity.Explosion) {
	idx := e.GetSpriteIndex()
	if idx < len(g.spriteStore.Explosions) {
		spriteImg := g.spriteStore.Explosions[idx]
		if spriteImg != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(e.X*32)-g.cameraX, float64(e.Y*32)-g.cameraY)
			screen.DrawImage(spriteImg, opts)
		}
	}
}

// drawPowerUp рисует улучшение
func (g *Game) drawPowerUp(screen *ebiten.Image, p *entity.PowerUp) {
	spriteImg := g.spriteStore.PowerUps[p.Type]
	if spriteImg != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(p.X*32)-g.cameraX, float64(p.Y*32)-g.cameraY)
		screen.DrawImage(spriteImg, opts)
	}
}

// drawMenu рисует меню
func (g *Game) drawMenu(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 200})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	menuText := `
╔═══════════════════════════════════════════╗
║         💣 BOMBERMAN 💣                   ║
║                                           ║
║           [SPACE] - Начать игру           ║
║           [ESC] - Выход                   ║
║                                           ║
║  🎮 Управление:                           ║
║     WASD или Стрелки - Движение           ║
║     SPACE - Поставить бомбу               ║
║                                           ║
║  🎯 Цель: Уничтожь всех врагов!           ║
║  💣 Собирай улучшения:                    ║
║     🔴 Больше бомб                        ║
║     🔥 Дальше взрыв                       ║
║     ⚡ Быстрее бег                        ║
╚═══════════════════════════════════════════╝
`
	ebitenutil.DebugPrintAt(screen, menuText, screenWidth/2-220, 80)
}

// drawHUD рисует интерфейс
func (g *Game) drawHUD(screen *ebiten.Image) {
	// Фон панели
	vector.DrawFilledRect(screen, 0, 0, screenWidth, 40, color.RGBA{0, 0, 0, 180}, true)

	// Жизни
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("❤️ x %d", g.player.Lives), 10, 10)

	// Бомбы
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("💣 x %d", g.player.MaxBombs), 150, 10)

	// Огонь
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("🔥 x %d", g.player.FlameRange), 300, 10)

	// Счёт
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("⭐ Счёт: %d", g.score), 450, 10)

	// Уровень
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("📍 Уровень: %d", g.levelNum), 700, 10)

	// Враги
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("👾 Враги: %d", len(g.enemies)), 900, 10)
}

// drawPause рисует паузу
func (g *Game) drawPause(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 180})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	ebitenutil.DebugPrintAt(screen, `
╔═══════════════════════════════════╗
║          ⏸️ ПАУЗА                  ║
╠═══════════════════════════════════╣
║   [ESC] - Продолжить              ║
║   [SPACE] - Выйти в меню          ║
╚═══════════════════════════════════╝
`, screenWidth/2-180, screenHeight/2-80)
}

// drawGameOver рисует Game Over
func (g *Game) drawGameOver(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{80, 0, 0, 200})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(`
╔═══════════════════════════════════╗
║       💀 GAME OVER 💀             ║
╠═══════════════════════════════════╣
║     Счёт: %6d                      ║
║     Уровень: %2d                    ║
║                                   ║
║     [SPACE] - Заново              ║
║     [ESC] - Выход                 ║
╚═══════════════════════════════════╝
`, g.score, g.levelNum), screenWidth/2-180, screenHeight/2-100)
}

// drawLevelComplete рисует завершение уровня
func (g *Game) drawLevelComplete(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 80, 0, 180})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(`
╔═══════════════════════════════════╗
║     ✅ УРОВЕНЬ %d ПРОЙДЕН! ✅      ║
╠═══════════════════════════════════╣
║     Счёт: %6d                      ║
║                                   ║
║     [SPACE] - Следующий уровень   ║
╚═══════════════════════════════════╝
`, g.levelNum, g.score), screenWidth/2-180, screenHeight/2-80)
}

// drawVictory рисует победу
func (g *Game) drawVictory(screen *ebiten.Image) {
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 80, 0, 180})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(overlay, op)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(`
╔═══════════════════════════════════╗
║     🎉 ПОБЕДА! 🎉                 ║
╠═══════════════════════════════════╣
║     Все уровни пройдены!           ║
║     Финальный счёт: %6d            ║
║                                   ║
║     [SPACE] - Играть снова        ║
║     [ESC] - Выход                 ║
╚═══════════════════════════════════╝
`, g.score), screenWidth/2-180, screenHeight/2-100)
}

// Layout размер экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
