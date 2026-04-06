# 🧪 ОТЧЁТ — День 97 (7 апреля 2026, Вторник)

**Дата:** 7 апреля 2026
**День челленджа:** 97
**Время работы:** ~4 часа
**Проект:** blogAPI (repo) + Go365
**Версия:** тесты v1.0.0

---

## 🎯 Цель дня

Написать comprehensive unit-тесты для Blog API, покрыть ключевые слои приложения тестами и получить работающую CI-готовную кодовую базу.

---

## ✅ Выполненные задачи

### 1. 🧪 Тесты для pkg/utils (slug generation) — 100% coverage
**Файл:** `pkg/utils/slug_test.go` (91 строка, 14 тест-кейсов)

**Что тестируется:**
- Латиница, кириллица, смешанный текст
- Числа в slug'ах
- Удаление специальных символов
- Множественные пробелы → один дефис
- Обрезка краёв, пустая строка, только спецсимволы
- Ё-буква, подчёркивания

```bash
=== RUN   TestGenerateSlug
--- PASS: TestGenerateSlug (0.00s)
    --- PASS: 14 подтестов, все прошли
```

---

### 2. 📦 Тесты для internal/model — 100% coverage
**Файл:** `internal/model/model_test.go` (272 строки, 11 тестов)

**Что тестируется:**
| Модель | Тесты |
|--------|-------|
| `User` | `NewUser` конструктор, `FullName()` |
| `Post` | `NewPost` конструктор, `IsPublished()` |
| `Comment` | `NewComment` конструктор |
| `Tag` | `NewTag` конструктор |
| `RefreshToken` | `NewRefreshToken`, `IsExpired()`, `IsValid()` |

---

### 3. ⚙️ Тесты для internal/config — 80.4% coverage
**Файл:** `internal/config/config_test.go` (256 строк, 8 тестов)

**Что тестируется:**
- Загрузка дефолтных значений
- Кастомные переменные окружения
- Запрет дефолтного JWT-секрета в production
- Валидация BcryptCost (3 и 13 → ошибка)
- `IsDevelopment()` / `IsProduction()`
- `DSN()` для SQLite и PostgreSQL
- `CORSAllowOrigins` парсинг из строки

---

### 4. 🎭 Mock-репозитории для всех 5 слоёв
**Файл:** `internal/mocks/mock_repositories.go` (672 строки)

Реализованы 5 полноценных mock-репозиториев с интерфейсами:
- `MockUserRepository` — 8 методов
- `MockPostRepository` — 11 методов
- `MockTokenRepository` — 4 метода
- `MockCommentRepository` — 8 методов
- `MockTagRepository` — 12 методов

Каждый моки имеет:
- In-memory хранилище (map)
- `SetError()` / `ClearErrors()` для инжекции ошибок
- `Reset()` для сброса между тестами
- Пагинацию, индексы по slug/email и т.д.

---

### 5. 🔧 Тесты для internal/service — 65.2% coverage
**Файл:** `internal/service/service_test.go` (1216 строк, 39 тестов)

#### AuthService (12 тестов)
| Тест | Что проверяет |
|------|--------------|
| `TestAuthServiceRegister` | Регистрация нового пользователя |
| `TestAuthServiceRegisterDuplicateEmail` | Ошибка дубликата email |
| `TestAuthServiceRegisterDuplicateUsername` | Ошибка дубликата username |
| `TestAuthLogin` | Успешный логин + генерация токенов |
| `TestAuthLoginInvalidEmail` | Неверный email |
| `TestAuthLoginInvalidPassword` | Неверный пароль |
| `TestAuthLogout` | Логаут + отзыв токена |
| `TestAuthRefreshToken` | Обновление пары токенов |
| `TestAuthRefreshTokenInvalid` | Неверный refresh-токен |
| `TestAuthValidateToken` | Валидация JWT |
| `TestAuthValidateTokenWrongSecret` | JWT с неверным секретом |
| `TestAuthServiceGenerateTokens` | Прямая генерация токенов |
| `TestAuthServiceGetJWTExpiration` | Время жизни JWT |

#### PostService (11 тестов)
| Тест | Что проверяет |
|------|--------------|
| `TestPostServiceCreatePost` | Создание поста |
| `TestPostServiceCreatePostTitleTooShort` | Валидация заголовка |
| `TestPostServiceCreatePostContentTooShort` | Валидация контента |
| `TestPostServiceCreatePostAuthorNotFound` | Несуществующий автор |
| `TestPostServiceGetPostByID` | Получение поста + счётчик просмотров |
| `TestPostServiceGetPostByIDNotFound` | Пост не найден |
| `TestPostServiceGetPosts` | Пагинация постов |
| `TestPostServiceUpdatePost` | Обновление поста автором |
| `TestPostServiceUpdatePostWrongAuthor` | Запрет редактирования чужого поста |
| `TestPostServiceDeletePost` | Удаление поста |
| `TestPostServicePublishPost` | Публикация поста |

#### CommentService (9 тестов)
| Тест | Что проверяет |
|------|--------------|
| `TestCommentServiceCreateComment` | Создание комментария |
| `TestCommentServiceCreateCommentContentTooShort` | Валидация контента |
| `TestCommentServiceCreateReply` | Вложенный комментарий (reply) |
| `TestCommentServiceGetCommentsByPost` | Список комментариев с пагинацией |
| `TestCommentServiceUpdateComment` | Обновление комментария |
| `TestCommentServiceUpdateCommentWrongAuthor` | Запрет редактирования чужого |
| `TestCommentServiceDeleteComment` | Удаление комментария |
| `TestCommentServiceLikeComment` | Лайк комментария |
| `TestCommentServiceUnlikeComment` | Снятие лайка |

#### TagService (8 тестов)
| Тест | Что проверяет |
|------|--------------|
| `TestTagServiceCreateTag` | Создание тега |
| `TestTagServiceCreateTagTooShort` | Валидация имени |
| `TestTagServiceCreateTagDuplicate` | Дубликат тега |
| `TestTagServiceGetTagBySlug` | Получение по slug |
| `TestTagServiceGetTags` | Список тегов |
| `TestTagServiceUpdateTag` | Обновление тега |
| `TestTagServiceDeleteTag` | Удаление тега |
| `TestTagServiceAddTagsToPost` | Добавление тегов к посту |
| `TestTagServiceAddTagsToPostPostNotFound` | Несуществующий пост |

---

### 6. 🔧 Обновление Makefile

**Новые команды:**
```bash
make test            # Запуск тестов (internal + pkg)
make test-coverage   # Тесты + HTML coverage report
make test-cover-view # Тесты + текстовый coverage в терминале
```

---

### 7. 🐛 Баг-фикс

**`cmd/api/main.go`:** Заменён `%w` на `%v` в `log.Fatalf` (не поддерживает error-wrapping).

---

## 📊 Статистика дня

| Метрика | Значение |
|---------|----------|
| **Новых файлов** | 5 |
| **Строк тестового кода** | 2507 |
| **Тест-кейсов** | 49 |
| **Mock-репозиториев** | 5 |
| **Mock-методов** | 43 |
| **Покрытие model** | 100.0% |
| **Покрытие utils** | 100.0% |
| **Покрытие config** | 80.4% |
| **Покрытие service** | 65.2% |
| **Все тесты проходят** | ✅ 49/49 |

---

## 🔍 Технические детали

### Архитектура тестов

```
internal/mocks/mock_repositories.go   ← Mock-репозитории (in-memory maps)
internal/model/model_test.go          ← Тесты моделей (конструкторы, хелперы)
internal/config/config_test.go        ← Тесты конфига (env vars, валидация)
internal/service/service_test.go      ← Тесты сервисов (с моками репозиториев)
pkg/utils/slug_test.go                ← Тесты утилит (slug generation)
```

### Паттерн моков

Каждый mock-репозиторий реализует:
1. **In-memory хранилище** — `map[int64]*Model` + индексы
2. **`SetError(method, err)`** — инжекция ошибок для негативных сценариев
3. **`ClearErrors()`** — сброс ошибок между тестами
4. **`Reset()`** — полный сброс состояния

### Что НЕ покрыто тестами

| Пакет | Причина |
|-------|---------|
| `cmd/api` | Интеграционный тест, требует БД |
| `internal/database` | Зависит от PostgreSQL |
| `internal/repository` | PostgreSQL-специфичный код |
| `internal/handler` | HTTP handler'ы (нужен httptest) |
| `internal/middleware` | HTTP middleware (нужен httptest) |
| `pkg/logger` | Логгер (stdout/stderr) |

---

## 💭 Рефлексия

### Что получилось
- ✅ 49 тестов за один день — отличный темп
- ✅ 100% покрытие model и utils
- ✅ Mock-репозитории переиспользуемы для будущих тестов
- ✅ Все сервис-тесты покрывают happy path + edge cases
- ✅ Makefile обновлён с тремя test-командами

### Что можно улучшить
- ⬜ Добавить HTTP handler тесты (httptest)
- ⬜ Добавить repository интеграционные тесты (testcontainers-go)
- ⬜ Покрыть тестами middleware (JWT, CORS, Logger)
- ⬜ Добавить benchmarks для критичных участков

### Уроки дня
1. Моки через in-memory maps — это просто и эффективно
2. Table-driven tests — стандарт де-факто в Go
3. Тесты на сервисы с моками репозиториев дают лучшее соотношение effort/coverage
4. 65% покрытия сервисов — хороший старт, но цель — 80%+

---

## 🎯 План на следующий день (День 98)
1. Написать HTTP handler тесты (httptest) для auth/post handlers
2. Покрыть тестами middleware слой (JWT middleware)
3. Добавить benchmarks для slug generation и валидации
4. Довести общее покрытие до 80%+

---

## 🔗 Ссылки
- **Go365:** https://github.com/Folombas/Go365
- **Blog API repo:** https://github.com/Folombas/blogAPI
- **Testing in Go:** https://go.dev/doc/tutorial/add-a-test
- **Table-driven tests:** https://github.com/golang/go/wiki/TableDrivenTests

---

**День 97 завершён!** 🎉
**Фокус на Go до конца 2026 года!** 🐍
**49 тестов написаны, 2507 строк тестового кода, все зелёные!** 💪
