package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory store for users.
// Note: In a real application, this would be a database.
var (
	users      = make(map[string]User)
	usersMutex = &sync.RWMutex{}
)

// RegisterUser handles the creation of a new user.
func RegisterUser(c *gin.Context) {
	var newUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Gender   string `json:"gender,omitempty"`
		Email    string `json:"email,omitempty"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Check if username already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}
	}

	user := User{
		ID:       uuid.New().String(),
		Username: newUser.Username,
		Gender:   newUser.Gender,
		Email:    newUser.Email,
	}

	if err := user.HashPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	users[user.ID] = user

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"userId":  user.ID,
	})
}

// LoginUser handles user authentication.
func LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	usersMutex.RLock()
	defer usersMutex.RUnlock()

	// Find user by username
	var foundUser *User
	for _, user := range users {
		if user.Username == credentials.Username {
			u := user
			foundUser = &u
			break
		}
	}

	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check password
	if err := foundUser.CheckPassword(credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"userId":  foundUser.ID,
		"user": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"gender":   foundUser.Gender,
			"email":    foundUser.Email,
			"bio":      foundUser.Bio,
			"avatarUrl": foundUser.AvatarURL,
		},
		// In a real app, you would return a JWT token here
	})
}

// GetProfile handles retrieving user profile information.
func GetProfile(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	usersMutex.RLock()
	defer usersMutex.RUnlock()

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"gender":   user.Gender,
		"email":    user.Email,
		"bio":      user.Bio,
		"avatarUrl": user.AvatarURL,
	})
}

// UpdateProfile handles updating user profile information.
func UpdateProfile(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var updates struct {
		Gender    string `json:"gender,omitempty"`
		Email     string `json:"email,omitempty"`
		Bio       string `json:"bio,omitempty"`
		AvatarURL string `json:"avatarUrl,omitempty"`
	}

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields if provided
	if updates.Gender != "" {
		user.Gender = updates.Gender
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.Bio != "" {
		user.Bio = updates.Bio
	}
	if updates.AvatarURL != "" {
		user.AvatarURL = updates.AvatarURL
	}

	users[userID] = user

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"gender":   user.Gender,
			"email":    user.Email,
			"bio":      user.Bio,
			"avatarUrl": user.AvatarURL,
		},
	})
}