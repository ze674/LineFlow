package main

import (
	"LineFlow/internal/adapter/storage/mem"
	"LineFlow/internal/service/productType"
)

func main() {
	storage := mem.New()
	productTypeService := productType.NewService(storage)

	pr, err := productTypeService.GetProductTypeByID(1)
	if err != nil {
		panic(err)
	}

	println(pr.Name)

}
