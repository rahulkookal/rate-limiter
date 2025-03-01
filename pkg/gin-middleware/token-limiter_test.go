package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestAuthTokenBucketRateLimiterMiddleware tests the token-based rate limiter
func TestAuthTokenBucketRateLimiterMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router with the rate limiter (2 requests per second)
	rate := 2
	interval := time.Second
	router := gin.New()
	router.Use(authTokenBucketRateLimiterMiddleware(rate, interval))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Allowed"})
	})

	// Helper function to send a request with a token
	sendRequest := func(token string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", token)
		router.ServeHTTP(w, req)
		return w
	}

	// Test: Missing token (should return 400)
	respMissingToken := sendRequest("")
	assert.Equal(t, 400, respMissingToken.Code, "Missing token should return 400")

	// Test: Token-based rate limiting
	token := "user-token-123"

	resp1 := sendRequest(token)
	resp2 := sendRequest(token)
	resp3 := sendRequest(token) // Should be blocked

	// Assertions
	assert.Equal(t, 200, resp1.Code, "First request should be allowed")
	assert.Equal(t, 200, resp2.Code, "Second request should be allowed")
	assert.Equal(t, 429, resp3.Code, "Third request should be rate-limited")

	// Wait for token refill
	time.Sleep(interval)

	// New request after interval should be allowed
	resp4 := sendRequest(token)
	assert.Equal(t, 200, resp4.Code, "Request after refill should be allowed")
}
