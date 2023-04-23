package controller

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	//test := c.Params("TEST")
	// call servic
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
