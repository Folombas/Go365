# План разработки — Go64

**Дата**: 5 марта 2026 года  
**Проект**: [qwen_test](https://github.com/Folombas/qwen_test)

---

## 🎯 Цель сессии

Продолжить разработку веб-приложения qwen_test, добавив новую интерактивную фичу.

---

## 📋 Задача: Страница праздников

### Описание
Добавить кнопку **«Какой праздник сегодня»** в левый нижний угол главного экрана.

### Требования

#### Кнопка вызова
- **Расположение**: левый нижний угол
- **Стиль**: Material Design, в единой стилистике с другими кнопками
- **Размер**: минимум 48x48px (touch-оптимизация)
- **Иконка/текст**: 🎉 или текст «Какой праздник сегодня»

#### Страница праздников
- **Навигация**: переход по клику на кнопку (отдельная страница или modal/overlay)
- **Дизайн**:
  - Красивая, адаптивная вёрстка
  - Стильные шрифты (Google Fonts или системные)
  - Соответствие текущей теме (тёмная/светлая)
- **Контент**:
  - Список праздников на сегодня (5 марта)
  - Международные, российские, профессиональные
- **Адаптивность**:
  - Мобильные (≤480px)
  - Планшеты (481-1024px)
  - Десктоп (>1024px)

#### Возврат на главную
- Кнопка «Назад» или крестик закрытия
- Анимация перехода

---

## 🛠 Техническая реализация

### Backend (Go)
```go
// Новый handler для страницы праздников
func holidaysHandler(w http.ResponseWriter, r *http.Request) {
    // Рендер шаблона holidays.html
}
```

### Frontend
- **HTML**: отдельный template `holidays.html`
- **CSS**: стили для страницы праздников (адаптивные)
- **JavaScript**: обработчик клика, возможно fetch API для загрузки данных

### Данные о праздниках
- **Вариант 1**: захардкоженный список для 5 марта
- **Вариант 2**: внешний API (например, Calend.ru или similar)
- **Вариант 3**: JSON-файл с праздниками

---

## ✅ Критерии готовности

- [x] Кнопка отображается в левом нижнем углу
- [x] По клику открывается страница праздников
- [x] Страница адаптивная и красивая
- [x] Список праздников актуален для 5 марта
- [x] Работает возврат на главную
- [x] Сохраняется тема (тёмная/светлая)
- [x] Код закоммичен и отправлен в репозиторий

---

## 📝 Заметки

> **Важно**: Сегодня я продолжу работу над существующим проектом, а не начинаю с нуля.
> Это уже вторая сессия в рамках одного проекта — тактика работает! 💪

---

## 🚀 Коммиты (заполняется по ходу работы)

```
commit 80741e7 - Go64: Add holidays page feature

- Added /holidays route with handler for March 5th holidays
- Created holidaysTmpl template with responsive design
- Added holidays button in bottom-left corner of main page
- Implemented 5 holidays: Efficiency Day, Archivist Day, Sports Day, Pushkin Birthday, Orthodox holiday
- Added theme toggle support (dark/light) with localStorage persistence
- Fully responsive design: mobile (≤480px), tablet (481-768px), desktop (>769px)
- Material Design styling with gradient backgrounds and animations
- Back button navigation to main page
- Google Fonts (Montserrat, Open Sans) for typography
```
