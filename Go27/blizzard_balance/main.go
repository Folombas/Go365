package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ========== DISCLAIMER: ХУДОЖЕСТВЕННЫЙ ВЫМЫСЕЛ ==========
// Blizzard Balance — это арт-проект о балансе между реальностью и мечтой.
// Все персонажи, диалоги и события являются метафорами пути в IT.
// Любые совпадения с реальными людьми или ситуациями случайны.
// Это не автобиография, а программная притча о выборе.
// ======================================================

// Player представляет Гошу с его дилеммами
type Player struct {
	Name                string
	// Основные характеристики (0-100)
	Stamina             int  // Выносливость (физическая)
	Focus               int  // Фокус (ментальный)
	GoKnowledge         int  // Знание Go и экосистемы
	Motivation          int  // Мотивация к изменениям
	MentalHealth        int  // Психическое здоровье
	// Ресурсы
	Money               int  // Финансы (рубли)
	Energy              int  // Энергия (0-100)
	// Статистика
	DaysAsCourier       int
	LinesOfCodeWritten  int
	VideosEdited        int
	// Флаги состояний
	IsInBlizzard        bool  // В метели?
	IsCoding            bool  // Кодит?
	IsEditing           bool  // Монтирует видео?
	HasTropicalMemories bool  // Есть ли тропические воспоминания?
}

// NewPlayer создаёт нового игрока в состоянии "утренняя депрессия + метель"
func NewPlayer(name string) *Player {
	return &Player{
		Name:         name,
		Stamina:      65,
		Focus:        40,
		GoKnowledge:  45,
		Motivation:   35,
		MentalHealth: 50,
		Money:        5000,
		Energy:       70,
		DaysAsCourier: 8,
		HasTropicalMemories: true, // Были же отпуска в тёплых странах!
		IsInBlizzard: true,        // За окном метель
	}
}

// WorkAsCourier — работа в метели (выживание)
func (p *Player) WorkAsCourier() {
	fmt.Println("\n❄️ РАБОТА В МЕТЕЛИ:")
	
	orders := []struct {
		desc   string
		pay    int
		stamCost int
	}{
		{"Цветы через город (замерзли)", 300, 20},
		{"Тяжёлая техника в офис", 500, 25},
		{"Еда по снежным дворам", 250, 15},
		{"Срочный документ в больницу", 400, 18},
	}
	
	totalEarned := 0
	for _, order := range orders {
		if p.Stamina < 10 {
			fmt.Println("  ⚠️  Силы кончились. Не могу больше.")
			break
		}
		
		fmt.Printf("  🚴 %s\n", order.desc)
		p.Stamina -= order.stamCost
		p.Money += order.pay
		totalEarned += order.pay
		p.Energy -= 8
		time.Sleep(800 * time.Millisecond)
	}
	
	p.DaysAsCourier++
	fmt.Printf("\n  💰 Заработано: %d руб. Выносливость: %d/100\n", totalEarned, p.Stamina)
	fmt.Printf("  📅 Дней курьером: %d\n", p.DaysAsCourier)
	
	// Эффекты
	if p.Stamina < 30 {
		p.MentalHealth -= 10
		fmt.Println("  💔 Низкая выносливость бьёт по психике")
	}
}

// PayBills — оплата счетов (взрослая жизнь)
func (p *Player) PayBills() {
	fmt.Println("\n🧾 ОПЛАТА СЧЕТОВ:")
	
	bills := map[string]int{
		"Налоги":          1200,
		"Перевод маме":    2000,
		"Коммуналка":      800,
		"Интернет":        500,
	}
	
	total := 0
	for name, amount := range bills {
		if p.Money < amount {
			fmt.Printf("  ⚠️  Не хватает на %s (%d руб)\n", name, amount)
			continue
		}
		
		p.Money -= amount
		total += amount
		fmt.Printf("  📄 %s: -%d руб\n", name, amount)
		
		if name == "Перевод маме" {
			p.MentalHealth += 5 // Это важно для души
			fmt.Println("    ❤️  Мама будет рада (+5 к психике)")
		}
	}
	
	fmt.Printf("  🧮 Итого потрачено: %d руб. Осталось: %d руб\n", total, p.Money)
}

// StudyGo — изучение Go (инвестиция в будущее)
func (p *Player) StudyGo(hours int) {
	if hours <= 0 {
		return
	}
	
	fmt.Printf("\n📚 ИЗУЧЕНИЕ GO (%d часа):\n", hours)
	p.IsCoding = true
	
	topics := []string{
		"Горутины и параллелизм",
		"Интерфейсы и полиморфизм",
		"Работа с базами данных",
		"Тестирование (go test)",
		"Web (Echo/Gin)",
		"Оптимизация и профилирование",
		"GRPC и микросервисы",
		"Кэширование (Redis)",
		"Docker и деплой",
		"CI/CD пайплайны",
	}
	
	linesPerHour := 30
	totalLines := 0
	
	for h := 1; h <= hours; h++ {
		if p.Focus < 20 {
			fmt.Println("  ⚠️  Фокус на нуле. Невозможно учиться.")
			break
		}
		
		topic := topics[rand.Intn(len(topics))]
		lines := linesPerHour + rand.Intn(20)
		totalLines += lines
		
		fmt.Printf("  🕐 Час %d: %s (+%d строк кода)\n", h, topic, lines)
		
		// Прогресс
		p.GoKnowledge += 8 + rand.Intn(5)
		p.LinesOfCodeWritten += lines
		p.Focus -= 5
		p.Energy -= 10
		
		// Случайная мотивация
		if rand.Intn(10) < 3 {
			p.printGoMotivation()
		}
		
		time.Sleep(700 * time.Millisecond)
	}
	
	p.IsCoding = false
	fmt.Printf("\n  ✅ Итого: %d строк кода. Знание Go: %d/100\n", totalLines, p.GoKnowledge)
	
	// Долгосрочный эффект
	if hours >= 2 {
		p.Motivation += 15
		fmt.Println("  🚀 Длительная учёба даёт +15 к мотивации!")
	}
}

// EditVacationVideo — монтаж видео из отпуска (перезагрузка)
func (p *Player) EditVacationVideo(minutes int) {
	if !p.HasTropicalMemories {
		fmt.Println("\n😔 Нет тропических воспоминаний для монтажа.")
		fmt.Println("  Совет: сначала нужно где-то побывать!")
		return
	}
	
	if minutes < 15 {
		fmt.Println("\n⏱️  Монтаж меньше 15 минут не имеет терапевтического эффекта!")
		return
	}
	
	fmt.Printf("\n🎬 МОНТАЖ ВИДЕО ИЗ ОТПУСКА (%d минут):\n", minutes)
	p.IsEditing = true
	
	// 10 причин, почему это НЕ лишнее
	reasons := []string{
		"1. Контекстное переключение: с 'снег/работа' на 'солнце/отпуск'",
		"2. Творческий процесс без давления дедлайнов и клиентов",
		"3. Визуализация цели: напоминание, ради чего всё затевалось",
		"4. Дофамин от создания: 'я сделал красиво!'",
		"5. Психологическая разгрузка: мозг отдыхает от кода и заказов",
		"6. Навык монтажа = +1 к digital-скиллам (никогда не лишнее)",
		"7. Контент для блога/соцсетей = потенциальная аудитория",
		"8. Практика английского: монтаж для иностранных подписчиков",
		"9. Терапия цветом: тропическая палитра vs зимняя монохромность",
		"10. Баланс: нельзя только страдать, надо иногда 'включать лето'",
	}
	
	fmt.Println("  🌴 Процесс:")
	steps := []string{
		"Отбор кадров с пляжа",
		"Цветокоррекция (делаем море ярче)",
		"Наложение лёгкой музыки",
		"Добавление субтитров",
		"Создание интро/аутро",
		"Экспорт и загрузка",
	}
	
	for i, step := range steps {
		if i >= minutes/5 { // Упрощённая прогрессия
			break
		}
		fmt.Printf("  ✨ Шаг %d: %s\n", i+1, step)
		time.Sleep(600 * time.Millisecond)
	}
	
	// Эффекты
	p.MentalHealth += 15
	p.Energy += 10
	p.VideosEdited++
	p.Focus += 5  // Перезагрузили мозг
	
	fmt.Printf("\n  🧠 Эффект: +15 к психике, +10 к энергии\n")
	fmt.Printf("  🎥 Смонтировано видео: %d\n", p.VideosEdited)
	
	// Вывод 2-3 случайных причин
	fmt.Println("\n  💡 Почему это НЕ лишнее:")
	for i := 0; i < 3; i++ {
		fmt.Printf("  %s\n", reasons[rand.Intn(len(reasons))])
	}
	
	p.IsEditing = false
}

// printGoMotivation — 10 мотивационных фраз для изучения Go
func (p *Player) printGoMotivation() {
	phrases := []string{
		"🚀 GO: Каждый пакет — кирпич в фундаменте тёплого офиса.",
		"🔥 Простота Go — не для слабых. Она для тех, кто хочет решать задачи, а не языковые головоломки.",
		"💎 Знание горутин сегодня = параллельная обработка зарплат завтра.",
		"🎯 100 дней Go = 100% больше шансов сменить -18°C на +22°C в офисе.",
		"⚡ Пока другие спорят о лучшем языке, ты уже пишешь на том, что работает в Google, Uber и Docker.",
		"🧠 Go учит думать о данных и потоках, а не о синтаксическом сахаре. Это мышление архитектора.",
		"🛡️ Статическая типизация Go = страховка от багов в 3 часа ночи перед дедлайном.",
		"📈 Рынок: Go-разработчики нужны. Не там, где 'и так сойдёт', а там, где масштаб и надёжность.",
		"🏆 Выучить Go — значит получить ключ от backend-мира, где всё серьезно и оплачивается соответственно.",
		"🌅 Завтра: ты либо снова в метели с заказами, либо пишешь код для системы, которую используют миллионы.",
	}
	
	fmt.Printf("  💫 МОТИВАЦИЯ: %s\n", phrases[rand.Intn(len(phrases))])
	p.Motivation += 5
}

// DailyChoice — ключевой выбор дня
func (p *Player) DailyChoice() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("  ВРЕМЯ ВЫБОРА. СЕГОДНЯ ТЫ МОЖЕШЬ:")
	fmt.Println("  1) 🏃 БЕЖАТЬ В МЕТЕЛЬ (заработать, но потратить силы)")
	fmt.Println("  2) 🐹 УЧИТЬ GO (инвестировать в будущее, но без денег сегодня)")
	fmt.Println("  3) 🎬 МОНТИРОВАТЬ ВИДЕО (восстановить психику, но потратить время)")
	fmt.Println("  4) 💸 ОПЛАТИТЬ СЧЕТА (взрослая жизнь зовёт)")
	fmt.Println("  5) 😴 ОТДОХНУТЬ (восстановить энергию)")
	fmt.Println(strings.Repeat("═", 60))
	
	// В реальной игре здесь был бы ввод, но симулируем выбор на основе состояния
	choice := p.makeSmartChoice()
	
	switch choice {
	case 1:
		p.WorkAsCourier()
	case 2:
		hours := 1 + rand.Intn(3)
		p.StudyGo(hours)
	case 3:
		minutes := 15 + rand.Intn(45)
		p.EditVacationVideo(minutes)
	case 4:
		p.PayBills()
	case 5:
		p.rest()
	}
}

// makeSmartChoice — ИИ выбора на основе текущего состояния
func (p *Player) makeSmartChoice() int {
	// Приоритеты в зависимости от состояния
	if p.Money < 2000 {
		return 1 // Срочно нужны деньги
	}
	if p.MentalHealth < 40 {
		if p.HasTropicalMemories && rand.Intn(10) < 7 {
			return 3 // Срочно нужно подлечить психику
		}
		return 5 // Или просто отдохнуть
	}
	if p.GoKnowledge < 60 && p.Energy > 40 {
		return 2 // Учиться, пока есть силы
	}
	if p.Energy < 30 {
		return 5 // Срочно отдыхать
	}
	
	// Иначе случайный выбор (но не оплата счетов, если всё оплачено)
	choices := []int{1, 2, 3, 5}
	return choices[rand.Intn(len(choices))]
}

// rest — отдых и восстановление
func (p *Player) rest() {
	fmt.Println("\n😴 ОТДЫХ И ВОССТАНОВЛЕНИЕ:")
	
	activities := []string{
		"Чай с мамой на кухне",
		"Короткий сон",
		"Медитация под шум метели",
		"Просмотр смешных видео",
		"Лёгкая растяжка",
	}
	
	for i := 0; i < 3; i++ {
		activity := activities[rand.Intn(len(activities))]
		fmt.Printf("  ✨ %s\n", activity)
		
		p.Energy += 15
		p.MentalHealth += 5
		p.Stamina += 10
		
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Printf("\n  🔋 Энергия: +45, Психика: +15, Выносливость: +30\n")
}

// DisplayStatus — отображение статуса
func (p *Player) DisplayStatus() {
	fmt.Println("\n" + strings.Repeat("█", 60))
	fmt.Println("📊 BLIZZARD BALANCE — СТАТУС ИГРОКА:")
	fmt.Printf("  Игрок: %s\n", p.Name)
	fmt.Printf("  📅 День: 27 января 2026 | Курьерских дней: %d\n", p.DaysAsCourier)
	
	// Основные характеристики
	fmt.Println("\n  🧬 ХАРАКТЕРИСТИКИ:")
	fmt.Printf("    Выносливость: [%s] %d/100\n", progressBar(p.Stamina), p.Stamina)
	fmt.Printf("    Фокус:        [%s] %d/100\n", progressBar(p.Focus), p.Focus)
	fmt.Printf("    Знание Go:    [%s] %d/100\n", progressBar(p.GoKnowledge), p.GoKnowledge)
	fmt.Printf("    Мотивация:    [%s] %d/100\n", progressBar(p.Motivation), p.Motivation)
	fmt.Printf("    Психика:      [%s] %d/100\n", progressBar(p.MentalHealth), p.MentalHealth)
	
	// Ресурсы
	fmt.Println("\n  📦 РЕСУРСЫ:")
	fmt.Printf("    Энергия:      [%s] %d/100\n", progressBar(p.Energy), p.Energy)
	fmt.Printf("    Деньги:       %d руб.\n", p.Money)
	
	// Достижения
	fmt.Println("\n  🏆 ДОСТИЖЕНИЯ:")
	fmt.Printf("    Написано строк кода: %d\n", p.LinesOfCodeWritten)
	fmt.Printf("    Смонтировано видео:  %d\n", p.VideosEdited)
	
	// Состояния
	fmt.Println("\n  🎭 СОСТОЯНИЯ:")
	states := []string{}
	if p.IsInBlizzard {
		states = append(states, "❄️ В метели")
	}
	if p.IsCoding {
		states = append(states, "💻 Кодит")
	}
	if p.IsEditing {
		states = append(states, "🎬 Монтирует")
	}
	if p.HasTropicalMemories {
		states = append(states, "🌴 Есть тёплые воспоминания")
	}
	if len(states) == 0 {
		states = append(states, "😐 Обычное состояние")
	}
	for _, state := range states {
		fmt.Printf("    %s\n", state)
	}
	
	// Совет системы
	fmt.Println("\n  💡 СОВЕТ СИСТЕМЫ:")
	if p.Money < 1000 {
		fmt.Println("    ❗ Срочно нужны деньги! Выбирай 'РАБОТА В МЕТЕЛИ'")
	} else if p.MentalHealth < 40 {
		fmt.Println("    ❗ Психика на пределе! Выбирай 'МОНТАЖ ВИДЕО' или 'ОТДЫХ'")
	} else if p.GoKnowledge < 50 {
		fmt.Println("    🎯 Фокус на Go! Меньше 50 знаний — это риск остаться в метели")
	} else if p.Energy < 40 {
		fmt.Println("    ⚡ Нужна энергия! Отдохни перед завтрашним днём")
	} else {
		fmt.Println("    ✅ Баланс в норме. Двигайся к цели, но не забывай о себе")
	}
	
	fmt.Println(strings.Repeat("█", 60))
}

// progressBar — создаёт текстовый прогресс-бар
func progressBar(value int) string {
	const width = 20
	filled := (value * width) / 100
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return bar
}

// printFinalThoughts — итоговые мысли дня
func (p *Player) printFinalThoughts() {
	fmt.Println("\n" + strings.Repeat("✨", 60))
	fmt.Println("  ИТОГИ ДНЯ 27.01.2026:")
	fmt.Println()
	
	if p.LinesOfCodeWritten > 50 {
		fmt.Println("  🎉 Ты написал больше 50 строк кода! Это реальный прогресс.")
		fmt.Println("    Каждая строка — шаг от метели к тёплому офису.")
	}
	
	if p.VideosEdited > 0 {
		fmt.Println("  🎬 Ты смонтировал видео из отпуска! И это правильно.")
		fmt.Println("    Психика важнее. Нельзя только страдать — надо иногда вспоминать лето.")
	}
	
	if p.Money > 3000 {
		fmt.Println("  💰 Финансовая подушка есть. Можно меньше паниковать.")
	} else {
		fmt.Println("  ⚠️  Денег маловато. Но помни: это временно.")
	}
	
	if p.GoKnowledge > 60 {
		fmt.Println("  🚀 Знание Go перевалило за 60! Ты уже на пути к junior-позиции.")
		fmt.Println("    Обновляй резюме, скоро собеседования.")
	} else if p.GoKnowledge > 40 {
		fmt.Println("  📚 Знание Go растёт. Продолжай в том же духе.")
		fmt.Println("    Идеальный баланс: 2 часа кода в день = рост без выгорания.")
	}
	
	fmt.Println()
	fmt.Println("  ФИНАЛЬНАЯ МЫСЛЬ:")
	fmt.Println("  Баланс — не про то, чтобы всё делать понемногу.")
	fmt.Println("  Баланс — про то, чтобы не сломаться, пока идешь к цели.")
	fmt.Println("  Иногда нужно учить Go. Иногда — монтировать видео из Таиланда.")
	fmt.Println("  Главное — не забывать, куда и зачем идешь.")
	fmt.Println()
	fmt.Println("  Завтра будет новый день. И новый выбор.")
	fmt.Println(strings.Repeat("✨", 60))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("❄️ BLIZZARD BALANCE: Snowstorm vs Code — День 27")
	fmt.Println("   Симулятор баланса между выживанием и мечтой")
	fmt.Println(strings.Repeat("─", 60))
	
	// Создаём игрока
	player := NewPlayer("Гоша")
	player.DisplayStatus()
	
	// Симуляция дня (5 ключевых выборов)
	fmt.Println("\n🌅 НАЧАЛО ДНЯ:")
	fmt.Println("  Утро. Метель. Депрессия. Но есть выбор.")
	
	for i := 1; i <= 5; i++ {
		fmt.Printf("\n🕐 ВЫБОР #%d:\n", i)
		player.DailyChoice()
		player.DisplayStatus()
		
		// Пауза между выборами
		if i < 5 {
			fmt.Println("\n  ... время проходит ...")
			time.Sleep(1 * time.Second)
		}
	}
	
	// Итоги дня
	player.printFinalThoughts()
	
	// Специальный вывод: если игрок в балансе
	if player.MentalHealth > 60 && player.GoKnowledge > 50 && player.Money > 2000 {
		fmt.Println("\n🏆 ДОСТИЖЕНИЕ РАЗБЛОКИРОВАНО: 'ГАРМОНИЧНЫЙ БАЛАНС'")
		fmt.Println("   Ты нашёл золотую середину между выживанием и развитием.")
		fmt.Println("   Так держать! Именно этот баланс приведёт тебя к цели.")
	}
	
	fmt.Println("\n🎮 Игра продолжается... Завтра снова выбор.")
}
