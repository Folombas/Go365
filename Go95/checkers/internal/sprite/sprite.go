// Package sprite - генерация спрайтов для Checkers
// Go365 Day 95 - Checkers
package sprite

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteStore хранит все спрайты
type SpriteStore struct {
	WhitePiece *ebiten.Image
	BlackPiece *ebiten.Image
	WhiteKing  *ebiten.Image
	BlackKing  *ebiten.Image
	LightTile  *ebiten.Image
	DarkTile   *ebiten.Image
	Highlight  *ebiten.Image
	ValidMove  *ebiten.Image
}

// NewSpriteStore создаёт хранилище спрайтов
func NewSpriteStore() *SpriteStore {
	ss := &SpriteStore{}
	ss.generateAll()
	return ss
}

func (ss *SpriteStore) generateAll() {
	ss.WhitePiece = ss.createPieceSprite(color.RGBA{240, 240, 240, 255}, color.RGBA{200, 200, 200, 255}, false)
	ss.BlackPiece = ss.createPieceSprite(color.RGBA{40, 40, 40, 255}, color.RGBA{80, 80, 80, 255}, false)
	ss.WhiteKing = ss.createPieceSprite(color.RGBA{240, 240, 240, 255}, color.RGBA{200, 200, 200, 255}, true)
	ss.BlackKing = ss.createPieceSprite(color.RGBA{40, 40, 40, 255}, color.RGBA{80, 80, 80, 255}, true)
	ss.LightTile = ss.createTileSprite(color.RGBA{240, 210, 170, 255})
	ss.DarkTile = ss.createTileSprite(color.RGBA{120, 80, 50, 255})
	ss.Highlight = ss.createHighlightSprite()
	ss.ValidMove = ss.createValidMoveSprite()
}

func (ss *SpriteStore) createPieceSprite(baseColor, shadowColor color.RGBA, isKing bool) *ebiten.Image {
	w, h := 60, 60
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	centerX, centerY := 30, 30
	radius := 25

	// Тень
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - (centerY + 3)
			dist := dx*dx + dy*dy
			if dist <= radius*radius && dist > (radius-3)*(radius-3) {
				img.Set(x, y, shadowColor)
			}
		}
	}

	// Основная шашка
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - centerY
			dist := dx*dx + dy*dy

			if dist <= radius*radius {
				// Градиент для объёма
				normalDist := math.Sqrt(float64(dist))
				brightness := 1.0 - normalDist/float64(radius)*0.3

				r := uint8(float64(baseColor.R) * brightness)
				g := uint8(float64(baseColor.G) * brightness)
				b := uint8(float64(baseColor.B) * brightness)

				img.Set(x, y, color.RGBA{r, g, b, baseColor.A})
			}
		}
	}

	// Внутренний круг
	innerRadius := radius - 6
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - centerY
			dist := dx*dx + dy*dy

			if dist <= innerRadius*innerRadius && dist > (innerRadius-2)*(innerRadius-2) {
				img.Set(x, y, shadowColor)
			}
		}
	}

	// Корона для дамки
	if isKing {
		crownColor := color.RGBA{255, 215, 0, 255}
		// Корона (упрощённая)
		for i := 0; i < 5; i++ {
			angle := float64(i) * math.Pi / 6
			x := centerX + int(math.Sin(angle)*12)
			y := centerY - 8 + int(math.Cos(angle)*4)
			if x >= 0 && x < w && y >= 0 && y < h {
				img.Set(x, y, crownColor)
			}
		}
		// Вершины короны
		for i := 0; i < 3; i++ {
			x := centerX - 8 + i*8
			y := centerY - 12
			if x >= 0 && x < w && y >= 0 && y < h {
				img.Set(x, y, crownColor)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createTileSprite(c color.RGBA) *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Лёгкая текстура дерева
			noise := (x*7 + y*13) % 20 - 10
			r := clamp(int(c.R) + noise)
			g := clamp(int(c.G) + noise)
			b := clamp(int(c.B) + noise)
			img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), c.A})
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createHighlightSprite() *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	highlightColor := color.RGBA{100, 255, 100, 100}

	// Полупрозрачная подсветка
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, highlightColor)
		}
	}

	// Рамка
	borderColor := color.RGBA{0, 255, 0, 200}
	for i := 0; i < 80; i++ {
		img.Set(i, 0, borderColor)
		img.Set(i, 79, borderColor)
		img.Set(0, i, borderColor)
		img.Set(79, i, borderColor)
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createValidMoveSprite() *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	dotColor := color.RGBA{100, 255, 100, 150}
	centerX, centerY := 40, 40
	radius := 15

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - centerY
			dist := dx*dx + dy*dy
			if dist <= radius*radius {
				img.Set(x, y, dotColor)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func clamp(val int) int {
	if val < 0 {
		return 0
	}
	if val > 255 {
		return 255
	}
	return val
}
