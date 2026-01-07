package main

import (
	"fmt"
	"strings"
	"time"
)

// TrollDeflector - система защиты от троллей и фокусировки на коде
type TrollDeflector struct {
	FocusLevel      int
	TrollsBlocked   int
	CodeLinesWritten int
	IsActivated     bool
}

// MorningRitual - утренний ритуал Гоши
type MorningRitual struct {
	WakeUpTime      string
	PhoneModel      string
	AlarmLocation   string
	GoodDeedsDone   int
	SnowfallEnjoyed bool
}

func main() {
	fmt.Println("🐐 ГОША 07.01.2026: ФОКУС ВМЕСТО ТРОЛЛЕЙ")
	fmt.Println(strings.Repeat("=", 50))

	// Часть 1: Утренний ритуал
	fmt.Println("\n🌅 УТРЕННИЙ РИТУАЛ:")
	fmt.Println(strings.Repeat("-", 30))

	ritual := MorningRitual{
		WakeUpTime:    "07:00",
		PhoneModel:    "Honor 10x Lite",
		AlarmLocation: "верхняя полка шкафа",
		GoodDeedsDone: 3,
		SnowfallEnjoyed: true,
	}

	ritual.Perform()
	time.Sleep(500 * time.Millisecond)

	// Часть 2: Защита от троллей и фокусировка
	fmt.Println("\n🛡️  СИСТЕМА ЗАЩИТЫ АКТИВИРОВАНА:")
	fmt.Println(strings.Repeat("-", 30))

	deflector := &TrollDeflector{
		FocusLevel:    75,
		IsActivated:   true,
	}

	// Имитация троллей в сети
	trollAttacks := []string{
		"Лол, опять за компом? У меня тут фото из Дубая!",
		"Go - узкоспециализированный язык, иди на Rust переучивайся!",
		"Чё, зубров в Беловежской пуще не хочешь посмотреть?",
		"Ты вообще живёшь или только код пишешь?",
		"Смотрю твой гитхаб... смешно, конечно",
	}

	for i, attack := range trollAttacks {
		fmt.Printf("\n👹 Тролль %d: \"%s\"\n", i+1, attack)
		response := deflector.Deflect(attack)
		fmt.Printf("   🛡️  Ответ: %s\n", response)
		time.Sleep(300 * time.Millisecond)
	}

	// Часть 3: Итоги дня
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 ИТОГИ ДНЯ 07.01.2026:")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Printf("✅ Утренний ритуал выполнен: %s\n", ritual.WakeUpTime)
	fmt.Printf("✅ Добрых дел сделано: %d\n", ritual.GoodDeedsDone)
	fmt.Printf("✅ Троллей отклонено: %d\n", deflector.TrollsBlocked)
	fmt.Printf("✅ Уровень фокуса: %d%%\n", deflector.FocusLevel)
	fmt.Printf("✅ Написано строк кода: %d\n", deflector.CodeLinesWritten)

	if deflector.FocusLevel >= 90 {
		fmt.Println("\n🏆 ДОСТИЖЕНИЕ: 'Непробиваемая концентрация' получено!")
	}

	fmt.Println("\n🎯 МОРАЛЬ: Пусть другие фоткают зубров,")
	fmt.Println("    а твои зубы будут сжиматься от концентрации на коде!")
	fmt.Println(strings.Repeat("=", 50))
}

// Perform - выполнение утреннего ритуала
func (r *MorningRitual) Perform() {
	steps := []struct{
		action string
		emoji  string
	}{
		{"Будильник на " + r.PhoneModel + " на " + r.AlarmLocation, "⏰"},
		{"Пришлось встать и тянуться (не лёг обратно!)", "🛌➡️🚶"},
		{"Бритьё и умывание", "🧔➡️👶"},
		{"Прогулка до Ozon/WB", "🚶❄️"},
		{"Забрал " + fmt.Sprintf("%d", r.GoodDeedsDone) + " посылок для друзей", "📦📦📦"},
	}

	for _, step := range steps {
		fmt.Printf("%s %s\n", step.emoji, step.action)
	}

	if r.SnowfallEnjoyed {
		fmt.Println("❄️  Насладился пушистым снежком!")
	}
}

// Deflect - отклонение атаки тролля и фокусировка на коде
func (td *TrollDeflector) Deflect(trollMessage string) string {
	td.TrollsBlocked++

	// Увеличиваем фокус при каждом отклонённом тролле
	td.FocusLevel += 5
	if td.FocusLevel > 100 {
		td.FocusLevel = 100
	}

	// Пишем код вместо ответа троллю
	codeLines := 10 + (td.TrollsBlocked * 5)
	td.CodeLinesWritten += codeLines

	responses := []string{
		"Игнорирую... пишу интерфейс на Go",
		"Спасибо за мотивацию! Добавил новую функцию",
		"Твой комментарий скомпилирован в 0 байт полезной информации",
		"Переадресую твою энергию в горутину...",
		"Зафиксировал твоё мнение в /dev/null",
	}

	// Выбираем ответ в зависимости от уровня фокуса
	responseIndex := td.TrollsBlocked % len(responses)

	return fmt.Sprintf("%s (+%d строк кода, фокус: %d%%)",
		responses[responseIndex], codeLines, td.FocusLevel)
}
