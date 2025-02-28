package mem

import (
	"LineFlow/internal/entity"
	"fmt"
	"time"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) GetProductTypeByID(id uint) (*entity.ProductType, error) {
	if id == 1 {
		return &entity.ProductType{
			ID:   1,
			Name: "test",
			GTIN: "1234",
		}, nil
	}
	return nil, fmt.Errorf("not found")
}

func (s *Storage) GetAllTasks() ([]*entity.Task, error) {
	tasks := []*entity.Task{
		{
			ID:             1,
			ProductTypeID:  1,
			ProductionDate: time.Now(),
			Batch:          "123",
			Status:         entity.TaskStatusNew,
		},
		{
			ID:             2,
			ProductTypeID:  5,
			ProductionDate: time.Now(),
			Batch:          "456",
			Status:         entity.TaskStatusNew,
		},
	}

	return tasks, nil
}
