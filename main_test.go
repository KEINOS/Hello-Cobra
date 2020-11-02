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

	var ActualCode int
	myExit := func(code int) {
		ActualCode = code
	}

	osExit = myExit
	var actualMsg = capturer.CaptureOutput(func() {
		main()
	})

	var expectCode = 0
	assert.Equal(t, ActualCode, expectCode, "Unexpected exit code.")
	assert.Contains(t, actualMsg, "About:", "Should contain help message")
}
