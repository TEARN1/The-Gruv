package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory store for events
// Note: In a real application, this would be a database
var (
	events      = make(map[string]Event)
	eventsMutex = &sync.RWMutex{}
)

// RegisterEventRoutes registers all event-related routes
func RegisterEventRoutes(router *gin.Engine) {
	// Routes expect to be accessed via proxy at /api/events/*
	// So /events maps to /api/events/events through the gateway
	router.GET("/events", GetEvents)
	router.POST("/events", CreateEvent)
	router.GET("/events/:id", GetEventByID)
	router.PUT("/events/:id", UpdateEvent)
	router.DELETE("/events/:id", DeleteEvent)
}

// GetEvents returns all events
func GetEvents(c *gin.Context) {
	eventsMutex.RLock()
	defer eventsMutex.RUnlock()

	eventList := make([]Event, 0, len(events))
	for _, event := range events {
		eventList = append(eventList, event)
	}

	c.JSON(http.StatusOK, gin.H{
		"events": eventList,
	})
}

// CreateEvent creates a new event
func CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Parse time strings
	startTime, err := time.Parse("2006-01-02T15:04", req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format. Use YYYY-MM-DDTHH:MM"})
		return
	}

	endTime, err := time.Parse("2006-01-02T15:04", req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format. Use YYYY-MM-DDTHH:MM"})
		return
	}

	// Validate that end time is after start time
	if endTime.Before(startTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	event := Event{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		Location:    req.Location,
		RSVPEnabled: req.RSVPEnabled,
		ImageURL:    req.ImageURL,
		Category:    req.Category,
		CreatedBy:   "system", // TODO: Get from authentication context
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	events[event.ID] = event

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

// GetEventByID returns a specific event by ID
func GetEventByID(c *gin.Context) {
	id := c.Param("id")

	eventsMutex.RLock()
	defer eventsMutex.RUnlock()

	event, exists := events[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

// UpdateEvent updates an existing event
func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	event, exists := events[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Parse time strings
	startTime, err := time.Parse("2006-01-02T15:04", req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format. Use YYYY-MM-DDTHH:MM"})
		return
	}

	endTime, err := time.Parse("2006-01-02T15:04", req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format. Use YYYY-MM-DDTHH:MM"})
		return
	}

	// Update the event
	event.Title = req.Title
	event.Description = req.Description
	event.StartTime = startTime
	event.EndTime = endTime
	event.Location = req.Location
	event.RSVPEnabled = req.RSVPEnabled
	event.ImageURL = req.ImageURL
	event.Category = req.Category
	event.UpdatedAt = time.Now()

	events[id] = event

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

// DeleteEvent deletes an event
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	_, exists := events[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	delete(events, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}