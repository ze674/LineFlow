package entity

import "time"

const (
	TaskStatusNew        TaskStatus = "new"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID             uint
	ProductTypeID  uint       // Тип продукции
	ProductionDate time.Time  // Дата производства
	Batch          string     // Номер партии
	Status         TaskStatus // Статус задачи
}

type TaskStatus string // Статус задачи
