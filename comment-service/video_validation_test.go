package main

import (
	"testing"
	"time"
)

func TestVideoValidationService_GetMaxVideoDuration(t *testing.T) {
	service := NewVideoValidationService()

	tests := []struct {
		name           string
		accountAge     time.Duration
		expectedDuration int
		expectedCanUpload bool
	}{
		{
			name:           "Account too young (17 months)",
			accountAge:     time.Hour * 24 * 30 * 17, // 17 months
			expectedDuration: 0,
			expectedCanUpload: false,
		},
		{
			name:           "Account 18 months old",
			accountAge:     time.Hour * 24 * 30 * 18, // 18 months  
			expectedDuration: 5,
			expectedCanUpload: true,
		},
		{
			name:           "Account 24 months old",
			accountAge:     time.Hour * 24 * 30 * 24, // 24 months (18 + 6)
			expectedDuration: 10,
			expectedCanUpload: true,
		},
		{
			name:           "Account 30 months old",
			accountAge:     time.Hour * 24 * 30 * 30, // 30 months (18 + 12)
			expectedDuration: 15,
			expectedCanUpload: true,
		},
		{
			name:           "Account 36 months old (max)",
			accountAge:     time.Hour * 24 * 30 * 36, // 36 months (18 + 18)
			expectedDuration: 20,
			expectedCanUpload: true,
		},
		{
			name:           "Account older than 36 months",
			accountAge:     time.Hour * 24 * 30 * 48, // 48 months
			expectedDuration: 20, // Capped at 20
			expectedCanUpload: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountCreatedAt := time.Now().Add(-tt.accountAge)
			duration, canUpload := service.GetMaxVideoDuration(accountCreatedAt)
			
			if duration != tt.expectedDuration {
				// Debug info
				accountAge := time.Now().Sub(accountCreatedAt)
				accountAgeDays := int(accountAge.Hours() / 24)
				accountAgeMonths := accountAgeDays / 30
				t.Logf("Debug: accountAgeMonths=%d, accountAgeDays=%d", accountAgeMonths, accountAgeDays)
				t.Errorf("GetMaxVideoDuration() duration = %v, want %v", duration, tt.expectedDuration)
			}
			if canUpload != tt.expectedCanUpload {
				t.Errorf("GetMaxVideoDuration() canUpload = %v, want %v", canUpload, tt.expectedCanUpload)
			}
		})
	}
}

func TestVideoValidationService_ValidateVideoUpload(t *testing.T) {
	service := NewVideoValidationService()

	tests := []struct {
		name           string
		accountAge     time.Duration
		videoDuration  int
		expectedValid  bool
		expectedMaxDur int
	}{
		{
			name:          "Valid upload for 18 month old account",
			accountAge:    time.Hour * 24 * 30 * 18, // 18 months
			videoDuration: 5,
			expectedValid: true,
			expectedMaxDur: 5,
		},
		{
			name:          "Invalid upload - too young account",
			accountAge:    time.Hour * 24 * 30 * 17, // 17 months
			videoDuration: 5,
			expectedValid: false,
			expectedMaxDur: 0,
		},
		{
			name:          "Invalid upload - video too long",
			accountAge:    time.Hour * 24 * 30 * 18, // 18 months
			videoDuration: 10, // More than allowed 5 seconds
			expectedValid: false,
			expectedMaxDur: 5,
		},
		{
			name:          "Valid upload for 30 month old account",
			accountAge:    time.Hour * 24 * 30 * 30, // 30 months
			videoDuration: 15,
			expectedValid: true,
			expectedMaxDur: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountCreatedAt := time.Now().Add(-tt.accountAge)
			valid, _, maxDur := service.ValidateVideoUpload(accountCreatedAt, tt.videoDuration)
			
			if valid != tt.expectedValid {
				t.Errorf("ValidateVideoUpload() valid = %v, want %v", valid, tt.expectedValid)
			}
			if maxDur != tt.expectedMaxDur {
				t.Errorf("ValidateVideoUpload() maxDur = %v, want %v", maxDur, tt.expectedMaxDur)
			}
		})
	}
}