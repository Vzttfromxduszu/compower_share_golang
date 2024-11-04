package router

import (
	"github.com/Vzttfromxduszu/compower_share_golang/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetIndex)
	r.GET("/user/userlist", service.GetUserList)
	return r
}
