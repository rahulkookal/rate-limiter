package middleware

import "sync"

// RateLimiter structure
type rateLimiter struct {
	requests map[string]int
	mu       sync.Mutex
}

// NewRateLimiter creates a rate limiter
func newRateLimiter() *rateLimiter {
	return &rateLimiter{
		requests: make(map[string]int),
	}
}
