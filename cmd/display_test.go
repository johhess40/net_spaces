package cmd

import (
	"testing"
)

func TestExecute(t *testing.T) {
	Execute()
}

func TestDisplay(t *testing.T) {
	err := displayCmd.Flags().Parse([]string{"display-name"})
	if err != nil {
		return
	}
}
