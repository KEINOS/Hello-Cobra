package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloCmd(t *testing.T) {
	var helloCmd = createHelloCmd()
	var argsTmp = []string{}
	var buffTmp = new(bytes.Buffer)

	helloCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	helloCmd.SetArgs(argsTmp) // set command args

	// Run `hello` command!
	if err := helloCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'helloCmd.Execute()'.", "Error msg: %v", err)
	}

	var expect = "Hello, world!\n"
	var actual = buffTmp.String() // resotre buffer
	assert.Equal(t, expect, actual,
		"Command 'hello' should return 'Hello, world!'.",
	)
}

func Test_helloCmd_Help(t *testing.T) {
	var helloCmd = createHelloCmd()
	var argsTmp = []string{"--help"}
	var buffTmp = new(bytes.Buffer)

	helloCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	helloCmd.SetArgs(argsTmp) // set command args

	// Run `hello` command!
	if err := helloCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'helloCmd.Execute()'.", "Error msg: %v", err)
	}

	var contains = "'hello' is a command that simply displays the \"Hello, world!\"."
	var result = buffTmp.String() // resotre buffer
	assert.Contains(t, result, contains, "The command didn't include the required help message for 'hello' in it's result.")
}
