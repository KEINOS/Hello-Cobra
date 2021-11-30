/*
Package main.
*/
package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdroot"
	"golang.org/x/mod/semver"
)

// FAILURE is an alias of exit status code to ease read.
const FAILURE = 1

// OsExit is a copy of `os.Exit` to ease mocking during test.
//     Ref: https://stackoverflow.com/a/40801733/8367711
//
// We found useful that `os.Exit` should be called only in the main package.
// And functions or methods in other packages should return an error rather than
// exiting the app. Doing so, it's a lot easier to test and trace.
var OsExit = os.Exit

// Version info via `-ldflags`.
// These values should be set via `-ldflags` option during build.
//
// Sample command to build:
//     $ VER_APP="$(git describe --tag)" // v1.0.0-alpha-g65a8e0c
//     $ go build -ldflags="-X 'main.version=${VER_APP}'" -o foo ./hello-cobra/
//     $ ./foo --version
//     hello-cobra version 1.0.0-alpha (g65a8e0c)
var (
	// The application version to display.
	version string
	// The short hash of the commit. This value will be used as build ver as well.
	commit string
)

// DebugReadBuildInfo is a copy of debug.ReadBuildInfo to ease testing.
var DebugReadBuildInfo = debug.ReadBuildInfo

// ----------------------------------------------------------------------------
//  Main
// ----------------------------------------------------------------------------

func main() {
	// Initialize the app then executes the designated command.
	if err := cmdroot.New(GetVersion()).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		OsExit(FAILURE)
	}
}

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// GetVersion returns the app version without the "v" prefix.
func GetVersion() string {
	result := ""

	// Version via commit tag on `go install`
	if buildInfo, ok := DebugReadBuildInfo(); ok {
		result = buildInfo.Main.Version
	}

	// Version via build flag
	if version != "" {
		result = version
	}

	return parseVersion(result)
}

func parseVersion(input string) string {
	if input == "" {
		return "(unknown)"
	}

	build := commit

	// Convert v1.2.3 (12345) --> v1.2.3+12345
	result := strings.ReplaceAll(input, " ", "")
	result = strings.ReplaceAll(result, "(", "+")
	result = strings.ReplaceAll(result, ")", "")
	result = strings.TrimLeft(result, "v. ") // Trim prefixes and add later

	if result[0] != 'v' {
		result = "v" + result // Add 'v' prefix for Go semver parse compatibility
	}

	// Convert git tag's format to Go semver style.
	// v1.2.3-4-5678 --> v1.2.3-4+5678
	if preCount := strings.Count(result, "-"); preCount > 1 {
		result = strings.ReplaceAll(result, "-", "+")
		result = strings.Replace(result, "+", "-", preCount-1)
	}

	// Get meta data as commit hash
	if build == "" {
		build = strings.ReplaceAll(semver.Build(result), "+", "")
	}

	// Trim meta data
	if result = semver.Canonical(result); result == "" {
		// Return raw input without "v" prefix if not Go semver compatible format.
		return strings.TrimLeft(input, "v. ")
	}

	if build != "" {
		build = fmt.Sprintf(" (%v)", build)
	}

	return fmt.Sprint(strings.TrimLeft(result, "v. "), build)
}
