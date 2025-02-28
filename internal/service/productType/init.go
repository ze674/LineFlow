package productType

import "LineFlow/internal/entity"

type productTypeRepository interface {
	GetProductTypeByID(id uint) (*entity.ProductType, error)
}

type Service struct {
	repo productTypeRepository
}

func NewService(repo productTypeRepository) *Service {
	return &Service{
		repo: repo,
	}
}
