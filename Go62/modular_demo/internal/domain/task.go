package domain

import "time"

// Task представляет бизнес-сущность задачи
type Task struct {
	ID        int
	Text      string
	Completed bool
	CreatedAt time.Time
}

// NewTask создаёт новую задачу (фабричный метод)
func NewTask(id int, text string) *Task {
	return &Task{
		ID:        id,
		Text:      text,
		Completed: false,
		CreatedAt: time.Now(),
	}
}
