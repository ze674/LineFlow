package entity

import "time"

type ProductInstance struct {
	ID         uint
	TaskID     uint
	Code       string
	Valid      bool
	Defected   bool
	DetectedAt time.Time
	ScannedAt  time.Time
	RejectedAt time.Time
}

func NewProductInstance(TaskID uint, code string) *ProductInstance {
	return &ProductInstance{
		TaskID:     TaskID,
		Code:       code,
		DetectedAt: time.Now(),
	}
}
