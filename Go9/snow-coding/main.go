package main

import (
	"fmt"
	"strings"
	"time"
)

// SnowDay - структура для дня в метель
type SnowDay struct {
	Date           string
	SnowLevel      string
	PassExpired    bool
	CarsInSnow     int
	CodeHours      float64
	FamilyTime     float64
	TrollMessages  int
	OldFanMessages int
	IgnoredTrolls  int
}

// FamilyPriority - приоритеты семьи
type FamilyPriority struct {
	MorningTea      bool
	MealsTogether   int
	Conversations   int
	HelpWithHome    bool
	EveningTime     bool
}

func main() {
	fmt.Println("❄️  ДЕНЬ 09.01.2026: ПРОГРАММИРОВАНИЕ В МЕТЕЛЬ")
	fmt.Println(strings.Repeat("=", 60))

	// Инициализируем день
	snowDay := SnowDay{
		Date:          "09.01.2026",
		SnowLevel:     "снег по колено, метель",
		PassExpired:   true,
		CarsInSnow:    8,
		CodeHours:     6.5,
		FamilyTime:    4.2,
		TrollMessages: 3,
		OldFanMessages: 1,
		IgnoredTrolls: 4,
	}

	// Приоритеты семьи
	family := FamilyPriority{
		MorningTea:    true,
		MealsTogether: 3,
		Conversations: 5,
		HelpWithHome:  true,
		EveningTime:   true,
	}

	// Выводим статистику дня
	printDayStats(snowDay, family)
	printSnowObservation()
	printTrollDeflection(snowDay)
	printFamilyMoments(family)
	printWisdomLessons()
	printGamificationResults(snowDay, family)
}

func printDayStats(snowDay SnowDay, family FamilyPriority) {
	fmt.Println("\n📊 СТАТИСТИКА ДНЯ:")
	fmt.Println(strings.Repeat("-", 60))

	items := []struct {
		label string
		value string
	}{
		{"Дата", snowDay.Date},
		{"Погода", "❄️  " + snowDay.SnowLevel},
		{"Проездной", map[bool]string{true: "❌ Истёк в полночь", false: "✅ Активен"}[snowDay.PassExpired]},
		{"Иномарок в сугробах", fmt.Sprintf("%d шт.", snowDay.CarsInSnow)},
		{"Часов кодинга", fmt.Sprintf("%.1f ч.", snowDay.CodeHours)},
		{"Времени с семьёй", fmt.Sprintf("%.1f ч.", snowDay.FamilyTime)},
		{"Сообщений от троллей", fmt.Sprintf("%d шт.", snowDay.TrollMessages)},
		{"Сообщений от 'олдового' фаната", fmt.Sprintf("%d шт.", snowDay.OldFanMessages)},
		{"Проигнорировано троллей", fmt.Sprintf("%d шт.", snowDay.IgnoredTrolls)},
	}

	for _, item := range items {
		fmt.Printf("   %-30s: %s\n", item.label, item.value)
	}
}

func printSnowObservation() {
	fmt.Println("\n👀 НАБЛЮДЕНИЕ ЗА ОКНОМ:")
	fmt.Println(strings.Repeat("-", 60))

	observations := []string{
		"Соседи откапывают иномарки лопатами",
		"Снег идёт не переставая",
		"Сугробы растут на глазах",
		"Дворники не справляются",
		"Машины буксуют на месте",
	}

	for i, obs := range observations {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("   %d. %s\n", i+1, obs)
	}

	fmt.Println("\n   💭 Мысль: 'Идеальный день, чтобы не выходить и кодить на Go'")
}

func printTrollDeflection(snowDay SnowDay) {
	fmt.Println("\n🛡️  ОБРАБОТКА ТРОЛЛЕЙ:")
	fmt.Println(strings.Repeat("-", 60))

	// Тролль 1
	fmt.Println("👹 Тролль 1: 'Гоша, ты что, опять дома сидишь?'")
	fmt.Println("   🎯 Реакция: Удалить сообщение")
	fmt.Println("   💡 Мысль: 'Не тратить миллисекунды на оправдания'")
	time.Sleep(300 * time.Millisecond)

	// Тролль 2
	fmt.Println("\n👹 Тролль 2: 'Когда уже найдёшь работу?'")
	fmt.Println("   🎯 Реакция: Игнорировать")
	fmt.Println("   💡 Мысль: 'Лучше написать 10 строк кода'")
	time.Sleep(300 * time.Millisecond)

	// Тролль 3
	fmt.Println("\n👹 Тролль 3: 'Другие уже senior, а ты...'")
	fmt.Println("   🎯 Реакция: Удалить и добавить в чёрный список")
	fmt.Println("   💡 Мысль: 'Каждый тролль = +5 минут к фокусу'")
	time.Sleep(300 * time.Millisecond)

	// Старый фанат
	fmt.Println("\n👴 'Олдовый' фанат (15 лет слежения):")
	fmt.Println("   'Ты тратишь миллисекунды на удаление сообщений троллей'")
	fmt.Println("   🎯 Реакция: Удалить сообщение")
	fmt.Println("   💡 Мысль: 'Даже фанатам иногда нужно давать отдых'")
	time.Sleep(300 * time.Millisecond)

	fmt.Printf("\n   📈 Итого игнорировано: %d/4 сообщений\n", snowDay.IgnoredTrolls)
}

func printFamilyMoments(family FamilyPriority) {
	fmt.Println("\n👨‍👩‍👦 СЕМЕЙНЫЕ ПРИОРИТЕТЫ:")
	fmt.Println(strings.Repeat("-", 60))

	moments := []struct {
		activity string
		status   string
		emoji    string
	}{
		{"Утренний чай с беляшом", "✅ За окном наблюдали метель", "☕"},
		{"Совместные приёмы пищи", fmt.Sprintf("✅ %d раза", family.MealsTogether), "🍽️"},
		{"Беседы с родными", fmt.Sprintf("✅ %d разговора", family.Conversations), "💬"},
		{"Помощь по дому", "✅ Помог с уборкой", "🧹"},
		{"Вечер вместе", "✅ Провели вечер за играми", "🎮"},
	}

	for _, moment := range moments {
		fmt.Printf("   %s %-25s: %s\n", moment.emoji, moment.activity, moment.status)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("\n   💭 Мысль: 'Семья > Работа. Всегда.'")
}

func printWisdomLessons() {
	fmt.Println("\n🧠 МУДРОСТЬ ДНЯ:")
	fmt.Println(strings.Repeat("-", 60))

	lessons := []string{
		"1. Не выходить в метель без необходимости — мудро",
		"2. Игнорировать троллей — экономия энергии",
		"3. Проводить время с семьёй — инвестиция в счастье",
		"4. Программировать в тишине — продуктивно",
		"5. Чай с беляшом > Кофе в офисе",
		"6. Снег за окном > Пробки на дорогах",
		"7. Домашний уют > Офисный стрес",
		"8. Личное развитие > Оправдания перед другими",
	}

	for _, lesson := range lessons {
		fmt.Println("   " + lesson)
		time.Sleep(150 * time.Millisecond)
	}
}

func printGamificationResults(snowDay SnowDay, family FamilyPriority) {
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ДНЯ:")
	fmt.Println(strings.Repeat("-", 60))

	// Рассчитываем очки
	codePoints := int(snowDay.CodeHours * 100)
	familyPoints := int(snowDay.FamilyTime * 150)
	trollPoints := snowDay.IgnoredTrolls * 50
	snowBonus := 200 // бонус за снежный день
	passPenalty := -100 // штраф за истёкший проездной

	if snowDay.PassExpired {
		passPenalty = -100
	} else {
		passPenalty = 0
	}

	totalPoints := codePoints + familyPoints + trollPoints + snowBonus + passPenalty

	fmt.Printf("   🖥️  Часы кодинга (%.1f ч.)  : +%d очков\n", snowDay.CodeHours, codePoints)
	fmt.Printf("   👨‍👩‍👦 Время с семьёй (%.1f ч.) : +%d очков\n", snowDay.FamilyTime, familyPoints)
	fmt.Printf("   🛡️  Игнорировано троллей (%d)  : +%d очков\n", snowDay.IgnoredTrolls, trollPoints)
	fmt.Printf("   ❄️  Бонус снежного дня       : +%d очков\n", snowBonus)

	if snowDay.PassExpired {
		fmt.Printf("   🚫 Штраф истёкшего проездного : %d очков\n", passPenalty)
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("   🏆 ИТОГО: %d очков\n", totalPoints)

	// Уровень дня
	fmt.Println("\n📊 УРОВЕНЬ ДНЯ:")
	level := calculateLevel(totalPoints)
	fmt.Printf("   %s\n", level)

	// Достижения
	fmt.Println("\n🏅 ДОСТИЖЕНИЯ ДНЯ:")
	achievements := []string{
		"✅ 'Снежный программист' (кодил в метель)",
		"✅ 'Непреклонный' (проигнорировал 4 тролля)",
		"✅ 'Семьянин' (провёл >4 часов с семьёй)",
		"✅ 'Мудрый экономист' (не тратил на транспорт)",
		"✅ 'Чайный ценитель' (пил горячий чай, глядя в окно на метель и снежные сугробы)",
	}

	for _, achievement := range achievements {
		fmt.Println("   " + achievement)
	}
}

func calculateLevel(points int) string {
	switch {
	case points >= 1500:
		return "🚀 КОСМИЧЕСКАЯ МУДРОСТЬ"
	case points >= 1200:
		return "🔥 ИДЕАЛЬНЫЙ БАЛАНС"
	case points >= 900:
		return "⭐ МУДРЫЙ ВЫБОР"
	case points >= 600:
		return "🌱 ХОРОШИЙ ДЕНЬ"
	default:
		return "📝 НАЧАЛО ПУТИ"
	}
}
