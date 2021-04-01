package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": GetHostName(),
	})
}

func IPInfo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": GetHostName(),
		"info": GetIPInfo(),
	})
}