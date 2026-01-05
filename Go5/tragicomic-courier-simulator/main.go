package main

import (
	"fmt"
	"tragicomic-courier-simulator/courier"
	"tragicomic-courier-simulator/game"
)

func main() {
	fmt.Println("🎮 ГОША-DAILY SIMULATOR 2026: Январское уныние")
	fmt.Println("==============================================")
	fmt.Println("День 5/365. Легенда: Поиск жирненького заказа")
	fmt.Println()

	// Инициализация Гоши
	g := courier.NewGosha()

	// Прогресс дня
	fmt.Println("🌅 УТРО:")
	g.WakeUp("Honor 10x Lite")
	g.Shave("тупая бритва", false) // без лосьона
	g.TakeShower("контрастный")
	g.EatBreakfast("чай + бутерброд")

	// Поиск заказов
	fmt.Println("\n🚇 ДЕНЬ: Мониторю заказы в электричках...")
	success := g.HuntForOrders()

	if !success {
		fmt.Println("\n💔 Результат: 0 заказов")
		fmt.Println("📦 Забираю мамины заказы из ПВЗ WB")
		g.GoHome()
	}

	// Вечернее обучение
	fmt.Println("\n🌙 ВЕЧЕР: Программирование дома")
	progress := game.CodingSession(90, 5) // 90 минут, 5-й день

	// Отчет
	fmt.Println("\n📊 ОТЧЕТ ДНЯ:")
	fmt.Printf("Заработано: %d руб\n", g.Earnings())
	fmt.Printf("Прогресс Go: %d/100 XP\n", progress.XP)
	fmt.Printf("Настроение: %s\n", g.Mood())
	fmt.Printf("Энергия: %d%%\n", g.Energy())

	// Геймификация
	level := game.CalculateLevel(progress.XP)
	fmt.Printf("\n🎯 УРОВЕНЬ: %d - %s\n", level.Number, level.Title)
	fmt.Printf("До след. уровня: %d XP\n", level.XPToNext)

	if progress.NewSkill != "" {
		fmt.Printf("🎁 Новый навык: %s\n", progress.NewSkill)
	}

	// Мотивация
	if g.Earnings() < 500 {
		fmt.Println("\n💡 ВЫВОД: Заказ за 200р через всю Москву? Не, не слышал.")
		fmt.Println("   Код не обесценивается в январе. Учи Go дальше!")
	}

	// Задание на завтра
	tomorrow := game.DailyChallenge(6)
	fmt.Printf("\n📝 ЗАДАНИЕ НА ЗАВТРА (День 6): %s\n", tomorrow)
}
