package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🎮 ГОША: Temptation Resistance Simulator")
	fmt.Println("========================================")
	fmt.Println("Дата: 24 января 2026 года, суббота")
	fmt.Println("Миссия: Не поддаться искушению и выучить Go!")
	fmt.Println()

	// Создаем игрока
	player := NewPlayer("Гоша")
	player.DisplayStatus()

	// Главный игровой цикл
	game := NewGame()
	game.StartDay(player)

	// Симуляция дня
	events := []string{
		"Утро: Проснулся в 8:00",
		"Ритуал: Бритьё, умывание, завтрак",
		"Прогулка: Выход на мороз -15°C",
		"Работа: Заказик с Речного Вокзала → Восток Москвы (+500₽)",
		"Транспорт: Трамвай → Электричка → Метро → Автобус",
		"Возвращение: Тёплая квартира, горячий чай",
		"Вечер: Время для программирования...",
	}

	for _, event := range events {
		fmt.Printf("\n⚡ %s\n", event)
		time.Sleep(1 * time.Second)

		// Шанс возникновения искушения
		if game.CheckTemptation() {
			temptation := GenerateTemptation()
			player.HandleTemptation(temptation)
		}

		// Шанс получить мотивацию
		if game.CheckMotivation() {
			motivation := GetRandomMotivation()
			player.ReceiveMotivation(motivation)
		}
	}

	// Финальная битва с искушением
	fmt.Println("\n🔥 ФИНАЛЬНАЯ БИТВА С ИСКУШЕНИЕМ!")
	finalTemptation := Temptation{
		Name:        "CapCut Видеомонтаж",
		Power:       95,
		Description: "Установи программу для видеомонтажа! Монтируй видосы с тропическими странами!",
	}

	battleResult := player.FinalBattle(finalTemptation)

	// Итоги дня
	game.EndDay(player, battleResult)

	// Показать достижения
	player.ShowAchievements()

	// Сохранение прогресса
	SaveProgress(player)
}
