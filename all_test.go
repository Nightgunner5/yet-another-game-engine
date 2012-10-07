// This file makes sure all the tests are run with a simple "go test"
package ignore

import (
	"os"
	"os/exec"
	"testing"
)

func TestAll(t *testing.T) {
	// having /... twice makes sure that only subdirectories will run
	// tests, preventing an infinite recursive loop.
	command := exec.Command("go", "test", "github.com/Nightgunner5/yet-another-game-engine/.../...", "-v", "-bench", ".")

	command.Stdin, command.Stdout, command.Stderr = os.Stdin, os.Stdout, os.Stderr

	if command.Run() == nil {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
