package middleware

import (
	"sync"
	"time"
)

// TokenBucket holds the rate limit state per IP/AuthToken
type tokenBucket struct {
	tokens    int
	lastCheck time.Time
	mu        sync.Mutex
}

// RateLimiter stores rate limiters for each IP/AuthToken
type ginRateLimiter struct {
	rate     int
	interval time.Duration
	buckets  map[string]*tokenBucket
	mu       sync.Mutex
}

func newGinRateLimiter(rate int, interval time.Duration) *ginRateLimiter {
	return &ginRateLimiter{
		rate:     rate,
		interval: interval,
		buckets:  make(map[string]*tokenBucket),
	}
}

// Allow checks and refills tokens
func (rl *ginRateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Get or create token bucket for the IP
	bucket, exists := rl.buckets[ip]
	if !exists {
		bucket = &tokenBucket{tokens: rl.rate, lastCheck: time.Now()}
		rl.buckets[ip] = bucket
	}

	// Lock bucket for safe updates
	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(bucket.lastCheck)
	newTokens := int(elapsed / rl.interval * time.Duration(rl.rate))
	if newTokens > 0 {
		bucket.tokens = min(bucket.tokens+newTokens, rl.rate)
		bucket.lastCheck = now
	}

	// Allow request if tokens are available
	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}

	// No tokens left, reject request
	return false
}

// min helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
