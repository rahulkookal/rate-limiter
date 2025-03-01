package cmd

import (
	"time"

	"github.com/rahulkookal/rate-limiter/examples"
	middleware "github.com/rahulkookal/rate-limiter/pkg/gin-middleware"
	"github.com/spf13/cobra"
)

// serverCmd represents the command to start the Gin server with rate limiting.
var serverCmd = &cobra.Command{
	Use:   "gin-middleware",
	Short: "Run the rate-limited Gin server",
	Long: `Start a Gin server with built-in rate limiting middleware. 
Allows IP-based or token-based rate limiting with configurable settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := middleware.RateLimiterConfig{
			Mode:     mode,
			Rate:     rate,
			Interval: interval,
		}
		examples.RunServer(config)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Allow customization through CLI flags
	serverCmd.Flags().StringVarP(&mode, "mode", "m", "ip", "Rate limiting mode (ip/token)")
	serverCmd.Flags().IntVarP(&rate, "rate", "r", 5, "Number of requests allowed per interval")
	serverCmd.Flags().DurationVarP(&interval, "interval", "i", 10*time.Second, "Time interval for rate limiting")
}
