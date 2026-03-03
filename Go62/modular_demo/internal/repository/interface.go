package repository

import "modular_demo/internal/domain"

// TaskRepository определяет контракт хранилища задач
type TaskRepository interface {
	Save(task *domain.Task) error
	FindByID(id int) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id int) error
}
