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
	}

	if err := user.HashPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	users[user.ID] = user

	// Generate JWT token
	token, err := GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"userId":  user.ID,
		"token":   token,
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

	// Generate JWT token
	token, err := GenerateJWT(foundUser.ID, foundUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"userId":  foundUser.ID,
		"token":   token,
	})
}

// GetUserProfile handles retrieving user profile information
func GetUserProfile(c *gin.Context) {
	// Get user ID from JWT claims (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	usersMutex.RLock()
	defer usersMutex.RUnlock()

	// Find user by ID
	user, exists := users[userID.(string)]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}