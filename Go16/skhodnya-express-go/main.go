package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🚂 Сходненский экспресс: Заказ отменён, но не дух!")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Дата: 16 января 2026 года")
	fmt.Println("Место действия: Москва - Сходня - Москва")
	fmt.Println()

	// Утренний ритуал
	fmt.Println("🌅 УТРЕННИЙ РИТУАЛ 16.01.2026:")
	ritual := []string{
		"Проснуться в 10:00",
		"Принять душ, ощутить свежесть",
		"Поздний завтрак: гречка с маслом",
		"Чай на кухне, мониторинг заказов",
		"ЗАКАЗ ПРИШЁЛ! 1000+ рублей: мешочки с песком для массажа",
	}

	for i, step := range ritual {
		fmt.Printf(" %d. %s\n", i+1, step)
		time.Sleep(400 * time.Millisecond)
	}

	// Путь на платформу
	fmt.Println("\n🚶 ПУТЬ К ЭЛЕКТРИЧКЕ:")
	fmt.Println("  Радостный выбежал из дома")
	fmt.Println("  Пешком от подъезда до платформы Ховрино")
	fmt.Println("  Электричка Москва → Сходня")
	time.Sleep(1 * time.Second)

	// Чтение книги
	fmt.Println("\n📚 ПОГРУЖЕНИЕ В КНИГУ:")
	fmt.Println("  Книга: 'Сам себе психолог' (серия софт-скиллов)")
	fmt.Println("  Бумажная версия - тактильное удовольствие")
	fmt.Println("  Полное погружение... смартфон забыт!")
	time.Sleep(1 * time.Second)

	// Отмена заказа - психологическая драма
	fmt.Println("\n💥 ПОВОРОТ СЮЖЕТА - ОТМЕНА ЗАКАЗА:")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("  Достаю телефон...")
	time.Sleep(700 * time.Millisecond)

	fmt.Println("  УВЕДОМЛЕНИЕ:")
	fmt.Println("  ─────────────────────────────────────")
	fmt.Println("  'Извините, ваш заказ был отменён.'")
	fmt.Println("  'Приносим извинения за неудобства.'")
	fmt.Println("  'На адрес забора товара можно вам не ехать.'")
	fmt.Println("  ─────────────────────────────────────")
	time.Sleep(1 * time.Second)

	fmt.Println("\n😱 ЭМОЦИОНАЛЬНАЯ РЕАКЦИЯ:")
	reactions := []string{
		"'Что, японский городовой?!'",
		"'Я был в шоке!'",
		"'Ведь я уже почти приехал на платформу Сходня!'",
		"Паника на 3 секунды...",
	}

	for _, reaction := range reactions {
		fmt.Printf("  • %s\n", reaction)
		time.Sleep(600 * time.Millisecond)
	}

	// Быстрая реакция
	fmt.Println("\n⚡ БЫСТРАЯ РЕАКЦИЯ:")
	fmt.Println("  Электричка остановилась на платформе Сходня")
	fmt.Println("  Кликнул в приложении 'Откликнуться'")
	fmt.Println("  Забрал другой заказ: КОРМ ДЛЯ ЖИВОТНЫХ")
	fmt.Println("  В последний момент выпрыгнул из электрички!")
	time.Sleep(1 * time.Second)

	// Путь до склада - трагикомичная часть
	fmt.Println("\n🥾 ЭПИЧЕСКИЙ ПУТЬ ДО СКЛАДА:")

	adventureSteps := []struct{
		step        string
		difficulty  int
		description string
	}{
		{"2 км пешком", 8, "По сугробам и непочищенным дорожкам"},
		{"Таможенный терминал", 6, "Крупный складской комплекс"},
		{"КПП", 5, "Стойка охраны, серьёзные лица"},
		{"Пропуск", 4, "Паспорт → охранник → пропуск"},
		{"Территория склада", 7, "Ещё 1 км по промышленной зоне"},
	}

	totalDifficulty := 0
	for i, step := range adventureSteps {
		fmt.Printf(" %d. %s (сложность: %d/10)\n", i+1, step.step, step.difficulty)
		fmt.Printf("    %s\n", step.description)
		totalDifficulty += step.difficulty
		time.Sleep(700 * time.Millisecond)
	}

	// Успешное завершение
	fmt.Println("\n✅ ФИНАЛЬНЫЙ АКТ:")
	fmt.Println("  Забрал мешок с кормом")
	fmt.Println("  Доставил на Юго-Западную")
	fmt.Println("  Получил оплату")
	fmt.Println()
	fmt.Println("  🎬 ИТОГ: Гоша - молодец! Справился с неожиданностью.")

	// Геймификация
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ЭПИЗОДА:")

	type Achievement struct {
		name   string
		points int
		emoji  string
	}

	achievements := []Achievement{
		{"Утренняя дисциплина", 25, "⏰"},
		{"Поймать жирный заказ", 75, "💰"},
		{"Чтение вместо скроллинга", 40, "📚"},
		{"Шоковая устойчивость к отмене", 65, "💥"},
		{"Быстрая реакция (новый заказ)", 90, "⚡"},
		{"Эпический путь (3 км по снегу)", 80, "🥾"},
		{"Взаимодействие с системой (КПП)", 30, "🛡️"},
		{"Успешная доставка", 100, "✅"},
	}

	totalPoints := 0
	for _, a := range achievements {
		fmt.Printf(" %s %s: +%d очков\n", a.emoji, a.name, a.points)
		totalPoints += a.points
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("\n📊 ВСЕГО ОЧКОВ: %d\n", totalPoints)

	level := 1
	if totalPoints > 200 {
		level = 2
	}
	if totalPoints > 350 {
		level = 3
	}
	if totalPoints > 500 {
		level = 4
	}

	levelNames := map[int]string{
		1: "Новичок в адаптации",
		2: "Опытный импровизатор",
		3: "Мастер непредвиденных ситуаций",
		4: "Легенда доставок",
	}

	fmt.Printf(" 🎯 УРОВЕНЬ: %d (%s)\n", level, levelNames[level])

	if totalPoints > 400 {
		fmt.Println(" 🏆 СЕКРЕТНОЕ ДОСТИЖЕНИЕ: 'Феникс доставок' - возродился из отмены!")
	}

	// Психологический анализ
	fmt.Println("\n🧠 ПСИХОЛОГИЧЕСКИЙ АНАЛИЗ СИТУАЦИИ:")

	psychologicalInsights := []string{
		"1. Отмена заказа = внешний фактор, неподконтрольный Гоше",
		"2. Шоковая реакция = естественная эмоция на неожиданность",
		"3. Быстрая адаптация = навык, который тренируется",
		"4. Чтение книги = прокачка софт-скиллов в движении",
		"5. Преодоление сугробов = метафора преодоления трудностей",
		"6. Успешная доставка = доказательство resilience (устойчивости)",
	}

	for _, insight := range psychologicalInsights {
		fmt.Printf(" %s\n", insight)
		time.Sleep(500 * time.Millisecond)
	}

	// Связь с программированием
	fmt.Println("\n💻 АНАЛОГИЯ С ПРОГРАММИРОВАНИЕМ НА GO:")

	analogies := []struct{
		situation   string
		programming string
	}{
		{"Отмена заказа", "Баг в продакшене или изменение требований"},
		{"Быстрая реакция", "Быстрое исправление кода или рефакторинг"},
		{"Путь по сугробам", "Работа с legacy кодом или документацией"},
		{"КПП и пропуск", "Code review и требования к коммитам"},
		{"Доставка корма", "Успешный деплой фичи"},
		{"Чтение книги", "Изучение документации Go"},
	}

	for i, analogy := range analogies {
		fmt.Printf(" %d. %s → %s\n", i+1, analogy.situation, analogy.programming)
		time.Sleep(400 * time.Millisecond)
	}

	// Философские выводы
	fmt.Println("\n💭 ФИЛОСОФСКИЕ ВЫВОДЫ ДНЯ:")

	philosophy := []string{
		"• Неожиданности - это не конец пути, а поворот на нём",
		"• Бумажная книга в 2026 году - акт цифрового детокса",
		"• 3 км по снегу закаляют не только тело, но и характер",
		"• Отменённый заказ может привести к лучшему заказу",
		"• Быстрая реакция ценится больше, чем идеальный план",
		"• Каждая трудность - это история для будущего резюме",
	}

	for _, p := range philosophy {
		fmt.Printf(" %s\n", p)
		time.Sleep(500 * time.Millisecond)
	}

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("                    DISCLAIMER")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n" + strings.Repeat("🚀", 60))
	fmt.Println("      ИСТОРИЯ ЗАВЕРШЕНА. УРОК УСВОЕН.")
	fmt.Println(strings.Repeat("🚀", 60))
}
