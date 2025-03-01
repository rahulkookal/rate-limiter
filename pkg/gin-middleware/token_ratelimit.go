package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// TokenRateLimiter limits requests per API token
func tokenRateLimiterMiddleware(limit int, interval time.Duration) gin.HandlerFunc {
	limiter := newRateLimiter()

	go func() {
		for {
			time.Sleep(interval)
			limiter.mu.Lock()
			limiter.requests = make(map[string]int) // Reset counters
			limiter.mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Missing API token"})
			return
		}

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		if limiter.requests[token] >= limit {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}

		limiter.requests[token]++
		c.Next()
	}
}
