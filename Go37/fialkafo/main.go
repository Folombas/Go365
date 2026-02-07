package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// ================== СТРУКТУРЫ ДАННЫХ ==================

type Cargo struct {
	ID          string
	Content     string
	Destination string
	Status      string
	Value       int
}

type Truck struct {
	ID      string
	Driver  string
	Company string
	Speed   int
	channel chan *Cargo
	quit    chan bool
}

type Dispatcher struct {
	sync.Mutex
	trucks     map[string]*Truck
	companies  map[string]int // рейтинги компаний
	cargoQueue chan *Cargo
	delivered  chan *Cargo
	failed     chan *Cargo
	chat       []string
	money      int
	dayCount   int
	goXP       int
}

// ================== КОМПАНИИ-ЯЗЫКИ ==================

var companies = map[string]string{
	"Go Logistics":        "🚚 Транспортная логистика (горутины, каналы)",
	"Java Trails":         "🛤️ Железные дороги (мощные составы, enterprise)",
	"Rust Garage Lab":     "🔧 Гараж-лаборатория (безопасные вездеходы)",
	"Python Data Gardens": "🌿 Теплицы данных (AI/ML, быстрорастущее)",
	"JavaScript Circus":   "🎪 Цирк-шапито (динамичное, вездесущее)",
	"C++ Foundry":         "🏭 Литейный цех (высокооптимизированные движки)",
}

// ================== ИНИЦИАЛИЗАЦИЯ ==================

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		trucks:     make(map[string]*Truck),
		companies:  make(map[string]int),
		cargoQueue: make(chan *Cargo, 50),
		delivered:  make(chan *Cargo, 30),
		failed:     make(chan *Cargo, 10),
		money:      1000,
		goXP:       350,
		dayCount:   365, // год работы
	}

	// Инициализируем рейтинги компаний
	for company := range companies {
		d.companies[company] = rand.Intn(100)
	}
	d.companies["Go Logistics"] = 85 // Go специально повыше

	// Создаем фуры для Go Logistics
	drivers := []string{"Петрович", "Семёныч", "Михалыч", "Иваныч", "Фёдорыч"}
	for i := 0; i < 5; i++ {
		truck := &Truck{
			ID:      fmt.Sprintf("GO-FURA-%03d", i+1),
			Driver:  drivers[i],
			Company: "Go Logistics",
			Speed:   rand.Intn(30) + 60,
			channel: make(chan *Cargo, 3),
			quit:    make(chan bool),
		}
		d.trucks[truck.ID] = truck
		go truck.work(d.delivered, d.failed)
	}

	return d
}

// ================== РАБОТА ФУРЫ ==================

func (t *Truck) work(delivered, failed chan<- *Cargo) {
	for {
		select {
		case cargo := <-t.channel:
			// Симуляция доставки с возможными проблемами
			time.Sleep(time.Duration(1000/t.Speed) * time.Millisecond * 10)

			// 15% шанс на проблему
			if rand.Intn(100) < 15 {
				// Проблемы Go Logistics
				problems := []string{
					"туман на трассе",
					"поломка двигателя",
					"пробка на кольцевой",
					"проверка ГАИ",
					"размытая дорога",
				}
				cargo.Status = problems[rand.Intn(len(problems))]
				failed <- cargo
			} else {
				cargo.Status = "доставлено"
				delivered <- cargo
			}

		case <-t.quit:
			return
		}
	}
}

// ================== РАБОЧИЙ ДЕНЬ ==================

func (d *Dispatcher) workDay(dayNumber int) {
	fmt.Printf("\n🌅 ДЕНЬ %d РАБОТЫ В GO LOGISTICS\n", dayNumber)
	fmt.Println(strings.Repeat("=", 50))

	// Генерация грузов на день
	go d.generateCargo(dayNumber)

	// Распределение грузов
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		d.distributeCargo()
	}()

	// Сбор результатов
	deliveredToday := 0
	failedToday := 0
	earnedToday := 0

	timeout := time.After(5 * time.Second)
	for {
		select {
		case cargo := <-d.delivered:
			deliveredToday++
			earnedToday += cargo.Value
			d.money += cargo.Value
			d.goXP += 10
			d.addToChat(fmt.Sprintf("✅ %s доставил %s → %s",
				strings.Split(cargo.ID, "-")[2], cargo.Content, cargo.Destination))

		case cargo := <-d.failed:
			failedToday++
			d.addToChat(fmt.Sprintf("❌ Сбой: %s - %s", cargo.Content, cargo.Status))
			// Retry через 2 дня
			go func(c *Cargo) {
				time.Sleep(2 * time.Second)
				d.cargoQueue <- c
			}(cargo)

		case <-timeout:
			// День окончен
			fmt.Printf("\n📊 ИТОГИ ДНЯ:\n")
			fmt.Printf("   📦 Доставлено: %d грузов\n", deliveredToday)
			fmt.Printf("   ❌ Сбоев: %d\n", failedToday)
			fmt.Printf("   💰 Заработано: %d $\n", earnedToday)
			fmt.Printf("   🧠 Go XP: %d (+%d)\n", d.goXP, deliveredToday*10)
			fmt.Printf("   🏦 Общий капитал: %d $\n", d.money)

			// Проверка на предложения от других компаний
			if dayNumber%30 == 0 {
				d.checkOtherCompanies(dayNumber)
			}

			return
		}
	}
}

func (d *Dispatcher) generateCargo(dayNumber int) {
	destinations := []string{
		"микросервис Auth",
		"сервис Пользователей",
		"платежный шлюз",
		"сервис Уведомлений",
		"хранилище S3",
		"база данных Postgres",
		"кэш Redis",
		"очередь Kafka",
	}

	contents := []string{
		"сессия пользователя",
		"транзакция оплаты",
		"логи аудита",
		"метрики производительности",
		"конфигурация сервиса",
		"сообщение в чате",
		"файл изображения",
		"документ PDF",
	}

	for i := 0; i < 15; i++ {
		cargo := &Cargo{
			ID:          fmt.Sprintf("CRG-%s-%04d", time.Now().Format("0102"), i+1),
			Content:     contents[rand.Intn(len(contents))],
			Destination: destinations[rand.Intn(len(destinations))],
			Value:       rand.Intn(200) + 50,
			Status:      "в обработке",
		}
		d.cargoQueue <- cargo
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
	close(d.cargoQueue)
}

func (d *Dispatcher) distributeCargo() {
	for cargo := range d.cargoQueue {
		// Ищем свободную фуру
		assigned := false
		for _, truck := range d.trucks {
			if truck.Company == "Go Logistics" {
				select {
				case truck.channel <- cargo:
					assigned = true
					break
				default:
					continue
				}
			}
			if assigned {
				break
			}
		}

		if !assigned {
			// Все фуры заняты - ждем
			time.Sleep(100 * time.Millisecond)
			d.cargoQueue <- cargo
		}
	}
}

// ================== ПРЕДЛОЖЕНИЯ ОТ ДРУГИХ КОМПАНИЙ ==================

func (d *Dispatcher) checkOtherCompanies(dayNumber int) {
	fmt.Println("\n📬 ПРЕДЛОЖЕНИЯ ОТ ДРУГИХ КОМПАНИЙ:")
	fmt.Println(strings.Repeat("-", 40))

	offers := []struct {
		company string
		salary  int
		message string
	}{
		{
			company: "Java Trails",
			salary:  180000,
			message: "Стальные магистрали нуждаются в опытном диспетчере!",
		},
		{
			company: "Rust Garage Lab",
			salary:  220000,
			message: "Строим будущее безопасного транспорта!",
		},
		{
			company: "Python Data Gardens",
			salary:  160000,
			message: "Теплицы данных ждут своего садовника!",
		},
	}

	for _, offer := range offers {
		fmt.Printf("\n🏢 %s\n", offer.company)
		fmt.Printf("   💰 ЗП: %d руб.\n", offer.salary)
		fmt.Printf("   📝 %s\n", offer.message)
		fmt.Printf("   🤔 Принять предложение? (сейчас: %d руб.)\n", d.money/12)

		// Гоша каждый раз отказывается
		d.addToChat(fmt.Sprintf("❌ Отклонено предложение от %s", offer.company))
	}

	fmt.Println("\n🎯 Гоша: 'Нет, спасибо. Я остаюсь с фурами-горутинами и дорогами-каналами.'")
	fmt.Println("        'Go Logistics - это мой выбор. Горутины доставят.'")
}

// ================== ИГРОВАЯ МЕХАНИКА ==================

func (d *Dispatcher) addToChat(message string) {
	d.Lock()
	d.chat = append(d.chat, fmt.Sprintf("[%s] %s",
		time.Now().Format("15:04"), message))

	// Ограничиваем историю чата
	if len(d.chat) > 20 {
		d.chat = d.chat[1:]
	}
	d.Unlock()
}

func (d *Dispatcher) printChat() {
	fmt.Println("\n💬 ЧАТ ДИСПЕТЧЕРСКОЙ:")
	fmt.Println(strings.Repeat("-", 40))
	for _, msg := range d.chat {
		fmt.Println(msg)
	}
}

// ================== ОСНОВНАЯ ПРОГРАММА ==================

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("🚚 FIALKA FO: ГОД В GO LOGISTICS")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("Диспетчер: Гоша")
	fmt.Println("Стаж: 1 год в компании")
	fmt.Println("Специализация: Управление фурами-горутинами")

	// Инициализация
	dispatcher := NewDispatcher()

	// Прошлые 30 дней (симуляция)
	for day := 335; day <= 365; day++ {
		dispatcher.cargoQueue = make(chan *Cargo, 50)
		dispatcher.workDay(day)

		// Каждые 7 дней показываем чат
		if day%7 == 0 {
			dispatcher.printChat()
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Итоги года
	fmt.Println("\n" + strings.Repeat("⭐", 60))
	fmt.Println("🎉 ГОД В GO LOGISTICS ЗАВЕРШЁН!")
	fmt.Println(strings.Repeat("⭐", 60))
	fmt.Printf("\n📈 ФИНАЛЬНАЯ СТАТИСТИКА:\n")
	fmt.Printf("   📦 Всего обработано грузов: ~%d\n", 365*10)
	fmt.Printf("   🧠 Накопленный Go XP: %d\n", dispatcher.goXP)
	fmt.Printf("   💰 Капитал компании: %d $\n", dispatcher.money)
	fmt.Printf("   🚛 Фур в парке: %d\n", len(dispatcher.trucks))
	fmt.Printf("   ❌ Отклонено предложений: 12\n")

	// Философский итог
	fmt.Println("\n" + strings.Repeat("🛣️", 40))
	fmt.Println("ГОША: 'Год назад я выбирал между компаниями...'")
	fmt.Println("      'Java Trails обещала стабильность рельсов.'")
	fmt.Println("      'Rust Garage Lab - инновационные вездеходы.'")
	fmt.Println("      'Но я выбрал дороги. Обычные, асфальтовые.'")
	fmt.Println("      'Где каждая фура - горутина, каждый маршрут - канал.'")
	fmt.Println("      'И знаете что? Грузы доставляются.'")
	fmt.Println("      'Не всегда быстро. Не без пробок и поломок.'")
	fmt.Println("      'Но доставляются. Изо дня в день.'")
	fmt.Println("      'И завтра будут новые грузы. Новые маршруты.'")
	fmt.Println("      'А я буду здесь. В диспетчерской.'")
	fmt.Println("      'Потому что это мой выбор. И он правильный.'")
	fmt.Println(strings.Repeat("🛣️", 40))

	// Обзор всех компаний
	fmt.Println("\n🏢 ОБЗОР IT-КОМПАНИЙ 2026:")
	fmt.Println(strings.Repeat("-", 40))
	for company, description := range companies {
		rating := dispatcher.companies[company]
		fmt.Printf("\n%s\n", company)
		fmt.Printf("   %s\n", description)
		fmt.Printf("   📊 Рейтинг: %d/100\n", rating)
	}

	fmt.Println("\n" + strings.Repeat("★", 60))
	fmt.Println("DISCLAIMER: Все персонажи вымышлены.")
	fmt.Println("Образовательная программа для изучения принципов Go.")
	fmt.Println(strings.Repeat("★", 60))
}
