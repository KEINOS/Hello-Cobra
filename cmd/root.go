/*
Package cmd rootCmd is a mother command of all it's child(subcommands).

Each subcommands should add(register) themselfs in their `init()` via
`rootCmd.AddCommand()`. See other command packages' `init()` func.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
===============================================================================
 Root command
===============================================================================
*/

// createRootCmd creates the command for root.
func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "Hello-Cobra",
		Short: "Sample command of Cobra usage.",
		Long: `About:
  Simple CLI app to see how Cobra works to create commands.`,
	}

	/* Additional flags of root command */
	// toggle flag does nothing.
	cmd.Flags().BoolP("toggle", "t", false, "Global command option but does nothing.")

	return cmd
}

// rootCmd is the mother command of it's child. Each child/sub command must
// register themself in their `init()`. See: hello.go
var rootCmd = createRootCmd()

// init runs on every app initialization regardless of whether sub command is
// specified or not.
func init() {
}

/*
===============================================================================
 Exported functions of cmd package. https://tour.golang.org/basics/3
===============================================================================
 Define functions that are used through out the `cmd` and `main` package here.
*/

// EchoStdErrIfError is an STDERR wrappter. Returns a bool to let main.go take
// controll to exit. Note that its input type is `error` and not `string`.
func EchoStdErrIfError(msg error) int {
	if msg != nil {
		fmt.Fprintln(os.Stderr, msg)
		return 1
	}
	return 0
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// Returns `error` when it fails to execute.
func Execute() error {
	return rootCmd.Execute()
}
