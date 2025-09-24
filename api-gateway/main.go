package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// reverseProxyWithPath creates a reverse proxy handler that strips a prefix from the path
func reverseProxyWithPath(target, prefix string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(c *gin.Context) {
		// Strip the prefix from the path
		originalPath := c.Request.URL.Path
		c.Request.URL.Path = strings.TrimPrefix(originalPath, prefix)
		
		// Update the request host to the target's host
		c.Request.Host = targetURL.Host
		proxy.ServeHTTP(c.Writer, c.Request)
		
		// Restore the original path
		c.Request.URL.Path = originalPath
	}
}

func main() {
	router := gin.Default()

	// Service URLs - use localhost for testing, Docker service names for production
	userServiceURL := "http://localhost:8081"
	eventsServiceURL := "http://localhost:8082"
	collaborationServiceURL := "http://localhost:8083"
	webServiceURL := "http://localhost:8084"

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

	// Route group for events service
	eventsGroup := router.Group("/api/events")
	{
		eventsGroup.Any("/*proxyPath", reverseProxyWithPath(eventsServiceURL, "/api/events"))
	}

	// Route group for collaboration service
	collaborationGroup := router.Group("/api/collaboration")
	{
		collaborationGroup.Any("/*proxyPath", reverseProxy(collaborationServiceURL))
	}

	// Route for the web interface (main page)
	router.Any("/", reverseProxy(webServiceURL))
	router.Any("/static/*proxyPath", reverseProxy(webServiceURL))

	// The real-time service will be handled separately due to WebSockets

	port := "8080"
	fmt.Printf("API Gateway listening on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}