package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Показать баланс и XP",
	Run: func(cmd *cobra.Command, args []string) {
		startDate := time.Date(2026, 1, 18, 0, 0, 0, 0, time.UTC)
		days := int(time.Since(startDate).Hours() / 24)

		fmt.Println("📊 СТАТУС ПЕРСОНАЖА:")
		fmt.Printf("Дней без CapCut (осознанных): %d\n", days)
		fmt.Printf("XP за учёбу: ~%d\n", days*10)
		fmt.Printf("XP за монтаж: ~%d\n", days*2) // для примера
		fmt.Printf("Баланс: %d%% учёбы, %d%% отдыха\n", 80, 20)
		fmt.Println("Совет: переключайся, чтобы не выгореть!")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
