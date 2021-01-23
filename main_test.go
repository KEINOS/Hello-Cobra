package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

var thresholdCoverage = 0.9

func Test_main(t *testing.T) {
	// Save current function and restore at the end
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var expectExitCode = 0
	var expectMsg = "Usage:" // Help msg includes this string
	var actualMsg string
	var actualExitCode int

	// Mock of "osExit" to capture the exit-status-code.
	var myExit = func(code int) {
		actualExitCode = code
	}

	// Assign the mock
	osExit = myExit
	// Run main() to capture STDOUT message and its exit-status-code
	actualMsg = capturer.CaptureOutput(func() {
		main()
	})

	// Assertion
	assert.Equal(t, actualExitCode, expectExitCode, "Unexpected exit code.")
	assert.Contains(t, actualMsg, expectMsg, "Should contain help message")
}
