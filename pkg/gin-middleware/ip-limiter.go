package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func iPBucketRateLimiterMiddleware(rate int, interval time.Duration) gin.HandlerFunc {
	limiter := newGinRateLimiter(rate, interval)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.Allow(ip) {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}
