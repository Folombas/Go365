// Package level - генерация уровней для Bomberman
// Go365 Day 95 - Bomberman
package level

import (
	"math/rand"
)

const (
	TileEmpty     = 0
	TileSolid     = 1
	TileBreakable = 2
)

// Level данные уровня
type Level struct {
	Grid      [][]int
	Width     int
	Height    int
	PlayerX   int
	PlayerY   int
	Enemies   []EnemySpawn
	PowerUps  []PowerUpSpawn
}

// EnemySpawn точка появления врага
type EnemySpawn struct {
	X, Y int
	Type string
}

// PowerUpSpawn точка появления улучшения
type PowerUpSpawn struct {
	X, Y int
	Type string
}

// GenerateLevel генерирует уровень
func GenerateLevel(levelNum int, rng *rand.Rand) *Level {
	// Размер поля увеличивается с уровнем
	baseSize := 11 + levelNum*2
	if baseSize > 21 {
		baseSize = 21
	}

	l := &Level{
		Width:  baseSize,
		Height: baseSize,
		Grid:   make([][]int, baseSize),
	}

	// Инициализация сетки
	for y := 0; y < baseSize; y++ {
		l.Grid[y] = make([]int, baseSize)
	}

	// Генерация
	l.generateBorders()
	l.generateSolidWalls(rng)
	l.generateBreakableWalls(rng)
	l.ensurePlayerPath(rng)
	l.spawnPlayer()
	l.spawnEnemies(rng, levelNum)
	l.spawnPowerUps(rng, levelNum)

	return l
}

func (l *Level) generateBorders() {
	for x := 0; x < l.Width; x++ {
		l.Grid[0][x] = TileSolid
		l.Grid[l.Height-1][x] = TileSolid
	}
	for y := 0; y < l.Height; y++ {
		l.Grid[y][0] = TileSolid
		l.Grid[y][l.Width-1] = TileSolid
	}
}

func (l *Level) generateSolidWalls(rng *rand.Rand) {
	// Фиксированная сетка твёрдых стен через каждые 2 клетки
	for y := 1; y < l.Height-1; y += 2 {
		for x := 1; x < l.Width-1; x += 2 {
			// Не ставим стены в зоне игрока
			if x <= 2 && y <= 2 {
				continue
			}
			l.Grid[y][x] = TileSolid
		}
	}
}

func (l *Level) generateBreakableWalls(rng *rand.Rand) {
	for y := 1; y < l.Height-1; y++ {
		for x := 1; x < l.Width-1; x++ {
			// Не ставим разрушимые стены на твёрдых стенах и в зоне игрока
			if l.Grid[y][x] == TileSolid {
				continue
			}
			if x <= 2 && y <= 2 {
				continue
			}

			// 40% шанс появления разрушимой стены
			if rng.Intn(100) < 40 {
				l.Grid[y][x] = TileBreakable
			}
		}
	}
}

func (l *Level) ensurePlayerPath(rng *rand.Rand) {
	// Гарантируем, что зона игрока свободна
	for y := 0; y <= 3; y++ {
		for x := 0; x <= 3; x++ {
			if l.Grid[y][x] != TileSolid {
				l.Grid[y][x] = TileEmpty
			}
		}
	}
}

func (l *Level) spawnPlayer() {
	l.PlayerX = 1
	l.PlayerY = 1
}

func (l *Level) spawnEnemies(rng *rand.Rand, levelNum int) {
	enemyTypes := []string{"balloon", "ghost", "slime"}
	enemyCount := 3 + levelNum*2

	for i := 0; i < enemyCount; i++ {
		for attempts := 0; attempts < 100; attempts++ {
			x := 4 + rng.Intn(l.Width-6)
			y := 4 + rng.Intn(l.Height-6)

			if l.Grid[y][x] == TileEmpty {
				// Проверка расстояния от игрока
				dist := abs(x-l.PlayerX) + abs(y-l.PlayerY)
				if dist >= 5 {
					enemyType := enemyTypes[rng.Intn(len(enemyTypes))]
					l.Enemies = append(l.Enemies, EnemySpawn{
						X: x, Y: y, Type: enemyType,
					})
					break
				}
			}
		}
	}
}

func (l *Level) spawnPowerUps(rng *rand.Rand, levelNum int) {
	powerUpTypes := []string{"bomb_up", "flame_up", "speed_up"}
	powerUpCount := 3 + levelNum

	for i := 0; i < powerUpCount; i++ {
		for attempts := 0; attempts < 100; attempts++ {
			x := 1 + rng.Intn(l.Width-2)
			y := 1 + rng.Intn(l.Height-2)

			if l.Grid[y][x] == TileEmpty {
				powerUpType := powerUpTypes[rng.Intn(len(powerUpTypes))]
				l.PowerUps = append(l.PowerUps, PowerUpSpawn{
					X: x, Y: y, Type: powerUpType,
				})
				break
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsWalkable проверяет, можно ли пройти через клетку
func (l *Level) IsWalkable(x, y int) bool {
	if x < 0 || x >= l.Width || y < 0 || y >= l.Height {
		return false
	}
	return l.Grid[y][x] == TileEmpty
}

// BreakWall разрушает стену
func (l *Level) BreakWall(x, y int) {
	if x >= 0 && x < l.Width && y >= 0 && y < l.Height {
		if l.Grid[y][x] == TileBreakable {
			l.Grid[y][x] = TileEmpty
		}
	}
}

// IsBreakable проверяет, является ли стена разрушимой
func (l *Level) IsBreakable(x, y int) bool {
	if x < 0 || x >= l.Width || y < 0 || y >= l.Height {
		return false
	}
	return l.Grid[y][x] == TileBreakable
}
