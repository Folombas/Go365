# 🚀 EBITENGINE: Быстрый старт для разработчика MonoGame

**Дата:** 22 марта 2026 года
**День челленджа:** 82
**Тема:** Переход с C# + MonoGame на Go + Ebitengine

---

## 📋 СОДЕРЖАНИЕ

1. [Почему Ebitengine?](#почему-ebitengine)
2. [Установка и настройка](#установка-и-настройка)
3. [Сравнение: MonoGame vs Ebitengine](#сравнение-monogame-vs-ebitengine)
4. [Перенос кода с C# на Go](#перенос-кода-с-c-на-go)
5. [Перенос ассетов](#перенос-ассетов)
6. [Первые шаги](#первые-шаги)
7. [Ресурсы для обучения](#ресурсы-для-обучения)

---

## 🎯 ПОЧЕМУ EBITENGINE?

### Ты уже знаешь MonoGame — Ebitengine будет проще!

| Что ты знаешь из MonoGame | Аналог в Ebitengine |
|---------------------------|---------------------|
| `Game` class | `Game` interface |
| `Update(GameTime)` | `Update() error` |
| `Draw(GameTime)` | `Draw(screen *Image)` |
| `Content.Load<T>()` | `ebitenutil.NewImageFromFile()` |
| `SpriteBatch.Draw()` | `screen.DrawImage()` |
| `Keyboard.GetState()` | `ebiten.IsKeyPressed()` |
| `GraphicsDevice.Clear()` | `screen.Fill()` |

**Ebitengine проще — меньше бойлерплейта!**

---

## ⚙️ УСТАНОВКА И НАСТРОЙКА

### Шаг 1: Установи Go (если ещё не установлен)

```bash
# Проверить версию
go version

# Должно быть: go version go1.25.x ...
# Если нет — скачать с https://go.dev/dl/
```

### Шаг 2: Создать новый проект

```bash
# Создать директорию проекта
mkdir go_mario
cd go_mario

# Инициализировать Go модуль
go mod init go_mario

# Установить Ebitengine
go get github.com/hajimehoshi/ebiten/v2
```

### Шаг 3: Создать первый файл

Создай `main.go`:

```go
package main

import (
    "log"
    "github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
    // Игровая логика (60 раз в секунду)
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Отрисовка
    screen.Fill(ebiten.ColorBlack)
}

func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}

func main() {
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("My Go Game")
    
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
```

### Шаг 4: Запустить игру

```bash
go run main.go
```

**Игра запущена! 🎉**

---

## 📊 СРАВНЕНИЕ: MONOGAME VS EBITENGINE

### Игровой цикл

**MonoGame (C#):**
```csharp
public class Game1 : Game
{
    protected override void Update(GameTime gameTime)
    {
        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;
        // Логика
    }
    
    protected override void Draw(GameTime gameTime)
    {
        GraphicsDevice.Clear(Color.CornflowerBlue);
        // Отрисовка
    }
}
```

**Ebitengine (Go):**
```go
type Game struct{}

func (g *Game) Update() error {
    // dt фиксированный (1/60 секунды)
    // Логика
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // screen уже очищается автоматически
    // Отрисовка
}

func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}
```

**В Go меньше кода!**

---

### Загрузка спрайтов

**MonoGame:**
```csharp
// Требуется Content Pipeline (.mgcb файл)
Texture2D playerSprite;

protected override void LoadContent()
{
    playerSprite = Content.Load<Texture2D>("Sprites/player");
}
```

**Ebitengine:**
```go
// Напрямую из PNG, без Content Pipeline!
var playerSprite *ebiten.Image

func init() {
    playerSprite, _ = ebitenutil.NewImageFromFile("assets/player.png")
}
```

**В Go проще — нет Content Pipeline!**

---

### Отрисовка спрайта

**MonoGame:**
```csharp
protected override void Draw(GameTime gameTime)
{
    GraphicsDevice.Clear(Color.CornflowerBlue);
    
    spriteBatch.Begin();
    spriteBatch.Draw(playerSprite, position, Color.White);
    spriteBatch.End();
}
```

**Ebitengine:**
```go
func (g *Game) Draw(screen *ebiten.Image) {
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(position.X, position.Y)
    screen.DrawImage(playerSprite, opts)
}
```

**Похоже, но в Go проще!**

---

### Ввод с клавиатуры

**MonoGame:**
```csharp
KeyboardState keys = Keyboard.GetState();

if (keys.IsKeyDown(Keys.Right))
{
    position.X += speed;
}
```

**Ebitengine:**
```go
if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
    position.X += speed
}
```

**В Go короче!**

---

### Звуки

**MonoGame:**
```csharp
SoundEffect sound;
SoundEffectInstance instance;

sound = Content.Load<SoundEffect>("sounds/jump");
instance = sound.CreateInstance();
instance.Play();
```

**Ebitengine:**
```go
import "github.com/hajimehoshi/ebiten/v2/audio"

var jumpSound *audio.Player

func init() {
    f, _ := os.Open("sounds/jump.wav")
    jumpSound, _ = audio.NewContext(44100).NewPlayerFromReader(f)
}

// Воспроизведение
jumpSound.Rewind()
jumpSound.Play()
```

**Похожая сложность!**

---

## 🔄 ПЕРЕНОС КОДА С C# НА GO

### Классы → Структуры

**C#:**
```csharp
public class Player : GameObject
{
    public float X { get; set; }
    public float Y { get; set; }
    public int Health { get; set; }
    
    public void Jump() { ... }
    public void TakeDamage(int amount) { ... }
}
```

**Go:**
```go
type Player struct {
    X, Y   float64
    Health int
}

func (p *Player) Jump() { ... }
func (p *Player) TakeDamage(amount int) { ... }
```

**В Go нет наследования — используем композицию!**

---

### Наследование → Композиция

**C# (наследование):**
```csharp
public class Enemy : GameObject
{
    public int Damage { get; set; }
}

public class Goomba : Enemy
{
    public void Walk() { ... }
}
```

**Go (композиция):**
```go
type GameObject struct {
    X, Y float64
}

type Enemy struct {
    GameObject
    Damage int
}

type Goomba struct {
    Enemy
}

func (g *Goomba) Walk() { ... }
```

---

### Исключения → Обработка ошибок

**C#:**
```csharp
try
{
    var texture = Content.Load<Texture2D>("sprite");
}
catch (Exception e)
{
    Console.WriteLine($"Error: {e.Message}");
}
```

**Go:**
```go
sprite, err := ebitenutil.NewImageFromFile("sprite.png")
if err != nil {
    log.Printf("Error: %v", err)
    // Создать заглушку
    sprite = createPlaceholder()
}
```

**В Go ошибки — значения, обрабатываются явно!**

---

### Async/Await → Горутины

**C#:**
```csharp
public async Task LoadAssetsAsync()
{
    await Task.Run(() => LoadHeavyAssets());
    OnAssetsLoaded();
}
```

**Go:**
```go
go func() {
    loadHeavyAssets()
    onAssetsLoaded()
}()
```

**Горутины проще и легче Task!**

---

## 🎨 ПЕРЕНОС АССЕТОВ

### Все ассеты с opengameart.org работают!

**Структура проекта mario_monogame:**
```
mario_monogame/assets/
├── PNG/
│   ├── Players/
│   ├── Enemies/
│   ├── Items/
│   └── ...
```

**Структура проекта go_mario:**
```
go_mario/assets/
├── PNG/
│   ├── Players/     ← те же файлы!
│   ├── Enemies/     ← те же файлы!
│   ├── Items/       ← те же файлы!
│   └── ...
```

**Просто скопируй папку assets!**

```bash
# Из корня projects
cp -r mario_monogame/assets go_mario/
```

---

### Загрузка ассетов в Go

**Вместо Content Pipeline:**

```go
// Загрузка всех спрайтов игрока
func loadPlayerSprites() map[string]*ebiten.Image {
    sprites := make(map[string]*ebiten.Image)
    
    // Из mario_monogame/assets/PNG/Players/
    sprites["stand"], _ = ebitenutil.NewImageFromFile(
        "assets/PNG/Players/128x256/Green/alienGreen_stand.png")
    sprites["walk1"], _ = ebitenutil.NewImageFromFile(
        "assets/PNG/Players/128x256/Green/alienGreen_walk1.png")
    sprites["walk2"], _ = ebitenutil.NewImageFromFile(
        "assets/PNG/Players/128x256/Green/alienGreen_walk2.png")
    
    return sprites
}
```

---

## 🎮 ПЕРВЫЕ ШАГИ

### День 1: Hello World

```go
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
)

type Game struct{}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
    ebitenutil.DebugPrint(screen, "Hello, Ebitengine!")
}

func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}

func main() {
    ebiten.RunGame(&Game{})
}
```

---

### День 2: Движение квадрата

```go
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "log"
)

type Game struct {
    x, y float64
}

func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.x += 4
    }
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        g.x -= 4
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Рисуем квадрат
    for i := 0; i < 32; i++ {
        for j := 0; j < 32; j++ {
            screen.SetPixel(int(g.x)+i, int(g.y)+j, color.RGBA{0, 255, 0, 255})
        }
    }
}

func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}

func main() {
    ebiten.RunGame(&Game{x: 300, y: 200})
}
```

---

### День 3: Загрузка спрайта

```go
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
)

type Game struct {
    sprite *ebiten.Image
    x, y   float64
}

func NewGame() *Game {
    sprite, err := ebitenutil.NewImageFromFile("assets/player.png")
    if err != nil {
        log.Fatal(err)
    }
    return &Game{sprite: sprite, x: 300, y: 200}
}

func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.x += 4
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(g.x, g.y)
    screen.DrawImage(g.sprite, opts)
}

func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}

func main() {
    ebiten.RunGame(NewGame())
}
```

---

### День 4: Гравитация и прыжки

```go
type Player struct {
    x, y       float64
    vx, vy     float64
    grounded   bool
    sprite     *ebiten.Image
}

const (
    gravity   = 0.5
    jumpForce = -12
    moveSpeed = 4
)

func (p *Player) Update() {
    // Горизонтальное движение
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        p.vx = moveSpeed
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        p.vx = -moveSpeed
    } else {
        p.vx = 0
    }
    
    // Прыжок
    if (ebiten.IsKeyPressed(ebiten.KeySpace) || 
        ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && p.grounded {
        p.vy = jumpForce
        p.grounded = false
    }
    
    // Гравитация
    p.vy += gravity
    
    // Обновление позиции
    p.x += p.vx
    p.y += p.vy
    
    // Простая коллизия с "землёй"
    if p.y > 400 {
        p.y = 400
        p.vy = 0
        p.grounded = true
    }
}
```

---

### День 5: Камера

```go
type Camera struct {
    x, y float64
}

func (c *Camera) Follow(target *Player, screenWidth, screenHeight int) {
    // Камера следует за игроком
    targetX := target.x - float64(screenWidth)/2
    targetY := target.y - float64(screenHeight)/2
    
    // Плавное движение камеры
    c.x += (targetX - c.x) * 0.1
    c.y += (targetY - c.y) * 0.1
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Применяем трансформацию камеры
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(-g.camera.x, -g.camera.y)
    
    // Рисуем мир с учётом камеры
    g.world.Draw(screen, opts)
    g.player.Draw(screen, opts)
}
```

---

## 📚 РЕСУРСЫ ДЛЯ ОБУЧЕНИЯ

### Официальные ресурсы

| Ресурс | Ссылка |
|--------|--------|
| Ebitengine Official | https://ebitengine.org/ |
| Documentation | https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2 |
| Examples | https://ebitengine.org/en/examples/ |
| GitHub | https://github.com/hajimehoshi/ebiten |

### Туториалы

| Туториал | Ссылка |
|----------|--------|
| Ebitengine Tutorial (официальный) | https://ebitengine.org/en/tutorial/ |
| Go Game Development | https://github.com/quasilyte/awesome-go-game-development |
| Making Games with Ebiten | https://www.youtube.com/results?search_query=ebitengine+tutorial |

### Книги

| Книга | Автор |
|-------|-------|
| Game Programming Patterns | Robert Nystrom |
| Learning Go | Jon Bodner |

### Discord-каналы

- Ebitengine Discord: https://discord.gg/ebitengine
- Go Gophers: https://invite.slack.golangbridge.org/

---

## 🎯 ПЛАН ПЕРЕХОДА С MONOGAME НА EBITENGINE

### Неделя 1: Основы

- [ ] Установить Go и Ebitengine
- [ ] Пройти официальный туториал
- [ ] Сделать Pong/Arkanoid
- [ ] Изучить загрузку спрайтов

### Неделя 2: Перенос go_mario

- [ ] Скопировать ассеты из mario_monogame
- [ ] Реализовать игрока с движением
- [ ] Добавить гравитацию и прыжки
- [ ] Сделать камеру

### Неделя 3: Геймплей

- [ ] Добавить платформы
- [ ] Добавить врагов
- [ ] Добавить сбор предметов
- [ ] Сделать систему жизней

### Неделя 4: Полировка

- [ ] Добавить звуки
- [ ] Добавить меню
- [ ] Добавить частицы
- [ ] Собрать для WebAssembly

---

## 💡 СОВЕТЫ ДЛЯ БЫВШЕГО C# РАЗРАБОТЧИКА

### 1. Забудь про классы

**C#:** `class Player : GameObject`
**Go:** `type Player struct { GameObject }`

Композиция вместо наследования!

### 2. Забудь про исключения

**C#:** `try { ... } catch (Exception e)`
**Go:** `if err != nil { return err }`

Ошибки — значения!

### 3. Забудь про LINQ

**C#:** `items.Where(x => x.Active).ToList()`
**Go:**
```go
var result []Item
for _, item := range items {
    if item.Active {
        result = append(result, item)
    }
}
```

Явные циклы проще!

### 4. Используй горутины

**C#:** `await Task.Run(...)`
**Go:** `go func() { ... }()`

Горутины легче и проще!

### 5. Наслаждайся простотой

**C#:** 50 строк бойлерплейта
**Go:** 10 строк чистого кода

---

## 🏆 ЗАКЛЮЧЕНИЕ

**Ebitengine — отличный выбор для разработчика с опытом MonoGame!**

✅ Похожий API (Update/Draw)
✅ Проще (нет Content Pipeline)
✅ Быстрее (мгновенная компиляция)
✅ Легче (один бинарник)
✅ Кроссплатформенный (Web, Desktop, Mobile)

**Все твои ассеты с opengameart.org работают без изменений!**

**Переходи на Go + Ebitengine и не распыляйся!** 🐍

---

**Go365 Challenge** — День 82 из 365

**Девиз:** Ebitengine проще MonoGame! Go лучше C# для 2D игр! 🎮
