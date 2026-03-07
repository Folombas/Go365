# Go365 Day 65 — March 6, 2026

## 📋 Обзор дня

**Тема:** Телеграм-бот Quiz Bot — Добавление вопросов собеседования и рефакторинг структуры

**Цель:** Интегрировать вопросы с собеседований из qwen_test и улучшить архитектуру бота

---

## 🎯 Задачи дня

### ✅ Выполнено:

1. **Расчёт дня Go365**
   - 6 марта 2026 = День 65 (31 день января + 28 дней февраля + 6 дней марта)
   - Создана папка `Go365/Go65/`

2. **Добавлена кнопка "Вопросы к собеседованию"**
   - Кнопка: "💼 Вопросы к собеседованию - Gopher, Go Offer!"
   - Интеграция с вопросами из `qwen_test/questions.json`
   - 120 вопросов собеседования добавлено в бота

3. **Рефакторинг структуры проекта**
   - Принята Standard Go Project Layout
   - Созданы директории: `cmd/`, `internal/`, `configs/`, `deploy/`

4. **Коммиты и пуш на GitHub**
   - quiz_bot: Initial commit + interview questions feature
   - Go365: Day 65 documentation

---

## 📁 Изменения в quiz_bot

### Новые файлы:
```
quiz_bot/
├── cmd/quiz-bot/main.go          # Точка входа
├── internal/
│   ├── bot/bot.go                # Логика бота
│   └── models/models.go          # Модели данных
├── configs/
│   ├── questions.json            # 70 вопросов викторины
│   └── interview_questions.json  # 120 вопросов собеседования
├── deploy/
│   ├── Dockerfile
│   └── docker-compose.yml
├── Makefile
└── README.md                     # Обновлённая документация
```

### Новые функции:
- **Команда `/interview`** — вопрос собеседования
- **Кнопка `cmd_interview`** — запуск вопросов собеседования
- **Отдельный трек прогресса** — `InterviewAsked` для вопросов собеседования
- **Callback `interview_<id>_<opt>`** — обработка ответов

### Обновлённые функции:
- **`/reset`** — теперь сбрасывает и вопросы собеседования тоже
- **`/help`** — добавлена информация о `/interview`
- **Главное меню** — добавлена новая кнопка

---

## 📊 Статистика вопросов

| Категория | Количество | ID вопросов |
|-----------|------------|-------------|
| Викторина | 70 | 31-100 |
| Собеседование | 120 | 1-120 |
| **Всего** | **190** | — |

### Темы вопросов собеседования:
- ООП в Go (наследование, инкапсуляция, полиморфизм)
- Пакеты и модули
- Типы данных (int, uint, rune, byte)
- Преобразования типов (strconv, atoi)
- iota и константы
- Строки и Unicode
- Массивы, слайсы, map

---

## 🔧 Технические детали

### Изменения в коде:

**UserData struct:**
```go
type UserData struct {
    TotalEXP       int   `json:"total_exp"`
    CorrectAnswers int   `json:"correct_answers"`
    WrongAnswers   int   `json:"wrong_answers"`
    Level          int   `json:"level"`
    AskedQuestions []int `json:"asked_questions"`
    InterviewAsked []int `json:"interview_asked"`  // НОВОЕ!
}
```

### Git коммиты:
```
dc5917b  Initial commit: Go quiz telegram bot with questions
f70ae08  Add version constant (v0.2.0) and display it in startup log
4a85bbe  Refactor: adopt Standard Go Project Layout
```

---

## 🚀 Запуск бота

```bash
cd quiz_bot
make run
# или
go run cmd/quiz-bot/main.go -config configs/config.dev.yaml
```

---

## 📝 Итоги дня

**Достигнуто:**
- ✅ 120 вопросов собеседования интегрировано
- ✅ Структура проекта приведена к стандарту Go
- ✅ Документация обновлена
- ✅ Код закоммичен и запушен на GitHub

**Следующий шаг:**
- Добавить SQLite базу данных
- Добавить структурированное логирование
- Добавить rate limiting
- Написать unit-тесты

---

## 🔗 Ссылки

- Репозиторий бота: https://github.com/Folombas/quiz_bot
- Репозиторий Go365: https://github.com/Folombas/Go365
- Вопросы собеседования: `qwen_test/questions.json`
