package cmd

import (
	"time"

	"github.com/rahulkookal/rate-limiter/examples"
	"github.com/rahulkookal/rate-limiter/pkg/middleware"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the rate-limited Gin server",
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
