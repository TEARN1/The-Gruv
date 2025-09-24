package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// reverseProxy creates a reverse proxy handler for the given target URL.
func reverseProxy(target string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(c *gin.Context) {
		// Update the request host to the target's host
		c.Request.Host = targetURL.Host
		
		// Extract the proxy path and use it as the new URL path
		proxyPath := c.Param("proxyPath")
		c.Request.URL.Path = proxyPath
		
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	// Service URLs - support both Docker and local development
	userServiceURL := getEnvOrDefault("USER_SERVICE_URL", "http://localhost:8081")
	collaborationServiceURL := getEnvOrDefault("COLLABORATION_SERVICE_URL", "http://localhost:8083")

	// Health check for the gateway itself
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "API Gateway",
			"status":  "UP",
		})
	})

	// Route group for user service
	userGroup := router.Group("/api/users")
	{
		userGroup.Any("/*proxyPath", reverseProxy(userServiceURL))
	}

	// Route group for collaboration service
	collaborationGroup := router.Group("/api/collaboration")
	{
		collaborationGroup.Any("/*proxyPath", reverseProxy(collaborationServiceURL))
	}

	// The real-time service will be handled separately due to WebSockets

	port := "8080"
	fmt.Printf("API Gateway listening on port %s\n", port)
	fmt.Printf("User Service URL: %s\n", userServiceURL)
	fmt.Printf("Collaboration Service URL: %s\n", collaborationServiceURL)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnvOrDefault returns the environment variable value or a default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}