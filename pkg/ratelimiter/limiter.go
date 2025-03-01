package ratelimiter

import (
	"sync"
	"time"
)

// Limiter struct manages rate limiting using a token bucket algorithm.
type Limiter struct {
	rate      int           // Maximum requests allowed per duration
	interval  time.Duration // Time window for rate limiting
	tokens    int           // Available tokens
	lastCheck time.Time     // Last time the bucket was checked
	mutex     sync.Mutex    // Ensures thread safety
}

// NewLimiter creates a new rate limiter with the given rate and time interval.
func NewLimiter(rate int, interval time.Duration) *Limiter {
	return &Limiter{
		rate:      rate,
		interval:  interval,
		tokens:    rate, // Start with a full bucket
		lastCheck: time.Now(),
	}
}

// Allow checks if a request can proceed or should be rate limited.
func (l *Limiter) Allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(l.lastCheck)

	// Add new tokens based on elapsed time
	newTokens := int(elapsed/l.interval) * l.rate
	if newTokens > 0 {
		l.tokens = min(l.tokens+newTokens, l.rate) // Cap at max rate
		l.lastCheck = now
	}

	// If tokens are available, allow request
	if l.tokens > 0 {
		l.tokens--
		return true
	}

	// No tokens left, reject request
	return false
}

// min is a helper function to get the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
