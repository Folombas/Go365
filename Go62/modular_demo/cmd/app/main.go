package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"modular_demo/internal/handler"
	"modular_demo/internal/repository"
	"modular_demo/internal/service"
	"modular_demo/pkg/logger"
)

func main() {
	log := logger.New()
	log.Info("Запуск приложения")

	// Инициализация зависимостей
	repo := repository.NewMemoryRepository()
	taskService := service.NewTaskService(repo)
	cliHandler := handler.NewCLIHandler(taskService, log)

	// Запуск интерфейса командной строки
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в менеджер задач!")
	fmt.Println("Доступные команды: add <текст>, list, complete <id>, exit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}
		cliHandler.Handle(input)
	}
}
