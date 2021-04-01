package main

import (
	"github.com/gin-gonic/gin"

	"svcMESH_test/utils"
)

func main() {
	router := gin.Default()
	router.GET("/ping", utils.Ping)
	router.GET("/info", utils.IPInfo)
	router.Run(":8080")
}