/*
Package cmd helloExtended.go defaines `ext` command which is an extended version
of `hello` with various features and their tests.

It's a sub command of `hello` (grand child of `root`) as well.
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// TypePropertyHelloExt defines the data structure to store option/flag values
// for `hello ext` command.
type TypePropertyHelloExt struct {
	greetTo   string
	isReverse bool
}

// propertyHello holds property values of `ext` command.
var propertyHello = &TypePropertyHelloExt{}

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// createHelloExtCmd creates the `ext` command. See the `init()`
func createHelloExtCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ext [args]", // command name and additional usage.
		Short: "Extended 'hello' command.",
		Long: `About:
  'ext' is a sub command of 'hello' which displays "Hello, world!" in various ways.
  `,
		Example: `
  Hello-Cobra hello ext
  Hello-Cobra hello ext foo

  Hello-Cobra hello ext -h
  Hello-Cobra hello ext --help

  Hello-Cobra hello ext --reverse foo bar
  `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHelloExt(cmd, args)
		},
	}

	// Define flags for `ext` command.
	// Such as: it's variable type, where to save, flag names (long and short),
	// it's default value and the help description.
	cmd.Flags().
		BoolVarP(
			&propertyHello.isReverse, "reverse", "r", false, "Reverses the output.",
		)
	cmd.Flags().
		StringVarP(
			&propertyHello.greetTo, "who", "w", "", "Sets who to greet.",
		)

	return cmd
}

// getMsgToGreet returns the regular greeting msg from the flag value and args.
func getMsgToGreet(args []string) string {
	var to = "world" // Default

	var greetTo = propertyHello.greetTo                        // get flag value
	var argJoined = strings.TrimSpace(strings.Join(args, " ")) // get arg value

	if "" != greetTo && "" != argJoined {
		to = greetTo + " and " + argJoined
	} else {
		if "" != greetTo {
			to = greetTo
		}

		if "" != argJoined {
			to = argJoined
		}
	}

	return "Hello, " + to + "!"
}

// init runs on app initialization.
// Regardless of whether a command is specified or not.
func init() {
	// Add `ext` command under `hello` as a child. We create and pass the command
	// directly to `AddCommand()` and not exposing as a variable like `hello` does
	// since `ext` won't have a child command.
	helloCmd.AddCommand(createHelloExtCmd())
}

// reverseString reverses/flip the input string
func reverseString(input string) string {
	var msgTmp string

	for _, v := range input {
		msgTmp = string(v) + msgTmp
	}

	return msgTmp
}

// runHelloExt is the main function of `ext`.
func runHelloExt(cmd *cobra.Command, args []string) error {

	var greetings = getMsgToGreet(args)

	if propertyHello.isReverse {
		greetings = reverseString(greetings)
	}

	// Output result
	cmd.Println(greetings)
	return nil
}