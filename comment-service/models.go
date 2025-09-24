package main

import (
	"time"
)

// Comment represents a comment in the system with support for threaded replies
type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	VideoURL  string    `json:"videoUrl,omitempty"` // Optional video content
	ParentID  *string   `json:"parentId,omitempty"` // For threaded replies
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Report represents a reported comment
type Report struct {
	ID        string    `json:"id"`
	CommentID string    `json:"commentId"`
	ReporterID string    `json:"reporterId"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"createdAt"`
}

// VideoValidation holds video validation rules
type VideoValidation struct {
	MaxDurationSeconds int `json:"maxDurationSeconds"`
}