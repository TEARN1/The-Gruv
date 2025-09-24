package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := InitDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer CloseDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Define a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "User Service",
			"status":  "UP",
		})
	})

	// Authentication routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", SignupUser)
		authGroup.POST("/login", LoginUser)
		authGroup.GET("/profile", AuthMiddleware(), GetProfile)
	}

	// Start the server
	port := "8081" // Different port from the API Gateway
	fmt.Printf("User Service listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}