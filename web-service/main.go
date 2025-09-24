package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")
	
	// Serve HTML templates
	router.LoadHTMLGlob("templates/*")

	// Route for the main page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "The Gruv - Events Platform",
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "Web Service",
			"status":  "UP",
		})
	})

	// Start the server
	port := "8084"
	fmt.Printf("Web Service listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}