package main

import "time"

// Event represents an event in the system
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Location    Location  `json:"location"`
	RSVPEnabled bool      `json:"rsvp_enabled"`
	ImageURL    string    `json:"image_url"`
	Category    string    `json:"category"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Location represents event location details
type Location struct {
	Type    string `json:"type"` // "physical" or "online"
	Address string `json:"address,omitempty"`
	URL     string `json:"url,omitempty"`
}

// CreateEventRequest represents the request payload for creating an event
type CreateEventRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	StartTime   string   `json:"start_time" binding:"required"`
	EndTime     string   `json:"end_time" binding:"required"`
	Location    Location `json:"location"`
	RSVPEnabled bool     `json:"rsvp_enabled"`
	ImageURL    string   `json:"image_url"`
	Category    string   `json:"category"`
}