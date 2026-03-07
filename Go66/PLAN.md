# Go365 Day 66 — План на 7 марта 2026

## 🎯 Главная цель

Превратить телеграм-бота Quiz Bot в **production-ready приложение** с правильной архитектурой, базой данных, логированием и тестами.

---

## 📋 Задачи на день

### 🔴 Критические (обязательно)

- [ ] **SQLite база данных**
  - [ ] Создать схему БД (users, user_quiz_progress, user_interview_progress)
  - [ ] Добавить миграции
  - [ ] Написать repository pattern для работы с пользователями
  - [ ] Перенести данные из JSON в SQLite
  - [ ] Протестировать сохранение/загрузку

- [ ] **Structured Logging**
  - [ ] Интегрировать `log/slog` (Go 1.21+)
  - [ ] Добавить уровни: debug, info, warn, error
  - [ ] Настроить JSON формат для продакшена
  - [ ] Добавить контекст (chatID, username) в логи

- [ ] **YAML Конфигурация**
  - [ ] Создать `configs/config.dev.yaml`
  - [ ] Создать `configs/config.prod.yaml`
  - [ ] Вынести настройки: токен, БД, логи, rate limit
  - [ ] Поддержка переменных окружения для секретов

---

### 🟡 Важные (желательно)

- [ ] **Rate Limiting**
  - [ ] Реализовать token bucket algorithm
  - [ ] Настроить лимиты на пользователя
  - [ ] Защита от спама и DDoS
  - [ ] Логирование превышений

- [ ] **Unit Tests**
  - [ ] Тесты на config package
  - [ ] Тесты на logger package
  - [ ] Тесты на ratelimit package
  - [ ] Тесты на models package
  - [ ] Покрытие > 50%

- [ ] **UI/UX Улучшения**
  - [ ] Menu Button с командами
  - [ ] Reply Keyboard (постоянные кнопки)
  - [ ] Кнопка "Следующий вопрос" после ответа
  - [ ] Команда `/keyboard` для возврата клавиатуры

---

### 🟢 Дополнительные (если останется время)

- [ ] **Docker**
  - [ ] Создать Dockerfile
  - [ ] Создать docker-compose.yml
  - [ ] Протестировать запуск в контейнере

- [ ] **CI/CD**
  - [ ] GitHub Actions для тестов
  - [ ] Auto-build при пуше

- [ ] **Мониторинг**
  - [ ] Prometheus metrics
  - [ ] Health check endpoint

- [ ] **Админка**
  - [ ] Команды для добавления вопросов
  - [ ] Статистика по пользователям

---

## ⏰ Расписание на день

| Время | Задача | Статус |
|-------|--------|--------|
| 09:00 - 10:00 | Планирование, анализ текущей архитектуры | ⬜ |
| 10:00 - 12:00 | YAML конфигурация + Structured logging | ⬜ |
| 12:00 - 13:00 | Обед | ⬜ |
| 13:00 - 15:00 | SQLite база данных + миграции | ⬜ |
| 15:00 - 16:00 | Rate limiting | ⬜ |
| 16:00 - 17:00 | Unit tests | ⬜ |
| 17:00 - 18:00 | UI/UX улучшения (Menu, Keyboard) | ⬜ |
| 18:00 - 19:00 | Тестирование, багфиксы | ⬜ |
| 19:00 - 20:00 | Документация, коммиты, пуш | ⬜ |

---

## 📝 Критерии приёмки

### ✅ День считается успешным, если:

1. [ ] Бот запускается с YAML конфигом
2. [ ] Данные сохраняются в SQLite (не JSON)
3. [ ] Логи структурированные, с уровнями
4. [ ] Rate limiting работает (30 запросов/мин)
5. [ ] 20+ unit тестов проходят
6. [ ] Menu Button показывает команды
7. [ ] Reply Keyboard отображается под полем ввода
8. [ ] Кнопка "Следующий вопрос" работает
9. [ ] Все коммиты запушены на GitHub
10. [ ] Документация в Go365/Go66 обновлена

---

## 🎯 Ожидаемый результат

**К концу дня бот должен иметь:**

```
quiz_bot/
├── cmd/quiz-bot/main.go        # Точка входа
├── internal/
│   ├── bot/                    # Логика бота
│   ├── config/                 # YAML конфигурация
│   ├── logger/                 # Structured logging
│   ├── models/                 # Модели данных
│   ├── ratelimit/              # Rate limiting
│   └── storage/                # SQLite + миграции
├── configs/
│   ├── config.dev.yaml
│   ├── config.prod.yaml
│   └── *.json
├── data/quiz_bot.db
├── deploy/
├── Makefile
└── README.md
```

**Версия:** v0.3.0

**Статус:** Production Ready ✅

---

## 🔗 Полезные ссылки

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [log/slog documentation](https://pkg.go.dev/log/slog)
- [gopkg.in/yaml.v3](https://github.com/go-yaml/yaml)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate)

---

## 📝 Заметки

- Не забыть обновить `.gitignore` для новых файлов
- Сделать бэкап старых данных перед миграцией
- Протестировать на реальном боте перед пушем
- Обновить README.md с новой структурой

---

**План создан:** 7 марта 2026, 00:00  
**Go365 Day:** 66  
**Статус:** ⏳ В процессе
