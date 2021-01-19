package cmd

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_helloCmd(t *testing.T) {
	var expect = "Hello, world!\n"
	var actual = capturer.CaptureStdout(func() {
		var cmdTmp = &cobra.Command{}
		var argsTmp = []string{}
		helloCmd.Run(cmdTmp, argsTmp)
	})
	assert.Equal(t, expect, actual, "The two words should be the same.")
}

func Test_sayHello(t *testing.T) {
	var expect = "Hello, world!\n"
	var actual = capturer.CaptureStdout(func() {
		sayHello()
	})

	assert.Equal(t, expect, actual, "The two words should be the same.")
}
