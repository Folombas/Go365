package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ========== DISCLAIMER: ХУДОЖЕСТВЕННЫЙ ВЫМЫСЕЛ ==========
// Все персонажи в Daily Code Life Story вымышлены.
// Все истории, сюжеты, диалоги и события являются плодом
// творческой обработки и метафорой пути в IT.
// Любые совпадения с реальными людьми или ситуациями
// случайны и непреднамерены.
// Это не дневник, а арт1факт — программный код как искусство.
// ======================================================

// Player представляет нашего героя Гошу, игрока в игре под названием "Карьера"
type Player struct {
	Name              string
	Role              string          // "courier" | "gopher"
	Focus             int             // 0-100: Уровень концентрации
	Stamina           int             // 0-100: Выносливость для заказов
	Knowledge         int             // 0-100: Знание Go и экосистемы
	Motivation        int             // 0-100: Уровень мотивации
	DaysAsCourier     int             // Счётчик дней в роли курьера
	N8nAutomation     bool            // Флаг: освоен ли n8n
	MotivationalQueue []string        // Очередь мотивационных фраз
}

// NewPlayer создаёт нового игрока в состоянии "Утро понедельника"
func NewPlayer(name string) *Player {
	p := &Player{
		Name:          name,
		Role:          "courier",
		Focus:         30, // С утра депрессия и мороз
		Stamina:       85, // Тело ещё работает
		Knowledge:     42, // Какие-то знания уже есть
		Motivation:    25, // Очень низкая
		DaysAsCourier: 1,
		N8nAutomation: true, // Бот уже создан!
	}
	// Заполняем очередь мотивации
	p.MotivationalQueue = []string{
		"🚀 Каждая строчка кода на Go — это гвоздь в крышку гроба твоей курьерской сумки.",
		"🔥 Холод на улице временный. Холод в незнании синтаксиса — вечный.",
		"🎯 Фокус на Go сегодня — тёплый офис завтра. Не распыляйся.",
		"⚡ Время, потраченное на бары и сериалы, можно конвертировать в время, оплаченное долларами удалёнки.",
		"🏆 GO — не просто язык. Это твой щит от оледеневших тротуаров и билет в элиту, которая пишет историю будущего.",
		"💡 Знание экосистемы Go (горутины, интерфейсы, тесты) — это суперсила, за которую компании готовы платить в 5 раз больше, чем за доставку цветов.",
		"🧠 Изучая Go, ты не учишь синтаксис. Ты строишь новую нейронную сеть в своей голове — архитектора систем, а не исполнителя поручений.",
		"🛡️ Каждый пакет, который ты освоил (net/http, context, database/sql) — это броня против увольнения. Разнообразь свой стек.",
		"📈 Рынок не ждёт. Пока ты смотришь видосы из отпуска, другой Гоша уже выучил goroutine и устроился на 300к. Беги быстрее.",
		"🎮 Прокачка навыка Go — самая честная игра. Ввод = вывод. Время за клавиатурой = рост зарплаты. Веди статистику, ставь рекорды.",
	}
	return p
}

// MorningDepression симулирует утреннее состояние
func (p *Player) MorningDepression() {
	fmt.Println("\n❄️════════════════════════════════════════════❄️")
	fmt.Println("  26 ЯНВАРЯ 2026. 11:00. МОСКВА. -18°C.")
	fmt.Println("  ЖЁСТКАЯ ДЕПРЕССИЯ АКТИВИЗИРОВАЛАСЬ.")
	fmt.Println("❄️════════════════════════════════════════════❄️")
	time.Sleep(2 * time.Second)
	fmt.Printf("\n%s лежит. Будильник на Honor 10x Lite давно отзвенел.\n", p.Name)
	fmt.Println("Мысли: 'Вставать? Зачем? Опять эти сугробы, эти коробки...'")
	time.Sleep(2 * time.Second)

	// Рандомный шанс подняться
	if rand.Intn(100) < p.Stamina { // Чем выше выносливость, тем больше шансов
		fmt.Println("\n💥 ВНУТРЕННИЙ ДИАЛОГ:")
		fmt.Println("  'НЕТ. ХВАТИТ. СЕГОДНЯ — ПЕРЕЛОМ.'")
		p.GetMotivation()
		p.Focus += 20
		fmt.Println("  *С силой поднимается с кровати*")
	} else {
		fmt.Println("\n💔 Силы покидают. День пропал.")
		p.Motivation -= 30
	}
}

// GetMotivation выдаёт случайную мотивационную фразу
func (p *Player) GetMotivation() {
	if len(p.MotivationalQueue) == 0 {
		return
	}
	idx := rand.Intn(len(p.MotivationalQueue))
	fmt.Printf("\n💫 МОТИВАЦИЯ: %s\n", p.MotivationalQueue[idx])
	p.Motivation += 15
	if p.Motivation > 100 {
		p.Motivation = 100
	}
}

// MakeChoice предлагает игроку ключевой выбор
func (p *Player) MakeChoice() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("  ПЕРЕКРЁСТОК СУДЬБЫ. КУДА ПОВЕРНЁШЬ?")
	fmt.Println("  1) 🏃 БЕЖАТЬ В МЕТЕЛЬ С ЗАКАЗАМИ (роль: courier)")
	fmt.Println("  2) 🐹 УЧИТЬ GO И ЭКОСИСТЕМУ (роль: gopher)")
	fmt.Println("  3) 🤖 ДОРАБОТАТЬ ТЕЛЕГРАМ-БОТА (n8n + ИИ)")
	fmt.Println(strings.Repeat("═", 60))

	// В реальной игре здесь был бы ввод, но симулируем выбор
	choice := rand.Intn(100)

	switch {
	case choice < 50: // 50% шанс - привычная зона комфорта
		fmt.Printf("  %s выбирает вариант 1: 'Пока ноги ходят...'\n", p.Name)
		p.WorkAsCourier()
	case choice < 85: // 35% - фокус на цели
		fmt.Printf("  %s выбирает вариант 2: 'ХВАТИТ! Пора браться за ум!'\n", p.Name)
		p.StudyGo()
	default: // 15% - отвлечение на новую игрушку
		fmt.Printf("  %s выбирает вариант 3: 'Ой, а что если боту добавить...'\n", p.Name)
		p.TinkerWithN8n()
	}
}

// WorkAsCourier - механика работы курьером
func (p *Player) WorkAsCourier() {
	p.Role = "courier"
	fmt.Println("\n🌨️ МОСКОВСКАЯ МЕТЕЛЬ:")
	deliveries := []string{
		"Тяжёлая коробка с техникой в офис на 'Москва-Сити'",
		"Хрупкий букет роз через весь город",
		"10 кг корма для собаки в спальный район",
		"Срочный документ в больницу",
	}

	earned := 0
	for _, d := range deliveries {
		fmt.Printf("  🚴 Доставка: %s\n", d)
		p.Stamina -= rand.Intn(15)
		earned += rand.Intn(500) + 200
		time.Sleep(1 * time.Second)

		if p.Stamina <= 20 {
			fmt.Println("  ⚠️  Силы на исходе. Ноги не идут.")
			break
		}
	}

	p.DaysAsCourier++
	fmt.Printf("  💰 Заработано: %d руб. Выносливость: %d/100\n", earned, p.Stamina)
	p.Motivation -= 10 // Работа угнетает
	p.Focus -= 5       // Мысли размываются
}

// StudyGo - механика изучения Go
func (p *Player) StudyGo() {
	p.Role = "gopher"
	fmt.Println("\n🔥 АБСОЛЮТНЫЙ ФОКУС НА GO:")
	topics := []string{
		"Concurrency: Горутины и каналы",
		"Интерфейсы и композиция",
		"Тестирование (go test, testify)",
		"Работа с базами (sqlx, migrations)",
		"Web (Echo, Gin, middleware)",
		"GRPC и protobuf",
		"Оптимизация и профилирование",
	}

	for i := 0; i < 3; i++ { // Изучаем 3 темы за сессию
		topic := topics[rand.Intn(len(topics))]
		fmt.Printf("  📚 Тема: %s\n", topic)

		// Шанс понять тему зависит от текущего уровня знаний
		understanding := p.Knowledge + rand.Intn(30)
		switch {
		case understanding > 80:
			fmt.Println("     ✅ Озарение! Всё становится на свои места.")
			p.Knowledge += 15
		case understanding > 50:
			fmt.Println("     👍 Стабильный прогресс. Нужно практиковать.")
			p.Knowledge += 8
		default:
			fmt.Println("     🤯 Сложно. Нужно перечитать документацию.")
			p.Knowledge += 3
		}

		p.Focus += 10
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("  🧠 Знание Go: %d/100. Фокус: %d/100\n", p.Knowledge, p.Focus)
	p.Motivation += 20 // Учёба мотивирует
	p.GetMotivation()  // Дополнительная мотивация
}

// TinkerWithN8n - механика возни с автоматизацией
func (p *Player) TinkerWithN8n() {
	fmt.Println("\n🤖 УВЛЕЧЕНИЕ N8N:")

	// 10 фраз, обосновывающих пользу n8n
	n8nBenefits := []string{
		"1) n8n — это мост между 'просто кодом' и бизнес-логикой. Ты учишься видеть процессы.",
		"2) Автоматизация = эффективность. Этот навык ценится в любом IT-проекте.",
		"3) Умение быстро создать бота или интеграцию — конкурентное преимущество на фоне чистых кодеров.",
		"4) Понимание webhook, API, JSON — это фундамент современной веб-разработки, который ты осваиваешь на практике.",
		"5) Это +1 пункт в резюме: 'Low-code automation (n8n)'. HR это заметит.",
		"6) Созданный бот — первый шаг к микросервисной архитектуре в твоей голове.",
		"7) Ты учишься декомпозировать задачу: 'бот с ИИ' = Telegram API + AI provider + логика + хранилище.",
		"8) Docker, где работает n8n, — стандарт индустрии. Ты уже в контейнерах!",
		"9) Это не отвлечение, если через неделю этот бот автоматизирует часть твоей курьерской отчётности.",
		"10) В будущем эти знания позволят тебе быстро прототипировать идеи, а не ждать бэкенд-разработчика.",
	}

	// Выводим 2-3 случайные фразы
	for i := 0; i < 2+rand.Intn(2); i++ {
		fmt.Printf("  💡 %s\n", n8nBenefits[rand.Intn(len(n8nBenefits))])
		time.Sleep(1 * time.Second)
	}

	fmt.Println("  🎭 'Очередное увлечение?' — спросит внутренний критик.")
	fmt.Println("  'НЕТ. ЭТО ИНВЕСТИЦИЯ В ШИРОТУ ВЗГЛЯДА', — ответишь ты.")

	p.N8nAutomation = true
	p.Knowledge += 5  // Широкий кругозор
	p.Focus -= 10     // Но фокус рассеивается
	p.Motivation += 5 // Приятно видеть работающий результат
}

// DisplayStatus показывает полный статус игрока с элементами геймификации
func (p *Player) DisplayStatus() {
	fmt.Println("\n" + strings.Repeat("█", 60))
	fmt.Println("🎮 CODE COURIER SIMULATOR — СТАТИСТИКА ПЕРСОНАЖА:")
	fmt.Printf("  Имя: %s\n", p.Name)
	fmt.Printf("  Роль: %s", p.Role)
	if p.Role == "courier" {
		fmt.Printf(" (%d дней)\n", p.DaysAsCourier)
	} else {
		fmt.Println(" (учится!)")
	}

	// Прогресс-бары
	fmt.Printf("  Фокус:       [%s] %d/100\n", progressBar(p.Focus), p.Focus)
	fmt.Printf("  Выносливость:[%s] %d/100\n", progressBar(p.Stamina), p.Stamina)
	fmt.Printf("  Знание Go:   [%s] %d/100\n", progressBar(p.Knowledge), p.Knowledge)
	fmt.Printf("  Мотивация:   [%s] %d/100\n", progressBar(p.Motivation), p.Motivation)

	// Критические состояния
	msgs := []string{}
	if p.Stamina < 20 {
		msgs = append(msgs, "⚠️  ТЕЛО НА ГРАНИ. Нужен отдых.")
	}
	if p.Knowledge > 70 {
		msgs = append(msgs, "🚀 Уровень знаний достаточен для первого коммерческого проекта!")
	}
	if p.Motivation < 30 {
		msgs = append(msgs, "💔 МОТИВАЦИЯ НА НУЛЕ. Вспомни, зачем начал.")
	}
	if p.N8nAutomation {
		msgs = append(msgs, "🤖 Автоматизатор (n8n) разблокирован!")
	}

	for _, msg := range msgs {
		fmt.Printf("  %s\n", msg)
	}

	// Совет от "внутреннего ментора"
	fmt.Println("\n💡 СОВЕТ НА ЗАВТРА:")
	if p.Knowledge < 50 {
		fmt.Println("  Сфокусируйся на основах Go. 2 часа кода > 8 часов доставок.")
	} else if p.DaysAsCourier > 5 {
		fmt.Println("  Хватит бегать. Пора обновить резюме и идти на собеседование.")
	} else {
		fmt.Println("  Баланс — ключ. Учи Go, но не забывай про здоровье.")
	}

	fmt.Println(strings.Repeat("█", 60))
}

// Вспомогательная функция для прогресс-бара
func progressBar(value int) string {
	width := 20
	filled := (value * width) / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎮 ЗАПУСК СИМУЛЯТОРА 'CODE COURIER'")
	fmt.Println("  Год: 2026. Место: Москва. Ставка: вся жизнь.")

	player := NewPlayer("Гоша")

	// Игровой цикл дня
	player.MorningDepression()
	player.DisplayStatus()

	// 3 ключевых выбора за день
	for i := 0; i < 3; i++ {
		player.MakeChoice()
		player.DisplayStatus()
		time.Sleep(2 * time.Second)
	}

	// Финальный монолог
	fmt.Println("\n🌙 ВЕЧЕР. 23:47. ОКОНЧАНИЕ ИГРОВОГО ДНЯ.")
	fmt.Println(strings.Repeat("~", 50))
	if player.Knowledge > player.DaysAsCourier*10 {
		fmt.Println("  Сегодня был вклад в будущее. Знания растут.")
		fmt.Println("  Каждая изученная тема — кирпич в фундаменте тёплого офиса.")
	} else {
		fmt.Println("  Сегодня было выживание. Мороз, заказы, усталость.")
		fmt.Println("  Но даже в выживании есть выбор. Завтра можешь выбрать иначе.")
	}

	fmt.Println("\n  n8n-бот ждёт новых команд. Go-туториал открыт в соседней вкладке.")
	fmt.Println("  Завтра снова будильник. Снова выбор.")
	fmt.Println("\n  ИГРА ПРОДОЛЖАЕТСЯ...")
	fmt.Println(strings.Repeat("~", 50))
}
