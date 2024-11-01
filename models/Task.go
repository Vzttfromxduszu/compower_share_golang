package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int    `json:"id"`
	Computation string `json:"computation"`
	ArrivalTime string `json:"arrival_time"`
}

func (b *Task) TableNametask() string {
	return "task"
}
