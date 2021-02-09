package cmd

import (
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

		confAppDummy  conf.TConfigApp
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
		confAppDummy = conf.TConfigApp{
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
