package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloCmd(t *testing.T) {
	var helloCmd = createHelloCmd()
	var buffTmp = new(bytes.Buffer)
	var argsTmp = []string{"hello"}

	helloCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	helloCmd.SetArgs(argsTmp) // set command args

	helloCmd.ExecuteC() // Run `hello` command!

	var expect = "Hello, world!\n"
	var actual = buffTmp.String() // resotre buffer
	assert.Equal(t, expect, actual, "The two words should be the same.")
}
