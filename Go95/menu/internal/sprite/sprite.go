// Package sprite - генерация спрайтов для Menu
// Go365 Day 95 - Menu
package sprite

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteStore хранит все спрайты
type SpriteStore struct {
	BombermanIcon *ebiten.Image
	CheckersIcon  *ebiten.Image
	PuzzleIcon    *ebiten.Image
	ButtonBg      *ebiten.Image
	ButtonHover   *ebiten.Image
	Background    *ebiten.Image
	TitleText     *ebiten.Image
}

// NewSpriteStore создаёт хранилище спрайтов
func NewSpriteStore() *SpriteStore {
	ss := &SpriteStore{}
	ss.generateAll()
	return ss
}

func (ss *SpriteStore) generateAll() {
	ss.BombermanIcon = ss.createBombermanIcon()
	ss.CheckersIcon = ss.createCheckersIcon()
	ss.PuzzleIcon = ss.createPuzzleIcon()
	ss.ButtonBg = ss.createButtonBg()
	ss.ButtonHover = ss.createButtonHover()
	ss.Background = ss.createBackground()
}

func (ss *SpriteStore) createBombermanIcon() *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Бомба
	bombColor := color.RGBA{30, 30, 30, 255}
	fuseColor := color.RGBA{255, 200, 50, 255}
	fireColor := color.RGBA{255, 100, 30, 255}

	// Круглая бомба
	centerX, centerY := 40, 45
	radius := 25
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - centerY
			dist := dx*dx + dy*dy
			if dist <= radius*radius {
				img.Set(x, y, bombColor)
			}
		}
	}

	// Блик
	for y := 25; y < 35; y++ {
		for x := 28; x < 38; x++ {
			dx := x - centerX
			dy := y - centerY
			if dx*dx+dy*dy <= radius*radius {
				img.Set(x, y, color.RGBA{80, 80, 80, 255})
			}
		}
	}

	// Фитиль
	img.Set(40, 18, fuseColor)
	img.Set(39, 16, fuseColor)
	img.Set(41, 16, fuseColor)

	// Огонь
	img.Set(40, 14, fireColor)
	img.Set(38, 12, fireColor)
	img.Set(42, 12, fireColor)
	img.Set(40, 10, color.RGBA{255, 255, 100, 255})

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createCheckersIcon() *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Доска 4x4
	tileSize := 15
	offsetX, offsetY := 10, 10

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			x := offsetX + col*tileSize
			y := offsetY + row*tileSize

			if (row+col)%2 == 0 {
				for yy := y; yy < y+tileSize; yy++ {
					for xx := x; xx < x+tileSize; xx++ {
						img.Set(xx, yy, color.RGBA{240, 210, 170, 255})
					}
				}
			} else {
				for yy := y; yy < y+tileSize; yy++ {
					for xx := x; xx < x+tileSize; xx++ {
						img.Set(xx, yy, color.RGBA{120, 80, 50, 255})
					}
				}
			}
		}
	}

	// Шашки
	// Белая
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - 25
			dy := y - 20
			if dx*dx+dy*dy <= 49 {
				img.Set(x, y, color.RGBA{240, 240, 240, 255})
			}
		}
	}

	// Чёрная
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - 55
			dy := y - 60
			if dx*dx+dy*dy <= 49 {
				img.Set(x, y, color.RGBA{40, 40, 40, 255})
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createPuzzleIcon() *ebiten.Image {
	w, h := 80, 80
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Плитки 3x3
	tileSize := 22
	offsetX, offsetY := 8, 8

	numbers := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	hues := []float64{0, 30, 60, 120, 180, 240, 270, 300}
	hueIdx := 0

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			x := offsetX + col*tileSize
			y := offsetY + row*tileSize
			num := numbers[row][col]

			if num == 0 {
				// Пустая клетка
				for yy := y; yy < y+tileSize; yy++ {
					for xx := x; xx < x+tileSize; xx++ {
						img.Set(xx, yy, color.RGBA{60, 60, 70, 255})
					}
				}
			} else {
				// Цветная плитка
				h := hues[hueIdx%len(hues)]
				c := hslToRGB(h, 0.6, 0.5)
				for yy := y; yy < y+tileSize; yy++ {
					for xx := x; xx < x+tileSize; xx++ {
						img.Set(xx, yy, c)
					}
				}
				hueIdx++
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createButtonBg() *ebiten.Image {
	w, h := 300, 100
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Фон кнопки
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Градиент
			ratio := float64(y) / float64(h)
			r := uint8(60 + ratio*20)
			g := uint8(70 + ratio*15)
			b := uint8(100 + ratio*20)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Рамка
	borderColor := color.RGBA{100, 120, 160, 255}
	for i := 0; i < w; i++ {
		img.Set(i, 0, borderColor)
		img.Set(i, h-1, borderColor)
	}
	for i := 0; i < h; i++ {
		img.Set(0, i, borderColor)
		img.Set(w-1, i, borderColor)
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createButtonHover() *ebiten.Image {
	w, h := 300, 100
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Фон кнопки (ярче)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ratio := float64(y) / float64(h)
			r := uint8(80 + ratio*30)
			g := uint8(100 + ratio*20)
			b := uint8(140 + ratio*30)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Рамка (ярче)
	borderColor := color.RGBA{150, 180, 255, 255}
	for i := 0; i < w; i++ {
		img.Set(i, 0, borderColor)
		img.Set(i, h-1, borderColor)
	}
	for i := 0; i < h; i++ {
		img.Set(0, i, borderColor)
		img.Set(w-1, i, borderColor)
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createBackground() *ebiten.Image {
	w, h := 1280, 720
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Градиентный фон
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ratio := float64(y) / float64(h)
			r := uint8(30 + ratio*20)
			g := uint8(40 + ratio*15)
			b := uint8(70 + ratio*30)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Декоративные частицы
	for i := 0; i < 50; i++ {
		px := (i * 137) % w
		py := (i * 89) % h
		size := 2 + (i % 3)
		alpha := uint8(50 + (i%50))

		for dy := 0; dy < size; dy++ {
			for dx := 0; dx < size; dx++ {
				if px+dx < w && py+dy < h {
					img.Set(px+dx, py+dy, color.RGBA{100, 150, 255, alpha})
				}
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

// HSL to RGB
func hslToRGB(h, s, l float64) color.RGBA {
	var r, g, b float64

	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		hNorm := h / 360.0
		r = hueToRGB(p, q, hNorm+1.0/3.0)
		g = hueToRGB(p, q, hNorm)
		b = hueToRGB(p, q, hNorm-1.0/3.0)
	}

	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}
