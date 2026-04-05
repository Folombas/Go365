// Package puzzle - логика пазла (sliding puzzle 15)
// Go365 Day 95 - Puzzle
package puzzle

import (
	"math/rand"
)

// Puzzle пазл 4x4
type Puzzle struct {
	Grid    [4][4]int
	EmptyR  int
	EmptyC  int
	Moves   int
	Solved  bool
}

// NewPuzzle создаёт новый пазл
func NewPuzzle() *Puzzle {
	p := &Puzzle{
		EmptyR: 3,
		EmptyC: 3,
	}
	p.setupSolved()
	return p
}

// setupSolved расставляет пазл в решённое состояние
func (p *Puzzle) setupSolved() {
	num := 1
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if r == 3 && c == 3 {
				p.Grid[r][c] = 0
			} else {
				p.Grid[r][c] = num
				num++
			}
		}
	}
	p.EmptyR = 3
	p.EmptyC = 3
	p.Moves = 0
	p.Solved = false
}

// Shuffle перемешивает пазл
func (p *Puzzle) Shuffle(moves int, rng *rand.Rand) {
	// Делаем случайные ходы из решённого состояния
	// Это гарантирует решаемость
	directions := []struct{ dr, dc int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	lastDir := -1

	for i := 0; i < moves; i++ {
		validDirs := make([]int, 0)

		for dirIdx, dir := range directions {
			// Не ходить обратно сразу
			if dirIdx == lastDir {
				continue
			}

			nr := p.EmptyR + dir.dr
			nc := p.EmptyC + dir.dc

			if nr >= 0 && nr < 4 && nc >= 0 && nc < 4 {
				validDirs = append(validDirs, dirIdx)
			}
		}

		if len(validDirs) == 0 {
			continue
		}

		dirIdx := validDirs[rng.Intn(len(validDirs))]
		dir := directions[dirIdx]

		nr := p.EmptyR + dir.dr
		nc := p.EmptyC + dir.dc

		// Меняем местами
		p.Grid[p.EmptyR][p.EmptyC] = p.Grid[nr][nc]
		p.Grid[nr][nc] = 0
		p.EmptyR = nr
		p.EmptyC = nc

		// Обратное направление для lastDir
		opposite := map[int]int{0: 1, 1: 0, 2: 3, 3: 2}
		lastDir = opposite[dirIdx]
	}

	p.Moves = 0
	p.Solved = false
}

// CanMove проверяет, можно ли сдвинуть плитку
func (p *Puzzle) CanMove(row, col int) bool {
	// Проверяем, соседняя ли она с пустой
	dr := row - p.EmptyR
	dc := col - p.EmptyC

	return (dr == 1 && dc == 0) || (dr == -1 && dc == 0) ||
		(dr == 0 && dc == 1) || (dr == 0 && dc == -1)
}

// Move сдвигает плитку
func (p *Puzzle) Move(row, col int) bool {
	if !p.CanMove(row, col) {
		return false
	}

	// Меняем местами
	p.Grid[p.EmptyR][p.EmptyC] = p.Grid[row][col]
	p.Grid[row][col] = 0
	p.EmptyR = row
	p.EmptyC = col
	p.Moves++

	// Проверка на решение
	p.Solved = p.checkSolved()

	return true
}

// checkSolved проверяет, решён ли пазл
func (p *Puzzle) checkSolved() bool {
	num := 1
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if r == 3 && c == 3 {
				if p.Grid[r][c] != 0 {
					return false
				}
			} else {
				if p.Grid[r][c] != num {
					return false
				}
				num++
			}
		}
	}
	return true
}

// GetTile получает значение плитки
func (p *Puzzle) GetTile(row, col int) int {
	if row < 0 || row >= 4 || col < 0 || col >= 4 {
		return -1
	}
	return p.Grid[row][col]
}

// IsEmpty проверяет, пустая ли клетка
func (p *Puzzle) IsEmpty(row, col int) bool {
	return row == p.EmptyR && col == p.EmptyC
}

// Reset сбрасывает пазл
func (p *Puzzle) Reset() {
	p.setupSolved()
}

// GetGrid получает копию сетки
func (p *Puzzle) GetGrid() [4][4]int {
	return p.Grid
}
