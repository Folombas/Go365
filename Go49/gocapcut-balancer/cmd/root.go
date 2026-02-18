package cmd

import (
	"fmt"
	"os"

	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocapcut-balancer",
	Short: "Баланс между Go и видео-монтажом",
	Long: `Инструмент для учёта времени, потраченного на изучение Go и на монтаж видео.
Позволяет не выгорать и следить за прогрессом.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Запусти gocapcut-balancer --help для списка команд")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Здесь будут подкоманды
}
