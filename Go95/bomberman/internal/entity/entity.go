// Package entity - игровые сущности для Bomberman
// Go365 Day 95 - Bomberman
package entity

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"bomberman/internal/sprite"
)

// Player игрок
type Player struct {
	X, Y         float64
	GridX, GridY int
	Speed        float64
	MaxBombs     int
	FlameRange   int
	Invincible   int
	Lives        int
	Direction    string // up, down, left, right
	AnimFrame    float64
	spriteStore  *sprite.SpriteStore
}

// NewPlayer создаёт нового игрока
func NewPlayer(x, y int, ss *sprite.SpriteStore) *Player {
	return &Player{
		X:         float64(x) * 32,
		Y:         float64(y) * 32,
		GridX:     x,
		GridY:     y,
		Speed:     3.0,
		MaxBombs:  1,
		FlameRange: 2,
		Lives:     3,
		Direction: "down",
		spriteStore: ss,
	}
}

// Update обновляет состояние игрока
func (p *Player) Update() {
	if p.Invincible > 0 {
		p.Invincible--
	}
	p.AnimFrame += 0.15
}

// Move двигает игрока
func (p *Player) Move(dx, dy float64) {
	p.X += dx * p.Speed
	p.Y += dy * p.Speed

	// Обновляем позицию в сетке
	p.GridX = int(math.Round(p.X / 32))
	p.GridY = int(math.Round(p.Y / 32))

	// Определяем направление
	if dx > 0 {
		p.Direction = "right"
	} else if dx < 0 {
		p.Direction = "left"
	} else if dy > 0 {
		p.Direction = "down"
	} else if dy < 0 {
		p.Direction = "up"
	}
}

// GetSprite получает спрайт игрока
func (p *Player) GetSprite() *ebiten.Image {
	return p.spriteStore.Bomberman[p.Direction]
}

// Draw рисует игрока
func (p *Player) Draw(screen *ebiten.Image, camX, camY float64) {
	if p.Invincible > 0 && p.Invincible%10 < 5 {
		return // Мигание
	}

	spriteImg := p.GetSprite()
	if spriteImg != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(p.X-camX, p.Y-camY)
		screen.DrawImage(spriteImg, opts)
	}
}

// Bomb бомба
type Bomb struct {
	X, Y       int
	Timer      int
	FlameRange int
	Active     bool
}

// NewBomb создаёт бомбу
func NewBomb(x, y, flameRange int) *Bomb {
	return &Bomb{
		X: x, Y: y,
		Timer:      120, // 2 секунды при 60 FPS
		FlameRange: flameRange,
		Active:     true,
	}
}

// Update обновляет бомбу
func (b *Bomb) Update() {
	b.Timer--
	if b.Timer <= 0 {
		b.Active = false
	}
}

// GetFrame получает кадр анимации
func (b *Bomb) GetFrame() string {
	elapsed := 120 - b.Timer
	if elapsed < 40 {
		return "bomb_0"
	} else if elapsed < 80 {
		return "bomb_1"
	}
	return "bomb_2"
}

// Explosion взрыв
type Explosion struct {
	X, Y   int
	Frame  int
	MaxFrame int
	Active bool
}

// NewExplosion создаёт взрыв
func NewExplosion(x, y int) *Explosion {
	return &Explosion{
		X: x, Y: y,
		Frame:    0,
		MaxFrame: 30,
		Active:   true,
	}
}

// Update обновляет взрыв
func (e *Explosion) Update() {
	e.Frame++
	if e.Frame >= e.MaxFrame {
		e.Active = false
	}
}

// GetSpriteIndex получает индекс спрайта
func (e *Explosion) GetSpriteIndex() int {
	return (e.Frame / 8) % 4
}

// Enemy враг
type Enemy struct {
	X, Y        float64
	GridX, GridY int
	Type        string
	Speed       float64
	DirX, DirY  float64
	ChangeTimer int
	Active      bool
	spriteStore *sprite.SpriteStore
}

// NewEnemy создаёт врага
func NewEnemy(x, y int, enemyType string, ss *sprite.SpriteStore) *Enemy {
	e := &Enemy{
		X: float64(x) * 32,
		Y: float64(y) * 32,
		GridX: x,
		GridY: y,
		Type: enemyType,
		Active: true,
		ChangeTimer: 60,
		spriteStore: ss,
	}

	// Настройка скорости по типу
	switch enemyType {
	case "balloon":
		e.Speed = 2.0
	case "ghost":
		e.Speed = 1.5
	case "slime":
		e.Speed = 1.0
	}

	// Случайное направление
	e.randomDirection()
	return e
}

// Update обновляет врага
func (e *Enemy) Update() {
	e.X += e.DirX * e.Speed
	e.Y += e.DirY * e.Speed

	e.GridX = int(math.Round(e.X / 32))
	e.GridY = int(math.Round(e.Y / 32))

	e.ChangeTimer--
	if e.ChangeTimer <= 0 {
		e.randomDirection()
		e.ChangeTimer = 60 + int(math.Abs(float64(e.GridX*e.GridY)))%60
	}
}

// RandomDirection случайное направление (экспортировано)
func (e *Enemy) RandomDirection() {
	e.randomDirection()
}

func (e *Enemy) randomDirection() {
	directions := []struct{ dx, dy float64 }{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
	dir := directions[int(math.Abs(float64(e.GridX+e.GridY)))%4]
	e.DirX = dir.dx
	e.DirY = dir.dy
}

// GetSprite получает спрайт врага
func (e *Enemy) GetSprite() *ebiten.Image {
	return e.spriteStore.Enemies[e.Type]
}

// Draw рисует врага
func (e *Enemy) Draw(screen *ebiten.Image, camX, camY float64) {
	spriteImg := e.GetSprite()
	if spriteImg != nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(e.X-camX, e.Y-camY)
		screen.DrawImage(spriteImg, opts)
	}
}

// PowerUp улучшение
type PowerUp struct {
	X, Y   int
	Type   string
	Active bool
}

// NewPowerUp создаёт улучшение
func NewPowerUp(x, y int, powerUpType string) *PowerUp {
	return &PowerUp{
		X: x, Y: y,
		Type:   powerUpType,
		Active: true,
	}
}
