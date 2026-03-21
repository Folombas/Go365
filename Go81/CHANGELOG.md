# 📝 CHANGELOG — Day 81 (21 марта 2026)

**Дата:** 21 марта 2026 года
**День челленджа:** 81
**Проект:** mario_monogame — 2D Platformer на MonoGame

---

## 🎯 Цель дня

**Осознанное продолжение Go365 с временным отклонением на C# + MonoGame.**

Работа над проектом mario_monogame, анализ различий между Go+Ebitengine и C#+MonoGame, рефлексия о важности фокусировки.

---

## ✅ Выполненные задачи

### 1. Работа над mario_monogame

**Проект:** 2D-платформер на MonoGame с кроссплатформенной поддержкой

**Структура проекта:**
```
mario_monogame/
├── mario_monogame.sln          # Решение Visual Studio
├── mario_monogame.Core/        # Общая игровая логика
├── mario_monogame.DesktopGL/   # Desktop (Linux/Mac/Windows)
├── mario_monogame.WindowsDX/   # Windows DirectX
├── mario_monogame.Android/     # Android версия
├── mario_monogame.iOS/         # iOS версия
├── assets/                     # Игровые ассеты
└── levels/                     # Файлы уровней
```

**Особенности проекта:**
- Кроссплатформенная архитектура (Desktop, Mobile)
- Разделение на Core и платформо-зависимые проекты
- Использование Content Pipeline для ассетов
- Поддержка Android и iOS

**Ассеты в проекте:**
- `arcade_platformerV2.png` — тайлсет платформера
- `generic_platformer_tiles.png` — базовые тайлы
- `Platformer Art Complete Pack.zip` — полный пак арта
- `Platformer Pack Redux (360 assets).zip` — 360 ассетов
- `parallax-industrial-pack.zip` — параллакс фон
- `parallax_forest_pack.zip` — лесный фон
- `parallax_mountain_pack.zip` — горный фон
- `credits.zip`, `howl.png` — дополнительные ресурсы
- Шрифты: `bloody-modes-font.zip`, `smilen-font.zip`, `super-feel-font.zip`

---

### 2. Анализ: Go + Ebitengine vs C# + MonoGame

#### 📊 Детальное сравнение

| Критерий | Go + Ebitengine | C# + MonoGame |
|----------|-----------------|---------------|
| **Язык программирования** | Go (статическая, простота, горутины) | C# (статическая, ООП, LINQ, async/await) |
| **Парадигма** | Процедурная + композиция | Объектно-ориентированная |
| **Синтаксис** | Минималистичный, 25 ключевых слов | Богатый, много синтаксического сахара |
| **Типизация** | Статическая, структурные интерфейсы | Статическая, классы и интерфейсы |
| **Наследование** | Нет (только композиция) | Да (классы, наследование, полиморфизм) |
| **Исключения** | Нет (error handling через return) | Да (try-catch-finally) |
| **Garbage Collection** | Есть (конкурентный GC) | Есть (поколенческий GC) |

---

#### 🏗️ Архитектура и дизайн

| Аспект | Go + Ebitengine | C# + MonoGame |
|--------|-----------------|---------------|
| **Архитектура игры** | Простая, плоская структура | Иерархическая, ECS-подобная |
| **Игровой цикл** | `Update()`, `Draw()`, `Layout()` | `Update()`, `Draw()` в Game class |
| **Управление состоянием** | Ручное, через переменные | State pattern, enum states |
| **Компонентная система** | Ручная реализация | Можно использовать ECS библиотеки |
| **Система сцен** | Ручная реализация | SceneManager, ScreenManager |
| **Шаблон проектирования** | Composition over inheritance | Entity-Component, Strategy, Observer |

**Пример структуры на Go:**
```go
type Game struct {
    player *Player
    world  *World
    ui     *UI
}

func (g *Game) Update() error {
    g.player.Update()
    g.world.Update()
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.world.Draw(screen)
    g.player.Draw(screen)
    g.ui.Draw(screen)
}
```

**Пример структуры на C#:**
```csharp
public class Game1 : Game
{
    private Player player;
    private World world;
    private UIManager ui;
    
    protected override void Update(GameTime gameTime)
    {
        player.Update(gameTime);
        world.Update(gameTime);
        base.Update(gameTime);
    }
    
    protected override void Draw(GameTime gameTime)
    {
        GraphicsDevice.Clear(Color.CornflowerBlue);
        world.Draw(spriteBatch);
        player.Draw(spriteBatch);
        ui.Draw(spriteBatch);
        base.Draw(gameTime);
    }
}
```

---

#### ⚡ Производительность

| Метрика | Go + Ebitengine | C# + MonoGame |
|---------|-----------------|---------------|
| **Компиляция** | Мгновенная (< 1 сек) | Быстрая (1-5 сек) |
| **Время запуска** | < 1 секунды | 1-3 секунды |
| **Потребление памяти** | 20-50 MB | 100-200 MB |
| **Размер бинарника** | 5-15 MB (один файл) | 50-100 MB (с runtime) |
| **FPS в 2D** | 60+ стабильно | 60+ стабильно |
| **GC паузы** | Минимальные (< 1ms) | Заметные (1-5ms) |
| **Нативный код** | Да (компиляция в бинарник) | Нет (JIT компиляция) |

**Вывод:** Go выигрывает по потреблению памяти и размеру бинарника. C# может иметь фризы из-за GC.

---

#### 🛠️ Инструменты разработки

| Инструмент | Go + Ebitengine | C# + MonoGame |
|------------|-----------------|---------------|
| **IDE** | VS Code, GoLand, Vim | Visual Studio, Rider, VS Code |
| **Отладчик** | Delve, встроенный в IDE | Visual Studio Debugger, Rider |
| **Форматирование** | `go fmt` (автоматически) | ReSharper, Code Cleanup |
| **Тестирование** | Встроенное `go test` | xUnit, NUnit, MSTest |
| **Профилирование** | `go tool pprof` | Visual Studio Profiler, dotTrace |
| **Hot Reload** | Air, fresh | Visual Studio Hot Reload |
| **Linting** | `go vet`, `golint`, `golangci-lint` | Roslyn Analyzers, ReSharper |
| **Package Manager** | `go mod` (встроен) | NuGet (встроен в .NET) |

**Вывод:** C# имеет более мощные IDE и инструменты, но Go инструменты проще и быстрее.

---

#### 📦 Экосистема и библиотеки

| Категория | Go + Ebitengine | C# + MonoGame |
|-----------|-----------------|---------------|
| **Игровые библиотеки** | Ebitengine, Raylib-go, Pixel | MonoGame, FNA, Xenko |
| **Физика** | Chipmunk, Box2D-go | Farseer Physics, Box2D.XNA |
| **Аудио** | Встроенное в Ebitengine | MonoGame Audio, FMOD, Wwise |
| **UI** | Ebitengine UI, fyne | MonoGame.Extended, Nez |
| **Карты/тайлы** | Tiled + парсеры | Tiled + MonoGame.Extended |
| **Анимации** | Ручная реализация | MonoGame.Extended, Spine |
| **Частицы** | Ручная реализация | MonoGame.Extended, Paragon |
| **Сообщество** | Растущее, ~10k звёзд Ebitengine | Большое, ~10k звёзд MonoGame |
| **Туториалы** | ~100 качественных | ~1000+ различных |
| **Готовые ассеты** | Мало | Много (itch.io, Asset Store) |

**Вывод:** C# имеет значительно более богатую экосистему для игр.

---

#### 🌐 Кроссплатформенность

| Платформа | Go + Ebitengine | C# + MonoGame |
|-----------|-----------------|---------------|
| **Windows** | ✅ Отлично | ✅ Отлично |
| **Linux** | ✅ Отлично | ✅ Хорошо |
| **macOS** | ✅ Отлично | ✅ Хорошо |
| **WebAssembly** | ✅ Отлично (в браузере) | ⚠️ Ограниченно (Blazor) |
| **Android** | ✅ Хорошо | ✅ Хорошо |
| **iOS** | ✅ Хорошо | ✅ Хорошо |
| **Nintendo Switch** | ⚠️ Требуется порт | ✅ Через Unity порты |
| **PlayStation/Xbox** | ⚠️ Сложно | ⚠️ Сложно |

**Вывод:** Go лучше для WebAssembly. C# лучше для консольных портов.

---

#### 🚀 Деплой и дистрибуция

| Аспект | Go + Ebitengine | C# + MonoGame |
|--------|-----------------|---------------|
| **Бинарник** | Один файл | Один файл (self-contained) |
| **Зависимости** | Никаких | .NET runtime или включён |
| **Размер дистрибутива** | 5-15 MB | 50-100 MB |
| **Установка** | Копирование файла | Установка или self-contained |
| **Steam** | ✅ Легко | ✅ Легко |
| **itch.io** | ✅ Легко | ✅ Легко |
| **Google Play** | ✅ Хорошо | ✅ Хорошо |
| **App Store** | ✅ Хорошо | ✅ Хорошо |

**Вывод:** Go проще для дистрибуции благодаря одному файлу.

---

#### 📈 Порог входа и обучение

| Аспект | Go + Ebitengine | C# + MonoGame |
|--------|-----------------|---------------|
| **Сложность языка** | Низкая (25 ключевых слов) | Средняя (много концепций) |
| **Время до первой игры** | 1-2 часа | 2-4 часа |
| **Документация** | Хорошая, лаконичная | Отличная, подробная |
| **Примеры кода** | ~50 качественных | ~500+ различных |
| **Видео туториалы** | ~20 на YouTube | ~200+ на YouTube |
| **Книги** | Мало про игры | Много про game dev |
| **Курсы** | Мало | Много (Udemy, Coursera) |
| **Форумы** | Discord, Reddit r/golang | Discord, Reddit r/monogame |

**Вывод:** Go проще выучить, но C# имеет больше обучающих материалов.

---

#### 💰 Работа и карьера

| Аспект | Go + Ebitengine | C# + MonoGame |
|--------|-----------------|---------------|
| **Вакансии Go** | Backend, Cloud, DevOps | - |
| **Вакансии C#** | - | Game Dev, Enterprise, Unity |
| **Зарплаты Go** | Высокие ($80k-200k) | Средние-высокие ($60k-150k) |
| **Индустрия игр** | Инди, мобильные | Инди, AA, мобильные |
| **Перенос навыков** | Backend, микросервисы | Unity, enterprise .NET |
| **Спрос на рынке** | Растущий | Стабильный |

**Вывод:** Go лучше для backend карьеры. C# лучше для game dev карьеры.

---

### 3. Рефлексия о фокусировке

#### ⚠️ Риски распыления

1. **Потеря прогресса**
   - Переключение между языками замедляет обучение
   - Контекст теряется, нужно время на "разгон"

2. **Поверхностные знания**
   - Знание многих языков на среднем уровне
   - Отсутствие экспертизы в одном языке

3. **Упущенное время**
   - 2026 год — год тотального фокуса на Go
   - Каждое отвлечение — шаг назад от цели

4. **Разные парадигмы**
   - Go: простота, композиция, горутины
   - C#: ООП, наследование, LINQ, async/await
   - Переключение требует ментальной перенастройки

#### ✅ Преимущества текущего подхода

1. **Сравнительный анализ**
   - Понимание различий между языками
   - Осознанный выбор инструментов

2. **Расширение кругозора**
   - Знание разных подходов к game dev
   - Умение выбирать правильный инструмент

3. **Готовые проекты**
   - mario_monogame на C#
   - go_mario на Go
   - snake на Go

#### 🎯 Решение

**Фокус на Go до конца 2026 года!**

1. **Завершить текущие задачи в mario_monogame**
2. **Вернуться к go_mario (Ebitengine)**
3. **Развивать go_mario до конца года**
4. **Достичь продвинутого уровня в Go**
5. **Устроиться на Go-позицию**

---

## 📊 Статистика

### Проект mario_monogame
| Метрика | Значение |
|---------|----------|
| Платформа | MonoGame 3.8+ |
| Язык | C# .NET |
| Структура | Core + платформо-зависимые проекты |
| Поддержка | Desktop, Android, iOS |
| Ассеты | 10+ паков с графикой и звуками |

### Проект go_mario (для сравнения)
| Метрика | Значение |
|---------|----------|
| Платформа | Ebitengine 2.9+ |
| Язык | Go 1.25+ |
| Структура | Один пакет (main) |
| Поддержка | Desktop, Web, Mobile |
| Размер | ~1900 строк кода |

### Сравнение времени разработки
| Задача | Go + Ebitengine | C# + MonoGame |
|--------|-----------------|---------------|
| Создание проекта | 1 минута | 5 минут |
| Добавление спрайта | 5 минут | 10 минут (Content Pipeline) |
| Компиляция | < 1 секунды | 1-5 секунд |
| Деплой | Копирование файла | Публикация с runtime |

---

## 🔧 Технические детали

### Go + Ebitengine: минимальный код
```go
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "image/color"
)

type Game struct{}

func (g *Game) Update() error { return nil }
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{255, 0, 0, 255})
}
func (g *Game) Layout(w, h int) (int, int) {
    return 640, 480
}

func main() {
    ebiten.RunGame(&Game{})
}
```

### C# + MonoGame: минимальный код
```csharp
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

public class Game1 : Game
{
    GraphicsDeviceManager graphics;
    SpriteBatch spriteBatch;
    
    public Game1() {
        graphics = new GraphicsDeviceManager(this);
    }
    
    protected override void LoadContent() {
        spriteBatch = new SpriteBatch(GraphicsDevice);
    }
    
    protected override void Update(GameTime gameTime) {
        base.Update(gameTime);
    }
    
    protected override void Draw(GameTime gameTime) {
        GraphicsDevice.Clear(Color.Red);
        base.Draw(gameTime);
    }
}
```

**Вывод:** Go код короче и проще.

---

## 💭 Итоги

**Реализовано:**
- ✅ Оформлен отчёт за день 81
- ✅ Задокументирована работа над mario_monogame
- ✅ Проведён детальный анализ Go vs C#
- ✅ Осознаны риски распыления
- ✅ Зафиксировано решение о фокусе на Go

**Выводы:**
1. **Go проще и быстрее** для разработки 2D игр
2. **C# богаче инструментами** и экосистемой
3. **Go лучше для обучения** и быстрых прототипов
4. **C# лучше для сложных проектов** с командой
5. **Фокус на Go** — правильный выбор для 2026 года

**План:**
1. Завершить текущие задачи в mario_monogame
2. Вернуться к go_mario (Ebitengine)
3. Развивать go_mario до продвинутого уровня
4. Устроиться на Go-позицию до конца 2026

**День 81 завершён!** 🎉

**Фокус на Go до конца 2026 года!** 🐍

---

## 📝 Заметки

**Проблемы:**
1. Распыление между Go и C# замедляет прогресс
2. Нужно чётче разделять время между проектами
3. mario_monogame требует завершения текущих задач

**Решения:**
1. Установить дедлайн для mario_monogame
2. После завершения — полный фокус на Go
3. Вести ежедневный учёт времени

**Планы:**
1. Завершить базовый функционал mario_monogame
2. Вернуться к go_mario
3. Добавить врагов, бонусы, уровни
4. Опубликовать на itch.io

---

## 🔗 Ссылки

- **Репозиторий Go365:** https://github.com/Folombas/Go365
- **Репозиторий playgo:** https://github.com/Folombas/playgo
- **Ebitengine:** https://ebitengine.org/
- **MonoGame:** https://www.monogame.net/
- **Go365 Day 78:** ../Go78/CHANGELOG.md
- **Go365 Day 80:** ../Go80/CHANGELOG.md

---

**Девиз дня:** Осознанная работа с пониманием долгосрочных целей! 🎯

**Go365 Challenge** — День 81 из 365
