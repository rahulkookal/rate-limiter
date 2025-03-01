package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RateLimiterFactory returns the appropriate rate limiter middleware based on the provided configuration.
//
// Parameters:
//   - config (RateLimiterConfig): The configuration struct containing rate limiting parameters, including mode, rate, and interval.
//
// Returns:
//   - gin.HandlerFunc: The Gin middleware function for rate limiting.
//   - error: An error if the mode is invalid.
//
// Supported Modes:
//   - "ip"    → Uses IP-based rate limiting.
//   - "token" → Uses token-based rate limiting.
//
// Example Usage:
//
//	config := RateLimiterConfig{
//	    Mode:     "ip",
//	    Rate:     10,
//	    Interval: time.Minute,
//	}
//	limiter, err := RateLimiterFactory(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	r.Use(limiter)
func RateLimiterFactory(config RateLimiterConfig) (gin.HandlerFunc, error) {
	switch config.Mode {
	case "ip":
		return iPBucketRateLimiterMiddleware(config.Rate, config.Interval), nil
	case "token":
		return authTokenBucketRateLimiterMiddleware(config.Rate, config.Interval), nil
	default:
		return nil, fmt.Errorf("invalid rate limiter mode: %s", config.Mode)
	}
}
