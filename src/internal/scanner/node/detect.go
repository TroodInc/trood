package node

import "os"

func IsNodeProject() bool {
	_, err := os.Stat("package.json")
	return err == nil
}
