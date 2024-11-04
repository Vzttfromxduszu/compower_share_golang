package utils

import (
	"fmt"

	models "github.com/Vzttfromxduszu/compower_share_golang/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

func InitDB() {
	// some db initialization
	DB, _ := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	user := models.Userinfo{}
	DB.Find(&user)
}

func InitConfig() {
	// some config initialization
	viper.SetConfigName("app")
	viper.AddConfigPath("c:\\Users\\Vzt\\Desktop\\OA_code\\Golang_subject\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(viper.Get("app"))
	fmt.Println(viper.Get("mysql"))
}
