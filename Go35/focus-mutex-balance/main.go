package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"strings"
)

// AttentionResource представляет ресурс внимания Гоши
type AttentionResource struct {
	Owner       string    // Кто владеет вниманием сейчас
	mu          sync.Mutex
	renderQueue chan string
}

// MentalState представляет психическое состояние
type MentalState struct {
	DepressionLevel int
	Energy          int
	mu              sync.RWMutex
}

// DeliveryTask представляет задачу доставки
type DeliveryTask struct {
	From      string
	To        string
	Payment   int
	Completed bool
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🧠 ДЕНЬ 35: МЬЮТЕКСЫ ВНИМАНИЯ")
	fmt.Println(strings.Repeat("═", 50))

	// Часть 1: Утро с депрессией (как deadlock)
	fmt.Println("\n🌫️  ЧАСТЬ 1: УТРО 12:00 - ДЕПРЕССИЯ КАК DEADLOCK")
	fmt.Println(strings.Repeat("-", 35))

	mentalState := &MentalState{DepressionLevel: 70, Energy: 30}

	fmt.Printf("   Уровень депрессии: %d/100\n", mentalState.DepressionLevel)
	fmt.Printf("   Энергия: %d/100\n", mentalState.Energy)
	fmt.Println("   💭 Мысли в deadlock: 'Вставать? Не вставать? Зачем?'")

	// Пытаемся "разблокировать" депрессию
	mentalState.mu.Lock()
	mentalState.DepressionLevel -= 25
	mentalState.Energy += 20
	mentalState.mu.Unlock()

	time.Sleep(1 * time.Second)
	fmt.Println("   ✅ Совершил действие: умылся, пообедал")
	fmt.Printf("   📊 После: Депрессия=%d, Энергия=%d\n",
		mentalState.DepressionLevel, mentalState.Energy)

	// Часть 2: Доставка как mutex
	fmt.Println("\n🚚 ЧАСТЬ 2: ДОСТАВКА - МЬЮТЕКС В ДЕЙСТВИИ")
	fmt.Println(strings.Repeat("-", 35))

	delivery := DeliveryTask{
		From:    "Речной Вокзал",
		To:      "Центр Москвы",
		Payment: 600,
	}

	var deliveryMutex sync.Mutex
	var balance int
	var wg sync.WaitGroup

	fmt.Printf("   📍 Маршрут: %s → %s\n", delivery.From, delivery.To)
	fmt.Printf("   💰 Оплата: %d руб\n", delivery.Payment)

	// Симуляция доставки с мьютексом
	wg.Add(1)
	go func() {
		defer wg.Done()

		deliveryMutex.Lock()
		fmt.Println("   🔒 [Mutex Lock] Начало доставки...")

		// Имитация процесса доставки
		stages := []string{"Автобус", "Метро", "Пешая прогулка", "Кафе"}
		for i, stage := range stages {
			time.Sleep(time.Duration(300+rand.Intn(300)) * time.Millisecond)
			fmt.Printf("   🚌 Этап %d: %s\n", i+1, stage)
		}

		// Безопасное обновление баланса
		oldBalance := balance
		time.Sleep(50 * time.Millisecond) // Искусственная задержка для демонстрации
		balance = oldBalance + delivery.Payment
		delivery.Completed = true

		fmt.Println("   🔓 [Mutex Unlock] Доставка завершена!")
		deliveryMutex.Unlock()
	}()

	wg.Wait()
	fmt.Printf("\n   💵 Баланс после доставки: %d руб\n", balance)

	// Часть 3: Вечерний баланс - Go vs Монтаж
	fmt.Println("\n⚖️  ЧАСТЬ 3: ВЕЧЕР - GO И МОНТАЖ")
	fmt.Println(strings.Repeat("-", 35))

	attention := &AttentionResource{
		Owner:       "Никто",
		renderQueue: make(chan string, 3),
	}

	fmt.Println("   🤔 Решение: Заниматься и Go, и видеомонтажом")
	fmt.Println("   📹 Пока рендерится видео → пишу код на Go")

	// Горутина видеомонтажа (блокирующая операция)
	wg.Add(1)
	go func() {
		defer wg.Done()

		attention.mu.Lock()
		attention.Owner = "CapCut"
		fmt.Println("\n   🎬 [LOCK] CapCut: Начал рендеринг travel-видео")
		fmt.Println("      (это займет 5-10 минут, но можно писать код)")

		// Имитация рендеринга
		renderSteps := []string{"Обработка клипов", "Наложение эффектов", "Цветокоррекция", "Экспорт"}
		for _, step := range renderSteps {
			time.Sleep(1 * time.Second)
			attention.renderQueue <- fmt.Sprintf("Рендеринг: %s", step)
		}

		time.Sleep(1 * time.Second)
		attention.Owner = "Никто"
		fmt.Println("   ✅ [UNLOCK] CapCut: Рендеринг завершен!")
		attention.mu.Unlock()
	}()

	// Горутина программирования на Go
	wg.Add(1)
	go func() {
		defer wg.Done()

		time.Sleep(500 * time.Millisecond) // Ждем начала рендеринга

		// Пока рендерится видео - пишем код
		codeTasks := []string{
			"Изучаю sync.Mutex",
			"Пишу тесты",
			"Рефакторю код",
			"Читаю документацию",
		}

		for i := 0; i < 4; i++ {
			select {
			case renderStatus := <-attention.renderQueue:
				fmt.Printf("   📹 %s\n", renderStatus)
			default:
				// Рендеринг не блокирует написание кода
			}

			fmt.Printf("   💻 Go: %s\n", codeTasks[i])
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()

	// Финальная статистика
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("📊 ИТОГ ДНЯ:")
	fmt.Println(strings.Repeat("-", 50))

	achievements := []struct{
		Name string
		Value string
	}{
		{"Доставка выполнена", "✅ +600 руб"},
		{"Депрессия снижена", fmt.Sprintf("📉 %d → %d", 70, mentalState.DepressionLevel)},
		{"Код написан", "💻 4 задачи по Go"},
		{"Видео смонтировано", "🎬 1 travel-ролик"},
		{"Баланс внимания", "⚖️ Go + Монтаж = ✓"},
	}

	for _, ach := range achievements {
		fmt.Printf("   %-20s %s\n", ach.Name+":", ach.Value)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("🏆 ДЕНЬ ПРОЙДЕН. БАЛАНС НАЙДЕН.")
	fmt.Println(strings.Repeat("═", 50))
}
