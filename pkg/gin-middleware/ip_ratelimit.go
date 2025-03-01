package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// IPRateLimiterMiddleware limits requests per IP
func iPRateLimiterMiddleware(limit int, interval time.Duration) gin.HandlerFunc {
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
		ip := c.ClientIP()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		if limiter.requests[ip] >= limit {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}

		limiter.requests[ip]++
		c.Next()
	}
}
