# Go365 Day 66 — March 7, 2026

## 📋 Обзор дня

**Тема:** Quiz Bot — Production-Ready Features

**Цель:** Превратить бота в production-ready приложение с правильной архитектурой, базой данных, логированием и тестами

---

## 🎯 Задачи дня

### ✅ Выполнено:

1. **YAML Конфигурация** — централизованное управление настройками
2. **Structured Logging (slog)** — структурированное логирование
3. **SQLite + Миграции** — надёжное хранение данных
4. **Rate Limiting** — защита от спама (token bucket algorithm)
5. **Unit Tests** — 20+ тестов для ключевых компонентов
6. **Menu Button** — команды в меню Telegram
7. **Reply Keyboard** — постоянные кнопки под полем ввода
8. **/keyboard команда** — показать клавиатуру
9. **Кнопка "Следующий вопрос"** — улучшение UX

---

## 📁 Новая структура проекта

```
quiz_bot/
├── cmd/quiz-bot/
│   └── main.go                   # Точка входа (35 строк)
├── internal/
│   ├── bot/
│   │   └── bot.go                # Логика бота (695 строк)
│   ├── config/
│   │   ├── config.go             # YAML конфигурация
│   │   └── config_test.go        # 4 теста
│   ├── logger/
│   │   ├── logger.go             # slog wrapper
│   │   └── logger_test.go        # 7 тестов
│   ├── models/
│   │   ├── models.go             # Модели данных
│   │   └── models_test.go        # 4 теста
│   ├── ratelimit/
│   │   ├── ratelimit.go          # Token bucket
│   │   └── ratelimit_test.go     # 8 тестов
│   └── storage/
│       ├── storage.go            # SQLite + миграции
│       └── user_repository.go    # Репозиторий пользователей
├── configs/
│   ├── config.dev.yaml           # Dev окружение
│   ├── config.prod.yaml          # Prod окружение
│   ├── questions.json            # 70 вопросов викторины
│   └── interview_questions.json  # 120 вопросов собеседования
├── data/
│   └── quiz_bot.db               # SQLite база данных
├── deploy/
│   ├── Dockerfile
│   └── docker-compose.yml
├── Makefile
├── README.md
├── .env.example
└── .gitignore
```

---

## 🔧 Новые компоненты

### 1. YAML Конфигурация (`internal/config/config.go`)

**Файлы:**
- `configs/config.dev.yaml` — разработка
- `configs/config.prod.yaml` — продакшен

**Пример конфигурации:**
```yaml
bot:
  token_env_var: TELEGRAM_BOT_TOKEN
  max_connections: 100
  timeout: 60

database:
  type: sqlite
  sqlite:
    path: data/quiz_bot.db

log:
  level: debug
  format: text
  output: stdout

rate_limit:
  enabled: true
  requests_per_min: 30
  burst_size: 10
```

**Преимущества:**
- ✅ Разные конфиги для dev/prod
- ✅ Секреты в переменных окружения
- ✅ Легко менять без перекомпиляции

---

### 2. Structured Logging (`internal/logger/logger.go`)

**Используется:** Go 1.21+ `log/slog`

**Уровни логирования:**
- `debug` — отладочная информация
- `info` — обычная информация
- `warn` — предупреждения
- `error` — ошибки

**Пример логов:**
```
time=2026-03-07T03:27:22.625+03:00 level=INFO msg="Bot authorized" username=golang_free_bot version=0.3.0
time=2026-03-07T03:27:22.626+03:00 level=INFO msg="Database initialized" path=data/quiz_bot.db
time=2026-03-07T03:27:22.626+03:00 level=INFO msg="Rate limiting enabled" requests_per_min=30 burst_size=10
time=2026-03-07T03:27:22.627+03:00 level=INFO msg="Questions loaded" count=70
time=2026-03-07T03:27:22.627+03:00 level=INFO msg="Interview questions loaded" count=120
time=2026-03-07T03:27:22.627+03:00 level=INFO msg="Bot started"
time=2026-03-07T03:28:08.476+03:00 level=INFO msg="Command received" chat_id=362689512 command=start
```

**Форматы:**
- `text` — человекочитаемый (dev)
- `json` — для ELK/Grafana (prod)

---

### 3. SQLite База Данных (`internal/storage/`)

**Таблицы:**

**users:**
```sql
CREATE TABLE users (
    chat_id INTEGER PRIMARY KEY,
    total_exp INTEGER NOT NULL DEFAULT 0,
    correct_answers INTEGER NOT NULL DEFAULT 0,
    wrong_answers INTEGER NOT NULL DEFAULT 0,
    level INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**user_quiz_progress:**
```sql
CREATE TABLE user_quiz_progress (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    answered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES users(chat_id),
    UNIQUE(chat_id, question_id)
);
```

**user_interview_progress:**
```sql
CREATE TABLE user_interview_progress (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    answered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES users(chat_id),
    UNIQUE(chat_id, question_id)
);
```

**Миграции:**
- Автоматическое применение при старте
- Таблица `schema_migrations` для отслеживания
- Индексы для производительности

**Преимущества перед JSON:**
- ✅ ACID транзакции
- ✅ Нет потерь данных при сбое
- ✅ Быстрые запросы с индексами
- ✅ Легко масштабировать на PostgreSQL

---

### 4. Rate Limiting (`internal/ratelimit/ratelimit.go`)

**Алгоритм:** Token Bucket

**Настройки:**
```yaml
rate_limit:
  enabled: true
  requests_per_min: 30
  burst_size: 10
```

**Принцип работы:**
- У каждого пользователя свой bucket
- 30 запросов в минуту = 0.5 запроса в секунду
- Burst size = 10 (можно сделать 10 быстрых запросов)
- Токены пополняются со временем

**Защита:**
- ✅ От спама
- ✅ От DDoS
- ✅ От злоупотреблений

---

### 5. Unit Tests

**Покрытие:**

| Пакет | Файл тестов | Количество тестов |
|-------|-------------|-------------------|
| `config` | `config_test.go` | 4 |
| `logger` | `logger_test.go` | 7 |
| `models` | `models_test.go` | 4 |
| `ratelimit` | `ratelimit_test.go` | 8 |
| **Всего** | — | **23** |

**Запуск тестов:**
```bash
make test
# или
go test -v ./...
```

**Пример теста:**
```go
func TestRateLimiter_Allow(t *testing.T) {
    rl := NewRateLimiter(60, 5) // 1 запрос в секунду, burst 5
    chatID := int64(12345)

    // Первые 5 запросов должны пройти (burst)
    for i := 0; i < 5; i++ {
        if !rl.Allow(chatID) {
            t.Errorf("Request %d should be allowed", i)
        }
    }

    // 6-й запрос должен быть отклонён
    if rl.Allow(chatID) {
        t.Error("Request 6 should be rate limited")
    }
}
```

---

## 🎨 UI/UX Улучшения

### 1. Menu Button

**Команды в меню:**
```
🧠 Начать викторину
💼 Вопросы к собеседованию
📊 Моя статистика
🏆 Таблица лидеров
🔄 Сбросить прогресс
ℹ️ Помощь
⌨️ Показать клавиатуру
```

**Реализация:**
```go
func (b *Bot) setMenuCommands() error {
    commands := []tgbotapi.BotCommand{
        {Command: "quiz", Description: "🧠 Начать викторину"},
        {Command: "interview", Description: "💼 Вопросы к собеседованию"},
        // ...
    }
    cmdConfig := tgbotapi.NewSetMyCommands(commands...)
    _, err := b.api.Request(cmdConfig)
    return err
}
```

---

### 2. Reply Keyboard

**Постоянные кнопки под полем ввода:**
```
┌─────────────────────────────┐
│ 🧠 Начать викторину         │
├─────────────────────────────┤
│ 💼 Вопросы к собеседованию  │
├─────────────────────────────┤
│ 📊 Моя статистика │ 🏆 ...  │
├─────────────────────────────┤
│ 🔄 Сбросить │ ℹ️ Помощь     │
└─────────────────────────────┘
```

**Обработчик:**
```go
func (b *Bot) handleKeyboardButton(chatID int64, text string) {
    switch text {
    case "🧠 Начать викторину":
        b.handleQuizCommand(chatID)
    case "💼 Вопросы к собеседованию":
        b.handleInterviewCommand(chatID)
    // ...
    }
}
```

---

### 3. Команда `/keyboard`

**Назначение:** Показать клавиатуру тем, кто её скрыл

**Реализация:**
```go
func (b *Bot) sendKeyboard(chatID int64) {
    text := "⌨️ Вот твоя клавиатура с кнопками!"
    msg := tgbotapi.NewMessage(chatID, text)
    msg.ReplyMarkup = getMainKeyboard()
    b.send(msg)
}
```

---

### 4. Кнопка "Следующий вопрос"

**После ответа на вопрос:**
```
✅ Правильно! +5 EXP
┌─────────────────────────┐
│ ➡️ Следующий вопрос     │
└─────────────────────────┘
```

**Для собеседования:**
```
❌ Неправильно. Правильный ответ: Stringer
┌─────────────────────────────────────┐
│ ➡️ Следующий вопрос собеседования  │
└─────────────────────────────────────┘
```

---

## 📊 База данных в действии

**Просмотр данных:**
```bash
# Установить sqlite3
sudo apt install sqlite3

# Подключиться к базе
sqlite3 data/quiz_bot.db

# Показать пользователей
.mode column
.headers on
SELECT * FROM users;

# Результат:
chat_id    total_exp  correct_answers  wrong_answers  level  created_at           updated_at
---------  ---------  ---------------  -------------  -----  -------------------  -------------------
362689512  5          1                2              1      2026-03-07 00:35:44  2026-03-07 00:36:28
```

---

## 🔧 Исправления багов

### 1. "no such table: users"

**Проблема:** Миграции не применялись

**Решение:** Убран `embed.FS`, миграции применяются напрямую в коде

**Коммит:** `604c3f2`

---

### 2. Прогресс не записывался

**Проблема:** Таблицы `user_quiz_progress` были пустыми

**Решение:** Добавлен вызов `RecordAnswer()` при ответе на вопрос

**Коммит:** `ff760dc`

---

## 📦 Git Коммиты (March 7, 2026)

```
139e5f0  Major refactor: Add production-grade features
         - YAML configuration
         - Structured logging with slog
         - SQLite database with migrations
         - Rate limiting
         - Unit tests (20+ tests)

604c3f2  Fix: Database migration - use inline SQL instead of embed.FS

ff760dc  Fix: Record answered questions in database

022cb80  Feature: Add 'Next question' button after answering

37cbd88  Feature: Add Menu Button with commands

ebb8383  Feature: Add Reply Keyboard (persistent buttons)

a68ecec  Feature: Add /keyboard command to show Reply Keyboard
```

---

## 🚀 Запуск бота

### Разработка:
```bash
cd quiz_bot
make run
# или
go run cmd/quiz-bot/main.go -config configs/config.dev.yaml
```

### Продакшен:
```bash
make run-prod
# или
go run cmd/quiz-bot/main.go -config configs/config.prod.yaml
```

### Docker:
```bash
make docker-build
make docker-up
```

### Тесты:
```bash
make test
# С покрытием
make test-coverage
```

---

## 📈 Версии бота

| Версия | Дата | Изменения |
|--------|------|-----------|
| v0.1.0 | Mar 3 | Initial commit, 70 вопросов |
| v0.2.0 | Mar 6 | +120 вопросов собеседования |
| v0.3.0 | Mar 7 | Production-ready features |

---

## 📝 Итоги дня

**Достигнуто:**
- ✅ YAML конфигурация для dev/prod
- ✅ Structured logging с уровнями
- ✅ SQLite база с миграциями
- ✅ Rate limiting (token bucket)
- ✅ 23 unit-теста
- ✅ Menu Button с командами
- ✅ Reply Keyboard (постоянные кнопки)
- ✅ Команда `/keyboard`
- ✅ Кнопка "Следующий вопрос"
- ✅ Полная документация

**Статус проекта:**
- 🟢 Production-ready
- 🟢 Proper architecture
- 🟢 Well tested
- 🟢 Documented

**Следующие шаги:**
- 🔵 Admin panel для управления вопросами
- 🔵 Prometheus metrics
- 🔵 CI/CD pipeline
- 🔵 Больше вопросов

---

## 🔗 Ссылки

- Репозиторий бота: https://github.com/Folombas/quiz_bot
- Репозиторий Go365: https://github.com/Folombas/Go365
- Telegram бот: @golang_free_bot

---

## 📸 Скриншоты функционала

### Главное меню:
- Menu Button слева от поля ввода
- Reply Keyboard под полем ввода
- Inline-кнопки в сообщениях

### Логирование:
```
INFO Bot authorized username=golang_free_bot version=0.3.0
INFO Database initialized path=data/quiz_bot.db
INFO Rate limiting enabled requests_per_min=30 burst_size=10
INFO Questions loaded count=70
INFO Interview questions loaded count=120
INFO Bot started
```

### База данных:
```
sqlite> SELECT * FROM users;
chat_id    total_exp  correct_answers  wrong_answers  level
---------  ---------  ---------------  -------------  -----
362689512  5          1                2              1
```

---

**Go365 Day 66 — Complete! 🎉**
