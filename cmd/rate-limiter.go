package cmd

import (
	"time"

	"github.com/rahulkookal/rate-limiter/examples"
	"github.com/spf13/cobra"
)

// testCmd represents the command to test the rate limiter functionality.
var testCmd = &cobra.Command{
	Use:   "rate-limiter",
	Short: "Run a rate limiter test example",
	Long: `Execute a test example of the rate limiter. 
This simulates requests at a given rate within a specified time interval.`,
	Run: func(cmd *cobra.Command, args []string) {
		examples.Run(rate, interval)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Allow customization through CLI flags
	testCmd.Flags().IntVarP(&rate, "rate", "r", 5, "Number of requests allowed per interval")
	testCmd.Flags().DurationVarP(&interval, "interval", "i", time.Second, "Time interval for rate limiting")
}
