package examples

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahulkookal/rate-limiter/pkg/middleware"
)

// RunServer starts a Gin server with the selected rate-limiting configuration
func RunServer(config middleware.RateLimiterConfig) {
	r := gin.Default()

	// Get the appropriate middleware from the factory
	rateLimiter, err := middleware.RateLimiterFactory(config)
	if err != nil {
		log.Fatalf("❌ Error: %v\n", err)
	}

	// Apply middleware
	r.Use(rateLimiter)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": fmt.Sprintf("Welcome! Mode: %s, Rate Limit: %d req/%s", config.Mode, config.Rate, config.Interval)})
	})

	fmt.Printf("🚀 Server running on :8080 | Mode: %s | Rate: %d req/%s\n", config.Mode, config.Rate, config.Interval)
	r.Run(":8080")
}
