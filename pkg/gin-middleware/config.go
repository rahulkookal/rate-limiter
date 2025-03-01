package middleware

import "time"

// RateLimiterConfig holds the configuration for rate limiting.
//
// Fields:
//   - Mode     (string)        → Determines the rate limiting strategy. Options: "ip" or "token".
//   - Rate     (int)           → Maximum number of allowed requests within the specified interval.
//   - Interval (time.Duration) → The time window for rate limiting.
//
// Example Usage:
//
//	config := RateLimiterConfig{
//	    Mode:     "ip",
//	    Rate:     10,
//	    Interval: time.Minute,
//	}
type RateLimiterConfig struct {
	Mode     string        // "ip" or "token"
	Rate     int           // Number of requests allowed
	Interval time.Duration // Time window for rate limiting
}
