package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/troodinc/trood/internal/scanner/node"
)

func ScanCmd() *cobra.Command {
	var path string

	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan a project for dependency issues.",
		Long: `The scan command will check a Node.js project for:
  - Missing dependencies
  - Unused dependencies
  - Security vulnerabilities`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("please provide the path to scan")
			}
			path := args[0]
			return node.RunNodeScan(path)
		},
	}

	cmd.Flags().StringVarP(&path, "path", "p", "", "Path to the project to scan (defaults to current directory)")

	return cmd
}
