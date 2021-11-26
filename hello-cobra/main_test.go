//nolint:testpackage // To override osExit the test needs to be part of main
package main

import (
	"runtime/debug"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Tests for main's behavior
// ----------------------------------------------------------------------------

func Test_main(t *testing.T) {
	// Run main() and capture STDOUT message and its exit-status-code
	out := capturer.CaptureOutput(func() {
		main()
	})

	// Assertion
	expectMsg := "Usage:" // Help msg includes this string

	assert.Contains(t, out, expectMsg, "it should contain help message if no args")
}

func Test_unknown_command(t *testing.T) {
	// Mock the "OsExit" to capture the exit-status-code
	actualExitCode := 0 // this shuld turn to 1
	expectExitCode := 1

	recoverOsExit := MockOsExit(t, &actualExitCode)
	defer recoverOsExit()

	// Mock args
	recoverOsArgs := MockOsArgs(t, []string{
		"foobar", // unknown command
	})
	defer recoverOsArgs()

	// Capture STDERR
	out := capturer.CaptureStderr(func() {
		main()
	})

	assert.Contains(t, out, "Error: unknown command \"foobar\"")
	require.Equal(t, expectExitCode, actualExitCode,
		"unknown command should exit with status 1")
}

func Test_version_flag(t *testing.T) {
	// Mock args.
	recoverOsArgs := MockOsArgs(t, []string{
		"--version",
	})
	defer recoverOsArgs()

	out := capturer.CaptureOutput(func() {
		main()
	})

	assert.Contains(t, out, "hello-cobra version (devel)")
}

// ----------------------------------------------------------------------------
//  Function Test
// ----------------------------------------------------------------------------

func TestGetVersion_version_not_set(t *testing.T) {
	// Backup and restore Version before and after the test
	oldVersion := Version
	defer func() {
		Version = oldVersion
	}()

	Version = "" // Mock Version as empty

	// Backup and restore DebugReadBuildInfo
	oldDebugReadBuildInfo := DebugReadBuildInfo
	defer func() {
		DebugReadBuildInfo = oldDebugReadBuildInfo
	}()

	// Mock debug.BuildInfo to return false
	DebugReadBuildInfo = func() (info *debug.BuildInfo, ok bool) {
		return nil, false
	}

	expect := "(unknown)"
	actual := GetVersion()

	assert.Equal(t, expect, actual)
}

func TestGetVersion_version_via_build_info(t *testing.T) {
	// Backup and recover Version before and after the test
	oldVersion := Version
	defer func() {
		Version = oldVersion
	}()

	Version = "" // Mock Version as empty

	oldDebugReadBuildInfo := DebugReadBuildInfo
	defer func() {
		DebugReadBuildInfo = oldDebugReadBuildInfo
	}()

	expect := "foobar" // Mock build info's version

	// Mock DebugReadBuildInfo
	DebugReadBuildInfo = func() (info *debug.BuildInfo, ok bool) {
		i := &debug.BuildInfo{
			Main: debug.Module{
				Version: expect,
			},
		}

		return i, true
	}

	assert.Equal(t, expect, GetVersion())
}

func TestGetVersion_version_via_build_flag(t *testing.T) {
	// Backup and recover Version before and after the test
	oldVersion := Version
	defer func() {
		Version = oldVersion
	}()

	Version = "v0.0.0-" + t.Name()

	expect := Version
	actual := GetVersion()

	assert.Equal(t, expect, actual, "it should return the main.Version value if set")
}
