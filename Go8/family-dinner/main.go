package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MAX_STAT = 100
	START_FOCUS = 40
	START_HAPPINESS = 60
)

type Stats struct {
	Focus     int // Уровень концентрации на Go
	Wisdom    int // Мудрость против троллей
	Happiness int // Семейное счастье
}

type Character struct {
	Name  string
	Stats Stats
}

type Event struct {
	Description string
	Emoji       string
	Duration    time.Duration
}

// Value receiver: анализ ситуации БЕЗ изменения состояния
func (c Character) AnalyzeTroll(trollMessage string) string {
	patterns := []string{
		"фабрика ёлочных игрушек",
		"лучше бы валялся на диване",
		"ты ничего не добьёшься",
	}

	for _, pattern := range patterns {
		if strings.Contains(strings.ToLower(trollMessage), pattern) {
			return fmt.Sprintf("⚠️ ТРИГГЕР ОБНАРУЖЕН: '%s'", pattern)
		}
	}
	return "✅ ЧИСТАЯ КАРМА: тролль не опасен"
}

// Pointer receiver: изменение состояния при семейном ужине
func (c *Character) FamilyDinner(belyashiCount int, weather string) {
	fmt.Printf("\n%s❤️ СЕМЕЙНЫЙ УЖИН В МЕТЕЛЬ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   На улице: %s\n", weather)
	fmt.Printf("   Мама подаёт %d сочных беляшей по-домашнему\n", belyashiCount)
	fmt.Println("   За окном вьюга, но в доме — тепло и уют")

	// Восстановление ресурсов
	c.Stats.Happiness = min(MAX_STAT, c.Stats.Happiness+35)
	c.Stats.Focus = min(MAX_STAT, c.Stats.Focus+25)
	c.Stats.Wisdom += 15
}

// Value receiver: безопасное принятие решений
func (c Character) ShouldEngageTroll(trollMessage string) bool {
	fmt.Printf("\n%s🧠 ГОША ДУМАЕТ:%s\n", ansi("1;34"), ansi("0"))
	fmt.Printf("   \"Если я отвечу на '%s'...\"\n", trollMessage)
	fmt.Println("   • Потеряю 15 минут жизни")
	fmt.Println("   • Не допишу модуль про Value/Pointer Receivers")
	fmt.Println("   • Пропущу горячие беляши с родными")

	return strings.Contains(strings.ToLower(trollMessage), "диван") &&
		c.Stats.Focus < 50 // Только если фокус низкий!
}

// Pointer receiver: активация защиты от троллей
func (c *Character) ActivateFocusShield() {
	fmt.Printf("\n%s🛡️ АКТИВИРОВАН ЩИТ ФОКУСА:%s\n", ansi("1;33"), ansi("0"))
	fmt.Println("   1. Удаление переписки с троллем")
	fmt.Println("   2. Запуск Go-модуля Value Receivers")
	fmt.Println("   3. Включение фоновой метели за окном")

	c.Stats.Focus = min(MAX_STAT, c.Stats.Focus+40)
	c.Stats.Wisdom += 25
}

func NewGoSha() Character {
	return Character{
		Name: "Гоша",
		Stats: Stats{
			Focus:     START_FOCUS,
			Wisdom:    38, // Символично: возраст в 2026
			Happiness: START_HAPPINESS,
		},
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	goSha := NewGoSha()

	fmt.Printf("%s🌨️ 8 ЯНВАРЯ 2026 | МЕТЕЛЬ vs СЕМЕЙНЫЕ БЕЛЯШИ 🌨️%s\n", ansi("1;33"), ansi("0"))
	fmt.Println(strings.Repeat("═", 60))

	// Событие 1: СМС о метели
	weatherEvent := Event{
		Description: "Сильный снег, метель, гололедица. Ветер 15-18 м/с",
		Emoji:       "❄️",
		Duration:    24 * time.Hour,
	}
	fmt.Printf("%s📱 МЕТЕОЦЕНТР:%s %s\n", ansi("1;36"), ansi("0"), weatherEvent.Description)

	// Событие 2: Провокация тролля
	trollMessage := "Я на фабрике ёлочных игрушек! А ты наверное валяешься на диване, ха-ха!"
	fmt.Printf("\n%s💬 ТРОЛЛЬ ПИШЕТ:%s \"%s\"\n", ansi("1;31"), ansi("0"), trollMessage)

	// Анализ ситуации (value receiver)
	analysis := goSha.AnalyzeTroll(trollMessage)
	fmt.Printf("%s🔍 АНАЛИЗ:%s %s\n", ansi("1;34"), ansi("0"), analysis)

	// Принятие решения
	if goSha.ShouldEngageTroll(trollMessage) {
		fmt.Printf("%s😠 ГОША ОТВЕЧАЕТ...%s\n", ansi("1;31"), ansi("0"))
		time.Sleep(1 * time.Second)
		fmt.Printf("%s😱 ОШИБКА! СРАЗУ ЖЕ УДАЛЯЕТ СООБЩЕНИЕ%s\n", ansi("1;33"), ansi("0"))
	} else {
		fmt.Printf("%s🧘 ГОША ДЫШИТ ГЛУБОКО И ЗАКРЫВАЕТ ЧАТ%s\n", ansi("1;32"), ansi("0"))
	}

	// Активация защиты
	goSha.ActivateFocusShield()

	// Событие 3: Семейный ужин
	goSha.FamilyDinner(8, weatherEvent.Description)

	// Итоги дня
	printSummary(goSha, weatherEvent)
}

func printSummary(g Character, w Event) {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Printf("%s📊 ИТОГИ ДНЯ 8 ЯНВАРЯ:%s\n", ansi("1;36"), ansi("0"))

	fmt.Printf("🧠 Мудрость: %s%d/100%s\n", ansi("1;34"), g.Stats.Wisdom, ansi("0"))
	fmt.Println(progressBar(g.Stats.Wisdom, "💡"))

	fmt.Printf("🎯 Фокус на Go: %s%d/100%s\n", ansi("1;32"), g.Stats.Focus, ansi("0"))
	fmt.Println(progressBar(g.Stats.Focus, "🎯"))

	fmt.Printf("❤️ Семейное счастье: %s%d/100%s\n", ansi("1;35"), g.Stats.Happiness, ansi("0"))
	fmt.Println(progressBar(g.Stats.Happiness, "❤️"))

	fmt.Printf("\n%s✅ ВЫВОДЫ:%s\n", ansi("1;33"), ansi("0"))
	fmt.Println("   - Тролли: 1 попытка → 0 побед (щит фокуса работает!)")
	fmt.Println("   - Беляшей съедено: 2 (рецепт: фарш из мясной лавки + материнская любовь)")
	fmt.Println("   - Код написан: 42+ строк (тема: Value vs Pointer Receivers)")
	fmt.Printf("   - Метель использована: %sПОЛНЫЙ ПРОФИТ%s\n", ansi("1;32"), ansi("0"))
	fmt.Println("   - Жизненный урок: \"Диван опаснее фабрики ёлочных игрушек\"")
}

func progressBar(value int, emoji string) string {
	width := 30
	filled := value * width / MAX_STAT
	return fmt.Sprintf("   [%s%s] %d%%",
		strings.Repeat(emoji, filled),
		strings.Repeat("░", width-filled),
		value)
}

func ansi(code string) string {
	return "\033[" + code + "m"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
