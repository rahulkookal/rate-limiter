package examples

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulkookal/rate-limiter/middleware"
)

func RunServer(rate int, interval time.Duration) {
	r := gin.Default()

	// Apply rate limiting middleware (5 requests per 10 seconds per IP)
	r.Use(middleware.RateLimiterMiddleware(rate, interval))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.Run(":8080")
}
