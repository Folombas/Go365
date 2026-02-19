package cmd

import (
	"fmt"
	//"time"

	"github.com/spf13/cobra"
)

var studyCmd = &cobra.Command{
	Use:   "study [minutes]",
	Short: "Записать время, потраченное на изучение Go",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		minutes := args[0]
		fmt.Printf("📚 Записано %s минут изучения Go. XP +%s\n", minutes, minutes)
		// Здесь можно добавить сохранение в файл
	},
}

func init() {
	rootCmd.AddCommand(studyCmd)
}
