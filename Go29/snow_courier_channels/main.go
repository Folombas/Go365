package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Структура курьера Гоши
type Courier struct {
	Name        string
	Health      int     // 0-100
	Willpower   int     // 0-100
	Money       int     // рублей
	XP          int     // опыт программирования
	Deliveries  int     // выполнено доставок
	CurrentLoad float64 // текущая нагрузка (кг)
	MaxLoad     float64 // макс нагрузка
	Skills      []string
	Location    string
	InSnowstorm bool
	AngerLevel  int // уровень злости на систему
}

// Заказ на доставку
type Delivery struct {
	ID         int
	From       string
	To         string
	Distance   float64 // км
	Weight     float64 // кг
	BasePrice  int     // базовая цена
	FinalPrice int     // цена после "накрутки" сервиса
	CourierCut int     // что получает курьер
	Difficulty int     // сложность 0-100
	Completed  bool
	TimeLimit  time.Duration // время на выполнение
}

func main() {
	fmt.Println("❄️ День 29: БЕСЦЕННЫЙ ТРУД КУРЬЕРА В МЕТЕЛИ ❄️")
	fmt.Println("Эпическая история о мотивации, снеге и каналах Go")
	fmt.Println(strings.Repeat("=", 60) + "\n")

	rand.Seed(time.Now().UnixNano())

	// Инициализируем Гошу
	gosh := Courier{
		Name:        "Гоша-Курьер",
		Health:      85,
		Willpower:   65,
		Money:       1200,
		XP:          2800, // 28 дней уже учит Go
		Deliveries:  0,
		CurrentLoad: 0,
		MaxLoad:     15.0,
		Skills:      []string{"Golang Basics", "Functions", "Structs", "Goroutines"},
		Location:    "Домашняя кухня",
		InSnowstorm: false,
		AngerLevel:  40,
	}

	// Создаем каналы для событий
	ordersChan := make(chan Delivery, 5) // канал заказов (буферизованный)
	weatherChan := make(chan string)     // канал погоды
	priceHikeChan := make(chan int)      // канал повышения цен
	motivationChan := make(chan string)  // канал мотивации
	resultsChan := make(chan string, 20) // канал результатов

	var wg sync.WaitGroup

	// Горутина метели
	wg.Add(1)
	go func() {
		defer wg.Done()
		snowEvents := []string{
			"❄️ СНЕГОПАД УСИЛИВАЕТСЯ: Видимость падает!",
			"🌬️ ВЕТЕР ДО 20 М/С: Стоять тяжело!",
			"☃️ СУГРОБЫ 50 СМ: Приходится карабкаться!",
			"🚧 ДОРОГИ ЗАМЕТАЕТ: Карабкаешься по снежным сугробам!",
			"🌨️ МОКРЫЙ СНЕГ: Промокаешь насквозь!",
			"⛄ СНЕЖНАЯ БУРЯ: Кажется, это конец света!",
		}

		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(1500)+800) * time.Millisecond)
			event := snowEvents[rand.Intn(len(snowEvents))]
			weatherChan <- event
			gosh.Health -= rand.Intn(5) + 2
			gosh.Willpower -= rand.Intn(3) + 1
		}
		close(weatherChan)
	}()

	// Горутина повышения цен (циничная система)
	wg.Add(1)
	go func() {
		defer wg.Done()
		hikes := []int{150, 200, 180, 250, 120}

		for i := 0; i < 4; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond)
			hike := hikes[rand.Intn(len(hikes))]
			priceHikeChan <- hike
			gosh.AngerLevel += 15

			// Ирония системы
			resultsChan <- fmt.Sprintf("💰 СИСТЕМА: 'Цены подняты на %d%%!' (Но тебе всё равно дадут 500 рублей с копейками)", hike)
		}
		close(priceHikeChan)
	}()

	// Горутина мотивации (внутренний голос Гоши)
	wg.Add(1)
	go func() {
		defer wg.Done()
		motivations := []string{
			"💪 Я НЕ БРОШУ ИЗУЧЕНИЕ GO! Металь меня не сломит!",
			"👨‍💻 Лучше потрачу время на видео-курсы по Go!",
			"⚡ Каждый день с кодом - шаг от этой метели!",
			"🚀 Когда-нибудь я буду писать код в теплом офисе!",
			"🎯 CapCut подождёт! Сначала - карьера программиста!",
			"💼 Таскать коробки в метель за гроши? НЕТ! Учить Go!",
		}

		for i := 0; i < 6; i++ {
			time.Sleep(time.Duration(rand.Intn(2500)+1500) * time.Millisecond)
			motivationChan <- motivations[rand.Intn(len(motivations))]
			gosh.Willpower += 10
			gosh.XP += 50
		}
		close(motivationChan)
	}()

	// Горутина заказов
	wg.Add(1)
	go func() {
		defer wg.Done()
		deliveries := []Delivery{
			{1, "Склад на улице", "Новопеределкино", 25.5, 12.0, 800, 2000, 300, 60, false, 2 * time.Hour},
			{2, "Бизнес-центр", "Север Москвы", 18.2, 8.5, 600, 1800, 250, 45, false, 90 * time.Minute},
			{3, "Торговый центр", "Спортивная", 12.7, 5.0, 400, 1200, 200, 30, false, 60 * time.Minute},
			{4, "Аэропорт", "Центр", 32.0, 15.0, 1200, 3600, 400, 80, false, 150 * time.Minute},
			{5, "Склад", "Квартира клиента", 3.2, 20.0, 500, 1500, 180, 70, false, 45 * time.Minute},
		}

		for _, order := range deliveries {
			time.Sleep(time.Duration(rand.Intn(1200)+600) * time.Millisecond)
			ordersChan <- order
			resultsChan <- fmt.Sprintf("📦 ПОСТУПИЛ ЗАКАЗ: %s → %s (%.1fкг, %d руб. для тебя)",
				order.From, order.To, order.Weight, order.CourierCut)
		}
		close(ordersChan)
	}()

	// Главная горутина - Гоша принимает решения
	wg.Add(1)
	go func() {
		defer wg.Done()
		resultsChan <- fmt.Sprintf("👤 %s начинает день. Здоровье: %d%%, Решимость: %d%%",
			gosh.Name, gosh.Health, gosh.Willpower)
		resultsChan <- "🍝 Завтрак: макароны с котлеткой + горячий кисель"
		resultsChan <- "👕 Оделся в уличное, смотрит в окно на метель..."

		gosh.Location = "В пути"
		gosh.InSnowstorm = true

		deliveryCount := 0

		// Используем select для обработки событий из разных каналов
		for gosh.Health > 0 && gosh.Willpower > 20 && deliveryCount < 5 {
			select {
			case order, ok := <-ordersChan:
				if !ok {
					ordersChan = nil
				} else {
					// Обработка заказа
					time.Sleep(time.Duration(order.Difficulty*10) * time.Millisecond)

					// Шанс успеха зависит от здоровья и метели
					successChance := gosh.Health - order.Difficulty/2
					if rand.Intn(100) < successChance {
						gosh.Money += order.CourierCut
						gosh.Deliveries++
						gosh.CurrentLoad += order.Weight
						gosh.Health -= order.Difficulty / 10
						resultsChan <- fmt.Sprintf("✅ ЗАКАЗ #%d ДОСТАВЛЕН! +%d руб. (система взяла %d руб.)",
							order.ID, order.CourierCut, order.FinalPrice-order.CourierCut)

						// Злость растет от несправедливости
						unfairness := (order.FinalPrice - order.CourierCut) / 100
						gosh.AngerLevel += unfairness
					} else {
						gosh.Health -= 15
						resultsChan <- fmt.Sprintf("❌ ПРОВАЛ ЗАКАЗА #%d! Застрял в сугробе! -15%% здоровья", order.ID)
					}
					deliveryCount++
				}

			case weather, ok := <-weatherChan:
				if !ok {
					weatherChan = nil
				} else {
					resultsChan <- fmt.Sprintf("🌨️  %s", weather)
					// Каждое погодное событие укрепляет решимость
					if strings.Contains(weather, "СУГРОБ") {
						gosh.Willpower += 5
						resultsChan <- "💪 'Лазил по сугробам и ЗАРЁКСЯ не бросать обучение!'"
					}
				}

			case hike, ok := <-priceHikeChan:
				if !ok {
					priceHikeChan = nil
				} else {
					resultsChan <- fmt.Sprintf("📈 ФИКТИВНОЕ ПОВЫШЕНИЕ ЦЕН: на %d%%", hike)
					resultsChan <- "😡 'Львиную долю оставляют себе. Курьеру - крошки!'"

					// Критический уровень злости дает бонус к обучению
					if gosh.AngerLevel > 70 {
						gosh.XP += 100
						resultsChan <- "🔥 МИЗЕРНЫЕ ДОХОДЫ ДАЖЕ В ЛЮТУЮ МЕТЕЛЬ ТРАНСФОРМИРУЮТСЯ В МОТИВАЦИЮ: +100 XP к изучению Go!"
					}
				}

			case motivation, ok := <-motivationChan:
				if !ok {
					motivationChan = nil
				} else {
					resultsChan <- fmt.Sprintf("🧠 ВНУТРЕННИЙ ГОЛОС: %s", motivation)
					gosh.Skills = append(gosh.Skills, "Channels Basics")
				}

			case <-time.After(3 * time.Second):
				if gosh.AngerLevel > 50 {
					resultsChan <- "🤬 'Нет! Не буду ставить CapCut! Сначала выучу Go и устроюсь программистом!'"
				} else {
					resultsChan <- "🚶 Пробираюсь через метель к следующей точке..."
				}
			}

			// Проверка на предельное состояние
			if gosh.AngerLevel > 90 {
				resultsChan <- "⚡ ДОСТИГНУТ ПРЕДЕЛ ЗЛОСТИ! Решение принято окончательно!"
				resultsChan <- "🎯 БРОСАЮ ЭТУ РАБОТУ! СЕГОДНЯ ЖЕ УЧУ КАНАЛЫ В GO!"
				gosh.Willpower = 100
				break
			}
		}

		// Возвращение домой
		gosh.Location = "Домашняя кухня"
		gosh.InSnowstorm = false
		resultsChan <- "\n🏠 ВОЗВРАЩЕНИЕ ДОМОЙ:"
		resultsChan <- "🍲 Мама приготовила вкусный ужин"
		resultsChan <- "☕ Попил чаёк, согрелся"
		resultsChan <- "💻 СЕЛ ЗА КОМП ЗАНИМАТЬСЯ ПРОГРАММИРОВАНИЕМ"
	}()

	// Вывод результатов в реальном времени
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Читаем и выводим все события
	fmt.Println("📖 ХРОНИКА ДНЯ 29 (читается как триллер):")
	fmt.Println(strings.Repeat("-", 60))

	for event := range resultsChan {
		fmt.Println(event)
		time.Sleep(400 * time.Millisecond)
	}

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 ИТОГОВАЯ СТАТИСТИКА ДНЯ:")
	fmt.Printf("👤 Курьер: %s\n", gosh.Name)
	fmt.Printf("❤️  Здоровье: %d%%\n", gosh.Health)
	fmt.Printf("💪 Решимость: %d%%\n", gosh.Willpower)
	fmt.Printf("😡 Уровень злости на систему: %d/100\n", gosh.AngerLevel)
	fmt.Printf("💰 Заработано: %d руб.\n", gosh.Money)
	fmt.Printf("📦 Выполнено доставок: %d\n", gosh.Deliveries)
	fmt.Printf("🧠 Опыт программирования: %d XP\n", gosh.XP)
	fmt.Printf("🎯 Новые навыки: %v\n", gosh.Skills[len(gosh.Skills)-3:])

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎮 ГЕЙМИФИКАЦИОННЫЕ ДОСТИЖЕНИЯ:")
	if gosh.AngerLevel > 80 {
		fmt.Println("🏆 ТРОФЕЙ: 'ЗЛОСТЬ КАК МОТИВАЦИЯ' - +500 к целеустремленности")
	}
	if gosh.Willpower > 90 {
		fmt.Println("🏆 ТРОФЕЙ: 'НЕСЛОМИМЫЙ ДУХ' - выстоял в метель и не сдался")
	}
	if gosh.XP > 3000 {
		fmt.Println("🏆 ТРОФЕЙ: 'НЕОСТАНОВИМОЕ ОБУЧЕНИЕ' - 3000+ XP за марафон")
	}

	fmt.Println("\n💡 ФИНАЛЬНЫЙ ВЫВОД ДНЯ 29:")
	fmt.Println("Злость на несправедливую систему + метель + сугробы =")
	fmt.Println("= СТАЛЬНАЯ РЕШИМОСТЬ ВЫУЧИТЬ GO И СМЕНИТЬ ПРОФЕССИЮ!")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n🚀 ЗАВТРА: День 30! Месяц обучения Go почти завершен!")
}
