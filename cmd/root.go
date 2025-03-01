/*
Copyright © 2025 Rahul Kookal<rahulkookal@protonmail.com>
*/
package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

// Flags for customization
var rate int
var interval time.Duration

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rate-limiter",
	Short: "Rate limiter CLI",
	Long:  `A CLI to run different examples of the Go rate limiter.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
