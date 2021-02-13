package cmd

import (
	"strings"
	"testing"

	"github.com/KEINOS/Hello-Cobra/conf"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_loadConfig_UserDefinedPath_Fails(t *testing.T) {
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

	var capturedMsg string = capturer.CaptureStderr(func() {
		// Test user defined bad (non-existing) file path
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

	// exit code assertion
	assert.Equal(t, expectExitCode, actualExitCode,
		"If user defined path doesn't exist then should exit with 1. Captured STDERR: "+capturedMsg,
	)
	// containing error message assertion
	assert.Contains(t, strings.TrimSpace(capturedMsg), "Failed to read configuration file")
}

func Test_loadConfig_UseDefault(t *testing.T) {
	// Save current function in osExt
	oldOsExit := osExit
	// restore osExit at the end
	defer func() { osExit = oldOsExit }()

	var (
		expectExitCode int
		actualExitCode int
		expectFlag     bool
		actualFlag     bool

		confAppDummy  conf.TConfigApp
		confUserDummy struct {
			NameToGreet string `mapstructure:"name_to_greet"` // // Dont'f forget to define `mapstructure`
		}
	)

	// Assign mock of "osExit" to capture the exit-status-code.
	osExit = func(code int) {
		actualExitCode = 0 // If PathFileConf is empty then should not reach here.
	}

	var capturedMsg string = capturer.CaptureStderr(func() {
		// Test app defined non-existing file path
		confAppDummy = conf.TConfigApp{
			PathFileConf: "",
			PathDirConf:  ".",
			NameFileConf: "dummy_config",
			NameTypeConf: "json",
		}
		confUserDummy.NameToGreet = "bar"
		actualExitCode = 1
		expectExitCode = 1
		loadConfig(&confAppDummy, &confUserDummy)
	})

	// Exit code assertion
	assert.Equal(t, expectExitCode, actualExitCode,
		"If app defined conf file does not exist and using default then should not call 'osExit()'."+
			"Captured STDERR:"+capturedMsg,
	)
	// Default flag assertion
	expectFlag = true
	actualFlag = confAppDummy.IsUsingDefaultConf
	assert.Equal(t, expectFlag, actualFlag,
		"Property 'TConfigApp.IsUsingDefaultConf' should be true when using default.(Only when user didn't define)",
	)
}
