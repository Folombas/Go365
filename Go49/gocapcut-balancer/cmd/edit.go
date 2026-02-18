package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [minutes]",
	Short: "Записать время, потраченное на монтаж видео",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		minutes := args[0]
		fmt.Printf("🎬 Записано %s минут монтажа. Отдых мозга засчитан.\n", minutes)
		// Здесь можно добавить сохранение в файл
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
