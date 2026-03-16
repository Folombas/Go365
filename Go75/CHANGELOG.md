# CHANGELOG.md — Go75 (16 марта 2026)

## [0.2.0] — 2026-03-16

### Добавлено

#### 🎮 Игровая механика
- ✅ **Пиратский сундук** — спавнится случайно на карте
  - Коричневый корпус с золотой крышкой
  - Золотой замок на закрытом сундуке
  - Содержит 5 стрел
  - Открывается только золотым ключом
- ✅ **Золотой ключик** — спавнится случайно на карте
  - Золотой цвет (RGB 255,215,0)
  - Круглая головка + стержень с зубцами
  - Подбирается змейкой при столкновении
- ✅ **Стрелы** — оружие против жуков
  - Серебристо-серые с треугольным наконечником
  - Запуск: **SPACE** (пробел)
  - Летят в направлении движения змейки
  - Убивают жуков (+1 очко за убийство)
  - Счётчик стрел отображается в UI

#### 📋 Меню и интерфейсы
- ✅ **Стартовое меню** (GameState: Menu)
  - Заголовок: "SNAKE GAME"
  - Подзаголовок: "Go365 Go75 - Ebitengine"
  - Кнопка старта: ENTER или SPACE
  - Информация об управлении
  - Описание цели игры
- ✅ **Пауза** (GameState: Paused)
  - Активация: клавиша **P**
  - Полупрозрачный чёрный оверлей
  - Текст "PAUSED"
  - Подсказка "Press P to Continue"
- ✅ **Game Over экран** (GameState: GameOver)
  - Полупрозрачный красный оверлей
  - Текст "GAME OVER"
  - Финальный счёт
  - Количество врагов
  - Рестарт по ENTER

#### 🐍 Визуальные улучшения
- ✅ **Змеиный язычок** — тонкий раздвоенный красный язык
  - Направлен по движению змейки
  - Две веточки на конце (forked tongue)
- ✅ **Глаза змейки** — белые с чёрными зрачками
  - Позиция зависит от направления движения
- ✅ **Видимая граница поля** — серая рамка 2px по периметру

#### 🐛 Улучшения жуков
- ✅ **Грозный рот** — тёмный круг вместо клыков
- ✅ **Большие светящиеся глаза** — красные с эффектом glow
  - Многослойное свечение (outer/middle/inner)
  - Чёрные зрачки
  - Белые блики для демонического вида

### Изменено
- 📐 Уменьшен размер окна: 2560×1440 → **800×600** пикселей
- 📊 Сетка: 128×72 → **40×30** клеток
- 🐛 Убраны белые клыки у жуков (оставлен только рот)
- 🎯 Игра стала компактнее и играбельнее

### Технические детали
- Фреймворк: Ebitengine v2.9.9
- Язык: Go 1.x
- Размер окна: 800×600 пикселей
- Размер клетки: 20×20 пикселей
- GameState enum: Menu, Playing, Paused, GameOver

### Управление
| Клавиша | Действие |
|---------|----------|
| ↑↓←→ | Движение змейки |
| SPACE | Выстрел стрелой / Старт в меню |
| P | Пауза / Продолжить |
| ENTER | Старт в меню / Рестарт после Game Over |

### Игровой процесс
1. Найди **золотой ключик** 🗝️
2. Найди **пиратский сундук** 🏴‍☠️
3. Открой сундук с ключом → получи **5 стрел** 🏹
4. Стреляй в жуков (**SPACE**) 💀
5. Избегай:
   - Столкновения со стенами ☠️
   - Столкновения с хвостом ☠️
   - Жуков с клыками ☠️
   - Взрывов бомб ☠️

### Коммиты
- `a0bd7b5` — Add start menu and pause system
- `5bdd9bd` — Add pirate chest, golden key, and arrows shooting system
- `8e7fce1` — Add snake forked tongue
- `ca74e22` — Remove white teeth from bugs
- `ee55b0b` — Replace fangs with menacing mouth and two front teeth
- `edc80a4` — Add big scary glowing red eyes to bugs
- `7dda02d` — Add visible border around play area
- `c972a59` — Add scary fangs/teeth to bugs
- `6cb210e` — Add black bombs with sparking fuse
- `c431d30` — Increase bug size 1.5x and spawn 10 enemies
- `acb33c9` — Add crawling bug enemies with legs and antennae
- `92ad27a` — Add snake eyes with pupils
- `a4a24ab` — Reduce window size to 800x600
- `063e14e` — Set resolution to 2560x1440 (full screen)
- `acb33c9` — Add crawling bug enemies with legs and antennae
- `c2a860b` — Set resolution to 2560x1440 (full screen)
- `9a9feeb` — Fix snake movement speed

### Замечания
- Для сборки в Linux требуются системные библиотеки:
  ```bash
  sudo apt-get install -y libglfw3-dev libx11-dev libxrandr-dev libxi-dev \
    libxcursor-dev libxinerama-dev libxxf86vm-dev libxfixes-dev \
    libasound2-dev libpulse-dev
  ```
- Для Windows кросс-компиляция:
  ```bash
  GOOS=windows GOARCH=amd64 go build -o snake.exe
  ```
- Игра кроссплатформенная — работает на Linux, Windows, macOS, WebAssembly

---
**Go365 Challenge** — День 75 из 365 (16 марта 2026)

**Фокус:** Go + Ebitengine = геймдев на Go! 🎮🐍
