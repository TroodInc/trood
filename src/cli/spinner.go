package cli

import (
	"fmt"

	"time"

	"github.com/briandowns/spinner"
)

func RunWithSpinner(message string, task func() error) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" %s", message)
	s.Start()

	err := task()

	s.Stop()

	return err
}
