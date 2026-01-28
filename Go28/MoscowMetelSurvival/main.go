package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
	"strings"
)

// Геймифицированный персонаж Гоши
type GoshCharacter struct {
	Name          string
	Energy        int      // 0-100
	Motivation    int      // 0-100
	Money         int      // рубли
	XP            int      // опыт программирования
	Location      string
	Skills        []string
	CompletedTasks int
	CancelledTasks int
	LastMeal      string
}

func main() {
	fmt.Println("=== DAY 28: МОСКОВСКАЯ МЕТЕЛЬ ===")
	fmt.Println("Эпичная сага о Гоше с зимней депрессией")
	fmt.Println("Игра началась:", time.Now().Format("02.01.2006 15:04"))
	fmt.Println(strings.Repeat("=", 50) + "\n")

	// Инициализация персонажа
	gosh := GoshCharacter{
		Name:       "Гоша-Кодер",
		Energy:     35,  // спал мало
		Motivation: 20,  // зимняя депрессия
		Money:      2500,
		XP:         2700, // 27 дней уже учил Go
		Location:   "Кровать",
		Skills:     []string{"Golang Basics", "Functions", "Structs"},
		LastMeal:   "Ничего",
	}

	// События дня (параллельные горутины)
	var wg sync.WaitGroup
	eventChan := make(chan string, 10)

	// 1. Пробуждение (горутина)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		gosh.Location = "Кухня"
		gosh.Energy += 10
		gosh.LastMeal = "Завтрак с мамой"
		gosh.Motivation += 5
		eventChan <- "🌅 10:30 - Пробуждение после 7.5 часов сна. Энергия +10"
	}()

	// 2. Первый заказ (успешный!)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		gosh.Location = "МЦД-3: Химки → Авиамоторная"
		gosh.Money += 950
		gosh.CompletedTasks++
		gosh.XP += 50
		eventChan <- "🚚 УСПЕШНЫЙ ЗАКАЗ: Подвез подарок! +950₽ | +50 XP"
	}()

	// 3. Второй заказ (ПРОВАЛ!)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		gosh.Location = "Перово (ожидание)"
		gosh.CancelledTasks++
		gosh.Motivation -= 15
		eventChan <- "💔 КРИТИЧЕСКИЙ ПРОВАЛ: Заказ отменили! Мотивация -15"
	}()

	// 4. Возвращение домой
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(4 * time.Second)
		gosh.Location = "Дом, душевая"
		gosh.Energy += 25
		gosh.Motivation += 10
		eventChan <- "🚿 ТЕПЛЫЙ ДУШ: Смываю разочарование. Энергия +25, Мотивация +10"
	}()

	// 5. Программирование (финальный босс!)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		gosh.Location = "За компьютером"

		// Шанс успеха зависит от мотивации
		successChance := gosh.Motivation + rand.Intn(30)
		if successChance > 50 {
			gosh.XP += 100
			gosh.Skills = append(gosh.Skills, "Goroutines")
			gosh.Motivation += 30
			eventChan <- "💻 МОЩНЫЙ ПРОГРЕСС: Освоил горутины! +100 XP | +30 Мотивации"
		} else {
			gosh.Motivation -= 5
			eventChan <- "💻 СЛОЖНЫЙ ДЕНЬ: Код не компилируется... Но я продолжаю!"
		}
	}()

	// Запускаем все события
	go func() {
		wg.Wait()
		close(eventChan)
	}()

	// Выводим события в реальном времени
	fmt.Println("📜 ХРОНИКА СОБЫТИЙ ДНЯ:")
	for event := range eventChan {
		fmt.Println(event)
		time.Sleep(500 * time.Millisecond)
	}

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎮 ИТОГИ ДНЯ 28:")
	fmt.Printf("Персонаж: %s\n", gosh.Name)
	fmt.Printf("Уровень энергии: %d/100\n", gosh.Energy)
	fmt.Printf("Уровень мотивации: %d/100\n", gosh.Motivation)
	fmt.Printf("Баланс: %d₽ (заработано: 950₽)\n", gosh.Money)
	fmt.Printf("Опыт программирования: %d XP\n", gosh.XP)
	fmt.Printf("Выполнено задач: %d | Отменено: %d\n", gosh.CompletedTasks, gosh.CancelledTasks)
	fmt.Printf("Новые навыки: %v\n", gosh.Skills)
	fmt.Println("Текущая локация:", gosh.Location)

	fmt.Println("\n" + "🔥 БОСС ДНЯ ПОВЕРЖЕН: 'Зимняя Депрессия'")
	fmt.Println("Награда: +1 к дисциплине, +100 к вере в себя")
	fmt.Println(strings.Repeat("=", 50))

	// Финальная мотивация
	fmt.Println("\n💪 ЗАВТРА БУДЕТ ДЕНЬ 29!")
	fmt.Println("Каждый день с кодом - шаг к работе мечты.")
	fmt.Println("Сейчас 21:00 - идеальное время для ещё одного урока Go!")
}
