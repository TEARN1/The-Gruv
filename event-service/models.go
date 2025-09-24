package main

import (
	"time"
)

// Event represents an event in the system.
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatorID   string    `json:"creatorId"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Location    Location  `json:"location"`
	RSVPs       map[string]RSVPStatus `json:"rsvps,omitempty"` // userID -> RSVP status
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Location represents the location information for an event.
type Location struct {
	Type        LocationType `json:"type"`        // online, venue, address
	Name        string       `json:"name"`        // venue name or online platform
	Address     string       `json:"address"`     // physical address (for venue/address types)
	Coordinates *Coordinates `json:"coordinates,omitempty"` // GPS coordinates for directions
	URL         string       `json:"url,omitempty"`         // URL for online events or venue websites
}

// LocationType defines the type of location
type LocationType string

const (
	LocationOnline  LocationType = "online"
	LocationVenue   LocationType = "venue"
	LocationAddress LocationType = "address"
)

// Coordinates represents GPS coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// RSVPStatus represents the RSVP status of a user for an event
type RSVPStatus struct {
	Status    RSVPType  `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"userId"`
}

// RSVPType defines the possible RSVP responses
type RSVPType string

const (
	RSVPGoing    RSVPType = "going"
	RSVPNotGoing RSVPType = "not_going"
	RSVPMaybe    RSVPType = "maybe"
)

// CreateEventRequest represents the request payload for creating an event
type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime" binding:"required"`
	EndTime     time.Time `json:"endTime" binding:"required"`
	Location    Location  `json:"location" binding:"required"`
}

// UpdateEventRequest represents the request payload for updating an event
type UpdateEventRequest struct {
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	StartTime   *time.Time `json:"startTime,omitempty"`
	EndTime     *time.Time `json:"endTime,omitempty"`
	Location    *Location `json:"location,omitempty"`
}

// RSVPRequest represents the request payload for RSVP operations
type RSVPRequest struct {
	Status RSVPType `json:"status" binding:"required"`
}

// EventListResponse represents the response for listing events
type EventListResponse struct {
	Events []Event `json:"events"`
	Total  int     `json:"total"`
}