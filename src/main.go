package main

import (
	router "github.com/Vzttfromxduszu/compower_share_golang/router"
	utils "github.com/Vzttfromxduszu/compower_share_golang/utils"
)

func main() {
	utils.InitDB()
	utils.InitConfig()
	r := router.Router()
	r.Run(":8081")

}
