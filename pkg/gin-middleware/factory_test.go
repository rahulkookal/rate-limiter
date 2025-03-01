package middleware

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiterFactory(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode

	// Test IP-based rate limiter
	configIP := RateLimiterConfig{Mode: "ip", Rate: 5, Interval: time.Second}
	ipLimiter, err := RateLimiterFactory(configIP)
	assert.NoError(t, err, "IP limiter should not return an error")
	assert.NotNil(t, ipLimiter, "IP limiter should be returned")

	// Test Token-based rate limiter
	configToken := RateLimiterConfig{Mode: "token", Rate: 5, Interval: time.Second}
	tokenLimiter, err := RateLimiterFactory(configToken)
	assert.NoError(t, err, "Token limiter should not return an error")
	assert.NotNil(t, tokenLimiter, "Token limiter should be returned")

	// Test invalid mode
	configInvalid := RateLimiterConfig{Mode: "unknown", Rate: 5, Interval: time.Second}
	invalidLimiter, err := RateLimiterFactory(configInvalid)
	assert.Error(t, err, "Invalid mode should return an error")
	assert.Nil(t, invalidLimiter, "No limiter should be returned for an invalid mode")
}
