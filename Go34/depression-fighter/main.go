package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type MentalState struct {
	Energy     int32
	Motivation int32
	Focus      int32
	mu         sync.Mutex
}

type DailyTask struct {
	Name        string
	Description string
	Duration    time.Duration
	Reward      int32
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎭 БИТВА С ВНУТРЕННИМ ХАОСОМ")
	fmt.Println(strings.Repeat("═", 40))

	state := &MentalState{Energy: 25, Motivation: 20, Focus: 15}

	fmt.Println("🌅 УТРО 11:00: Пробуждение с депрессией")
	fmt.Printf("   Энергия: %d | Мотивация: %d | Фокус: %d\n",
		state.Energy, state.Motivation, state.Focus)

	tasks := []DailyTask{
		{"Пробуждение", "Встать с кровати", 2, 5},
		{"Завтрак", "Гречка с сосиской", 3, 15},
		{"Работа", "Доставка Химки→Зюзино", 5, 30},
		{"Мамин звонок", "Выбор картошки - пюре или круглая", 1, 20},
		{"Душ", "Горячий душ", 2, 25},
		{"Ужин", "Картошка-пюре с жареной рыбой", 4, 30},
		{"Изучение Go", "Пакет sync", 6, 50},
	}

	var wg sync.WaitGroup
	completedTasks := int32(0)
	var totalReward int32

	fmt.Println("\n📅 ЗАПУСК ДНЯ (параллельные задачи):")

	for i, task := range tasks {
		wg.Add(1)

		go func(taskNum int, task DailyTask) {
			defer wg.Done()
			time.Sleep(time.Duration(task.Duration) * 100 * time.Millisecond)

			state.mu.Lock()
			atomic.AddInt32(&state.Energy, task.Reward/3)
			atomic.AddInt32(&state.Motivation, task.Reward/2)
			atomic.AddInt32(&state.Focus, task.Reward/2)
			atomic.AddInt32(&completedTasks, 1)
			atomic.AddInt32(&totalReward, task.Reward)

			fmt.Printf("   ✅ [%d] %s (+%d)\n", taskNum+1, task.Name, task.Reward)
			state.mu.Unlock()
		}(i, task)

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	fmt.Println("\n" + strings.Repeat("═", 40))
	fmt.Println("📊 ИТОГ ДНЯ:")
	fmt.Printf("✅ Задач выполнено: %d/%d\n", completedTasks, len(tasks))
	fmt.Printf("🎯 Состояние: Э=%d М=%d Ф=%d\n", state.Energy, state.Motivation, state.Focus)
	fmt.Printf("💰 Всего очков опыта: %d\n", totalReward)

	fmt.Println("\n💡 СИНХРОНИЗАЦИЯ УСПЕШНА!")
	fmt.Println("   Завтра — новый день, новый код.")
}
