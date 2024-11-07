package models

import (
	utils "github.com/Vzttfromxduszu/compower_share_golang/utils"
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

func GetTaskList() []*Task {
	data := make([]*Task, 10)
	utils.DB.Find(&data)
	return data
}
