// Package sprite - генерация спрайтов для Puzzle
// Go365 Day 95 - Puzzle
package sprite

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteStore хранит все спрайты
type SpriteStore struct {
	Tiles       map[int]*ebiten.Image
	EmptyTile   *ebiten.Image
	Background  *ebiten.Image
	NumberFont  map[int]*ebiten.Image
}

// NewSpriteStore создаёт хранилище спрайтов
func NewSpriteStore() *SpriteStore {
	ss := &SpriteStore{
		Tiles:      make(map[int]*ebiten.Image),
		NumberFont: make(map[int]*ebiten.Image),
	}
	ss.generateAll()
	return ss
}

func (ss *SpriteStore) generateAll() {
	ss.generateTiles()
	ss.generateEmptyTile()
	ss.generateBackground()
	ss.generateNumberFont()
}

func (ss *SpriteStore) generateTiles() {
	// Генерируем тайлы для чисел 1-15
	for i := 1; i <= 15; i++ {
		ss.Tiles[i] = ss.createTileSprite(i)
	}
}

func (ss *SpriteStore) createTileSprite(number int) *ebiten.Image {
	w, h := 100, 100
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Цвет зависит от числа (градиент)
	hue := float64(number) * 24 // 360/15 = 24
	baseColor := hslToRGB(hue, 0.6, 0.5)
	lightColor := hslToRGB(hue, 0.6, 0.65)
	darkColor := hslToRGB(hue, 0.6, 0.35)

	// Основной фон
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Градиент для объёма
			distFromCenter := float64(x-w/2)*float64(x-w/2) + float64(y-h/2)*float64(y-h/2)
			maxDist := float64(w/2 * w/2)
			brightness := 1.0 - distFromCenter/maxDist*0.3

			r := uint8(float64(baseColor.R) * brightness)
			g := uint8(float64(baseColor.G) * brightness)
			b := uint8(float64(baseColor.B) * brightness)

			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Рамка
	for i := 0; i < w; i++ {
		img.Set(i, 0, lightColor)
		img.Set(i, 1, lightColor)
		img.Set(i, h-1, darkColor)
		img.Set(i, h-2, darkColor)
	}
	for i := 0; i < h; i++ {
		img.Set(0, i, lightColor)
		img.Set(1, i, lightColor)
		img.Set(w-1, i, darkColor)
		img.Set(w-2, i, darkColor)
	}

	// Число
	ss.drawNumberOnTile(img, number)

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) drawNumberOnTile(img *image.RGBA, number int) {
	numberStr := itoa(number)
	startX := 50 - len(numberStr)*12
	startY := 35

	numberColor := color.RGBA{255, 255, 255, 255}
	shadowColor := color.RGBA{0, 0, 0, 100}

	for _, ch := range numberStr {
		digit := int(ch - '0')
		ss.drawDigit(img, startX, startY, digit, numberColor, shadowColor)
		startX += 24
	}
}

func (ss *SpriteStore) drawDigit(img *image.RGBA, x, y int, digit int, c, shadow color.RGBA) {
	// Простой 5x7 шрифт
	font := getDigitPattern(digit)

	for row := 0; row < 7; row++ {
		for col := 0; col < 5; col++ {
			if font[row][col] == 1 {
				px := x + col*3
				py := y + row*3
				if px >= 0 && px < 95 && py >= 0 && py < 95 {
					img.Set(px, py, c)
					img.Set(px+1, py, c)
					img.Set(px, py+1, c)
					img.Set(px+1, py+1, c)

					// Тень
					img.Set(px+2, py+2, shadow)
				}
			}
		}
	}
}

func getDigitPattern(digit int) [7][5]int {
	patterns := map[int][7][5]int{
		0: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 1, 1},
			{1, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
		1: {
			{0, 0, 1, 0, 0},
			{0, 1, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 1, 1, 0},
		},
		2: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		3: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
		4: {
			{0, 0, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
		},
		5: {
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
		6: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
		7: {
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
		},
		8: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
		9: {
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	}

	return patterns[digit]
}

func (ss *SpriteStore) generateEmptyTile() {
	w, h := 100, 100
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	emptyColor := color.RGBA{60, 60, 70, 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, emptyColor)
		}
	}

	// Внутренняя рамка
	innerColor := color.RGBA{40, 40, 50, 255}
	for y := 5; y < 95; y++ {
		for x := 5; x < 95; x++ {
			img.Set(x, y, innerColor)
		}
	}

	ss.EmptyTile = ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateBackground() {
	w, h := 1280, 720
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Градиентный фон
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ratio := float64(y) / float64(h)
			r := uint8(40 + ratio*20)
			g := uint8(50 + ratio*15)
			b := uint8(80 + ratio*20)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	ss.Background = ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateNumberFont() {
	// Для HUD
	for i := 0; i <= 9; i++ {
		w, h := 20, 24
		img := image.NewRGBA(image.Rect(0, 0, w, h))

		c := color.RGBA{255, 255, 255, 255}
		ss.drawDigitSmall(img, 2, 2, i, c)

		ss.NumberFont[i] = ebiten.NewImageFromImage(img)
	}
}

func (ss *SpriteStore) drawDigitSmall(img *image.RGBA, x, y int, digit int, c color.RGBA) {
	font := getDigitPattern(digit)

	for row := 0; row < 7; row++ {
		for col := 0; col < 5; col++ {
			if font[row][col] == 1 {
				px := x + col*2
				py := y + row*2
				if px >= 0 && px < 18 && py >= 0 && py < 22 {
					img.Set(px, py, c)
					img.Set(px+1, py, c)
					img.Set(px, py+1, c)
					img.Set(px+1, py+1, c)
				}
			}
		}
	}
}

// HSL to RGB конвертация
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

// itoa простое преобразование числа в строку
func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	digits := make([]byte, 0)
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}

	return string(digits)
}
