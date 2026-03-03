package repository

import (
	"errors"
	"modular_demo/internal/domain"
	"sync"
)

// MemoryRepository реализует TaskRepository в памяти
type MemoryRepository struct {
	mu     sync.RWMutex
	tasks  map[int]*domain.Task
	nextID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks:  make(map[int]*domain.Task),
		nextID: 1,
	}
}

func (r *MemoryRepository) Save(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if task.ID == 0 {
		task.ID = r.nextID
		r.nextID++
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *MemoryRepository) FindByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *MemoryRepository) FindAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *MemoryRepository) Update(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[task.ID]; !ok {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *MemoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
