package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestIPBucketRateLimiterMiddleware tests the IP-based rate limiter
func TestIPBucketRateLimiterMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router with the rate limiter (2 requests per second)
	rate := 2
	interval := time.Second
	router := gin.New()
	router.Use(iPBucketRateLimiterMiddleware(rate, interval))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Allowed"})
	})

	// Create a test request function
	sendRequest := func(ip string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("X-Forwarded-For", ip) // Simulate real IP
		router.ServeHTTP(w, req)
		return w
	}

	// Simulate requests from the same IP
	ip := "192.168.1.1"
	resp1 := sendRequest(ip)
	resp2 := sendRequest(ip)
	resp3 := sendRequest(ip) // Should be blocked

	// Assertions
	assert.Equal(t, 200, resp1.Code, "First request should be allowed")
	assert.Equal(t, 200, resp2.Code, "Second request should be allowed")
	assert.Equal(t, 429, resp3.Code, "Third request should be rate-limited")

	// Wait for token refill
	time.Sleep(interval)

	// New request after interval should be allowed
	resp4 := sendRequest(ip)
	assert.Equal(t, 200, resp4.Code, "Request after refill should be allowed")
}
