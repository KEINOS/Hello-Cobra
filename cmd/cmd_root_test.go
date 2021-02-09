/*
Package cmd cmd_root_test.go
*/
package cmd

import (
	"errors"
	"testing"

	"github.com/KEINOS/Hello-Cobra/conf"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_loadConfigFail(t *testing.T) {
	// Save current function in osExt
	oldOsExit := osExit
	// restore osExit at the end
	defer func() { osExit = oldOsExit }()

	var (
		expectExitCode int
		actualExitCode int = 0 // This should turn into 1

		confAppDummy  conf.TypeConfigApp
		confUserDummy struct {
			NameToGreet string `mapstructure:"name_to_greet"` // // Dont'f forget to define `mapstructure`
		}
	)

	// Assign mock of "osExit" to capture the exit-status-code.
	osExit = func(code int) {
		actualExitCode = 1
	}

	var capturedMsg string = capturer.CaptureStdout(func() {
		// Test user defined bad file path
		confAppDummy = conf.TypeConfigApp{
			PathFileConf: "./foobar.json",
			PathDirConf:  "",
			NameFileConf: "",
			NameTypeConf: "",
		}
		confUserDummy.NameToGreet = "bar"
		expectExitCode = 1
		loadConfig(&confAppDummy, &confUserDummy)
	})

	// Assertion
	assert.Equal(t, expectExitCode, actualExitCode, "Msg: "+capturedMsg)
}

func TestEchoStdErrIfError(t *testing.T) {
	var (
		expectStatus int = 1
		actualStatus int
		expectMsg    string = "foo bar"
		actualMsg    string
		errorMsg     error = errors.New(expectMsg)
	)

	// Run the function and capture the STDERR msg and it's returned int value.
	actualMsg = capturer.CaptureStderr(func() {
		actualStatus = 0 // This should turn into 1
		actualStatus = EchoStdErrIfError(errorMsg)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg+"\n", actualMsg, "The two words should be the same.")
}

func TestEchoStdErrIfError_IsNil(t *testing.T) {
	var (
		expectStatus int    = 0
		actualStatus int    = 1 // This should turn into 0
		expectMsg    string = ""
		actualMsg    string
	)

	actualMsg = capturer.CaptureStderr(func() {
		actualStatus = EchoStdErrIfError(nil)
	})

	assert.Equal(t, expectStatus, actualStatus, "Error code should return 1")
	assert.Equal(t, expectMsg, actualMsg,
		"If the arg is nil it should not print enything to STDERR. STDERR Msg: "+actualMsg,
	)
}

func TestExecute(t *testing.T) {
	var (
		result   string
		contains string
	)

	result = capturer.CaptureStdout(func() {
		if err := Execute(); err != nil {
			assert.FailNowf(t, "Failed to execute 'root.Execute()'.", "Error msg: %v", err)
		}
	})
	contains = "A simple CLI app to see how Cobra works to create commands."

	assert.Contains(t, result, contains, "When no arg, should return help message.")
}
