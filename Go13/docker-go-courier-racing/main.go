package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🐳 Docker Go Courier Racing 🚀")
	fmt.Println("=================================")
	fmt.Println("Дата: 13 января 2026 года")
	fmt.Println("Время: 6:30 утра")
	fmt.Println()

	// Утренний ритуал
	fmt.Println("⏰ УТРЕННИЙ РИТУАЛ:")
	ritual := []string{
		"Потягушки, улыбка",
		"Зарядка: 8 отжиманий, 13 приседаний, 1.5 подтягивания",
		"Завтрак: пшёнка кашка",
		"Тройка заряжена (вчера) на 30 дней безлимит + Пригород",
		"Утренний душ и бритьё (сначала побрился, потом помылся)",
	}

	for i, step := range ritual {
		fmt.Printf(" %d. %s\n", i+1, step)
		time.Sleep(300 * time.Millisecond)
	}

	// Гонка за заказами
	fmt.Println("\n🏎️  ГОНКА ЗА ЗАКАЗАМИ:")
	fmt.Println("  Гоша пьет чай на кухне, мониторит пул заказов...")
	time.Sleep(1 * time.Second)

	fmt.Println("  🔥 Появился ЖИРНЫЙ ЗАКАЗ! 1500 рублей!")
	fmt.Println("  ...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  Исчез за 0.5 секунды! Быстрый курьер забрал!")

	reactions := []string{
		"😱 Не успел моргнуть!",
		"💔 Еще один упущен...",
		"⚡ Кто успел - тот молодец!",
		"🏃 Надо быть быстрее!",
	}

	fmt.Printf("  %s\n", reactions[rand.Intn(len(reactions))])

	// Доставки дня
	fmt.Println("\n📦 ДОСТАВКИ ДНЯ:")
	deliveries := []struct{
		from string
		to   string
		what string
	}{
		{"Речной вокзал", "Белорусская", "документы"},
		{"Центр", "Северо-Запад", "открытки"},
	}

	earnings := 0
	for i, d := range deliveries {
		price := 400 + rand.Intn(600)
		earnings += price
		fmt.Printf(" %d. %s → %s (%s): %d руб\n",
			i+1, d.from, d.to, d.what, price)
		time.Sleep(400 * time.Millisecond)
	}

	fmt.Printf("\n💰 ЗАРАБОТОК: %d рублей (тыщёнка с копейками)\n", earnings)

	// Книга в метро
	fmt.Println("\n📚 МОМЕНТ В МЕТРО:")
	fmt.Println("  Гоша достает бумажную книгу 'Docker Deep Dive'")
	fmt.Println("  (подарок на Новый год)")
	fmt.Println("  Читает несколько страниц...")
	time.Sleep(1 * time.Second)

	// Глубокая мысль
	fmt.Println("\n💭 ГЛУБОКАЯ МЫСЛЬ:")
	fmt.Println("  Docker написан на Go.")
	fmt.Println("  А я учу Go.")
	fmt.Println()
	fmt.Println("  Это говорит о том, что...")
	time.Sleep(1 * time.Second)

	insights := []string{
		"Go - язык для серьезных production-систем",
		"Docker доказал мощь Go в системном программировании",
		"Каждая строка моего кода - шаг к созданию чего-то великого",
		"Контейнеры Docker = контейнеры знаний в моей голове",
		"Учить Go сегодня = строить инфраструктуру завтра",
	}

	for _, insight := range insights {
		fmt.Printf("  • %s\n", insight)
		time.Sleep(600 * time.Millisecond)
	}

	// Психологическая драма
	fmt.Println("\n🎭 ПСИХОЛОГИЧЕСКАЯ ДРАМА:")

	innerDialogue := []string{
		"ГОЛОС 1: 'Ты опять проспал жирный заказ...'",
		"ГОЛОС 2: 'Зато я прочитал про Docker!'",
		"ГОЛОС 1: 'Деньги важнее книжек!'",
		"ГОЛОС 2: 'Знания важнее сиюминутной выгоды!'",
		"ГОЛОС 1: 'Ты никогда не станешь программистом...'",
		"ГОЛОС 2: 'Docker написан на Go. Я учу Go. Это не случайно.'",
	}

	for _, line := range innerDialogue {
		fmt.Printf("  %s\n", line)
		time.Sleep(700 * time.Millisecond)
	}

	// Геймификация
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ДНЯ:")

	type Achievement struct {
		name  string
		score int
	}

	achievements := []Achievement{
		{"Утренний ритуал выполнен", 25},
		{"Не унывать из-за упущенного заказа", 50},
		{"Выполнено 2 доставки", 75},
		{"Чтение в метро (прокачка знаний)", 100},
		{"Глубокая мысль о Docker/Go связи", 150},
	}

	totalScore := 0
	for _, a := range achievements {
		fmt.Printf("  ✅ %s: +%d очков\n", a.name, a.score)
		totalScore += a.score
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("\n📊 ИТОГО: %d очков\n", totalScore)
	level := totalScore/100 + 1
	fmt.Printf("  Уровень: %d\n", level)

	if totalScore > 300 {
		fmt.Println("  🏆 ДОСТИЖЕНИЕ: 'Мыслитель-курьер' разблокировано!")
	}

	// Статистика
	fmt.Println("\n📈 СТАТИСТИКА ДНЯ:")
	stats := map[string]interface{}{
		"Время подъема": "6:30",
		"Физическая активность": "8 отжиманий, 13 приседаний, 1.5 подтягиваний",
		"Завтрак": "пшенная каша",
		"Транспортная карта": "Тройка + Пригород (30 дней)",
		"Упущенных жирных заказов": rand.Intn(5) + 1,
		"Выполненных доставок": 2,
		"Заработано": fmt.Sprintf("%d руб", earnings),
		"Прочитано страниц": rand.Intn(10) + 2,
		"Уровень осознанности": "высокий",
	}

	for k, v := range stats {
		fmt.Printf("  %-25s: %v\n", k, v)
	}

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("⚠️  DISCLAIMER")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(`
Важное юридическое и этическое уведомление:

1. Все персонажи в историях про Гошу являются художественной выдумкой.
2. Все события, сюжеты и диалоги придуманы автором.
3. Любые совпадения с реальными людьми, живыми или мертвыми,
   являются случайными и непреднамеренными.
4. Все описанные ситуации - плод творческого воображения.
5. Москва, Химки, Тройка, Docker, Go - реально существуют,
   но их использование в сюжете - художественный приём.
6. Эта программа является художественным произведением,
   выполненным в стиле программного кода.
7. Автор не несет ответственности за эмоциональные потрясения,
   вызванные чтением кода.
8. Программируйте ответственно.

С уважением к вашему воображению,
Виртуальный сценарист системы "Гоша 2026".
`)

	fmt.Println("\n" + strings.Repeat("🚀", 50))
	fmt.Println("      КОНЕЦ ЭПИЗОДА. ДО ЗАВТРА!")
	fmt.Println(strings.Repeat("🚀", 50))
}
