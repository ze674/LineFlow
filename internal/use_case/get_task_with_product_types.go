package use_case

type GetTaskWithProductTypes struct {
	taskService        TaskService
	productTypeService ProductTypeService
}
