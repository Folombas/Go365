package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🚂 Путешествие в Апрелевку: Docker в пути")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Дата: 15 января 2026 года")
	fmt.Println("Время: 9:30 утра")
	fmt.Println()

	// Утренний ритуал
	fmt.Println("🌅 УТРЕННИЙ РИТУАЛ:")
	ritual := []string{
		"Подъем в 9:30",
		"Бритье",
		"Душ",
		"Вкусный завтрак",
	}

	for i, step := range ritual {
		fmt.Printf(" %d. %s\n", i+1, step)
		time.Sleep(300 * time.Millisecond)
	}

	// Жирный заказ
	fmt.Println("\n💰 ЖИРНЫЙ ЗАКАЗ:")
	fmt.Println("  Гоше повезло - поймал жирный заказ!")
	fmt.Println("  Маршрут: метро 'Верхние Лихоборы' → Апрелевка")
	fmt.Println("  Расстояние: 40+ км на электричках")
	time.Sleep(1 * time.Second)

	// Путешествие на электричке
	fmt.Println("\n🚆 ПУТЕШЕСТВИЕ НА ЭЛЕКТРИЧКЕ:")
	journey := []string{
		"Сел в электричку на вокзале",
		"Вагон полупустой",
		"Только бомж, лежащий на сиденьях в центре вагона",
		"Стук колес, мелькающие за окном дома",
		"Гоша достает книгу...",
	}

	for _, step := range journey {
		fmt.Printf("  • %s\n", step)
		time.Sleep(500 * time.Millisecond)
	}

	// Чтение книги
	fmt.Println("\n📚 ЧТЕНИЕ В ПУТИ:")
	fmt.Println("  Книга: 'Docker Deep Dive' (бумажная версия)")
	fmt.Println("  Автор: Nigel Poulton")
	fmt.Println("  Страниц прочитано: 15-20")
	fmt.Println("  Состояние: полное погружение")
	time.Sleep(1 * time.Second)

	// Осознание
	fmt.Println("\n💡 ОСОЗНАНИЕ:")
	fmt.Println("  Docker написан на Go.")
	fmt.Println("  Я учу Go.")
	fmt.Println("  Я читаю про Docker в электричке.")
	fmt.Println()

	realizations := []string{
		"Я не теряю время в дороге!",
		"Каждая поездка - это учебная сессия",
		"Дальние заказы = больше времени на обучение",
		"Электричка = мобильный учебный класс",
	}

	for _, realization := range realizations {
		fmt.Printf("  • %s\n", realization)
		time.Sleep(600 * time.Millisecond)
	}

	// Психологическая драма
	fmt.Println("\n🎭 ПСИХОЛОГИЧЕСКАЯ ДРАМА:")

	scenes := []struct{
		character string
		text      string
	}{
		{"ГОША", "Как же хорошо... Стук колес, книга, знания..."},
		{"БОМЖ", "*храпит на соседнем сиденье*"},
		{"ГОША", "Он спит, а я учусь. У каждого свой путь."},
		{"ВНУТРЕННИЙ ГОЛОС", "Ты мог бы спать дома в теплой постели..."},
		{"ГОША", "Но тогда я не читал бы про Docker!"},
		{"ВНУТРЕННИЙ ГОЛОС", "Зачем тебе это? Ты же просто курьер..."},
		{"ГОША", "Docker написан на Go. Я учу Go. Я на правильном пути."},
	}

	for _, scene := range scenes {
		if scene.character == "БОМЖ" {
			fmt.Printf("  👤 %s: %s\n", scene.character, scene.text)
		} else if scene.character == "ВНУТРЕННИЙ ГОЛОС" {
			fmt.Printf("  👻 %s: %s\n", scene.character, scene.text)
		} else {
			fmt.Printf("  👨 %s: %s\n", scene.character, scene.text)
		}
		time.Sleep(800 * time.Millisecond)
	}

	// Геймификация
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ПОЕЗДКИ:")

	type Achievement struct {
		name   string
		points int
	}

	achievements := []Achievement{
		{"Поймать жирный заказ", 100},
		{"Проехать 40+ км на электричках", 75},
		{"Читать книгу в движущемся транспорте", 60},
		{"Игнорировать отвлекающие факторы (бомж)", 40},
		{"Осознать связь Docker и Go", 150},
		{"Использовать время в пути для обучения", 125},
	}

	totalPoints := 0
	for _, a := range achievements {
		fmt.Printf("  ✅ %s: +%d очков\n", a.name, a.points)
		totalPoints += a.points
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("\n  📊 Всего очков: %d\n", totalPoints)

	// Уровень
	level := 1
	if totalPoints > 200 {
		level = 2
	}
	if totalPoints > 400 {
		level = 3
	}
	if totalPoints > 550 {
		level = 4
	}

	levelNames := map[int]string{
		1: "Новичок на рельсах",
		2: "Опытный путешественник",
		3: "Железнодорожный мудрец",
		4: "Мастер пути",
	}

	fmt.Printf("  🎯 Уровень: %d (%s)\n", level, levelNames[level])

	// Разблокированные достижения
	if totalPoints > 300 {
		fmt.Println("\n  🏆 ДОСТИЖЕНИЕ РАЗБЛОКИРОВАНО:")
		fmt.Println("     'Философ электрички' - находить смысл в рутине")
	}

	// Статистика поездки
	fmt.Println("\n📈 СТАТИСТИКА ПОЕЗДКИ:")
	stats := map[string]string{
		"Время в пути":          "≈2 часа",
		"Прочитано страниц":     "15-20",
		"Расстояние":            "40+ км",
		"Стоимость заказа":      "2000+ рублей",
		"Вагонов электрички":    "10",
		"Попутчиков в вагоне":   "1 (бомж)",
		"Качество концентрации": "высокое",
		"Уровень осознанности":  "максимальный",
	}

	for k, v := range stats {
		fmt.Printf("  %-25s: %s\n", k, v)
		time.Sleep(200 * time.Millisecond)
	}

	// Философские выводы
	fmt.Println("\n🧠 ФИЛОСОФСКИЕ ВЫВОДЫ:")

	philosophy := []string{
		"Поезд идет по рельсам - у меня есть путь (изучение Go)",
		"Каждая остановка - это новая тема для изучения",
		"Бомж в вагоне - как отвлекающие мысли, которые нужно игнорировать",
		"Чтение книги - инвестиция в себя, которая окупится",
		"Дальняя поездка - не наказание, а возможность",
		"Docker на Go - как знак, что я на правильном пути",
	}

	for _, p := range philosophy {
		fmt.Printf("  • %s\n", p)
		time.Sleep(500 * time.Millisecond)
	}

	// Disclaimer
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("                    DISCLAIMER")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n" + strings.Repeat("🚂", 60))
	fmt.Println("       ПОЕЗД ПРИБЫЛ В АПРЕЛЕВКУ. ДЕНЬ ЗАВЕРШЕН!")
	fmt.Println(strings.Repeat("🚂", 60))
}
