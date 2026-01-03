// Go3/life-reboot-2026/motivation/quotes.go
package motivation

import (
	"fmt"
	"math/rand"
	"time"
)

// Quote представляет мотивационную цитату
type Quote struct {
	Text     string
	Author   string
	Category string // "health", "programming", "focus", "family"
}

// GetDailyQuote возвращает цитату дня
func GetDailyQuote() Quote {
	quotes := []Quote{
		{
			Text:     "Отказаться от развлечений сегодня — значит получить профессию завтра",
			Author:   "Гоша Гофер",
			Category: "focus",
		},
		{
			Text:     "Каждый день без сигарет — это +1% к здоровью и +10% к фокусу",
			Author:   "Законы Гофера",
			Category: "health",
		},
		{
			Text:     "Семья поддержит тебя сегодня, когда ты строишь лучшее завтра",
			Author:   "Мудрость 2026",
			Category: "family",
		},
		{
			Text:     "Go — не просто язык, это билет в новую жизнь",
			Author:   "Сообщество Go",
			Category: "programming",
		},
		{
			Text:     "Зарядка утром = компиляция без ошибок днём",
			Author:   "Физика программиста",
			Category: "health",
		},
	}

	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}

// DisplayQuote красиво показывает цитату
func DisplayQuote(q Quote) {
	fmt.Println("\n✨ МОТИВАЦИЯ:")
	fmt.Println("  «" + q.Text + "»")
	fmt.Println("  — " + q.Author)
	fmt.Println("  Категория: " + q.Category)
}
