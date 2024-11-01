package models

import (
	"gorm.io/gorm"
)

type Userinfo struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Task     []Task `json:"task"`
}

func (b *Userinfo) TableNameuser() string {
	return "userinfo"
}
