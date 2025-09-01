package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"greeting": "Hello from Go + Gin"})
	})

	router.GET("/env", func(c *gin.Context) {
		dbUrl := os.Getenv("DATABASE_URL")
		c.JSON(http.StatusOK, gin.H{"DATABASE_URL": dbUrl})
	})

	router.Run(":8205")
}
