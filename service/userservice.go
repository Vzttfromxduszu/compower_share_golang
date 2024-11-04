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
