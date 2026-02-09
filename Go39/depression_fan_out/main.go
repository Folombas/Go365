package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// EnergySource - источник энергии (остатки воли Гоши)
type EnergySource struct {
	TotalEnergy int // Вся энергия, которая есть
	mu          sync.Mutex
}

// TransportTask - задача поездки на транспорте
type TransportTask struct {
	Route      string
	Duration   int    // в минутах
	EnergyCost int    // сколько энергии отнимает
	Reward     string // что дает
}

// Worker - воркер, который будет выполнять задачи
type Worker struct {
	ID      int
	Energy  int // персональный запас энергии
	Results chan<- string
}

func main() {
	fmt.Println("🚇 ДЕНЬ 39: FAN-OUT ПОСЛЕДНЕЙ ЭНЕРГИИ ПО ГОРОДУ")
	fmt.Println("================================================")
	fmt.Println("Состояние: Депрессивный понедельник, 9 февраля 2026")
	fmt.Println("Ресурсы: Последние крохи воли перед покупкой проездного")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	// Создаем источник энергии (последние силы Гоши)
	source := &EnergySource{
		TotalEnergy: 20, // Очень мало энергии
	}

	// Создаем задачи на завтра (после покупки проездного)
	tasks := []TransportTask{
		{"Метро: ветка БКЛ, полный круг", 45, 5, "Осмотр города из-под земли"},
		{"Автобус: от вокзала до парка", 30, 4, "Наблюдение за природой"},
		{"Трамвай: исторический маршрут", 25, 3, "Архитектурные впечатления"},
		{"Электричка: в пригород и обратно", 90, 8, "Смена обстановки"},
		{"Пешая прогулка: набережная", 40, 2, "Свежий воздух и движение"},
		{"Маршрутка: случайный номер", 20, 3, "Элемент неожиданности"},
	}

	// Создаем канал для задач
	taskChan := make(chan TransportTask, len(tasks))

	// Создаем канал для результатов
	resultChan := make(chan string, len(tasks))

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Fan-out: создаем пул воркеров (силы воли)
	workerCount := 3 // Сколько сил воли осталось у Гоши
	fmt.Printf("🔄 Запускаем FAN-OUT: %d воркера получают энергию из источника\n\n", workerCount)

	// Создаем воркеров и запускаем их
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		worker := &Worker{
			ID:      i,
			Energy:  0,
			Results: resultChan,
		}
		go worker.processTasks(&wg, taskChan, source)
	}

	// Отправляем задачи в канал
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan) // Закрываем канал, чтобы воркеры знали, когда остановиться

	// Запускаем сбор результатов
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Выводим результаты
	fmt.Println("📊 РЕЗУЛЬТАТЫ РАСПРЕДЕЛЕНИЯ ЭНЕРГИИ (FAN-OUT):")
	fmt.Println(strings.Repeat("-", 50))

	completedTasks := 0
	for result := range resultChan {
		fmt.Println(result)
		completedTasks++
	}

	// Итоги
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("✅ ВЫПОЛНЕНО ЗАДАЧ: %d из %d\n", completedTasks, len(tasks))
	fmt.Printf("⚡ ОСТАТОК ЭНЕРГИИ В ИСТОЧНИКЕ: %d\n", source.TotalEnergy)

	// Геймификация: уровень достижений
	achievement := ""
	switch {
	case completedTasks >= 5:
		achievement = "🏆 МАСТЕР FAN-OUT: Депрессия не помеха продуктивности!"
	case completedTasks >= 3:
		achievement = "🥈 СЕРЕБРЯНЫЙ УРОВЕНЬ: Хорошее распределение ресурсов"
	default:
		achievement = "🥉 БРОНЗА: Сегодня тяжело, но ты справился"
	}

	fmt.Println("\n" + achievement)

	// Мотивация на завтра
	fmt.Println("\n🎯 ЗАВТРАШНИЙ ПЛАН (после покупки проездного):")
	if source.TotalEnergy > 5 {
		fmt.Println("   Можно выполнить все запланированные поездки!")
	} else {
		fmt.Println("   Выбери 2-3 самых важных маршрута")
	}

	fmt.Println("\n" + getGoMotivation())
	fmt.Println("\n🌟 ГОША! Ты распределил последнюю энергию через Fan-out!")
	fmt.Println("   Завтра купишь проездной и реализуешь этот план на практике.")
}

// Метод для обработки задач воркером
func (w *Worker) processTasks(wg *sync.WaitGroup, tasks <-chan TransportTask, source *EnergySource) {
	defer wg.Done()

	for task := range tasks {
		// Пытаемся получить энергию из источника
		source.mu.Lock()
		if source.TotalEnergy >= task.EnergyCost {
			source.TotalEnergy -= task.EnergyCost
			w.Energy += task.EnergyCost
			source.mu.Unlock()

			// Имитируем выполнение задачи
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Отправляем результат
			w.Results <- fmt.Sprintf("👷 Воркер #%d выполнил: %s (-%d энергии, +%s)",
				w.ID, task.Route, task.EnergyCost, task.Reward)
		} else {
			source.mu.Unlock()
			w.Results <- fmt.Sprintf("⛔ Воркер #%d НЕ СМОГ: %s (не хватило энергии)",
				w.ID, task.Route)
		}
	}
}

func getGoMotivation() string {
	phrases := []string{
		"🔥 10 FAN-OUT МОТИВАЦИЙ ДЛЯ GO-РАЗРАБОТЧИКА:",
		"1. 🚀 FAN-OUT В GO = РАСПРЕДЕЛЕНИЕ НАГРУЗКИ МЕЖДУ ЯДРАМИ ПРОЦЕССОРА",
		"2. 💡 КАЖДАЯ ГОРУТИНА - ОТДЕЛЬНЫЙ ПОТОК ИСПОЛНЕНИЯ, КАК КАЖДАЯ ПОЕЗДКА - ШАНС ВЫЙТИ ИЗ ДЕПРЕССИИ",
		"3. 🎯 GO РУТИНЫ ЛЕГКОВЕСНЫ - КАК ЛЕГКОВЕСНЫ ТВОИ МИКРОДЕЙСТВИЯ ПРОТИВ АПАТИИ",
		"4. ⚡ FAN-OUT ЧЕРЕЗ КАНАЛЫ = БЕЗОПАСНОЕ РАСПРЕДЕЛЕНИЕ ЗАДАЧ МЕЖДУ ВОРКЕРАМИ",
		"5. 🛡️ SYNC.MUTEX ЗАЩИЩАЕТ ОБЩИЙ РЕСУРС (ЭНЕРГИЮ) ОТ КОНКУРЕНТНЫХ ПОТОКОВ",
		"6. 🔄 WAITGROUP ГАРАНТИРУЕТ - ВСЕ ВОРКЕРЫ ЗАВЕРШАТСЯ ПРЕЖДЕ ЧЕМ ПРОГРАММА",
		"7. 📈 FAN-OUT МАСШТАБИРУЕТСЯ - ДОБАВИЛ ЕЩЕ ВОРКЕРОВ = УВЕЛИЧИЛ ПРОПУСКНУЮ СПОСОБНОСТЬ",
		"8. 🎮 FAN-OUT КАК ГЕЙМПЛЕЙ: РАСПРЕДЕЛЯЙ РЕСУРСЫ МЕЖДУ ЮНИТАМИ ДЛЯ МАКСИМАЛЬНОЙ ЭФФЕКТИВНОСТИ",
		"9. 💰 FAN-OUT ЗНАНИЙ GO = ВЫСОКАЯ ЗАРПЛАТА РАЗРАБОТЧИКА В 2026",
		"10. 🏆 GO КАНАЛЫ + FAN-OUT = ПРОФЕССИОНАЛЬНЫЙ УРОВЕНЬ КОНКУРЕНТНОГО ПРОГРАММИРОВАНИЯ",
	}

	result := ""
	for _, phrase := range phrases {
		result += phrase + "\n"
	}

	result += "\n🎮 ГЕЙМИФИКАЦИЯ ДНЯ:\n"
	result += "   Уровень: 'Fan-out распределитель'\n"
	result += "   Персонаж: Go-разработчик в состоянии депрессии\n"
	result += "   Миссия: Распределить 20 единиц энергии на 6 транспортных маршрутов\n"
	result += "   Бонус: Каждая выполненная задача = +1 к навыку 'Concurrency Patterns'\n"
	result += "   Следующий уровень: Fan-in результатов завтрашних поездок"

	return result
}
