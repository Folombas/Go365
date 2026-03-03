package handler

import (
	"fmt"
	"modular_demo/internal/service"
	"modular_demo/pkg/logger"
	"strconv"
	"strings"
)

// CLIHandler обрабатывает команды из командной строки
type CLIHandler struct {
	taskService *service.TaskService
	log         *logger.Logger
}

func NewCLIHandler(taskService *service.TaskService, log *logger.Logger) *CLIHandler {
	return &CLIHandler{taskService: taskService, log: log}
}

func (h *CLIHandler) Handle(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	cmd := parts[0]
	switch cmd {
	case "add":
		if len(parts) < 2 {
			fmt.Println("Использование: add <текст задачи>")
			return
		}
		text := strings.Join(parts[1:], " ")
		task, err := h.taskService.AddTask(text)
		if err != nil {
			h.log.Error("Ошибка добавления задачи: %v", err)
			fmt.Println("Ошибка")
			return
		}
		fmt.Printf("Задача добавлена с ID %d\n", task.ID)

	case "list":
		tasks, err := h.taskService.ListTasks()
		if err != nil {
			h.log.Error("Ошибка получения задач: %v", err)
			fmt.Println("Ошибка")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("Список задач пуст")
			return
		}
		for _, t := range tasks {
			status := "[ ]"
			if t.Completed {
				status = "[✓]"
			}
			fmt.Printf("%d %s %s\n", t.ID, status, t.Text)
		}

	case "complete":
		if len(parts) != 2 {
			fmt.Println("Использование: complete <id>")
			return
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		if err := h.taskService.CompleteTask(id); err != nil {
			h.log.Error("Ошибка завершения задачи: %v", err)
			fmt.Println("Задача не найдена")
			return
		}
		fmt.Println("Задача отмечена выполненной")

	default:
		fmt.Println("Неизвестная команда. Доступно: add, list, complete, exit")
	}
}
