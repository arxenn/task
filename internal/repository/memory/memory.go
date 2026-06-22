package memrepo

import (
	"errors"
	"sync"

	"github.com/arxenn/tasks/internal/domain"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidID    = errors.New("invalid task ID")
)

type InMemoryRepository struct {
	mu     sync.RWMutex
	tasks  map[int]domain.Task
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tasks:  make(map[int]domain.Task),
		nextID: 1,
	}
}

func (r *InMemoryRepository) Add(t domain.Task) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++

	t.ID = id
	r.tasks[id] = t

	return id, nil
}

func (r *InMemoryRepository) List(filters domain.TaskFilters) ([]domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Task, 0)

	for _, task := range r.tasks {
		if !r.matchesFilters(task, filters) {
			continue
		}
		result = append(result, task)
	}

	return result, nil
}

func (r *InMemoryRepository) Update(id int, t domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if id <= 0 {
		return ErrInvalidID
	}

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	task := r.tasks[id]
	var doUpdate bool

	if t.Content != "" && t.Content != task.Content {
		task.Content = t.Content
		doUpdate = true
	}

	if t.Status != "" && t.Status != task.Status {
		task.Status = t.Status
		doUpdate = true
	}

	if t.Priority != "" && t.Priority != task.Priority {
		task.Priority = t.Priority
		doUpdate = true
	}

	if !doUpdate {
		return nil
	}

	r.tasks[id] = task
	return nil
}

func (r *InMemoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if id <= 0 {
		return ErrInvalidID
	}

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}

func (r *InMemoryRepository) matchesFilters(task domain.Task, filters domain.TaskFilters) bool {
	if filters.Status != "" && task.Status != filters.Status {
		return false
	}

	if filters.Priority != "" && task.Priority != filters.Priority {
		return false
	}

	return true
}
