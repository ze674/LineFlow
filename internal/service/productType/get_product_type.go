package productType

import (
	"LineFlow/internal/entity"
	"fmt"
)

func (pt *Service) GetProductTypeByID(id uint) (*entity.ProductType, error) {
	op := "service.productType.GetProductTypeByID"
	fmt.Printf("%s: %d\n", op, id)

	product, err := pt.repo.GetProductTypeByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return product, nil
}
