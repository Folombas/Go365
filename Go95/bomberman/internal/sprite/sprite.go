// Package sprite - генерация и загрузка спрайтов для Bomberman
// Go365 Day 95 - Bomberman
package sprite

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// SpriteStore хранит все спрайты
type SpriteStore struct {
	Bomberman    map[string]*ebiten.Image
	Walls        map[string]*ebiten.Image
	Bombs        map[string]*ebiten.Image
	Explosions   []*ebiten.Image
	PowerUps     map[string]*ebiten.Image
	Enemies      map[string]*ebiten.Image
	Menu         *ebiten.Image
}

// NewSpriteStore создаёт хранилище спрайтов
func NewSpriteStore() *SpriteStore {
	ss := &SpriteStore{
		Bomberman:  make(map[string]*ebiten.Image),
		Walls:      make(map[string]*ebiten.Image),
		Bombs:      make(map[string]*ebiten.Image),
		PowerUps:   make(map[string]*ebiten.Image),
		Enemies:    make(map[string]*ebiten.Image),
		Explosions: make([]*ebiten.Image, 4),
	}

	ss.generateAll()
	return ss
}

func (ss *SpriteStore) generateAll() {
	ss.generateBomberman()
	ss.generateWalls()
	ss.generateBombs()
	ss.generateExplosions()
	ss.generatePowerUps()
	ss.generateEnemies()
}

func (ss *SpriteStore) generateBomberman() {
	// 4 направления: up, down, left, right
	directions := []string{"up", "down", "left", "right"}
	for _, dir := range directions {
		ss.Bomberman[dir] = ss.createBombermanSprite(dir)
	}
}

func (ss *SpriteStore) createBombermanSprite(direction string) *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Цвета
	helmetColor := color.RGBA{255, 100, 100, 255}
	skinColor := color.RGBA{255, 220, 180, 255}
	bodyColor := color.RGBA{50, 100, 255, 255}
	bootColor := color.RGBA{80, 60, 40, 255}
	eyeColor := color.RGBA{30, 30, 30, 255}

	// Шлем (верхняя часть)
	for y := 2; y < 12; y++ {
		for x := 8; x < 24; x++ {
			img.Set(x, y, helmetColor)
		}
	}

	// Лицо
	for y := 10; y < 18; y++ {
		for x := 10; x < 22; x++ {
			img.Set(x, y, skinColor)
		}
	}

	// Глаза
	eyeY := 13
	switch direction {
	case "up":
		// Глаза не видны, виден затылок
		for y := 10; y < 18; y++ {
			for x := 10; x < 22; x++ {
				img.Set(x, y, helmetColor)
			}
		}
	case "down":
		img.Set(13, eyeY, eyeColor)
		img.Set(18, eyeY, eyeColor)
	case "left":
		img.Set(12, eyeY, eyeColor)
	case "right":
		img.Set(19, eyeY, eyeColor)
	}

	// Тело
	for y := 18; y < 26; y++ {
		for x := 8; x < 24; x++ {
			img.Set(x, y, bodyColor)
		}
	}

	// Ноги
	for y := 26; y < 32; y++ {
		for x := 10; x < 16; x++ {
			img.Set(x, y, bootColor)
		}
		for x := 16; x < 22; x++ {
			img.Set(x, y, bootColor)
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateWalls() {
	// Разрушимая стена
	ss.Walls["breakable"] = ss.createBreakableWall()
	// Неразрушимая стена
	ss.Walls["solid"] = ss.createSolidWall()
	// Пол
	ss.Walls["floor"] = ss.createFloor()
}

func (ss *SpriteStore) createBreakableWall() *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	brickColor := color.RGBA{180, 80, 60, 255}
	mortarColor := color.RGBA{150, 150, 150, 255}

	// Кирпичи
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			row := y / 8
			offset := (row % 2) * 8
			if (x+offset)%16 < 14 && y%8 < 6 {
				img.Set(x, y, brickColor)
			} else {
				img.Set(x, y, mortarColor)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createSolidWall() *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	stoneColor := color.RGBA{100, 100, 110, 255}
	darkColor := color.RGBA{80, 80, 90, 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (x+y)%5 < 3 {
				img.Set(x, y, stoneColor)
			} else {
				img.Set(x, y, darkColor)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createFloor() *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	floorColor := color.RGBA{120, 140, 80, 255}
	darkFloor := color.RGBA{100, 120, 70, 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (x/8+y/8)%2 == 0 {
				img.Set(x, y, floorColor)
			} else {
				img.Set(x, y, darkFloor)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateBombs() {
	ss.Bombs["bomb_0"] = ss.createBombSprite(0)
	ss.Bombs["bomb_1"] = ss.createBombSprite(1)
	ss.Bombs["bomb_2"] = ss.createBombSprite(2)
}

func (ss *SpriteStore) createBombSprite(frame int) *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	bombColor := color.RGBA{30, 30, 30, 255}
	highlightColor := color.RGBA{80, 80, 80, 255}
	fuseColor := color.RGBA{200, 150, 50, 255}

	// Пульсация
	pulse := frame % 2

	// Бомба (круг)
	centerX, centerY := 16, 18
	radius := 10 + pulse
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dx := x - centerX
			dy := y - centerY
			dist := dx*dx + dy*dy
			if dist <= radius*radius {
				if dist > (radius-2)*(radius-2) {
					img.Set(x, y, highlightColor)
				} else {
					img.Set(x, y, bombColor)
				}
			}
		}
	}

	// Фитиль
	img.Set(16, 6, fuseColor)
	img.Set(15, 5, fuseColor)
	img.Set(17, 5, fuseColor)
	img.Set(16, 4, color.RGBA{255, 200, 50, 255})

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateExplosions() {
	explosionColors := [][]color.RGBA{
		{{255, 255, 200, 255}, {255, 200, 50, 255}, {255, 100, 30, 255}, {200, 50, 20, 255}},
		{{255, 255, 255, 255}, {255, 255, 100, 255}, {255, 150, 50, 255}, {255, 80, 30, 255}},
		{{255, 255, 255, 255}, {255, 255, 200, 255}, {255, 200, 100, 255}, {255, 120, 50, 255}},
		{{255, 255, 255, 200}, {255, 255, 255, 150}, {255, 200, 150, 100}, {200, 150, 100, 50}},
	}

	for i := 0; i < 4; i++ {
		w, h := 32, 32
		img := image.NewRGBA(image.Rect(0, 0, w, h))
		centerX, centerY := 16, 16

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				dx := x - centerX
				dy := y - centerY
				dist := dx*dx + dy*dy
				maxDist := (12 + i*2) * (12 + i*2)

				if dist <= maxDist {
					colorIdx := 0
					if dist > maxDist*3/4 {
						colorIdx = 0
					} else if dist > maxDist/2 {
						colorIdx = 1
					} else if dist > maxDist/4 {
						colorIdx = 2
					} else {
						colorIdx = 3
					}
					img.Set(x, y, explosionColors[i][colorIdx])
				}
			}
		}

		ss.Explosions[i] = ebiten.NewImageFromImage(img)
	}
}

func (ss *SpriteStore) generatePowerUps() {
	ss.PowerUps["bomb_up"] = ss.createBombUpSprite()
	ss.PowerUps["flame_up"] = ss.createFlameUpSprite()
	ss.PowerUps["speed_up"] = ss.createSpeedUpSprite()
}

func (ss *SpriteStore) createBombUpSprite() *ebiten.Image {
	w, h := 24, 24
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	bombColor := color.RGBA{30, 30, 30, 255}
	fuseColor := color.RGBA{255, 200, 50, 255}

	// Маленькая бомба
	for y := 6; y < 20; y++ {
		for x := 4; x < 20; x++ {
			dx := x - 12
			dy := y - 14
			if dx*dx+dy*dy <= 36 {
				img.Set(x, y, bombColor)
			}
		}
	}

	// Фитиль
	img.Set(12, 4, fuseColor)
	img.Set(12, 3, color.RGBA{255, 100, 50, 255})

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createFlameUpSprite() *ebiten.Image {
	w, h := 24, 24
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	flameColors := []color.RGBA{
		{255, 255, 200, 255},
		{255, 200, 50, 255},
		{255, 100, 30, 255},
	}

	for y := 2; y < 22; y++ {
		for x := 4; x < 20; x++ {
			dist := 22 - y
			if dist > 0 && dist <= 20 {
				idx := (dist - 1) / 7
				if idx < len(flameColors) {
					img.Set(x, y, flameColors[idx])
				}
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) createSpeedUpSprite() *ebiten.Image {
	w, h := 24, 24
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	shoeColor := color.RGBA{255, 100, 100, 255}
	soleColor := color.RGBA{80, 80, 80, 255}

	// Кроссовок
	for y := 8; y < 18; y++ {
		for x := 4; x < 20; x++ {
			img.Set(x, y, shoeColor)
		}
	}
	for y := 18; y < 22; y++ {
		for x := 4; x < 20; x++ {
			img.Set(x, y, soleColor)
		}
	}

	// Молния
	for i := 0; i < 5; i++ {
		img.Set(8+i, 10+i, color.RGBA{255, 255, 0, 255})
	}

	return ebiten.NewImageFromImage(img)
}

func (ss *SpriteStore) generateEnemies() {
	// 3 типа врагов
	enemyTypes := []string{"balloon", "ghost", "slime"}
	for _, et := range enemyTypes {
		ss.Enemies[et] = ss.createEnemySprite(et)
	}
}

func (ss *SpriteStore) createEnemySprite(enemyType string) *ebiten.Image {
	w, h := 32, 32
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	switch enemyType {
	case "balloon":
		// Воздушный шар
		balloonColor := color.RGBA{255, 50, 50, 255}
		for y := 2; y < 20; y++ {
			for x := 6; x < 26; x++ {
				dx := x - 16
				dy := y - 12
				if dx*dx+dy*dy <= 64 {
					img.Set(x, y, balloonColor)
				}
			}
		}
		// Нить
		img.Set(16, 22, color.RGBA{200, 200, 200, 255})
		img.Set(16, 24, color.RGBA{200, 200, 200, 255})
		img.Set(16, 26, color.RGBA{200, 200, 200, 255})
		// Глаза
		img.Set(12, 10, color.RGBA{255, 255, 255, 255})
		img.Set(20, 10, color.RGBA{255, 255, 255, 255})

	case "ghost":
		// Призрак
		ghostColor := color.RGBA{200, 200, 255, 200}
		for y := 4; y < 24; y++ {
			for x := 6; x < 26; x++ {
				if y < 18 || (x+y)%3 != 0 {
					img.Set(x, y, ghostColor)
				}
			}
		}
		// Глаза
		img.Set(11, 12, color.RGBA{0, 0, 0, 255})
		img.Set(20, 12, color.RGBA{0, 0, 0, 255})

	case "slime":
		// Слайм
		slimeColor := color.RGBA{50, 200, 50, 255}
		for y := 12; y < 28; y++ {
			for x := 4; x < 28; x++ {
				dx := x - 16
				dy := y - 20
				if dx*dx/4+dy*dy <= 64 {
					img.Set(x, y, slimeColor)
				}
			}
		}
		// Глаза
		img.Set(11, 16, color.RGBA{255, 255, 255, 255})
		img.Set(20, 16, color.RGBA{255, 255, 255, 255})
		img.Set(12, 16, color.RGBA{0, 0, 0, 255})
		img.Set(21, 16, color.RGBA{0, 0, 0, 255})
	}

	return ebiten.NewImageFromImage(img)
}

// SaveSpritesToPNG сохраняет все спрайты в PNG файлы
func (ss *SpriteStore) SaveSpritesToPNG(baseDir string) error {
	os.MkdirAll(filepath.Join(baseDir, "sprites"), 0755)

	// Сохраняем бомбермена
	for dir, img := range ss.Bomberman {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", "bomberman_"+dir+".png"), img); err != nil {
			return err
		}
	}

	// Сохраняем стены
	for name, img := range ss.Walls {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", "wall_"+name+".png"), img); err != nil {
			return err
		}
	}

	// Сохраняем бомбы
	for name, img := range ss.Bombs {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", name+".png"), img); err != nil {
			return err
		}
	}

	// Сохраняем взрывы
	for i, img := range ss.Explosions {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", "explosion_"+string(rune('0'+i))+".png"), img); err != nil {
			return err
		}
	}

	// Сохраняем улучшения
	for name, img := range ss.PowerUps {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", "powerup_"+name+".png"), img); err != nil {
			return err
		}
	}

	// Сохраняем врагов
	for name, img := range ss.Enemies {
		if err := saveEbitenImageToPNG(filepath.Join(baseDir, "sprites", "enemy_"+name+".png"), img); err != nil {
			return err
		}
	}

	return nil
}

func saveEbitenImageToPNG(path string, img *ebiten.Image) error {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, rgba)
}

// LoadImage загружает изображение из файла
func LoadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	return img, err
}
