package middleware

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGinRateLimiter(t *testing.T) {
	rate := 2
	interval := time.Second
	limiter := newGinRateLimiter(rate, interval)

	// Test Token 1 (Simulate an IP or Auth Token)
	token1 := "user-1"

	// Allow first `rate` requests
	assert.True(t, limiter.Allow(token1), "First request should be allowed")
	assert.True(t, limiter.Allow(token1), "Second request should be allowed")

	// Exceed limit - should return false
	assert.False(t, limiter.Allow(token1), "Third request should be blocked")

	// Wait for interval to replenish tokens
	time.Sleep(interval)

	// New request after refill should be allowed
	assert.True(t, limiter.Allow(token1), "Request after refill should be allowed")

	// Test Token 2 (Different user/IP should have a separate bucket)
	token2 := "user-2"

	// Token2 should be allowed fresh since it's a separate limiter
	assert.True(t, limiter.Allow(token2), "New token should be allowed")
	assert.True(t, limiter.Allow(token2), "Second request for new token should be allowed")
	assert.False(t, limiter.Allow(token2), "Third request for new token should be blocked")
}
