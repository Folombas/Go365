package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Советчик по среде обитания CLI ===")
	fmt.Println("Анализируем твою систему...")

	os := runtime.GOOS
	fmt.Printf("Текущая операционная система: %s\n", os)

	// Инвентарь Гоши (овощная мудрость)
	vegetables := []string{
		"картошка даёт стабильность, как Linux",
		"свёкла — для глубокой настройки, как WSL",
		"горошек символизирует быстроту команд в консоли",
		"фасоль — модульность, как pip-пакеты",
		"кукуруза — масштабирование, как облачные технологии",
	}

	// Рекомендация по установке DeepSeek CLI
	var advice string
	switch os {
	case "linux":
		advice = "Ты уже в Linux! Смело ставь DeepSeek CLI прямо сюда, используя pip или другой пакетный менеджер. " +
			"Здесь ему будет так же уютно, как Qwen. Не забудь про виртуальное окружение."
	case "windows":
		advice = "Ты в Windows. Для лучшей совместимости рекомендуется установить DeepSeek CLI внутри WSL (например, Ubuntu). " +
			"Так оба твоих ИИ-помощника будут жить в одной экосистеме, и файлы будут доступны из Windows через \\\\wsl.localhost."
	case "darwin":
		advice = "macOS — тоже отличная среда. DeepSeek CLI можно ставить через pip, но убедись, что установлены все зависимости."
	default:
		advice = "Неизвестная ОС. Попробуй установить DeepSeek CLI через контейнер (Docker) для изоляции."
	}

	fmt.Println("\n🔮 Рекомендация:")
	fmt.Println(advice)

	// Добавим овощную мудрость
	fmt.Println("\n🧺 Сегодняшний урожай Гоши подсказывает:")
	fmt.Println("   -", vegetables[rand.Intn(len(vegetables))])

	// Проверим, нужен ли DeepSeek как запасной
	fmt.Println("\n⚡ Анализ токенов Qwen:")
	tokensLeft := rand.Intn(100)
	if tokensLeft < 30 {
		fmt.Printf("   У Qwen осталось всего %d токенов! Срочно ставь DeepSeek как резерв.\n", tokensLeft)
	} else {
		fmt.Printf("   У Qwen ещё %d токенов. Можно не спешить, но установка DeepSeek не помешает для сравнения.\n", tokensLeft)
	}

	fmt.Println("\nИди и программируй, Гоша! Да прибудет с тобой сила открытых ИИ-моделей.")
}
