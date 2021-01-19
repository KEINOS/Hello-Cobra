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

func TestEcho(t *testing.T) {
	var expect = "hoge"
	var actual = capturer.CaptureStdout(func() {
		Echo(expect)
	})

	assert.Equal(t, expect+"\n", actual, "The two words should be the same.")
}

func TestEchoIfError(t *testing.T) {
	var expectStatus = 1
	var actualStatus = 0 // This should turned into 1

	var expectMsg = "foo bar"
	var msg = errors.New(expectMsg)
	var actualMsg = capturer.CaptureStderr(func() {
		actualStatus = EchoIfError(msg)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg+"\n", actualMsg, "The two words should be the same.")
}

func TestEchoIfErrorIsNil(t *testing.T) {
	var expectStatus = 0
	var actualStatus = 1 // This should turned into 0

	var expectMsg = ""
	var actualMsg = capturer.CaptureStderr(func() {
		actualStatus = EchoIfError(nil)
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
