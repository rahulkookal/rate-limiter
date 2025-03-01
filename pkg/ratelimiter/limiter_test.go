package ratelimiter

import (
	"testing"
	"time"
)

func TestNewLimiter(t *testing.T) {
	rate := 5
	interval := time.Second
	lim := NewLimiter(rate, interval)

	if lim.rate != rate {
		t.Errorf("Expected rate %d, got %d", rate, lim.rate)
	}

	if lim.interval != interval {
		t.Errorf("Expected interval %v, got %v", interval, lim.interval)
	}

	if lim.tokens != rate {
		t.Errorf("Expected tokens %d, got %d", rate, lim.tokens)
	}
}

func TestAllow(t *testing.T) {
	lim := NewLimiter(2, time.Second)

	// First two requests should be allowed
	if !lim.Allow() {
		t.Errorf("First request should be allowed")
	}
	if !lim.Allow() {
		t.Errorf("Second request should be allowed")
	}

	// Third request should be denied since no tokens left
	if lim.Allow() {
		t.Errorf("Third request should be denied")
	}

	// Wait for tokens to replenish
	time.Sleep(time.Second)

	// Request should be allowed again
	if !lim.Allow() {
		t.Errorf("Request should be allowed after refill")
	}
}

func TestTokenRefill(t *testing.T) {
	lim := NewLimiter(3, 500*time.Millisecond)

	// Use all tokens
	for i := 0; i < 3; i++ {
		if !lim.Allow() {
			t.Errorf("Request %d should be allowed", i+1)
		}
	}

	// Should be denied since tokens are exhausted
	if lim.Allow() {
		t.Errorf("Request should be denied when tokens are empty")
	}

	// Wait for half the interval
	time.Sleep(500 * time.Millisecond)

	// One token should be refilled
	if !lim.Allow() {
		t.Errorf("Request should be allowed after partial refill")
	}
}
