# 📝 CHANGELOG — Day 79 (19 марта 2026)

**Дата:** 19 марта 2026 года  
**День челленджа:** 79  
**Проект:** Simple Snake (playgo/snake)

---

## 🎯 Цель дня

**Тотальная фокусировка на Go!** Визуальная полировка и рефакторинг Simple Snake.

---

## ✅ Выполненные задачи

### 1. Перенос проекта в playgo

**Репозиторий:** https://github.com/Folombas/playgo

**Изменения:**
- ✅ Скопирован проект simple snake из Go76
- ✅ Обновлён go.mod (module playgo/snake)
- ✅ Сделан коммит и пуш в репозиторий playgo

**Коммит:**
```
Add Simple Snake game from Go365 Go76 (v0.3.0)
```

---

### 2. Визуальная полировка (v0.4.0)

#### 🎨 Градиентный фон
- ✅ Плавный переход от тёмно-синего к чёрному
- ✅ Создаётся один раз при инициализации

#### ✨ Система частиц
- ✅ Структура Particle с физикой
- ✅ Частицы при всех значимых событиях

#### 📳 Тряска экрана
- ✅ Структура ScreenShake
- ✅ Разная интенсивность для событий

#### 💫 Пульсация объектов
- ✅ Монеты пульсируют
- ✅ Еда пульсирует при появлении

---

### 3. Рефакторинг кода (v0.5.0)

#### 📦 Новая структура проекта
```
snake/
├── main.go                    # Точка входа, координатор
├── go.mod
├── CHANGELOG.md
├── PLAN.md
└── internal/
    ├── game/                  # Игровая логика
    │   └── game.go
    ├── effects/               # Визуальные эффекты
    │   └── effects.go
    └── ui/                    # Отрисовка
        └── renderer.go
```

#### 🎮 Пакет game
**Файл:** `internal/game/game.go`

**Структуры:**
- `Game` — основная игровая структура
- `Config` — конфигурация игры
- `Point` — позиция на сетке
- `Enemy`, `Bomb`, `TreasureChest`, `Key`, `Coin`, `Arrow` — игровые объекты
- `Direction`, `GameState`, `Difficulty` — перечисления
- `GameEvent`, `GameEventType` — система событий

**Методы:**
- `NewGame()` — создание новой игры
- `StartGame()` — начало игры
- `Update()` — обновление состояния, возвращает события
- `UpdateDirection()` — обновление направления
- `ShootArrow()` — выстрел стрелой
- `Config()` — получение конфигурации

**События игры:**
```go
EventEatFood        // Поедание еды
EventCollectKey     // Сбор ключа
EventCollectCoin    // Сбор монеты
EventOpenChest      // Открытие сундука
EventEnemyKill      // Убийство врага
EventEnemyCollision // Столкновение с врагом
EventBombExplode    // Взрыв бомбы
EventBombCollision  // Столкновение с бомбой
EventWallCollision  // Удар о стену
EventSelfCollision  // Столкновение с хвостом
```

#### ✨ Пакет effects
**Файл:** `internal/effects/effects.go`

**Структуры:**
- `Particle` — частица для визуальных эффектов
- `ScreenShake` — тряска экрана
- `EffectSystem` — управление всеми эффектами

**Методы:**
- `NewEffectSystem()` — создание системы эффектов
- `SpawnParticles()` — создание частиц
- `Update()` — обновление всех эффектов
- `Draw()` — отрисовка частиц
- `TriggerShake()` — запуск тряски
- `CreateGradientBackground()` — создание градиентного фона
- `PulseScale()` — коэффициент пульсации
- `FoodPulseScale()` — пульсация еды

#### 🎨 Пакет ui
**Файл:** `internal/ui/renderer.go`

**Структура:**
- `Renderer` — рендерер игры

**Методы:**
- `NewRenderer()` — создание рендерера
- `DrawMenu()` — отрисовка меню
- `DrawDifficultySelection()` — выбор сложности
- `DrawGame()` — отрисовка игры
- `DrawPauseOverlay()` — оверлей паузы
- `DrawGameOverOverlay()` — оверлей конца игры
- `drawSnakeEyes()`, `drawSnakeTongue()` — глаза и язык змейки
- `drawFood()`, `drawEnemy()`, `drawBomb()` — еда, враги, бомбы
- `drawChest()`, `drawKey()`, `drawCoin()`, `drawArrow()` — предметы

---

#### 🔄 Обновлённый main.go
**До:** ~1383 строки (монолит)  
**После:** ~213 строк (координатор)

**Структура App:**
```go
type App struct {
    game       *game.Game
    effects    *effects.EffectSystem
    renderer   *ui.Renderer
    background *ebiten.Image
}
```

**Методы:**
- `NewApp()` — создание приложения
- `Update()` — обработка ввода и обновление
- `Draw()` — отрисовка
- `Layout()` — размер экрана

**Обработка событий:**
```go
events := a.game.Update()
for _, event := range events {
    switch event.Type {
    case game.EventEatFood:
        a.effects.SpawnParticles(x, y, 10, color.RGBA{255, 100, 0, 255}, 2)
    case game.EventEnemyKill:
        a.effects.SpawnParticles(x, y, 25, color.RGBA{128, 0, 128, 255}, 4)
        a.effects.TriggerShake(5, 20)
    // ...
    }
}
```

---

## 📊 Статистика

### Код
| Метрика | До рефакторинга | После | Изменение |
|---------|-----------------|-------|-----------|
| Файлов | 1 | 4 | +3 |
| Строк кода | ~1383 | ~1820 | +437 |
| main.go строк | ~1383 | ~213 | -1170 |
| Пакетов | 1 | 4 | +3 |

### Структура проекта
```
До:
main.go (1383 строки)

После:
main.go (213 строк)
internal/game/game.go (~550 строк)
internal/effects/effects.go (~180 строк)
internal/ui/renderer.go (~615 строк)
```

### Производительность
- Градиент создаётся один раз при инициализации
- Частицы обновляются только когда активны
- События обрабатываются централизованно

---

## 🔧 Технические детали

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

### Коммиты
- `Go79: Add visual polish to Simple Snake (v0.4.0)`
- `Go79: Refactor code into packages (game, effects, ui)`

---

## 💭 Итоги

**Реализовано:**
- ✅ Проект перенесён в playgo/snake
- ✅ Градиентный фон реализован
- ✅ Система частиц добавлена
- ✅ Тряска экрана работает
- ✅ Пульсация объектов добавлена
- ✅ **Код разделён на пакеты**
- ✅ main.go уменьшен с 1383 до 213 строк

**Влияние:**
- Игра стала визуально более отзывчивой
- **Код стал легче поддерживать и расширять**
- **Разделение ответственности между пакетами**
- **Упростилось тестирование отдельных компонентов**

**День 79 завершён!** 🎉

**Фокус на Go до конца 2026 года!** 🐍

---

## 📝 Заметки

**Проблемы:**
1. Нет спрайтов — векторная графика
2. Нет звуковых эффектов
3. Нет таблицы рекордов

**Планы:**
1. Добавить звуковые эффекты (программная генерация)
2. Таблица рекордов (сохранение в файл)
3. Добавить новые типы бонусов

---

## 🔗 Ссылки

- **Репозиторий Go365:** https://github.com/Folombas/Go365
- **Репозиторий playgo:** https://github.com/Folombas/playgo
- **Ebitengine:** https://ebitengine.org/

---

**Девиз дня:** Тотальная фокусировка на Go! Никакого распыления! 💪

**Go365 Challenge** — День 79 из 365
