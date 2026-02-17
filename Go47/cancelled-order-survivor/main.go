package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	dataFile  = "survivor_stats.txt"
	startDate = "2026-01-18"
	cancelMsg = "Извините за неудобства. На адрес вы можете не ехать"
)

type Stats struct {
	LastRunDate string
	TotalDays   int
	CancelCount int
	TotalXP     int
}

func main() {
	// Настройка логирования в файл
	logFile, err := os.OpenFile("cancelled_order.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ошибка создания лог-файла:", err)
		return
	}
	defer logFile.Close()
	logger := slog.New(slog.NewTextHandler(logFile, nil))

	fmt.Println("🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂")
	fmt.Println("      ДЕНЬ 47: РАССТРОЕН ПОСЛЕ ОТМЕНЫ ЗАКАЗА     ")
	fmt.Println("🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂🚂")
	fmt.Println()
	fmt.Println("📚 Тема дня: Standard Library: I/O & File Handling: slog and regexp")
	fmt.Println()

	// Загружаем статистику
	stats := loadStats()

	// Подсчёт дней без CapCut
	start, _ := time.Parse("2006-01-02", startDate)
	today := time.Now()
	days := int(today.Sub(start).Hours() / 24)
	stats.TotalDays = days

	// Логируем события дня
	logger.Info("Начало дня", "time", today.Format(time.RFC3339))
	logger.Info("Попытка взять заказ", "status", "откликнулся, но не прозвонил")
	logger.Info("Маршрут", "from", "Ховрино", "to", "Химки", "transport", "электричка")
	logger.Info("Турникет обесточен", "action", "прошёл без закрытия, риск блокировки карты")
	logger.Info("Доехал до Молжаниново", "action", "закрыл поездку")
	logger.Info("Вернулся в Химки", "action", "обнаружил отмену заказа")
	logger.Info("Сообщение об отмене", "text", cancelMsg)
	logger.Info("Расстроился, просидел в кафе 1 час", "заказы", "ничего стоящего")
	logger.Info("Вернулся домой в 17:00", "ужин", "с мамой", "перевод", "2000 руб")
	logger.Info("Мотивация учить Go", "уровень", "возросла")

	// Парсим сообщение об отмене с помощью regexp
	re := regexp.MustCompile(`Извините за неудобства\. На адрес вы можете не ехать`)
	if re.MatchString(cancelMsg) {
		stats.CancelCount++
		fmt.Println("🔴 Обнаружено саркастичное сообщение об отмене! (совпадение по regexp)")
	}

	// Начисляем XP
	stats.TotalXP += 10 + stats.CancelCount*5

	// Сохраняем статистику
	stats.LastRunDate = today.Format("2006-01-02")
	saveStats(stats)

	// Вывод легенды
	printLegend(stats)
	printMotivation()
	printPraise(stats)
	printDeepSeekBoost()
}

func loadStats() *Stats {
	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		return &Stats{}
	} else if err != nil {
		fmt.Println("⚠️ Ошибка чтения файла:", err)
		return &Stats{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stats Stats
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			continue
		}
		stats.LastRunDate = parts[0]
		stats.TotalDays, _ = strconv.Atoi(parts[1])
		stats.CancelCount, _ = strconv.Atoi(parts[2])
		stats.TotalXP, _ = strconv.Atoi(parts[3])
	}
	return &stats
}

func saveStats(stats *Stats) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("⚠️ Ошибка сохранения файла:", err)
		return
	}
	defer file.Close()

	line := fmt.Sprintf("%s|%d|%d|%d\n",
		stats.LastRunDate,
		stats.TotalDays,
		stats.CancelCount,
		stats.TotalXP)
	file.WriteString(line)
}

func printLegend(stats *Stats) {
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║        🚆 ЛЕГЕНДА ДНЯ: ЭЛЕКТРИЧКА И ОТМЕНА 🚆           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  17 февраля 2026. Гоша встал в 11 утра. Умылся,           ║\n")
	fmt.Printf("║  покушал, попил чайку. Откликнулся на заказ, но не        ║\n")
	fmt.Printf("║  прозвонил. Оделся, вышел, дошёл до Ховрино, сел на       ║\n")
	fmt.Printf("║  электричку до Химок. Турникет был обесточен — побоялся,  ║\n")
	fmt.Printf("║  что заблокируют карту. Доехал до Молжаниново, и там уже  ║\n")
	fmt.Printf("║  закрыл поездку, вернулся в Химки. Отмена заказа!         ║\n")
	fmt.Printf("║  Только такое сообщение: «Извините за неудобства...»      ║\n")
	fmt.Printf("║  Гоша расстроился, час просидел в кафе, но ничего не      ║\n")
	fmt.Printf("║  нашёл. В 17:00 поехал домой, поужинал, перевёл маме      ║\n")
	fmt.Printf("║  2000 руб. Мотивация учить Go только выросла!             ║\n")
	fmt.Printf("║                                                           ║\n")
	fmt.Printf("║  Дней без CapCut: %d                                      ║\n", stats.TotalDays)
	fmt.Printf("║  Отмен заказов: %d                                        ║\n", stats.CancelCount)
	fmt.Printf("║  Всего XP: %d                                             ║\n", stats.TotalXP)
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printMotivation() {
	fmt.Println("💬 5 МОТИВАТОРОВ ДЛЯ ВЫЖИВШИХ:")
	motivation := []string{
		"1. Отмена заказа — не отмена твоей цели обучиться языку Go и Go-стеку.",
		"2. Турникет без электричества — как баг в коде: временно, но преодолимо.",
		"3. Каждая неудачная поездка приближает к успешному коммиту.",
		"4. Маме перевёл денег — карму поднял, XP капают.",
		"5. Час в кафе без заказов — час для мыслей о Go.",
	}
	for _, m := range motivation {
		fmt.Println(m)
	}
	fmt.Println()
}

func printPraise(stats *Stats) {
	fmt.Println("🎖️  ПОХВАЛА ДНЯ:")
	fmt.Printf("   ГОША! Ты прошёл %d дней без CapCut! Даже отмена заказа не сломила тебя.\n", stats.TotalDays)
	fmt.Printf("   XP: %d. Ты становишься сильнее с каждым днём.\n", stats.TotalXP)
	fmt.Println("   Помни: DeepSeek V4 уже в пути, и он ускорит твою прокачку!")
	fmt.Println()
}

func printDeepSeekBoost() {
	fmt.Println("🤖 DEEPSEEK V4 НА ПОДХОДЕ!")
	fmt.Println("   Со дня на день выйдет новая версия, которая:")
	fmt.Println("   • Проанализирует весь твой код за секунды")
	fmt.Println("   • Объяснит сложные концепции Go на пальцах")
	fmt.Println("   • Напишет тесты и микросервисы под твоё железо")
	fmt.Println("   • И всё это бесплатно, даже на Honor 10x Lite")
	fmt.Println("   Держись, Гоша! Прокачка выйдет на новый уровень!")
	fmt.Println()
}
