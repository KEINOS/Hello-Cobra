/*
Package cmd root
*/
package cmd

import (
	"errors"
	"testing"

	"github.com/KEINOS/Hello-Cobra/util"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_loadConfigFail(t *testing.T) {
	// Save current function and restore at the end
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var expectExitCode int
	var actualExitCode int = 0 // This should turn into 1
	var ConfAppDummy util.TypeConfigApp
	var ConfUserDummy struct {
		NameToGreet string `mapstructure:"name_to_greet"` // // Dont'f forget to define `mapstructure`
	}
	// Assign mock of "osExit" to capture the exit-status-code.
	osExit = func(code int) {
		actualExitCode = 1
	}

	var actualMsg = capturer.CaptureStdout(func() {
		// Test user defined bad file path
		ConfAppDummy = util.TypeConfigApp{
			PathFileConf: "./foobar.json",
			PathDirConf:  "",
			NameFileConf: "",
			NameTypeConf: "",
		}
		ConfUserDummy.NameToGreet = "bar"
		expectExitCode = 1
		loadConfig(&ConfAppDummy, &ConfUserDummy)
	})

	// Assertion
	assert.Equal(t, expectExitCode, actualExitCode, "Msg: "+actualMsg)
}

func TestEchoStdErrIfError(t *testing.T) {
	var expectStatus int = 1
	var actualStatus int
	var expectMsg string = "foo bar"
	var errorMsg error = errors.New(expectMsg)

	// Run the function and capture the STDERR msg and it's returned int value.
	var actualMsg string = capturer.CaptureStderr(func() {
		actualStatus = 0 // This should turn into 1
		actualStatus = EchoStdErrIfError(errorMsg)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg+"\n", actualMsg, "The two words should be the same.")
}

func TestEchoStdErrIfError_IsNil(t *testing.T) {
	var expectStatus = 0
	var actualStatus = 1 // This should turned into 0

	var expectMsg = ""
	var actualMsg = capturer.CaptureStderr(func() {
		actualStatus = EchoStdErrIfError(nil)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg, actualMsg,
		"If the arg is nil it should not print enything to STDERR. STDERR Msg: "+actualMsg,
	)
}

func TestExecute(t *testing.T) {
	var actual = capturer.CaptureStdout(func() {
		Execute()
	})
	var contains = "A simple CLI app to see how Cobra works to create commands."
	assert.Contains(t, actual, contains, "When no arg, should return help message.")
}
