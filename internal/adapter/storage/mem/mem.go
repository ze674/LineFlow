package mem

import (
	"LineFlow/internal/entity"
	"fmt"
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
