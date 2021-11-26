/*
Package main.
*/
package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdroot"
)

// FAILURE is an alias of exit status code to ease read.
const FAILURE = 1

// OsExit is a copy of `os.Exit` to ease mocking during test.
//     Ref: https://stackoverflow.com/a/40801733/8367711
//
// We found useful that `os.Exit` should be called only in main package.
// And functions or methods in other packages should return an error rather than
// exiting the app. To do so, it was way easy to test them.
var OsExit = os.Exit

// Version is the application version to display. This value should be set via
// `-ldflags` during build.
//
// Sample command to build:
//
//     $ VER_APP="v1.0.0"
//     $ go build -ldflags="-X 'main.Version=${VER_APP}'" -o foo ./hello-cobra/
//     $ ./foo --version
//     hello-cobra version v1.0.0
var Version string

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

// GetVersion returns the app version.
func GetVersion() string {
	// Via build flag
	if Version != "" {
		return Version
	}

	// Via commit tag
	if buildInfo, ok := DebugReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}

	return "(unknown)"
}
