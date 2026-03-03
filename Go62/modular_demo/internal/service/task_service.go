package service

import (
	"modular_demo/internal/domain"
	"modular_demo/internal/repository"
)

// TaskService содержит бизнес-логику работы с задачами
type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// AddTask добавляет новую задачу
func (s *TaskService) AddTask(text string) (*domain.Task, error) {
	task := domain.NewTask(0, text)
	if err := s.repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}

// ListTasks возвращает все задачи
func (s *TaskService) ListTasks() ([]*domain.Task, error) {
	return s.repo.FindAll()
}

// CompleteTask отмечает задачу выполненной
func (s *TaskService) CompleteTask(id int) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	return s.repo.Update(task)
}

// GetTask возвращает задачу по ID
func (s *TaskService) GetTask(id int) (*domain.Task, error) {
	return s.repo.FindByID(id)
}
