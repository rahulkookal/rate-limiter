package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// TokenRateLimiterMiddleware applies the token bucket rate limiter to API tokens
func authTokenBucketRateLimiterMiddleware(rate int, interval time.Duration) gin.HandlerFunc {
	limiter := newGinRateLimiter(rate, interval)

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Missing API token"})
			return
		}

		if !limiter.Allow(token) {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}

		c.Next()
	}
}
