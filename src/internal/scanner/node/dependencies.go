package node

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func CheckMissingDependencies(path string) error {
	cmd := exec.Command("npm", "install", "--dry-run")
	cmd.Dir = path
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), "added") {
		fmt.Println("\n\n⚠️ Missing dependencies detected")
		fmt.Println(string(output))
	} else {
		fmt.Printf("✅ No missing dependencies found.")
	}

	return nil
}

func CheckUnusedDependencies(path string) error {
	cmd := exec.Command("npx", "depcheck")
	cmd.Dir = path
	output, _ := cmd.CombinedOutput()

	if len(output) > 0 {
		return errors.New(string(output))
	}

	return nil
}

func CheckSecurityIssues(path string) error {
	cmd := exec.Command("npm", "audit", "--json")
	cmd.Dir = path
	output, _ := cmd.CombinedOutput()

	if len(output) > 0 {
		fmt.Println("\n\n📢 Security vulnerabilities detected")
		fmt.Println(string(output))
	} else {
		fmt.Println("✅ No security vulnerabilities found.")
	}

	return nil
}
