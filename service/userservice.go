package service

import (
	"net/http"

	gin "github.com/gin-gonic/gin"

	// utils "github.com/Vzttfromxduszu/compower_share_golang/utils"
	models "github.com/Vzttfromxduszu/compower_share_golang/models"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.Userinfo, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func CreateUser(c *gin.Context) {
	newuser := models.Userinfo{} //定义一个实例，作为即将新添加的用户实例
	Username := c.Query("用户名")
	Password := c.Query("用户密码")
	Email := c.Query("用户邮箱")
	newuser.Username = Username
	newuser.Password = Password
	newuser.Email = Email
	models.CreateUser(newuser)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})

}
