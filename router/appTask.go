package router

import (
	"github.com/Vzttfromxduszu/compower_share_golang/service"
	"github.com/gin-gonic/gin"
)

func RouterTask() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetIndex)
	r.GET("/task/tasklist", service.GetTaskList)
	return r
}
