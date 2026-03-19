# 📝 CHANGELOG — Day 79 (19 марта 2026)

**Дата:** 19 марта 2026 года  
**День челленджа:** 79  
**Проект:** Simple Snake (playgo/snake)

---

## 🎯 Цель дня

**Тотальная фокусировка на Go!** Визуальная полировка игры Simple Snake.

### Контекст
- Проект перенесён из Go365/Go76 в playgo/snake
- Текущая версия: 0.4.0
- Цель: добавить визуальную полировку для улучшения восприятия игры

---

## ✅ Выполненные задачи

### 1. Перенос проекта в playgo

**Репозиторий:** https://github.com/Folombas/playgo

**Изменения:**
- ✅ Скопирован проект из Go365/Go76 в playgo/snake
- ✅ Обновлён go.mod (module playgo/snake)
- ✅ Сделан коммит и пуш в репозиторий

**Статус:** Проект теперь в общем репозитории игр playgo

---

### 2. Визуальная полировка (Go79)

#### 🎨 Градиентный фон
- ✅ Плавный переход от тёмно-синего (10, 10, 30) к чёрному (0, 0, 0)
- ✅ Создаётся один раз при инициализации игры
- ✅ Рисуется как фоновое изображение

**Функция:**
```go
func createGradientBackground() *ebiten.Image {
    gradient := ebiten.NewImage(screenWidth, screenHeight)
    for y := 0; y < screenHeight; y++ {
        ratio := float32(y) / float32(screenHeight)
        r := uint8(float32(10) * (1 - ratio))
        g := uint8(float32(10) * (1 - ratio))
        b := uint8(float32(30) * (1 - ratio))
        for x := 0; x < screenWidth; x++ {
            gradient.Set(x, y, color.RGBA{r, g, b, 255})
        }
    }
    return gradient
}
```

---

#### ✨ Система частиц
**Структура Particle:**
```go
type Particle struct {
    X, Y     float32    // Позиция
    VX, VY   float32    // Скорость
    Life     int        // Оставшаяся жизнь
    MaxLife  int        // Максимальная жизнь
    Color    color.RGBA // Цвет
    Size     float32    // Размер
    Gravity  float32    // Гравитация
}
```

**Методы:**
```go
// Создание частиц
func (g *Game) spawnParticles(x, y float32, count int, baseColor color.RGBA, spread float32)

// Обновление частиц (физика)
func (g *Game) updateParticles()

// Отрисовка частиц
func (g *Game) drawParticles(screen *ebiten.Image)
```

**Эффекты частиц:**

| Событие | Количество | Цвет | Интенсивность |
|---------|------------|------|---------------|
| 🍎 Поедание еды | 10 частиц | Оранжево-красный | Средняя |
| 🗝️ Сбор ключа | 15 частиц | Золотой | Средняя |
| 🪙 Сбор монеты | 15 частиц | Золотой | Средняя |
| 🏴‍☠️ Открытие сундука | 20 частиц | Золотой | Высокая |
| 💀 Убийство врага | 25 частиц | Фиолетовый | Высокая |
| 💣 Взрыв бомбы | 40 частиц | Оранжево-красный | Очень высокая |
| ☠️ Удар о стену | 20 частиц | Красный | Средняя |
| ☠️ Столкновение с врагом | 30 частиц | Фиолетовый | Высокая |

**Физика частиц:**
- Случайное направление разлёта
- Гравитация (0.1)
- Затухание по мере окончания жизни
- Альфа-канал уменьшается со временем

---

#### 📳 Тряска экрана (Screen Shake)

**Структура ScreenShake:**
```go
type ScreenShake struct {
    Intensity float32 // Сила тряски
    Duration  int     // Длительность в кадрах
    Timer     int     // Таймер обратного отсчёта
    Angle     float64 // Угол для синусоидального движения
}
```

**Методы:**
```go
// Обновление тряски
func (ss *ScreenShake) Update()

// Проверка активности
func (ss *ScreenShake) IsActive() bool

// Получение смещения
func (ss *ScreenShake) GetOffset() (float32, float32)

// Запуск тряски
func (ss *ScreenShake) Trigger(intensity float32, duration int)
```

**Эффекты тряски:**

| Событие | Интенсивность | Длительность |
|---------|--------------|--------------|
| ☠️ Удар о стену | 5 | 20 кадров |
| ☠️ Столкновение с врагом | 8 | 25 кадров |
| 💣 Взрыв бомбы | 10 | 30 кадров |
| 💀 Game Over | 5 | 20 кадров |

**Реализация:**
```go
func (ss *ScreenShake) GetOffset() (float32, float32) {
    if !ss.IsActive() {
        return 0, 0
    }
    offset := ss.Intensity * float32(ss.Timer) / float32(ss.Duration)
    dx := offset * float32(math.Sin(ss.Angle))
    dy := offset * float32(math.Cos(ss.Angle))
    return dx, dy
}
```

---

#### 💫 Пульсация объектов

**Монеты:**
```go
// Пульсация монеты
pulseScale := 1.0 + 0.1*math.Sin(coin.pulsePhase)
vector.DrawFilledCircle(screen, centerX, centerY, coinRadius*float32(pulseScale), ...)
```

**Еда:**
```go
// Анимация появления (пульсация в начале)
pulseScale := 1.0
if g.foodTimer > 0 {
    pulseScale = 1.0 + 0.3*math.Sin(float64(g.foodTimer)*math.Pi/10)
    g.foodTimer--
}
```

**Обновление пульсации:**
```go
func (g *Game) Update() error {
    // Обновление пульсации монет
    for i := range g.coins {
        g.coins[i].pulsePhase += 0.15
    }
    // ...
}
```

---

### 3. Изменения в коде

#### Обновлённая структура Game
```go
type Game struct {
    // ... существующие поля ...
    
    // Visual effects
    particles          []Particle
    screenShake        ScreenShake
    foodTimer          int
    backgroundGradient *ebiten.Image
}
```

#### Обновлённый метод Draw
```go
func (g *Game) Draw(screen *ebiten.Image) {
    // Применяем тряску экрана
    dx, dy := g.screenShake.GetOffset()

    // Рисуем градиентный фон
    screen.DrawImage(g.backgroundGradient, nil)

    // Создаём временную поверхность для игры
    gameScreen := ebiten.NewImage(screenWidth, screenHeight)

    // Рисуем игру на временную поверхность
    switch g.state {
    case Menu:
        g.drawMenu(gameScreen)
        g.drawParticles(gameScreen)
        screen.DrawImage(gameScreen, nil)
        // ...
    }

    // Рисуем игровую поверхность со смещением тряски
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(dx), float64(dy))
    screen.DrawImage(gameScreen, op)
}
```

---

## 📊 Статистика

### Код
| Метрика | До | После | Изменение |
|---------|-----|-------|-----------|
| Файлов | 1 | 1 | 0 |
| Строк кода | ~1356 | ~1383 | +27 |
| Структур | 8 | 10 | +2 |
| Методов | ~25 | ~30 | +5 |

### Новые структуры
- `Particle` — частица для визуальных эффектов
- `ScreenShake` — тряска экрана

### Новые методы
- `spawnParticles()` — создание частиц
- `updateParticles()` — обновление частиц
- `drawParticles()` — отрисовка частиц
- `createGradientBackground()` — создание фона

### Производительность
- Градиент создаётся один раз при инициализации
- Частицы обновляются только когда активны
- Тряска экрана работает только во время события

---

## 🔧 Технические детали

**Фреймворк:** Ebitengine v2.9.9  
**Язык:** Go 1.25.0  
**Размер окна:** 800×600 пикселей  
**Размер клетки:** 20×20 пикселей

### Сборка
```bash
cd playgo/snake
go build -o snake .
```

### Запуск
```bash
./snake
# или
go run .
```

---

## 💭 Итоги

**Реализовано:**
- ✅ Градиентный фон (тёмно-синий → чёрный)
- ✅ Система частиц (физика, затухание, гравитация)
- ✅ Тряска экрана (разная интенсивность для событий)
- ✅ Пульсация монет и ключей
- ✅ Анимация появления еды
- ✅ Частицы при всех значимых событиях игры

**Влияние:**
- Игра стала визуально более отзывчивой
- Частицы добавляют "сочности" (juiciness) геймплею
- Тряска экрана подчёркивает важные события
- Градиентный фон создаёт приятную атмосферу

**День 79 завершён!** 🎉

**Фокус на Go до конца 2026 года!** 🐍

---

## 📝 Заметки

**Проблемы:**
1. Один большой файл main.go — стоит разделить на пакеты
2. Нет спрайтов — вся графика векторная
3. Нет звуковых эффектов

**Оптимизации:**
1. Вынести эффекты в отдельный пакет (effects/particles.go)
2. Использовать пул объектов для частиц
3. Кэширование градиентов для разных состояний

**Планы:**
1. Добавить звуковые эффекты (программная генерация)
2. Реализовать таблицу рекордов
3. Разделить код на пакеты (game, effects, ui)
4. Добавить новые типы бонусов

---

## 🔗 Ссылки

- **Репозиторий Go365:** https://github.com/Folombas/Go365
- **Репозиторий playgo:** https://github.com/Folombas/playgo
- **Ebitengine:** https://ebitengine.org/
- **Ebitengine Vector:** https://ebitengine.org/en/packages/vector.html

---

**Девиз дня:** Тотальная фокусировка на Go! Никакого распыления! 💪

**Go365 Challenge** — День 79 из 365
