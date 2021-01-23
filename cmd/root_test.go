/*
Package cmd root
*/
package cmd

import (
	"errors"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestEchoStdErrIfError(t *testing.T) {
	var expectStatus int = 1
	var expectMsg string = "foo bar"
	var errorMsg error = errors.New(expectMsg)

	var actualStatus int = 0 // This should turn into 1
	var actualMsg string = capturer.CaptureStderr(func() {
		actualStatus = EchoStdErrIfError(errorMsg)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg+"\n", actualMsg, "The two words should be the same.")
}

func TestEchoStdErrIfErrorIsNil(t *testing.T) {
	var expectStatus = 0
	var actualStatus = 1 // This should turned into 0

	var expectMsg = ""
	var actualMsg = capturer.CaptureStderr(func() {
		actualStatus = EchoStdErrIfError(nil)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg, actualMsg, "The two words should be the same.")
}

func TestExecute(t *testing.T) {
	var actual = capturer.CaptureStdout(func() {
		Execute()
	})
	var contains = "Simple CLI app to see how Cobra works to create commands"
	assert.Contains(t, actual, contains, "When no arg, should return help message.")
}
