package task

import (
	"LineFlow/internal/entity"
	"fmt"
)

func (s *Service) GetAllTasks() ([]*entity.Task, error) {
	op := "service.task.GetAllTasks"
	fmt.Printf("%s\n", op)
	return s.repo.GetAllTasks()
}
