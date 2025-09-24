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

	// Define a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "Event Service",
			"status":  "UP",
		})
	})

	// Event routes
	api := router.Group("/api/events")
	{
		api.POST("", CreateEvent)           // Create a new event
		api.GET("", GetEvents)              // Get all events
		api.GET("/:id", GetEvent)           // Get event by ID
		api.PUT("/:id", UpdateEvent)        // Update event by ID
		api.DELETE("/:id", DeleteEvent)     // Delete event by ID
		
		// RSVP routes
		api.POST("/:id/rsvp", RSVPToEvent)      // RSVP to an event
		api.GET("/:id/rsvps", GetEventRSVPs)    // Get all RSVPs for an event
		api.DELETE("/:id/rsvp", RemoveRSVP)     // Remove RSVP from an event
	}

	// Start the server
	port := "8082" // Different port from other services
	fmt.Printf("Event Service listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}