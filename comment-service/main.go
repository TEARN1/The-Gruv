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

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "Comment Service",
			"status":  "UP",
		})
	})

	// Comment routes
	router.POST("/comments", CreateComment)
	router.GET("/comments", GetComments)
	router.GET("/comments/:id", GetComment)
	router.DELETE("/comments/:id", DeleteComment)
	
	// Comment reporting
	router.POST("/comments/:id/report", ReportComment)
	
	// Admin routes
	router.GET("/reports", GetReports)
	
	// Video validation
	router.GET("/users/:userId/video-validation", GetVideoValidation)

	// Start the server
	port := "8082"
	fmt.Printf("Comment Service listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}