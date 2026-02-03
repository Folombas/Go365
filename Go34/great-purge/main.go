package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Resource представляет ограниченный ресурс сервера
type Resource struct {
	Name     string
	UsedMB   int
	TotalMB  int
	Critical bool
}

// N8nCube представляет "кубик"-узел no-code платформы
type N8nCube struct {
	Name         string
	Description  string
	Dependencies []string
}

func main() {
	fmt.Println("🎮 СИМУЛЯТОР: ВЕЛИКАЯ ЧИСТКА СЕРВЕРА ГОШИ")
	fmt.Println(strings.Repeat("=", 50))

	// Конфигурация сервера за 7 рублей/день
	serverSpecs := map[string]int{
		"RAM_MB":       1024,
		"DISK_GB":      10,
		"COST_PER_DAY": 7,
	}

	fmt.Printf("💻 Конфигурация сервера:\n")
	fmt.Printf("   Память: %d MB\n", serverSpecs["RAM_MB"])
	fmt.Printf("   Диск: %d GB\n", serverSpecs["DISK_GB"])
	fmt.Printf("   Стоимость: %d руб/день\n\n", serverSpecs["COST_PER_DAY"])

	// Ресурсы, которые "съедает" n8n
	resources := []Resource{
		{"n8n Процесс", 450, serverSpecs["RAM_MB"], true},
		{"n8n Кэш", 120, serverSpecs["RAM_MB"], false},
		{"PostgreSQL", 220, serverSpecs["RAM_MB"], true},
		{"Очереди Redis", 85, serverSpecs["RAM_MB"], false},
	}

	// "Кубики" n8n, от которых зависит Гоша
	cubes := []N8nCube{
		{"HTTP Request", "Запросы к API", []string{"axios", "proxy"}},
		{"Telegram Trigger", "Запуск по сообщению", []string{"telegraf", "webhook"}},
		{"Condition", "Проверка условий", []string{"logic"}},
		{"Google Sheets", "Работа с таблицами", []string{"oauth", "api-client"}},
	}

	fmt.Println("🔍 АНАЛИЗ ТЕКУЩЕЙ СИТУАЦИИ:")
	fmt.Println(strings.Repeat("-", 30))

	// Показываем проблему
	totalUsed := 0
	for _, res := range resources {
		usagePercent := (res.UsedMB * 100) / res.TotalMB
		fmt.Printf("⚠️  %s: %d/%d MB (%d%%)\n",
			res.Name, res.UsedMB, res.TotalMB, usagePercent)
		totalUsed += res.UsedMB
	}

	fmt.Printf("\n📊 ИТОГО используется: %d/%d MB (%d%%)\n",
		totalUsed, serverSpecs["RAM_MB"], (totalUsed*100)/serverSpecs["RAM_MB"])

	if totalUsed > serverSpecs["RAM_MB"] {
		fmt.Println("💥 КРИТИЧЕСКОЕ ПЕРЕПОЛНЕНИЕ! Сервер работает на свопе.")
	}

	// Драматическая пауза
	fmt.Println("\n🤔 Гоша размышляет...")
	time.Sleep(2 * time.Second)

	// Процесс "удаления n8n"
	fmt.Println("\n🔥 НАЧИНАЕТСЯ ВЕЛИКАЯ ЧИСТКА...")
	fmt.Println(strings.Repeat("-", 30))

	for i, cube := range cubes {
		fmt.Printf("[%d/4] Удаляем кубик: %s\n", i+1, cube.Name)
		fmt.Printf("    Зависимости: %v\n", strings.Join(cube.Dependencies, ", "))

		// Имитация освобождения ресурсов
		freedMB := rand.Intn(100) + 50
		fmt.Printf("    🔓 Освобождается ~%d MB памяти\n", freedMB)

		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}

	// Установка чистого Go-окружения
	fmt.Println("\n🔄 УСТАНАВЛИВАЕМ ЧИСТУЮ СРЕДУ GO...")
	components := []string{
		"Go 1.25.6",
		"Git для контроля версий",
		"PostgreSQL (чистая версия)",
		"Redis (минимальная конфигурация)",
	}

	for _, comp := range components {
		fmt.Printf("   ✅ %s установлен\n", comp)
		time.Sleep(300 * time.Millisecond)
	}

	// Финальный результат
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🎉 ВЕЛИКАЯ ЧИСТКА ЗАВЕРШЕНА!")
	fmt.Println(strings.Repeat("=", 50))

	newResources := []Resource{
		{"Go Runtime", 45, serverSpecs["RAM_MB"], false},
		{"PostgreSQL", 110, serverSpecs["RAM_MB"], true},
		{"Telegram Bot", 28, serverSpecs["RAM_MB"], false},
		{"API Cache", 15, serverSpecs["RAM_MB"], false},
	}

	totalFreed := 0
	fmt.Println("\n📈 НОВАЯ КОНФИГУРАЦИЯ:")
	for _, res := range newResources {
		fmt.Printf("   🟢 %s: %d MB\n", res.Name, res.UsedMB)
		totalFreed += res.UsedMB
	}

	freeMemory := serverSpecs["RAM_MB"] - totalFreed
	fmt.Printf("\n💎 СВОБОДНО ПАМЯТИ: %d MB (до чистки: %d MB)\n",
		freeMemory, serverSpecs["RAM_MB"]-totalUsed)
	fmt.Printf("🚀 ПРОИЗВОДИТЕЛЬНОСТЬ ВЫРОСЛА НА: ~%d%%\n",
		((totalUsed-totalFreed)*100)/serverSpecs["RAM_MB"])

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🏆 ГОША СДЕЛАЛ ЭТО! ТЕПЕРЬ ТОЛЬКО GO И КОД!")
	fmt.Println(strings.Repeat("=", 50))

	// Мотивационное сообщение
	fmt.Println("\n📢 НАПОМИНАНИЕ ДНЯ:")
	messages := []string{
		"No-code даёт быстрое решение, но не глубину понимания.",
		"Go даёт контроль, производительность и карьеру.",
		"Каждая строка кода на Go — это инвестиция в себя.",
		"Сервер за 7 рублей теперь — твоя крепость, а не ограничение.",
	}

	for _, msg := range messages {
		fmt.Printf("   💭 %s\n", msg)
		time.Sleep(800 * time.Millisecond)
	}
}
