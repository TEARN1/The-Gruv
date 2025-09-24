package main

import (
	"time"
)

// VideoValidationService provides video validation logic
type VideoValidationService struct{}

// NewVideoValidationService creates a new video validation service
func NewVideoValidationService() *VideoValidationService {
	return &VideoValidationService{}
}

// GetMaxVideoDuration calculates the maximum allowed video duration based on account age
// Rules:
// - Accounts older than 18 months can upload video comments
// - Video duration starts at 5 seconds at 18 months 
// - Increases by 5 seconds every 6 months
// - Maximum of 20 seconds (36 months)
func (v *VideoValidationService) GetMaxVideoDuration(accountCreatedAt time.Time) (int, bool) {
	now := time.Now()
	accountAge := now.Sub(accountCreatedAt)
	
	// Convert to months using a more accurate calculation
	// Use 30 days per month as approximation
	accountAgeDays := int(accountAge.Hours() / 24)
	accountAgeMonths := accountAgeDays / 30
	
	// Must be at least 18 months old
	if accountAgeMonths < 18 {
		return 0, false
	}
	
	// Calculate allowed duration
	// Start at 5 seconds at 18 months, add 5 seconds every 6 months
	monthsSince18 := accountAgeMonths - 18
	additionalPeriods := monthsSince18 / 6
	maxDuration := 5 + (additionalPeriods * 5)
	
	// Cap at 20 seconds (which happens at 36 months: 18 + 18 months)
	if maxDuration > 20 {
		maxDuration = 20
	}
	
	return maxDuration, true
}

// ValidateVideoUpload validates if a video upload is allowed for the user
func (v *VideoValidationService) ValidateVideoUpload(accountCreatedAt time.Time, videoDurationSeconds int) (bool, string, int) {
	maxDuration, canUpload := v.GetMaxVideoDuration(accountCreatedAt)
	
	if !canUpload {
		return false, "Account must be at least 18 months old to upload video comments", 0
	}
	
	if videoDurationSeconds > maxDuration {
		return false, "Video duration exceeds allowed limit", maxDuration
	}
	
	return true, "", maxDuration
}