package main

import (
	"LineFlow/internal/adapter/storage/mem"
	"LineFlow/internal/service/productType"
	"LineFlow/internal/service/task"
	"LineFlow/utils/logger"
)

func main() {
	log, err := logger.New(false)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	log.Info("Hello, World!")
	storage := mem.New()
	taskService := task.NewService(storage)
	productTypeService := productType.NewService(storage)

	taskList, err := taskService.GetAllTasks()
	if err != nil {
		log.Error(err)
	}
	for _, t := range taskList {
		pt, err := productTypeService.GetProductTypeByID(t.ProductTypeID)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(pt)
	}

	log.Info("Bye, World!")

}
