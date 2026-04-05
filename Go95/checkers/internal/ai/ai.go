// Package ai - искусственный интеллект для Checkers
// Go365 Day 95 - Checkers
package ai

import (
	"math/rand"
	"checkers/internal/board"
)

// AI искусственный интеллект
type AI struct {
	depth     int
	rng       *rand.Rand
	isWhiteAI bool
}

// NewAI создаёт AI
func NewAI(isWhiteAI bool, depth int) *AI {
	return &AI{
		depth:     depth,
		rng:       rand.New(rand.NewSource(42)),
		isWhiteAI: isWhiteAI,
	}
}

// GetMove получает лучший ход
func (ai *AI) GetMove(b *board.Board) board.Move {
	moves := b.GetValidMoves(ai.isWhiteAI)

	if len(moves) == 0 {
		return board.Move{}
	}

	// Если есть обязательные взятия, выбираем из них
	hasCaptures := false
	for _, m := range moves {
		if m.CapturedRow >= 0 {
			hasCaptures = true
			break
		}
	}

	if hasCaptures {
		// Фильтруем только взятия
		captures := make([]board.Move, 0)
		for _, m := range moves {
			if m.CapturedRow >= 0 {
				captures = append(captures, m)
			}
		}

		// Выбираем лучшее взятие
		bestMove := captures[0]
		bestScore := -9999

		for _, m := range captures {
			clone := b.Clone()
			clone.MakeMove(m)
			score := ai.minimax(clone, ai.depth-1, false, -9999, 9999)
			if score > bestScore {
				bestScore = score
				bestMove = m
			}
		}

		return bestMove
	}

	// Обычный выбор через minimax
	bestMove := moves[0]
	bestScore := -9999

	for _, m := range moves {
		clone := b.Clone()
		clone.MakeMove(m)
		score := ai.minimax(clone, ai.depth-1, false, -9999, 9999)
		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	return bestMove
}

// minimax алгоритм минимакс с альфа-бета отсечением
func (ai *AI) minimax(b *board.Board, depth int, isMaximizing bool, alpha, beta int) int {
	if depth == 0 || b.IsGameOver(!ai.isWhiteAI) {
		return ai.evaluate(b)
	}

	var moves []board.Move
	if isMaximizing {
		moves = b.GetValidMoves(ai.isWhiteAI)
	} else {
		moves = b.GetValidMoves(!ai.isWhiteAI)
	}

	if len(moves) == 0 {
		if isMaximizing {
			return -1000 // Нет ходов - проигрыш
		}
		return 1000
	}

	if isMaximizing {
		maxEval := -9999
		for _, m := range moves {
			clone := b.Clone()
			clone.MakeMove(m)
			eval := ai.minimax(clone, depth-1, false, alpha, beta)
			if eval > maxEval {
				maxEval = eval
			}
			if eval > alpha {
				alpha = eval
			}
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := 9999
		for _, m := range moves {
			clone := b.Clone()
			clone.MakeMove(m)
			eval := ai.minimax(clone, depth-1, true, alpha, beta)
			if eval < minEval {
				minEval = eval
			}
			if eval < beta {
				beta = eval
			}
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

// evaluate оценивает позицию
func (ai *AI) evaluate(b *board.Board) int {
	score := 0

	whiteCount, blackCount := b.CountPieces()

	if ai.isWhiteAI {
		score = whiteCount - blackCount
	} else {
		score = blackCount - whiteCount
	}

	// Бонус за дамки
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := b.Get(row, col)
			if board.IsKing(piece) {
				if ai.isWhiteAI && board.IsWhite(piece) {
					score += 2
				} else if !ai.isWhiteAI && board.IsBlack(piece) {
					score += 2
				} else {
					score -= 2
				}
			}
		}
	}

	// Бонус за продвижение вперёд
	if ai.isWhiteAI {
		// Белые идут вверх (к row=0)
		for row := 0; row < 8; row++ {
			for col := 0; col < 8; col++ {
				piece := b.Get(row, col)
				if piece == board.White {
					score += (7 - row) // Чем ближе к дамке, тем лучше
				}
			}
		}
	} else {
		// Чёрные идут вниз (к row=7)
		for row := 0; row < 8; row++ {
			for col := 0; col < 8; col++ {
				piece := b.Get(row, col)
				if piece == board.Black {
					score += row
				}
			}
		}
	}

	// Бонус за центр
	for row := 2; row < 6; row++ {
		for col := 2; col < 6; col++ {
			piece := b.Get(row, col)
			if piece != board.Empty {
				if ai.isWhiteAI && board.IsWhite(piece) {
					score += 1
				} else if !ai.isWhiteAI && board.IsBlack(piece) {
					score += 1
				}
			}
		}
	}

	return score
}
