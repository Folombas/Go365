package main

import (
	"fmt"
	"os"
	"time"
)

// PastGlory хранит славу прошлого
type PastGlory struct {
	YearsAgo    int
	Venue       string
	Perks       []string
	FanLove     bool
	CurrentGigs bool
}

// PresentReality отражает текущую реальность Гоши
type PresentReality struct {
	WakeUpTime     string
	Activities     []string
	BudgetStatus   string
	LearningHours  int
	FriendsGigs    bool
	InvitedToGigs  bool
}

// FutureDream - мечты о будущем в IT
type FutureDream struct {
	TargetIndustry string
	DailyCoding    int
	Motivation     string
	ExpectedSalary string
}

// repeatString - вспомогательная функция для повторения строки
func repeatString(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func main() {
	fmt.Println("🎤 =========================================")
	fmt.Println("       ИСТОРИЯ ТРАНСФОРМАЦИИ ГОШИ")
	fmt.Println("      от рэп-фрика к IT-фрилансеру")
	fmt.Println("========================================== 🖥️\n")

	// Прошлое Гоши
	past := PastGlory{
		YearsAgo:    12,
		Venue:       "клубы на фрик-шоу",
		Perks:       []string{"гримёрка с закусками", "лимузины", "гастроли", "клипы"},
		FanLove:     true,
		CurrentGigs: false,
	}

	// Настоящее Гоши
	present := PresentReality{
		WakeUpTime:   "10:00",
		Activities:   []string{"умылся", "попил чайку", "прогулялся", "оплатил кварплату"},
		BudgetStatus: "последние деньги",
		LearningHours: 4,
		FriendsGigs:   true,
		InvitedToGigs: false,
	}

	// Будущее Гоши
	future := FutureDream{
		TargetIndustry: "IT",
		DailyCoding:    6,
		Motivation:     "самая крепкая надежда",
		ExpectedSalary: "достойная",
	}

	// Выводим историю
	fmt.Println("📜 ПРОШЛОЕ:")
	fmt.Printf("Лет назад: %d\n", past.YearsAgo)
	fmt.Printf("Сцена: %s\n", past.Venue)
	fmt.Println("Плюшки:")
	for i, perk := range past.Perks {
		fmt.Printf("  %d. %s\n", i+1, perk)
	}

	fmt.Println("\n⚡ НАСТОЯЩЕЕ:")
	fmt.Printf("Проснулся в: %s\n", present.WakeUpTime)
	fmt.Println("Действия:")
	for i, activity := range present.Activities {
		fmt.Printf("  %d. %s\n", i+1, activity)
	}
	fmt.Printf("Бюджет: %s\n", present.BudgetStatus)
	fmt.Printf("Учит Go: %d часа в день\n", present.LearningHours)

	// Трагикомичный момент
	fmt.Println("\n💔 ТРАГИКОМИЧНЫЙ ФАКТ:")
	if present.FriendsGigs && !present.InvitedToGigs {
		fmt.Println("Друзья выступают в клубах в эти выходные...")
		fmt.Println("Но Гошу опять НЕ позвали")
		time.Sleep(2 * time.Second)
		fmt.Println("Гоша всё понял. 🤔")
		time.Sleep(1 * time.Second)
		fmt.Println("Он не настоящий фрик. 😅")
		time.Sleep(1 * time.Second)
		fmt.Println("Значит, он — фрилансер-программист! 💻")
	}

	fmt.Println("\n🚀 БУДУЩЕЕ:")
	fmt.Printf("Цель: влететь в %s\n", future.TargetIndustry)
	fmt.Printf("Кодит: %d часов в день\n", future.DailyCoding)
	fmt.Printf("Мотивация: %s\n", future.Motivation)
	fmt.Printf("Ожидания: зарплата %s\n", future.ExpectedSalary)

	// Читаем прощальный рэп
	fmt.Println("\n🎶 ПРОЩАЛЬНЫЙ РЭП ГОШИ:")
	rap, err := os.ReadFile("lyrics/goodbye_rap.txt")
	if err != nil {
		fmt.Println("Рэп-текст потерялся, как и приглашения на выступления...")
	} else {
		fmt.Println(string(rap))
	}

	// Геймификация: выбор пути
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ: ВЫБОР ПУТИ")
	fmt.Println("Нажмите 1 для продолжения карьеры рэпера")
	fmt.Println("Нажмите 2 для вкатывания в IT")

	var choice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("❌ ОШИБКА: Путь 'рэп-фрик' больше не доступен")
		fmt.Println("   Причина: отсутствие приглашений на выступления")
	case 2:
		fmt.Println("✅ ПУТЬ 'IT-ФРИЛАНСЕР' АКТИВИРОВАН!")
		fmt.Println("   Загрузка знаний Go...")
		time.Sleep(1 * time.Second)
		fmt.Println("   Установка VS Code...")
		time.Sleep(1 * time.Second)
		fmt.Println("   Подключение к GitHub...")
		time.Sleep(1 * time.Second)
		fmt.Println("   🎉 ГОША ТЕПЕРЬ IT-СПЕЦИАЛИСТ!")
	default:
		fmt.Println("⚠  Гоша слишком долго выбирал...")
		fmt.Println("   Система автоматически выбрала IT!")
	}

	// Финальное сообщение
	fmt.Println("\n" + repeatString("=", 50))
	fmt.Println("МОРАЛЬ ИСТОРИИ:")
	fmt.Println("Когда закрывается одна дверь (клубная),")
	fmt.Println("открывается другая (терминал Go).")
	fmt.Println("Гоша выбрал код.")
	fmt.Println(repeatString("=", 50))
}
