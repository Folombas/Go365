package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var deepseekCmd = &cobra.Command{
	Use:   "deepseek",
	Short: "Показать, сколько дней осталось до выхода DeepSeek V4 (по слухам)",
	Run: func(cmd *cobra.Command, args []string) {
		estimatedRelease := time.Date(2026, 2, 20, 0, 0, 0, 0, time.UTC) // предположим
		daysLeft := int(time.Until(estimatedRelease).Hours() / 24)
		if daysLeft < 0 {
			fmt.Println("🎉 DeepSeek V4 уже вышел! Беги пробовать!")
		} else {
			fmt.Printf("🤖 До выхода DeepSeek V4 осталось примерно %d дней. Терпение!\n", daysLeft)
		}
	},
}

func init() {
	rootCmd.AddCommand(deepseekCmd)
}
