# CHANGELOG.md — Go75 (16 марта 2026)

## [0.1.0] — 2026-03-16

### Добавлено
- ✅ Инициализация проекта Simple Snake на Ebitengine
- ✅ Базовая реализация игры Змейка
- ✅ Управление стрелками клавиатуры
- ✅ Механика роста змейки при поедании пищи
- ✅ Обнаружение столкновений (стены, собственный хвост)
- ✅ Подсчёт очков
- ✅ Экран Game Over с возможностью перезапуска (Enter)
- ✅ Документация: README.md, PLAN.md

### Технические детали
- Фреймворк: Ebitengine v2.9.9
- Язык: Go 1.x
- Размер сетки: 20x20 клеток
- Размер клетки: 20x20 пикселей
- Размер окна: 400x400 пикселей

### Замечания
- Для сборки требуются системные библиотеки: libglfw3-dev, libx11-dev, и другие X11/Audio dev-пакеты
- Команда установки зависимостей:
  ```bash
  sudo apt-get install -y libglfw3-dev libx11-dev libxrandr-dev libxi-dev \
    libxcursor-dev libxinerama-dev libxxf86vm-dev libxfixes-dev \
    libasound2-dev libpulse-dev
  ```

### Коммиты
- `3c6f109` — Initial commit: Snake game implementation
- `97745d7` — Add README and PLAN documentation

---
**Go365 Challenge** — День 75 из 365
