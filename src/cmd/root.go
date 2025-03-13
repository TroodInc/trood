package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trood",
		Short: "A CLI tool for detecting and fixing issues in your projects.",
		Long: `Trood is an interactive command-line tool designed to scan project files,
identify potential issues, and provide actionable fixes.`,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	cmd.AddCommand(newFixCmd())

	return cmd
}
