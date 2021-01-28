/*
Package cmd root.go defines `root` command, the mother command of all the commands.

Define common things here. Also note that each child(sub commands) of root command
should be added/registered themself in their `init()` via `rootCmd.AddCommand()`.
See other command's `init()` func.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ===============================================================================
//  Constants
// ===============================================================================
const (
	// Exit status to ease read
	SUCCESS int = 0
	FAILURE int = 1
	// Default property values of the app
	NAME_FILE_CONF_DEFAULT = "config"
	NAME_EXT_CONF_DEFAULT  = "json"
)

// ===============================================================================
//  Root command
// ===============================================================================

// rootCmd is the mother command of it's child. Each child/sub command must
// register themself in their `init()`. See: hello.go
var rootCmd = createRootCmd()

// createRootCmd creates the `root` command.
func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		// The first word in "Use:" will be used as a command name.
		Use: "Hello-Cobra",
		// One-line description of the app. This will not appear on command help
		// but in the generated document.
		Short: "A sample of Cobra usage.",
		// Detailed description of the app.
		Long: `About:
  A simple CLI app to see how Cobra works to create commands.`,
		Example: `
  Hello-Cobra hello
  Hello-Cobra hello ext --reverse foo bar

  Hello-Cobra hello --help
  Hello-Cobra help hello
  Hello-Cobra help hello ext`,
	}

	// Additional flags of root command.
	// Flags() is the drop-in replacement for Go's flag package to implement
	// POSIX/GNU-style --flags. For the detailed usage see:
	//   https://github.com/spf13/pflag
	cmd.Flags().BoolP(
		"toggle",
		"t",
		false,
		"A flag for the main command but does nothing.",
	)

	return cmd
}

// init runs on app initialization.
// Regardless of whether a command is specified or not.
func init() {}

// ============================================================================
//  Common Functions.
// ============================================================================
// ----------------------------------------------------------------------------
//  Local functions.
//  Define functions that were used throughout the whole "cmd" package.
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
//  Exported functions of cmd package. ( https://tour.golang.org/basics/3 )
//  Define functions that were used outside this package such as `main` pkg.
//  These will work as `cmd` method.
//    Ex) cmd.EchoStdErrIfError(err)
// ----------------------------------------------------------------------------

// EchoStdErrIfError is an STDERR wrappter.
// Returns 0 for success and 1 for failure to let main.go take controll to exit.
// Note that its input arg type is `error` and not `string`.
func EchoStdErrIfError(msg error) int {
	if msg != nil {
		fmt.Fprintln(os.Stderr, msg)
		return FAILURE
	}
	return SUCCESS
}

// Execute adds all child commands to the root command and sets flags appropriately.
// Usualy `Execute` will be called by the `main.main()` and it only needs to happen
// once to the rootCmd.
// Returns `error` when it fails to execute.
func Execute() error {
	return rootCmd.Execute()
}
