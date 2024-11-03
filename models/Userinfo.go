package models

import (
	"gorm.io/gorm"
)

type Userinfo struct {
	gorm.Model
	UserID   int    `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tasks    []Task `gorm:"foreignKey:UserinfoID" json:"tasks"` // 一对多关系
}

func (b *Userinfo) TableNameuser() string {
	return "userinfo"
}
