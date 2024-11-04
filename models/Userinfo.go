package models

import (
	utils "github.com/Vzttfromxduszu/compower_share_golang/utils"
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

func GetUserList() []*Userinfo {
	data := make([]*Userinfo, 10)
	data = utils.DB.Find(&data)
	return data
}
