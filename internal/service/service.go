package service

import (
	"github.com/arxenn/tasks/internal/domain"
	"github.com/arxenn/tasks/internal/repository"
)

type Service interface {
	Add(content, priority string) (int, error)
	Done(id int) error
	List(name, priority, status string) ([]domain.Task, error)
	Delete(id int) error
}

type service struct {
	repo repository.Repository
}

func (s *service) Add(content, priority string) (int, error) {

	return 0, nil
}

func (s *service) Done(id int) error {
	return s.repo.Update(
		id,
		domain.Task{
			Status: domain.DoneTaskStatus,
		},
	)
}

func (s *service) List(name, priority, status string) ([]domain.Task, error) {
	return nil, nil
}

func (s *service) Delete(id int) error {
	return s.Delete(id)
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
