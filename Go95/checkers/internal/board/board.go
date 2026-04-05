// Package board - логика доски и игры в шашки
// Go365 Day 95 - Checkers
package board

import "fmt"

// PieceType тип шашки
type PieceType int

const (
	Empty PieceType = iota
	White
	Black
	WhiteKing
	BlackKing
)

// Board доска 8x8
type Board struct {
	Grid [8][8]PieceType
}

// NewBoard создаёт новую доску
func NewBoard() *Board {
	b := &Board{}
	b.setupInitial()
	return b
}

// setupInitial расставляет шашки
func (b *Board) setupInitial() {
	// Чёрные шашки (ряды 0-2)
	for row := 0; row < 3; row++ {
		for col := 0; col < 8; col++ {
			if (row+col)%2 == 1 {
				b.Grid[row][col] = Black
			}
		}
	}

	// Белые шашки (ряды 5-7)
	for row := 5; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if (row+col)%2 == 1 {
				b.Grid[row][col] = White
			}
		}
	}
}

// Get получает шашку
func (b *Board) Get(row, col int) PieceType {
	if row < 0 || row >= 8 || col < 0 || col >= 8 {
		return Empty
	}
	return b.Grid[row][col]
}

// Set ставит шашку
func (b *Board) Set(row, col int, piece PieceType) {
	if row >= 0 && row < 8 && col >= 0 && col < 8 {
		b.Grid[row][col] = piece
	}
}

// IsWhite проверяет, белая ли шашка
func IsWhite(p PieceType) bool {
	return p == White || p == WhiteKing
}

// IsBlack проверяет, чёрная ли шашка
func IsBlack(p PieceType) bool {
	return p == Black || p == BlackKing
}

// IsKing проверяет, дамка ли
func IsKing(p PieceType) bool {
	return p == WhiteKing || p == BlackKing
}

// IsOwnPiece проверяет, своя ли шашка
func IsOwnPiece(p PieceType, isWhiteTurn bool) bool {
	if isWhiteTurn {
		return IsWhite(p)
	}
	return IsBlack(p)
}

// IsEnemyPiece проверяет, вражеская ли шашка
func IsEnemyPiece(p PieceType, isWhiteTurn bool) bool {
	if p == Empty {
		return false
	}
	if isWhiteTurn {
		return IsBlack(p)
	}
	return IsWhite(p)
}

// Move ход
type Move struct {
	FromRow, FromCol int
	ToRow, ToCol     int
	CapturedRow      int // -1 если нет взятия
	CapturedCol      int
}

// String строковое представление
func (m Move) String() string {
	if m.CapturedRow >= 0 {
		return fmt.Sprintf("(%d,%d)->(%d,%d) x(%d,%d)",
			m.FromRow, m.FromCol, m.ToRow, m.ToCol, m.CapturedRow, m.CapturedCol)
	}
	return fmt.Sprintf("(%d,%d)->(%d,%d)", m.FromRow, m.FromCol, m.ToRow, m.ToCol)
}

// GetValidMoves получает все допустимые ходы для игрока
func (b *Board) GetValidMoves(isWhiteTurn bool) []Move {
	moves := make([]Move, 0)

	// Сначала ищем обязательные взятия
	captures := b.GetCaptures(isWhiteTurn)
	if len(captures) > 0 {
		return captures
	}

	// Обычные ходы
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := b.Grid[row][col]
			if IsOwnPiece(piece, isWhiteTurn) {
				moves = append(moves, b.GetPieceMoves(row, col, piece)...)
			}
		}
	}

	return moves
}

// GetCaptures получает все возможные взятия
func (b *Board) GetCaptures(isWhiteTurn bool) []Move {
	captures := make([]Move, 0)

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := b.Grid[row][col]
			if IsOwnPiece(piece, isWhiteTurn) {
				captures = append(captures, b.GetPieceCaptures(row, col, piece)...)
			}
		}
	}

	return captures
}

// GetPieceMoves ходы конкретной шашки
func (b *Board) GetPieceMoves(row, col int, piece PieceType) []Move {
	moves := make([]Move, 0)

	if IsKing(piece) {
		// Дамка ходит на любое количество клеток
		directions := []struct{ dr, dc int }{
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}

		for _, dir := range directions {
			for i := 1; i < 8; i++ {
				nr := row + dir.dr*i
				nc := col + dir.dc*i

				if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
					break
				}

				if b.Grid[nr][nc] != Empty {
					break
				}

				moves = append(moves, Move{
					FromRow: row, FromCol: col,
					ToRow: nr, ToCol: nc,
					CapturedRow: -1, CapturedCol: -1,
				})
			}
		}
	} else {
		// Обычная шашка ходит на 1 клетку вперёд
		directions := []struct{ dr, dc int }{}

		if IsWhite(piece) {
			directions = append(directions, struct{ dr, dc int }{-1, -1}, struct{ dr, dc int }{-1, 1})
		} else {
			directions = append(directions, struct{ dr, dc int }{1, -1}, struct{ dr, dc int }{1, 1})
		}

		for _, dir := range directions {
			nr := row + dir.dr
			nc := col + dir.dc

			if nr >= 0 && nr < 8 && nc >= 0 && nc < 8 && b.Grid[nr][nc] == Empty {
				moves = append(moves, Move{
					FromRow: row, FromCol: col,
					ToRow: nr, ToCol: nc,
					CapturedRow: -1, CapturedCol: -1,
				})
			}
		}
	}

	return moves
}

// GetPieceCaptures взятия конкретной шашки
func (b *Board) GetPieceCaptures(row, col int, piece PieceType) []Move {
	captures := make([]Move, 0)

	if IsKing(piece) {
		// Дамка бьёт через всё поле
		directions := []struct{ dr, dc int }{
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}

		for _, dir := range directions {
			enemyFound := false
			enemyRow, enemyCol := -1, -1

			for i := 1; i < 8; i++ {
				nr := row + dir.dr*i
				nc := col + dir.dc*i

				if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
					break
				}

				target := b.Grid[nr][nc]

				if !enemyFound {
					if IsEnemyPiece(target, IsWhite(piece)) {
						enemyFound = true
						enemyRow = nr
						enemyCol = nc
					} else if IsOwnPiece(target, IsWhite(piece)) {
						break
					}
				} else {
					if target == Empty {
						captures = append(captures, Move{
							FromRow: row, FromCol: col,
							ToRow: nr, ToCol: nc,
							CapturedRow: enemyRow, CapturedCol: enemyCol,
						})
					} else {
						break
					}
				}
			}
		}
	} else {
		// Обычная шашка бьёт на 2 клетки
		directions := []struct{ dr, dc int }{
			{-2, -2}, {-2, 2}, {2, -2}, {2, 2},
		}

		for _, dir := range directions {
			nr := row + dir.dr
			nc := col + dir.dc

			if nr < 0 || nr >= 8 || nc < 0 || nc >= 8 {
				continue
			}

			if b.Grid[nr][nc] != Empty {
				continue
			}

			// Проверяем шашку посередине
			midRow := row + dir.dr/2
			midCol := col + dir.dc/2
			midPiece := b.Grid[midRow][midCol]

			if IsEnemyPiece(midPiece, IsWhite(piece)) {
				captures = append(captures, Move{
					FromRow: row, FromCol: col,
					ToRow: nr, ToCol: nc,
					CapturedRow: midRow, CapturedCol: midCol,
				})
			}
		}
	}

	return captures
}

// MakeMove делает ход
func (b *Board) MakeMove(move Move) {
	piece := b.Grid[move.FromRow][move.FromCol]
	b.Grid[move.FromRow][move.FromCol] = Empty
	b.Grid[move.ToRow][move.ToCol] = piece

	// Удаление съеденной шашки
	if move.CapturedRow >= 0 {
		b.Grid[move.CapturedRow][move.CapturedCol] = Empty
	}

	// Превращение в дамку
	if piece == White && move.ToRow == 0 {
		b.Grid[move.ToRow][move.ToCol] = WhiteKing
	}
	if piece == Black && move.ToRow == 7 {
		b.Grid[move.ToRow][move.ToCol] = BlackKing
	}
}

// CountPieces считает шашки
func (b *Board) CountPieces() (whiteCount, blackCount int) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := b.Grid[row][col]
			if IsWhite(piece) {
				whiteCount++
			}
			if IsBlack(piece) {
				blackCount++
			}
		}
	}
	return
}

// IsGameOver проверяет, окончена ли игра
func (b *Board) IsGameOver(isWhiteTurn bool) bool {
	whiteCount, blackCount := b.CountPieces()
	if whiteCount == 0 || blackCount == 0 {
		return true
	}

	// Нет доступных ходов
	moves := b.GetValidMoves(isWhiteTurn)
	return len(moves) == 0
}

// GetWinner определяет победителя
func (b *Board) GetWinner() string {
	whiteCount, blackCount := b.CountPieces()
	if whiteCount == 0 {
		return "black"
	}
	if blackCount == 0 {
		return "white"
	}
	return ""
}

// Clone клонирует доску
func (b *Board) Clone() *Board {
	clone := &Board{}
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			clone.Grid[row][col] = b.Grid[row][col]
		}
	}
	return clone
}
