/*
Package cmd hello is a sub-command that echoes simply "Hello, world!".
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createHelloCmd creates/generates an instance of hello command.
// Generator function eases unit testing of a command rather than defining
// anonymous function in to a variable
func createHelloCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hello",
		Short: "Says hello to the world.",
		Long: `About:
  It will return "Hello, world!".
	  `,
		// One of the best practices in Cobra is to use `RunE` instead of `Run`
		// and return `error` only if an error occures. That will ease testing.
		// In that manner try not to use `os.Exit()` or `panic()` in the child
		// commands but return error instead and let the main package handle it.
		RunE: func(cmd *cobra.Command, args []string) error {
			return sayHello(cmd)
		},
	}

	// Return the created command.
	return cmd
}

// init runs on app initialization. Regardless of whether sub command was specified
// or not.
func init() {
	// Create and add "hello" command as a child of the root command(rootCmd).
	rootCmd.AddCommand(createHelloCmd())
}

// sayHello is the main function of "hello"(helloCmd).
func sayHello(cmd *cobra.Command) error {
	// Outputs "Hello, world!".
	// We use `cobra.Command`'s `fmt.Println` wrapper to ease testing. Which can
	// be changed it's output. See: hello_test.sh
	cmd.Println("Hello, world!")

	return nil
}
