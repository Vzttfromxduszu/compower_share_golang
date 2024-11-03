package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID      int    `gorm:"primaryKey" json:"task_id"`
	Computation string `json:"computation"`
	ArrivalTime string `json:"arrival_time"`
	UserinfoID  int    `gorm:"index;" json:"userinfo_id"` // 外键字段
}

func (b *Task) TableNametask() string {
	return "task"
}
