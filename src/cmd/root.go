package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trood",
	Short: "A CLI tool for detecting and fixing issues in your projects.",
	Long: `Trood is an interactive command-line tool designed to scan project files,
identify potential issues, and provide actionable fixes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
