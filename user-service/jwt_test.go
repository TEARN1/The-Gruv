package main

import (
	"testing"
	"time"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := "test-user-id"
	username := "testuser"

	// Generate a JWT token
	token, err := GenerateJWT(userID, username)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	if token == "" {
		t.Fatal("Generated token is empty")
	}

	// Validate the JWT token
	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected user ID %s, got %s", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}

	// Check if token expires in approximately 24 hours
	expectedExpiry := time.Now().Add(24 * time.Hour)
	if claims.ExpiresAt == nil {
		t.Fatal("Token should have an expiry time")
	}
	
	timeDiff := claims.ExpiresAt.Time.Sub(expectedExpiry)
	if timeDiff > time.Minute || timeDiff < -time.Minute {
		t.Errorf("Token expiry time is not around 24 hours from now")
	}
}

func TestValidateInvalidJWT(t *testing.T) {
	// Test with invalid token
	_, err := ValidateJWT("invalid.token.here")
	if err == nil {
		t.Fatal("Should have failed to validate invalid token")
	}

	// Test with empty token
	_, err = ValidateJWT("")
	if err == nil {
		t.Fatal("Should have failed to validate empty token")
	}
}