package node

import (
	"github.com/troodinc/trood/cli"
)

func RunNodeScan(path string) error {
	cli.RunWithSpinner("Checking for missing dependencies", func() error {
		return CheckMissingDependencies(path)
	})

	cli.RunWithSpinner("Checking for unused dependencies", func() error {
		return CheckUnusedDependencies(path)
	})

	cli.RunWithSpinner("Checking for security vulnerabilities", func() error {
		return CheckSecurityIssues(path)
	})

	return nil
}
