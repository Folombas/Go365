package main

import (
	"fmt"
	"strings"
	"time"
)

// PlayerStat представляет характеристику игрока
type PlayerStat struct {
	Name  string
	Value int
	Max   int
}

// GameCharacter представляет игрового персонажа (Гошу)
type GameCharacter struct {
	Name          string
	Level         int
	XP            int
	XPToNextLevel int
	Stats         map[string]PlayerStat
	Inventory     map[string]int
	Quests        []string
	Completed     []string
}

// NewGosha создает нового персонажа Гошу
func NewGosha() GameCharacter {
	return GameCharacter{
		Name:          "Гоша Golang Гофер",
		Level:         1,
		XP:            0,
		XPToNextLevel: 1000,
		Stats: map[string]PlayerStat{
			"go_skill":   {"Навык Go", 20, 100},
			"discipline": {"Дисциплина", 40, 100},
			"focus":      {"Фокус", 10, 100},
			"mood":       {"Настроение", 30, 100},
			"depression": {"Депрессия", 70, 100},
		},
		Inventory: map[string]int{
			"money":      0,
			"energy":     60,
			"motivation": 50,
		},
		Quests: []string{
			"Восстание из постели",
			"Гигиена воина",
			"Курьерский рейс",
			"Битва с коммуналкой",
			"Священный час Go",
		},
	}
}

// CompleteQuest выполняет квест и обновляет характеристики
func (gc *GameCharacter) CompleteQuest(questName string, xpReward int, statChanges map[string]int) {
	fmt.Printf("\n🎯 Квест: %s\n", questName)

	// Убираем из списка активных квестов
	for i, q := range gc.Quests {
		if q == questName {
			gc.Quests = append(gc.Quests[:i], gc.Quests[i+1:]...)
			break
		}
	}

	// Добавляем в выполненные
	gc.Completed = append(gc.Completed, questName)

	// Начисляем XP
	gc.XP += xpReward
	fmt.Printf("   Получено: %d XP\n", xpReward)

	// Применяем изменения характеристик
	for statKey, change := range statChanges {
		if stat, exists := gc.Stats[statKey]; exists {
			stat.Value += change
			if stat.Value < 0 {
				stat.Value = 0
			}
			if stat.Value > stat.Max {
				stat.Value = stat.Max
			}
			gc.Stats[statKey] = stat
			if change > 0 {
				fmt.Printf("   ↑ %s: +%d\n", stat.Name, change)
			} else {
				fmt.Printf("   ↓ %s: %d\n", stat.Name, change)
			}
		}
	}

	// Проверяем уровень
	if gc.XP >= gc.XPToNextLevel {
		gc.LevelUp()
	}
}

// LevelUp повышает уровень персонажа
func (gc *GameCharacter) LevelUp() {
	gc.Level++
	gc.XP -= gc.XPToNextLevel
	gc.XPToNextLevel = int(float64(gc.XPToNextLevel) * 1.5)

	fmt.Printf("\n⭐ УРОВЕНЬ ПОВЫШЕН! Теперь уровень %d ⭐\n", gc.Level)

	// Бонусы за уровень
	switch gc.Level {
	case 2:
		fmt.Println("   🎁 Получен: базовый синтаксис Go")
		gc.Stats["go_skill"] = PlayerStat{"Навык Go", 30, 100}
	case 3:
		fmt.Println("   🎁 Получен: понимание горутин")
		gc.Stats["focus"] = PlayerStat{"Фокус", 25, 100}
	}
}

// DisplayStats показывает текущие характеристики
func (gc *GameCharacter) DisplayStats() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("👤 %s | Уровень: %d | XP: %d/%d\n",
		gc.Name, gc.Level, gc.XP, gc.XPToNextLevel)
	fmt.Println(strings.Repeat("-", 50))

	// Характеристики
	fmt.Println("📊 ХАРАКТЕРИСТИКИ:")

	// Определяем порядок вывода характеристик
	statOrder := []string{"go_skill", "discipline", "focus", "mood", "depression"}

	for _, statKey := range statOrder {
		stat := gc.Stats[statKey]
		barLength := 20
		filled := int(float64(stat.Value) / float64(stat.Max) * float64(barLength))
		bar := strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)

		fmt.Printf("   %s: %s %d/%d\n", stat.Name, bar, stat.Value, stat.Max)
	}

	// Инвентарь
	fmt.Println("\n🎒 ИНВЕНТАРЬ:")
	for item, amount := range gc.Inventory {
		fmt.Printf("   %s: %d\n", strings.Title(item), amount)
	}

	// Квесты
	fmt.Println("\n📋 АКТИВНЫЕ КВЕСТЫ:")
	for _, quest := range gc.Quests {
		fmt.Printf("   • %s\n", quest)
	}

	fmt.Println("\n✅ ВЫПОЛНЕННЫЕ КВЕСТЫ:")
	for _, quest := range gc.Completed {
		fmt.Printf("   ✓ %s\n", quest)
	}

	fmt.Println(strings.Repeat("=", 50))
}

// DistractionCheck проверяет отвлекающие факторы
func (gc *GameCharacter) DistractionCheck(distraction string) bool {
	distractions := map[string]string{
		"n8n":        "Ночные боты подождут! Сначала стань Go-разработчиком!",
		"video_game": "Игры будут всегда. Вакансия Go-разработчика — нет!",
		"bar":        "Бары никуда не денутся. Твоя мотивация — может!",
		"series":     "Сериалы подождут. Интерфейсы в Go — нет!",
		"video_edit": "Смонтируешь отпуск, когда устроишься в офис!",
	}

	if warning, exists := distractions[distraction]; exists {
		fmt.Printf("\n🚫 ОПАСНОСТЬ: %s\n", warning)

		// Потеря фокуса
		if focusStat, exists := gc.Stats["focus"]; exists {
			focusStat.Value -= 5
			if focusStat.Value < 0 {
				focusStat.Value = 0
			}
			gc.Stats["focus"] = focusStat
		}
		return true
	}
	return false
}

func main() {
	fmt.Println("🎮 GoRising: Из курьера в Go-разработчики")
	fmt.Println("========================================")

	// Создаем персонажа
	gosha := NewGosha()
	gosha.DisplayStats()

	// Симуляция дня
	fmt.Println("\n🌅 УТРО 2 ФЕВРАЛЯ 2026:")

	// 1. Восстание из постели
	time.Sleep(1 * time.Second)
	gosha.CompleteQuest("Восстание из постели", 50, map[string]int{
		"discipline": 10,
		"mood":       5,
		"depression": -10,
		"energy":     -10,
	})

	// Проверка на отвлечение (n8n)
	gosha.DistractionCheck("n8n")

	// 2. Гигиена воина
	time.Sleep(800 * time.Millisecond)
	gosha.CompleteQuest("Гигиена воина", 30, map[string]int{
		"mood":       15,
		"depression": -5,
		"energy":     -5,
	})

	// 3. Курьерский рейс
	time.Sleep(1 * time.Second)
	gosha.CompleteQuest("Курьерский рейс", 100, map[string]int{
		"mood":       20, // красивые девушки в метро
		"depression": -15,
		"energy":     -30,
	})
	gosha.Inventory["money"] += 2000

	// Проверка на отвлечение (игры)
	gosha.DistractionCheck("video_game")

	// 4. Битва с коммуналкой
	time.Sleep(800 * time.Millisecond)
	gosha.CompleteQuest("Битва с коммуналки", 80, map[string]int{
		"discipline": 20,
		"mood":       -30, // последние деньги
		"depression": 20,
	})
	gosha.Inventory["money"] -= 10000

	// 5. Священный час Go
	time.Sleep(1 * time.Second)
	fmt.Println("\n🕖 19:00 - СВЯЩЕННЫЙ ЧАС GO")
	gosha.CompleteQuest("Священный час Go", 200, map[string]int{
		"go_skill":   25,
		"focus":      15,
		"discipline": 15,
		"mood":       10,
		"depression": -20,
		"energy":     -20,
	})

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("✨", 25))
	fmt.Println("          ИТОГИ ДНЯ")
	fmt.Println(strings.Repeat("✨", 25))

	gosha.DisplayStats()

	// Мотивационный итог
	if gosha.Stats["go_skill"].Value >= 50 {
		fmt.Println("\n🎊 ПОЗДРАВЛЯЕМ! Вы прошли порог в 50% навыка Go!")
		fmt.Println("   Скоро вы сможете претендовать на первую работу!")
	} else {
		fmt.Println("\n💪 ПРОДОЛЖАЙТЕ! Каждый день приближает к цели!")
		fmt.Println("   Завтра: интерфейсы и горутины!")
	}

	// Показываем мотивационные фразы
	showMotivationalPhrases()
}

func showMotivationalPhrases() {
	phrases := []string{
		"\n🔟 МОТИВАЦИОННЫХ ФРАЗ ДЛЯ GO-РАЗРАБОТЧИКА:\n",
		"1.  Каждый коммит в Go — это шаг от курьера к разработчику!",
		"2.  Интерфейсы в Go научат тебя интерфейсам в жизни!",
		"3.  Горутины легче видеоигр — они бесплатные и полезные!",
		"4.  go fmt твой код, пока жизнь не отформатировала тебя!",
		"5.  Тесты в Go проходят. Тесты в n8n подождут!",
		"6.  Пока ты учишь Go, твой будущий офис строится!",
		"7.  Каждая изученная структура — кирпич в твоей карьере!",
		"8.  Компилятор Go строгий, но справедливый. Как и рынок труда!",
		"9.  Твой первый Go-проект важнее сотого n8n-бота!",
		"10. Завтра ты скажешь спасибо сегодняшнему себе!",
	}

	for _, phrase := range phrases {
		fmt.Println(phrase)
		time.Sleep(300 * time.Millisecond)
	}
}
