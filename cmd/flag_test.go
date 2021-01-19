package cmd

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_flagCmd(t *testing.T) {
	var contains string
	var result = capturer.CaptureStdout(func() {
		var cmdTmp = &cobra.Command{}
		var argsTmp = []string{}
		flagCmd.Run(cmdTmp, argsTmp)
	})

	contains = "The value of Alpha: '' (empty. Option not provided)"
	assert.Contains(t, result, contains, "When no arg, it should display that no option was provided.")

	contains = "The value of Beta: false (Option not provided)"
	assert.Contains(t, result, contains, "When no arg, it should display that no option was provided."+result)
}

func Test_echoValueAlpha(t *testing.T) {
	var contains string
	var result string

	// Alpha variable is empty
	Alpha = ""
	contains = "empty. Option not provided"
	result = capturer.CaptureStdout(func() {
		echoValueAlpha()
	})
	assert.Contains(t, result, contains, "When variable Alpha is empty, it should display that it's empty.")

	// Alpha variable is empty
	Alpha = "foo bar buzz"
	contains = Alpha
	result = capturer.CaptureStdout(func() {
		echoValueAlpha()
	})
	assert.Contains(t, result, contains, "When variable Alpha is not empty, it should display its value.")
}

func Test_echoValueBeta(t *testing.T) {
	var contains string
	var result string

	// Beta variable is empty (--beta flag wasn't provided)
	contains = "The value of Beta: false (Option not provided)"
	result = capturer.CaptureStdout(func() {
		echoValueBeta()
	})
	assert.Contains(t, result, contains, "When variable Beta is false (default), it should display that option wasn't provided.")

	// Beta variable is true (--beta flag was provided)
	Beta = true
	contains = "The value of Beta: true"
	result = capturer.CaptureStdout(func() {
		echoValueBeta()
	})
	assert.Contains(t, result, contains, "When variable Beta is true, it should display it's true.")
}
