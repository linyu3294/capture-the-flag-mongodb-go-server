package controller
//

import (
	"github.com/gin-gonic/gin"
)



func AddPlayer(router *gin.Engine) {
	router.GET("/api/player", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

