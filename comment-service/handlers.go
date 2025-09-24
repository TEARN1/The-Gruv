package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// In-memory stores
// Note: In a real application, this would be a database
var (
	comments      = make(map[string]Comment)
	commentsMutex = &sync.RWMutex{}
	reports       = make(map[string]Report)
	reportsMutex  = &sync.RWMutex{}
	
	videoValidationService = NewVideoValidationService()
)

// User struct for API calls to user service
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

// getUserByID fetches user information from the user service
func getUserByID(userID string) (*User, error) {
	// In a real application, this would make an HTTP request to the user service
	// For now, we'll simulate it with a mock response
	
	// Mock user data for testing
	mockUser := &User{
		ID:        userID,
		Username:  "testuser",
		CreatedAt: time.Now().AddDate(-2, 0, 0), // 2 years ago
	}
	
	return mockUser, nil
}

// CreateComment handles creating a new comment
func CreateComment(c *gin.Context) {
	var newComment struct {
		UserID           string `json:"userId" binding:"required"`
		Content          string `json:"content" binding:"required"`
		VideoURL         string `json:"videoUrl,omitempty"`
		VideoDuration    int    `json:"videoDuration,omitempty"`
		ParentID         *string `json:"parentId,omitempty"`
	}

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Get user information to validate video upload
	user, err := getUserByID(newComment.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user information"})
		return
	}

	// Validate video upload if video is provided
	if newComment.VideoURL != "" && newComment.VideoDuration > 0 {
		canUpload, errorMsg, maxDuration := videoValidationService.ValidateVideoUpload(
			user.CreatedAt, 
			newComment.VideoDuration,
		)
		
		if !canUpload {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       errorMsg,
				"maxDuration": maxDuration,
			})
			return
		}
	}

	// Validate parent comment exists if this is a reply
	if newComment.ParentID != nil {
		commentsMutex.RLock()
		_, exists := comments[*newComment.ParentID]
		commentsMutex.RUnlock()
		
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parent comment not found"})
			return
		}
	}

	commentsMutex.Lock()
	defer commentsMutex.Unlock()

	comment := Comment{
		ID:        uuid.New().String(),
		UserID:    newComment.UserID,
		Content:   newComment.Content,
		VideoURL:  newComment.VideoURL,
		ParentID:  newComment.ParentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	comments[comment.ID] = comment

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"comment": comment,
	})
}

// GetComments retrieves comments (with optional parent filtering for threads)
func GetComments(c *gin.Context) {
	parentID := c.Query("parentId")
	
	commentsMutex.RLock()
	defer commentsMutex.RUnlock()

	var filteredComments []Comment
	for _, comment := range comments {
		// If parentId query param is provided, filter by it
		if parentID != "" {
			if (parentID == "null" || parentID == "") && comment.ParentID == nil {
				// Top-level comments
				filteredComments = append(filteredComments, comment)
			} else if comment.ParentID != nil && *comment.ParentID == parentID {
				// Replies to specific parent
				filteredComments = append(filteredComments, comment)
			}
		} else {
			// Return all comments
			filteredComments = append(filteredComments, comment)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": filteredComments,
		"total":    len(filteredComments),
	})
}

// GetComment retrieves a specific comment by ID
func GetComment(c *gin.Context) {
	commentID := c.Param("id")

	commentsMutex.RLock()
	comment, exists := comments[commentID]
	commentsMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

// DeleteComment handles comment deletion (by owner or admin)
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	requestingUserID := c.GetHeader("X-User-ID") // In practice, this would come from auth middleware
	
	if requestingUserID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID required"})
		return
	}

	commentsMutex.Lock()
	defer commentsMutex.Unlock()

	comment, exists := comments[commentID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check if user owns the comment or is admin (simplified check)
	isAdmin := c.GetHeader("X-Is-Admin") == "true"
	isOwner := comment.UserID == requestingUserID

	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this comment"})
		return
	}

	delete(comments, commentID)

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// ReportComment handles reporting a comment
func ReportComment(c *gin.Context) {
	commentID := c.Param("id")
	
	var reportData struct {
		ReporterID string `json:"reporterId" binding:"required"`
		Reason     string `json:"reason" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reportData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Verify comment exists
	commentsMutex.RLock()
	_, exists := comments[commentID]
	commentsMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	reportsMutex.Lock()
	defer reportsMutex.Unlock()

	report := Report{
		ID:         uuid.New().String(),
		CommentID:  commentID,
		ReporterID: reportData.ReporterID,
		Reason:     reportData.Reason,
		CreatedAt:  time.Now(),
	}

	reports[report.ID] = report

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment reported successfully",
		"report":  report,
	})
}

// GetVideoValidation returns video validation info for a user
func GetVideoValidation(c *gin.Context) {
	userID := c.Param("userId")
	
	user, err := getUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user information"})
		return
	}

	maxDuration, canUpload := videoValidationService.GetMaxVideoDuration(user.CreatedAt)
	
	c.JSON(http.StatusOK, gin.H{
		"canUpload":    canUpload,
		"maxDuration": maxDuration,
		"accountAge":  int(time.Since(user.CreatedAt).Hours() / (24 * 30.44)), // months
	})
}

// GetReports returns all reports (admin only)
func GetReports(c *gin.Context) {
	isAdmin := c.GetHeader("X-Is-Admin") == "true"
	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}

	reportsMutex.RLock()
	defer reportsMutex.RUnlock()

	var allReports []Report
	for _, report := range reports {
		allReports = append(allReports, report)
	}

	c.JSON(http.StatusOK, gin.H{
		"reports": allReports,
		"total":   len(allReports),
	})
}