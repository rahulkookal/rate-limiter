package examples

import (
	"fmt"
	"time"

	"github.com/rahulkookal/rate-limiter/pkg/ratelimiter"
)

// RunTest runs the rate limiter test with customizable parameters
func Run(rate int, interval time.Duration) {
	limiter := ratelimiter.NewLimiter(rate, interval)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("✅ Request allowed", i+1)
		} else {
			fmt.Println("❌ Request denied", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
