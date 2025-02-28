package task

import "LineFlow/internal/entity"

type taskRepository interface {
	GetAllTasks() ([]*entity.Task, error)
}

type Service struct {
	repo taskRepository
}

func NewService(repo taskRepository) *Service {
	return &Service{
		repo: repo,
	}
}
