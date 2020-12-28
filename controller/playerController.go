package controller

//

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPlayers(collection *mongo.Database, router *gin.Engine) {
	a, _ := collection.Collection("players").Find(context.TODO(), bson.M{"name" :"Yu Lin"})
	router.GET("/api/game/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": a,
		})
	})
}

func CreatePlayer(router *gin.Engine) {
	router.POST("/api/game/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, 世界",
		})
	})
}

func UpdatePlayer(router *gin.Engine) {
	router.PUT("/api/game/players/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, 世界",
		})
	})
}

func Deleteplayer(router *gin.Engine) {
	router.DELETE("/api/game/players/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, 世界",
		})
	})
}
