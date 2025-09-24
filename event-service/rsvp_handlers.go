package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RSVPToEvent handles RSVP requests to an event.
func RSVPToEvent(c *gin.Context) {
	eventID := c.Param("id")
	
	var req RSVPRequest
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

	// Check if event has already ended
	if time.Now().After(event.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot RSVP to a past event"})
		return
	}

	// Update or create RSVP
	rsvp := RSVPStatus{
		Status:    req.Status,
		Timestamp: time.Now(),
		UserID:    userID,
	}

	if event.RSVPs == nil {
		event.RSVPs = make(map[string]RSVPStatus)
	}
	
	event.RSVPs[userID] = rsvp
	event.UpdatedAt = time.Now()
	events[eventID] = event

	c.JSON(http.StatusOK, gin.H{
		"message": "RSVP updated successfully",
		"rsvp":    rsvp,
	})
}

// GetEventRSVPs handles getting all RSVPs for an event.
func GetEventRSVPs(c *gin.Context) {
	eventID := c.Param("id")

	eventsMutex.RLock()
	event, exists := events[eventID]
	eventsMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Convert RSVPs map to slice for consistent JSON output
	rsvps := make([]RSVPStatus, 0, len(event.RSVPs))
	for _, rsvp := range event.RSVPs {
		rsvps = append(rsvps, rsvp)
	}

	// Count RSVPs by status
	counts := make(map[RSVPType]int)
	for _, rsvp := range event.RSVPs {
		counts[rsvp.Status]++
	}

	c.JSON(http.StatusOK, gin.H{
		"eventId": eventID,
		"rsvps":   rsvps,
		"counts": counts,
		"total":  len(rsvps),
	})
}

// RemoveRSVP handles removing an RSVP from an event.
func RemoveRSVP(c *gin.Context) {
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

	// Check if user has an RSVP
	if _, hasRSVP := event.RSVPs[userID]; !hasRSVP {
		c.JSON(http.StatusNotFound, gin.H{"error": "No RSVP found for this user"})
		return
	}

	delete(event.RSVPs, userID)
	event.UpdatedAt = time.Now()
	events[eventID] = event

	c.JSON(http.StatusOK, gin.H{
		"message": "RSVP removed successfully",
	})
}