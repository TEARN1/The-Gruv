package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory store for events.
// Note: In a real application, this would be a database.
var (
	events      = make(map[string]Event)
	eventsMutex = &sync.RWMutex{}
)

// CreateEvent handles the creation of a new event.
func CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validate that end time is after start time
	if req.EndTime.Before(req.StartTime) || req.EndTime.Equal(req.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
		return
	}

	// Get creator ID from header (in a real app, this would come from JWT token)
	creatorID := c.GetHeader("X-User-ID")
	if creatorID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID required"})
		return
	}

	event := Event{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		CreatorID:   creatorID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		RSVPs:       make(map[string]RSVPStatus),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	eventsMutex.Lock()
	events[event.ID] = event
	eventsMutex.Unlock()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

// GetEvents handles listing all events.
func GetEvents(c *gin.Context) {
	eventsMutex.RLock()
	defer eventsMutex.RUnlock()

	eventsList := make([]Event, 0, len(events))
	for _, event := range events {
		eventsList = append(eventsList, event)
	}

	c.JSON(http.StatusOK, EventListResponse{
		Events: eventsList,
		Total:  len(eventsList),
	})
}

// GetEvent handles getting a specific event by ID.
func GetEvent(c *gin.Context) {
	eventID := c.Param("id")

	eventsMutex.RLock()
	event, exists := events[eventID]
	eventsMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

// UpdateEvent handles updating an existing event.
func UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")
	
	var req UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Get user ID from header
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID required"})
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	event, exists := events[eventID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Check if user is the creator
	if event.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only event creator can update the event"})
		return
	}

	// Update fields if provided
	if req.Title != nil {
		event.Title = *req.Title
	}
	if req.Description != nil {
		event.Description = *req.Description
	}
	if req.StartTime != nil {
		event.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		event.EndTime = *req.EndTime
	}
	if req.Location != nil {
		event.Location = *req.Location
	}

	// Validate that end time is after start time
	if event.EndTime.Before(event.StartTime) || event.EndTime.Equal(event.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
		return
	}

	event.UpdatedAt = time.Now()
	events[eventID] = event

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

// DeleteEvent handles deleting an event.
func DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")

	// Get user ID from header
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID required"})
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	event, exists := events[eventID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Check if user is the creator
	if event.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only event creator can delete the event"})
		return
	}

	delete(events, eventID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}