package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("🎮 TEMPTATION SIMULATOR 2.0")
	fmt.Println("══════════════════════════════════════")
	fmt.Println("Версия: 2.0 | День: 24 января 2026")
	fmt.Println("Миссия: Не поддаться искушениям и выучить Go!")
	fmt.Println()

	// Главное меню
	showMainMenu()
	
	// Создаем или загружаем игрока
	player := createOrLoadPlayer()
	if player == nil {
		return
	}
	
	// Создаем игру
	game := NewGame()
	
	// Запускаем игровой цикл
	runGameLoop(player, game)
}

// showMainMenu показывает главное меню
func showMainMenu() {
	fmt.Println("══════════════════════════════════════")
	fmt.Println("1. Новая игра")
	fmt.Println("2. Загрузить игру")
	fmt.Println("3. О игре")
	fmt.Println("0. Выход")
	fmt.Println()
	
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	
	switch choice {
	case "1":
		// Новая игра
	case "2":
		// Загрузка будет в createOrLoadPlayer
	case "3":
		showAbout()
		showMainMenu()
	case "0":
		fmt.Println("До встречи! Удачи в изучении Go! 🚀")
		os.Exit(0)
	default:
		fmt.Println("Неверный выбор. Попробуйте снова.")
		showMainMenu()
	}
}

// showAbout показывает информацию об игре
func showAbout() {
	fmt.Println("\n📖 О ИГРЕ")
	fmt.Println("══════════════════════════════════════")
	fmt.Println("Temptation Simulator - это симулятор борьбы с искушениями")
	fmt.Println("в процессе изучения языка программирования Go.")
	fmt.Println()
	fmt.Println("🎯 Цель:")
	fmt.Println("  - Сопротивляться искушениям")
	fmt.Println("  - Изучать Go и повышать навыки")
	fmt.Println("  - Развивать дерево навыков")
	fmt.Println("  - Выполнять ежедневные квесты")
	fmt.Println("  - Достичь уровня Go-Мастера!")
	fmt.Println()
	fmt.Println("🎮 Управление:")
	fmt.Println("  - Выбирайте действия из меню")
	fmt.Println("  - Следите за показателями (фокус, воля, дофамин)")
	fmt.Println("  - Улучшайте навыки за очки")
	fmt.Println("  - Сохраняйте прогресс")
	fmt.Println()
}

// createOrLoadPlayer создает или загружает игрока
func createOrLoadPlayer() *Player {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n══════════════════════════════════════")
	fmt.Println("1. Новый персонаж")
	fmt.Println("2. Загрузить сохранение")
	fmt.Print("Выбор: ")
	
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	
	if choice == "2" {
		saves := ListAvailableSaves()
		if len(saves) > 0 {
			fmt.Println("\nДоступные сохранения:")
			for i, save := range saves {
				fmt.Printf("%d. %s\n", i+1, save)
			}
			fmt.Print("Выберите сохранение: ")
			idx, _ := reader.ReadString('\n')
			idx = strings.TrimSpace(idx)
			
			// Простая конвертация в int
			for i, save := range saves {
				if fmt.Sprintf("%d", i+1) == idx {
					return LoadProgress(save)
				}
			}
		} else {
			fmt.Println("Нет доступных сохранений. Создаем нового персонажа...")
		}
	}
	
	// Создаем нового персонажа
	fmt.Print("\nВведите имя персонажа [Гоша]: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Гоша"
	}
	
	fmt.Println("\n✨ Создание персонажа:", name)
	return NewPlayer(name)
}

// runGameLoop запускает основной игровой цикл
func runGameLoop(player *Player, game *Game) {
	reader := bufio.NewReader(os.Stdin)
	
	// Начинаем день
	game.StartDay(player)
	player.DisplayStatus()
	
	// Игровой цикл по часам (16 часов = 1 день)
	hoursPlayed := 0
	totalHours := 16
	
	for hoursPlayed < totalHours {
		fmt.Printf("\n⏰ ЧАС %d/%d | %02d:00\n", hoursPlayed+1, totalHours, 8+hoursPlayed)
		fmt.Println("══════════════════════════════════════")
		
		// Показываем доступные действия
		showActionsMenu()
		
		fmt.Print("Выбор: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			// Изучение Go
			fmt.Print("Сколько минут учить Go? [30]: ")
			minutes, _ := reader.ReadString('\n')
			minutes = strings.TrimSpace(minutes)
			if minutes == "" {
				minutes = "30"
			}
			
			var mins int
			fmt.Sscanf(minutes, "%d", &mins)
			if mins <= 0 {
				mins = 30
			}
			
			game.StudyGoSession(player, mins)
			hoursPlayed += (mins / 60) + 1
			
		case "2":
			// Проверить квесты
			game.ShowQuestsMenu(player)
			
		case "3":
			// Дерево навыков
			game.UpgradeSkillMenu(player)
			
		case "4":
			// Отдых
			fmt.Print("Сколько минут отдыхать? [15]: ")
			minutes, _ := reader.ReadString('\n')
			minutes = strings.TrimSpace(minutes)
			if minutes == "" {
				minutes = "15"
			}
			
			var mins int
			fmt.Sscanf(minutes, "%d", &mins)
			if mins <= 0 {
				mins = 15
			}
			
			game.RestSession(player, mins)
			hoursPlayed += (mins / 60)
			
		case "5":
			// Статистика
			player.DisplayStatistics()
			
		case "6":
			// Сохранение
			game.SaveGameMenu(player)
			
		case "7":
			// Завершить день
			fmt.Println("\nЗавершить день и перейти к финальной битве? (y/n)")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSpace(confirm)
			
			if confirm == "y" || confirm == "Y" {
				hoursPlayed = totalHours
			}
			
		case "0":
			// Сохранить и выйти
			fmt.Println("Сохранение и выход...")
			SaveProgress(player, game)
			fmt.Println("До встречи! Удачи в изучении Go! 🚀")
			return
			
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
		
		// Случайные события
		handleRandomEvents(player, game)
		
		// Проверяем выполнение квеста на изучение Go
		if player.Quests.Quests[0].Progress >= 30 {
			player.Quests.Quests[0].Completed = true
			fmt.Println("\n✅ Квест выполнен: 30 минут Go!")
		}
	}
	
	// Финальная битва
	fmt.Println("\n🌙 ВЕЧЕР. Наступает время финальной битвы...")
	time.Sleep(2 * time.Second)
	
	bossTemptation := GenerateBossTemptation()
	battleResult := player.FinalBattle(bossTemptation)
	
	if battleResult {
		game.BossDefeated = true
	}
	
	// Завершаем день
	game.EndDay(player, battleResult)
	
	// Сохраняем прогресс
	fmt.Println("\n💾 Автосохранение после завершения дня...")
	SaveProgress(player, game)
	
	fmt.Println("\n🎮 День завершен! Спасибо за игру!")
	fmt.Println("🚀 Продолжай учить Go и достигнешь цели!")
}

// showActionsMenu показывает меню действий
func showActionsMenu() {
	fmt.Println("\n📋 ДЕЙСТВИЯ:")
	fmt.Println("  1. 📚 Учить Go")
	fmt.Println("  2. 📋 Проверить квесты")
	fmt.Println("  3. 🌳 Дерево навыков")
	fmt.Println("  4. 💤 Отдохнуть")
	fmt.Println("  5. 📊 Статистика")
	fmt.Println("  6. 💾 Сохранение")
	fmt.Println("  7. 🌙 Завершить день")
	fmt.Println("  0. 🚪 Сохранить и выйти")
}

// handleRandomEvents обрабатывает случайные события
func handleRandomEvents(player *Player, game *Game) {
	game.TotalTemptations++
	
	// Проверка на искушение
	if game.CheckTemptation() {
		temptation := GenerateTemptation()
		player.HandleTemptation(temptation)
		
		// Обновляем счетчик игры
		if player.Willpower > temptation.Power {
			game.HandleTemptationResisted(player, temptation)
		} else {
			game.HandleTemptationFailed(player, temptation)
		}
	}
	
	// Проверка на мотивацию
	if game.CheckMotivation() {
		motivation := GetRandomMotivation()
		player.ReceiveMotivation(motivation)
		player.AddExperience(motivation.XPBonus)
	}
	
	// Проверка на босса
	if game.CheckBossEncounter() {
		fmt.Println("\n⚠️  ВНИМАНИЕ! Появляется ОСОБО ОПАСНОЕ ИСКУШЕНИЕ!")
		time.Sleep(1 * time.Second)
		
		boss := GenerateBossTemptation()
		fmt.Printf("\n👹 БОСС: %s\n", boss.Name)
		fmt.Printf("Сила: %d%% | Описание: %s\n", boss.Power, boss.Description)
		
		// Автоматическая битва с боссом
		successChance := (player.Willpower * 2 + player.Focus) / 3
		fmt.Printf("Шанс на победу: %d%%\n", successChance)
		
		if randIntn(100) < successChance {
			fmt.Println("\n🎉 ПОБЕДА НАД БОССОМ!")
			player.AddExperience(boss.XPLoss * 2)
			player.Achievements = append(player.Achievements, fmt.Sprintf("Победил %s", boss.Name))
			game.BossDefeated = true
		} else {
			fmt.Println("\n💔 ПОРАЖЕНИЕ ОТ БОССА...")
			player.Focus = 20
			player.Willpower = 30
		}
	}
}

// Вспомогательная функция для генерации случайных чисел
func randIntn(n int) int {
	return int(time.Now().UnixNano() % int64(n))
}
