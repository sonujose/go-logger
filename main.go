package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := NewLogger()

	router := gin.New()
	router.Use(RequestLogger(logger))

	// Health endpoint
	router.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "up"}) })

	registerEndpoints(router)

	logger.Infof("Starting server on port %s...", getEnv("APP_PORT", "7006"))

	router.Run(":" + getEnv("APP_PORT", "7006"))
}

func registerEndpoints(router *gin.Engine) {
	apiV1 := router.Group("api")
	{
		apiV1.POST("/users", getUsers)
	}
}

func getUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"user1": "mike",
	})
}
