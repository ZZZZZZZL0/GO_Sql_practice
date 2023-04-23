package controller

import (
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, "1.0.0")
}
