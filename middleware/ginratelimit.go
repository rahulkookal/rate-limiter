package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulkookal/rate-limiter/pkg/ratelimiter"
)

// RateLimiterMiddleware is a Gin middleware that applies rate limiting per client IP.
func RateLimiterMiddleware(rate int, interval time.Duration) gin.HandlerFunc {
	limiterMap := make(map[string]*ratelimiter.Limiter)
	var mutex sync.Mutex

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		// Get or create a limiter for this IP
		mutex.Lock()
		limiter, exists := limiterMap[clientIP]
		if !exists {
			limiter = ratelimiter.NewLimiter(rate, interval)
			limiterMap[clientIP] = limiter
		}
		mutex.Unlock()

		// Check if the request is allowed
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
