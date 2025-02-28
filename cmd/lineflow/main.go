package main

import (
	"LineFlow/internal/adapter/storage/mem"
	"LineFlow/internal/service/productType"

	"LineFlow/utils/logger"
)

func main() {
	log, err := logger.New(false)
	if err != nil {
		panic(err)
	}

	log.Info("Hello, World!")
	storage := mem.New()
	productTypeService := productType.NewService(storage)

	pr, err := productTypeService.GetProductTypeByID(1)
	if err != nil {
		panic(err)
	}

	println(pr.Name)

}
