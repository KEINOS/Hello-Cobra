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

// createHelloExtCmd creates the `ext` command. See the `init()` function below.
func createHelloExtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ext [args]", // command name and additional usage.
		Short: "Extended 'hello' command.",
		Long: `About:
  'ext' is a sub command of 'hello' which displays "Hello, world!" in various ways.`,
		Example: `
  Hello-Cobra hello ext                   // Hello, world!
  Hello-Cobra hello ext foo               // Hello, foo!
  Hello-Cobra hello ext foo bar           // Hello, foo bar!
  Hello-Cobra hello ext --who foo bar     // Hello, foo and bar!
  Hello-Cobra hello ext --who "foo bar"   // Hello, foo bar!
  Hello-Cobra hello ext foo bar --reverse // !rab oof ,olleH

  Hello-Cobra hello ext -h
  Hello-Cobra hello ext --help`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHelloExt(cmd, args)
		},
	}

	// Define flags for `ext` command.
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
	var (
		to        = ConfUser.NameToGreet                       // Set default
		greetTo   = propertyHello.greetTo                      // get flag value
		argJoined = strings.TrimSpace(strings.Join(args, " ")) // get arg value
	)

	if greetTo != "" && argJoined != "" {
		to = greetTo + " and " + argJoined
	} else {
		if greetTo != "" {
			to = greetTo
		}

		if argJoined != "" {
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

// reverseString reverses/flip the input string.
func reverseString(input string) string {
	var msgTmp string

	for _, v := range input {
		msgTmp = string(v) + msgTmp
	}

	return msgTmp
}

// runHelloExt is the main function of `ext`.
func runHelloExt(cmd *cobra.Command, args []string) error {
	greetings := getMsgToGreet(args)

	if propertyHello.isReverse {
		greetings = reverseString(greetings)
	}

	// Output result
	cmd.Println(greetings)

	return nil
}
