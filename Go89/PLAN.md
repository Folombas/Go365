# 📋 PLAN — Day 89 (26 марта 2026)

**Дата:** 26 марта 2026 года (четверг)
**День челленджа:** 89
**Проект:** blogAPI — продолжение разработки

---

## 🎯 Цель дня

**Завершение базового функционала Blog API!**

Сегодня нужно закончить все основные CRUD-операции и подготовить проект к развёртыванию.

---

## 📋 Задачи на сегодня

### Приоритет 1: User Profile полный CRUD 🔴

- [ ] **User Repository**
  - [ ] GetProfile
  - [ ] UpdateProfile
  - [ ] ChangePassword
  - [ ] UploadAvatar (опционально)

- [ ] **User Service**
  - [ ] Бизнес-логика профиля
  - [ ] Валидация данных
  - [ ] Хэширование пароля (bcrypt)
  - [ ] Проверка старого пароля

- [ ] **User Handler**
  - [ ] `GET /api/v1/users/me` — получить профиль
  - [ ] `PUT /api/v1/users/me` — обновить профиль
  - [ ] `PUT /api/v1/users/me/password` — сменить пароль
  - [ ] `POST /api/v1/users/me/avatar` — загрузить аватар (опционально)

- [ ] **Интеграция с main.go**
  - [ ] Инициализация UserRepository
  - [ ] Инициализация UserService
  - [ ] Инициализация UserHandler
  - [ ] Добавление routes

---

### Приоритет 2: Unit-тесты 🟡

- [ ] **Настроить тестирование**
  - [ ] Создать папку `internal/service/testdata`
  - [ ] Написать helper-функции для тестов
  - [ ] Настроить test fixtures

- [ ] **Тесты на AuthService**
  - [ ] TestRegister_Success
  - [ ] TestRegister_EmailExists
  - [ ] TestRegister_UsernameExists
  - [ ] TestLogin_Success
  - [ ] TestLogin_InvalidCredentials
  - [ ] TestGenerateTokens

- [ ] **Тесты на PostService**
  - [ ] TestCreatePost_Success
  - [ ] TestCreatePost_Validation
  - [ ] TestUpdatePost_AuthorCheck
  - [ ] TestDeletePost_Permissions
  - [ ] TestSlugGeneration (кириллица + латиница)

- [ ] **Тесты на CommentService**
  - [ ] TestCreateComment_Success
  - [ ] TestCreateComment_WithParent
  - [ ] TestLikeComment
  - [ ] TestUnlikeComment

- [ ] **Тесты на TagService**
  - [ ] TestCreateTag_Success
  - [ ] TestCreateTag_Validation
  - [ ] TestUpdateTag
  - [ ] TestDeleteTag

- [ ] **Запуск тестов**
  - [ ] `make test` — все тесты
  - [ ] `make test-coverage` — покрытие
  - [ ] Покрытие > 70%

---

### Приоритет 3: Docker deployment 🟢

- [ ] **Протестировать docker-compose**
  - [ ] `make docker-up` — запуск
  - [ ] Проверка PostgreSQL
  - [ ] Проверка миграций
  - [ ] Проверка API (health check)
  - [ ] `make docker-down` — остановка

- [ ] **Исправить проблемы (если есть)**
  - [ ] Ошибки подключения к БД
  - [ ] Ошибки миграций
  - [ ] Ошибки компиляции в Docker

- [ ] **Документация Docker**
  - [ ] Обновить README
  - [ ] Добавить примеры использования
  - [ ] Переменные окружения

---

### Приоритет 4: Финальные штрихи ⚪

- [ ] **Обновить API_ENDPOINTS.md**
  - [ ] Добавить User endpoints
  - [ ] Примеры запросов
  - [ ] Примеры ответов

- [ ] **Swagger документация** (опционально)
  - [ ] Установить swag
  - [ ] Сгенерировать docs
  - [ ] Добавить swagger UI

- [ ] **Проверка компиляции**
  - [ ] `make build` — сборка
  - [ ] `./bin/blog-api` — запуск
  - [ ] Тестовый запрос

---

## 📊 Метрики успеха

- [ ] User Profile CRUD реализован
- [ ] Unit-тесты написаны (покрытие > 70%)
- [ ] Docker работает (`make docker-up` ✅)
- [ ] Все коммиты сделаны
- [ ] Изменения отправлены в репозиторий

---

## 🔗 Ресурсы

- **Репозиторий blogAPI:** https://github.com/Folombas/blogAPI
- **Репозиторий Go365:** https://github.com/Folombas/Go365
- **Go Testing:** https://go.dev/doc/tutorial/add-a-test
- **Testify:** https://github.com/stretchr/testify
- **Docker Compose:** https://docs.docker.com/compose/

---

## 📝 Заметки

**Важно:**
- Начинать с User Profile (это база для остальных функций)
- Тесты писать после реализации (TBD или после кода)
- Docker тестировать на чистой системе (без кэша)

**Время работы:** ~4-6 часов

---

**Девиз дня:** Только Go! Только фокус! Только прогресс! 🐍🎯

**Go365 Challenge** — День 89 из 365
