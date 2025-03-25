package scanner

import "github.com/troodinc/trood/internal/scanner/node"

func DetectEnvironment() string {
	if node.IsNodeProject() {
		return "nodejs"
	}

	return "unknown"
}
