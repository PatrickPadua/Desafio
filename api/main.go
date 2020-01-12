package main

import (
	"github.com/joho/godotenv"
	"github.com/gin-contrib/gzip"
	. "app/domain/suggestion"
	"github.com/gin-gonic/gin"
	"os"
)

var mongo = MongoDB{}

func init() {
	godotenv.Load()
	mongo.Server = os.Getenv("MONGODB_HOST")
	mongo.Database = os.Getenv("MONGODB_DATABASE")
	mongo.Connect()
}

func main() {
	
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	v1 := r.Group("/v1")
	{
		v1.GET("/data", ListSuggestions)
		v1.POST("/data", CreateSuggestion)
	}
	r.GET("/health", HealthCheck)
	r.Run(":8080")
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}