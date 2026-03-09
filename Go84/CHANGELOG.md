# 📝 CHANGELOG — Day 84 (10 марта 2026)

**Дата:** 10 марта 2026 года  
**День челленджа:** 84  
**Проект:** qwen_test — Социальные функции

---

## 🎯 Цель

Реализовать **социальные функции**: друзья, чат, дуэли, активность.

---

## ✅ Выполнено

### Backend (Go83)
- internal/social/service.go — SocialService
- internal/social/handlers.go — HTTP handlers
- internal/database/social_migrations.go — 5 таблиц БД

**Таблицы:**
- friends (друзья)
- friend_requests (запросы)
- messages (сообщения)
- challenges (дуэли)
- activity (активность)

**API (12 endpoints):**
- POST /api/social/friends/requests/send
- POST /api/social/friends/requests/accept
- POST /api/social/friends/requests/reject
- GET /api/social/friends/requests
- GET /api/social/friends
- POST /api/social/friends/remove
- POST /api/social/messages/send
- GET /api/social/messages
- GET /api/social/messages/unread
- POST /api/social/challenges/send
- GET /api/social/challenges
- GET /api/social/activity

### Frontend (Go84)
- static/social-store.js — SocialStore
- static/friends-component.js — FriendsComponent
- static/chat-component.js — ChatComponent
- static/activity-component.js — Activity + Challenges
- static/social-styles.css — Стили (500+ строк)

**Компоненты:**
- 👥 Друзья (запросы, поиск, онлайн)
- 💬 Чат (личные сообщения)
- 📜 Лента активности
- ⚔️ Дуэли (вызовы)

---

## 📊 Статистика

| Метрика | Значение |
|---------|----------|
| Файлов | 9 |
| Строк кода | ~2000 |
| Таблиц БД | 5 |
| API endpoints | 12 |
| Vue компонентов | 4 |

---

## 🎮 Навигация

Новые кнопки:
- **👥** — Друзья
- **💬** — Чат
- **📜** — Активность
- **⚔️** — Дуэли

---

## 🚀 Коммиты

**Go83:** `27226b0` — Социальные фичи (Backend)  
**Go84:** `54f00e2` — Социальные фичи (Frontend)

**Пуш:** https://github.com/Folombas/qwen_test

---

## 💭 Итоги

**Реализовано:**
- ✅ Друзья и запросы
- ✅ Личные сообщения
- ✅ Дуэли/вызовы
- ✅ Лента активности

**Влияние:**
- Социальное взаимодействие
- Соревновательный элемент
- Удержание пользователей

**День 84 завершён!** 🎉

---

## 🔮 Планы (Go85)

- [ ] Деплой на сервер
- [ ] HTTPS (Let's Encrypt)
- [ ] Email уведомления
- [ ] Больше вопросов (500+)
