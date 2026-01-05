package courier

import (
	"fmt"
	"math/rand"
	"time"
)

type Gosha struct {
	mood     string
	energy   int
	earnings int
	orders   int
	day      int
}

func NewGosha() *Gosha {
	return &Gosha{
		mood:   "сонный",
		energy: 85,
		day:    5,
	}
}

func (g *Gosha) WakeUp(phone string) {
	fmt.Printf("⏰ 8:00 - Будильник на %s\n", phone)
	fmt.Println("   'Ещё пять минуточек...'")
	g.mood = "не выспавшийся"
	time.Sleep(500 * time.Millisecond)
}

func (g *Gosha) Shave(razor string, hasLotion bool) {
	fmt.Printf("✂️ Бреюсь %s...\n", razor)
	if !hasLotion {
		fmt.Println("   Лосьон кончился. Бреюсь на сухую - ай!")
		g.energy -= 5
	} else {
		fmt.Println("   Освежился лосьоном!")
	}
	time.Sleep(300 * time.Millisecond)
}

func (g *Gosha) TakeShower(showerType string) {
	fmt.Printf("🚿 %s душ... Бррр-Ааа!\n", showerType)
	g.energy += 15
	g.mood = "бодрый"
	time.Sleep(400 * time.Millisecond)
}

func (g *Gosha) EatBreakfast(food string) {
	fmt.Printf("🍳 Завтрак: %s\n", food)
	g.energy += 10
	time.Sleep(300 * time.Millisecond)
}

func (g *Gosha) HuntForOrders() bool {
	fmt.Println("🔎 Смотрю цены в приложении...")
	time.Sleep(1 * time.Second)

	// Трагичная реальность января
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)

	if chance < 80 { // 80% шанс не найти заказ
		fmt.Println("   😢 Ховрино → Реутов: 600 руб (а в декабре было 2000!)")
		fmt.Println("   😭 Ховрино → Новые Черёмушки: 400 руб через всю Москву!")
		fmt.Println("   💡 Напоминание: 27 декабря возил подарок за 2000р!")
		g.mood = "философский"
		return false
	}

	// 20% шанс найти "жирненький"
	price := rand.Intn(1500) + 800
	fmt.Printf("   🎉 Нашёл заказ за %d руб!\n", price)
	g.earnings += price
	g.orders++
	g.mood = "счастливый"
	return true
}

func (g *Gosha) GoHome() {
	fmt.Println("🏠 Возвращаюсь домой к близким")
	fmt.Println("   'Хоть прогулялся... Завтра попробую снова'")
	g.energy -= 20
	time.Sleep(500 * time.Millisecond)
}

func (g *Gosha) Earnings() int {
	return g.earnings
}

func (g *Gosha) Mood() string {
	return g.mood
}

func (g *Gosha) Energy() int {
	return g.energy
}
