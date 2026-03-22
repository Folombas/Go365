# 📝 Отчёт о работе — День 84 (24 марта 2026)

**Дата:** 24 марта 2026 года  
**День челленджа:** 84  
**Время работы:** ~2 часа  
**Проект:** go_mario (playgo) + Go365  
**Версия:** 0.5.0

---

## 🎯 Цель дня

**Улучшение визуальных эффектов и геймплея!**

1. Система частиц для всех ключевых событий
2. Интеграция частиц в игровой процесс
3. Заготовка для powerups

---

## ✅ Выполненные задачи

### 1. Система частиц v2.0

**Добавлено 5 специализированных функций:**

#### spawnJumpParticles(x, y)
- **Событие:** Прыжок игрока
- **Цвет:** Серый (RGB 200, 200, 200) — пыль
- **Количество:** 8 частиц
- **Эффект:** Облачко пыли под ногами

#### spawnCoinParticles(x, y)
- **Событие:** Сбор монеты/разбивание блока
- **Цвет:** Золотой (RGB 255, 215, 0)
- **Количество:** 10 частиц
- **Эффект:** Всплеск вверх

#### spawnStompParticles(x, y)
- **Событие:** Уничтожение врага (stomp)
- **Цвет:** Коричневый (RGB 139, 69, 19)
- **Количество:** 12 частиц
- **Эффект:** Разлёт во все стороны

#### spawnHitParticles(x, y)
- **Событие:** Получение урона игроком
- **Цвет:** Красный (RGB 255, 50, 50)
- **Количество:** 15 частиц
- **Эффект:** Кровь/искры

#### spawnPowerupParticles(x, y)
- **Событие:** Получение бонуса
- **Цвет:** Розовый (RGB 255, 100, 100)
- **Количество:** 20 частиц
- **Эффект:** Круговой взрыв (360°)

---

### 2. Отрисовка частиц

**Добавлена функция `drawParticles()`:**

```go
func (g *Game) drawParticles(screen *ebiten.Image) {
    for _, p := range g.particles {
        vector.DrawFilledCircle(
            screen,
            float32(p.x),
            float32(p.y),
            p.size,
            p.color,
            true, // anti-alias
        )
    }
}
```

**Интеграция в игровой цикл:**
```go
func (g *Game) drawPlaying(screen *ebiten.Image) {
    screen.Fill(color.RGBA{100, 150, 200, 255})
    g.drawLevel(screen)
    g.drawPlayer(screen)
    g.drawParticles(screen)  // ← добавлено
    g.drawUI(screen)
}
```

---

### 3. Интеграция в геймплей

**Обновлены 4 функции игры:**

#### updatePlayer() — прыжок
```go
if keyPressed && p.onGround {
    p.vy = JumpForce
    p.onGround = false
    playSound(SoundJump)
    g.spawnJumpParticles(p.x+float64(p.width/2), p.y+float64(p.height))
}
```

#### hitBlock() — разбивание блока
```go
if tile == TileQuestion {
    g.level.tiles[x][y] = TileUsed
    g.player.coins++
    g.player.score += 200
    playSound(SoundCoin)
    g.spawnCoinParticles(float64(x*TileSize), float64(y*TileSize))
    g.spawnPowerupParticles(float64(x*TileSize+TileSize/2), float64(y*TileSize))
    
    // 10% шанс на powerup
    if rand.Float32() < 0.1 {
        // spawn powerup
    }
}
```

#### updateEnemies() — stomp врага
```go
if g.player.vy > 0 && playerBelowEnemy {
    enemy.squashed = true
    g.player.vy = -6  // bounce
    g.player.score += 100
    playSound(SoundStomp)
    g.spawnStompParticles(enemy.x+width/2, enemy.y+height/2)
}
```

#### playerHit() — получение урона
```go
if g.player.isBig {
    g.player.isBig = false
    g.player.isInvincible = true
    g.player.powerTimer = 120
    playSound(SoundHit)
    g.spawnHitParticles(g.player.x, g.player.y)
}
```

---

## 📊 Статистика

### Код

| Метрика | Значение | Изменения |
|---------|----------|-----------|
| Строк кода | 1442 | +78 |
| Функций добавлено | 6 | drawParticles + 5 spawn* |
| Вызовов частиц | 5 | в 4 функциях игры |
| Типов частиц | 5 | jump, coin, stomp, hit, powerup |

### Версии

| Версия | Изменения | Дата |
|--------|-----------|------|
| 0.1.0 | Базовая версия | 01.01.2026 |
| 0.2.0 | Платформы, враги, монеты | ~20.03.2026 |
| 0.3.0 | Спрайты игрока, врагов, монет | 22.03.2026 |
| 0.4.0 | Тайлы, анимация врагов, звуки | 23.03.2026 |
| 0.5.0 | Система частиц v2.0 | 24.03.2026 |

### Частицы (детально)

| Тип | Цвет (RGB) | Кол-во | Жизнь (кадры) | Скорость |
|-----|------------|--------|---------------|----------|
| Jump | 200,200,200 | 8 | 30-50 | ±2.5 |
| Coin | 255,215,0 | 10 | 40-60 | ±4 |
| Stomp | 139,69,19 | 12 | 25-40 | ±4.2 |
| Hit | 255,50,50 | 15 | 35-55 | ±4.2 |
| Powerup | 255,100,100 | 20 | 50-70 | 2.0 (круг) |

---

## 🔍 Технические детали

### Физика частиц

**Обновление в `updateParticles()`:**
```go
func (g *Game) updateParticles() {
    for i := len(g.particles) - 1; i >= 0; i-- {
        p := g.particles[i]
        p.x += p.vx
        p.y += p.vy
        p.vy += 0.2  // гравитация
        p.life--
        
        if p.life <= 0 {
            // удалить частицу
        }
    }
}
```

**Параметры:**
- Гравитация: 0.2 пикселя/кадр²
- Затухание: нет (частицы просто исчезают)
- Удаление: при life ≤ 0

### Генерация частиц

**Базовый паттерн:**
```go
for i := 0; i < count; i++ {
    g.particles = append(g.particles, &Particle{
        x: x, y: y,
        vx: rand(min, max),
        vy: rand(min, max),
        life: base + rand(variation),
        color: color,
        size: rand(minSize, maxSize),
    })
}
```

**Специальные эффекты:**

1. **Круговой взрыв (Powerup):**
   ```go
   angle := float64(i) * 2 * math.Pi / count
   vx := math.Cos(angle) * speed
   vy := math.Sin(angle) * speed
   ```

2. **Всплеск вверх (Coin):**
   ```go
   vy = -rand(5, 15) * 0.5  // только вверх
   ```

---

## 💭 Рефлексия

### Что получилось

- ✅ 5 типов частиц реализовано
- ✅ Частицы работают во всех ключевых моментах
- ✅ Отрисовка через vector.DrawFilledCircle
- ✅ Физика (гравитация, время жизни)
- ✅ Игра выглядит намного лучше

### Что можно улучшить

- ⬜ Добавить текстуры для частиц (спрайты)
- ⬜ Реализовать затухание (alpha blending)
- ⬜ Добавить больше вариаций
- ⬜ Оптимизировать удаление частиц

### Уроки дня

1. **Визуальная обратная связь важна** — игра feels better
2. **Простая физика работает** — гравитации достаточно
3. **Разные цвета = лучше восприятие** — игрок понимает что произошло
4. **Не переусердствовать** — 8-20 частиц достаточно

---

## 🎯 План на следующий день (День 85)

1. **Powerups**
   - Mushroom — увеличение игрока
   - Flower — огненная атака
   - Star — временная неуязвимость

2. **Улучшение ИИ врагов**
   - Piranha Plant из труб
   - Koopa с панцирем
   - Bee с полётом

3. **Оптимизация**
   - Pool частиц (не создавать каждый кадр)
   - Batch отрисовка

---

## 🔗 Ссылки

- **Go365:** https://github.com/Folombas/Go365
- **playgo:** https://github.com/Folombas/playgo
- **Ebitengine:** https://ebitengine.org/

---

**День 84 завершён!** 🎉

**Фокус на Go до конца 2026 года!** 🐍

**Никакого распыления! Только Go! Только Ebitengine!** 💪
