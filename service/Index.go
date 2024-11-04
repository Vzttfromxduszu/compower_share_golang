package service

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
