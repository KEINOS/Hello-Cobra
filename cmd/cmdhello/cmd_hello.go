/*
Package cmdhello defines the "hello" command.
*/
package cmdhello

import (
	"fmt"
	"strings"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello/cmdworld"
	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
//  Commnad Struct
// ----------------------------------------------------------------------------

// Command is the struct to hold cobra.Command and it's flag options.
type Command struct {
	*cobra.Command
	isReverse bool // flag for "--reverse" option
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns the pointer of the "hello" command's singleton object.
func New() *cobra.Command {
	cmdHello := &Command{
		&cobra.Command{
			Use:   "hello [name [name] ...]",
			Short: "Greets to the given arg.",
			Long: `About:
		  'hello' prints a greeting message to the given name in the arguments.`,
			Example: `
		  Hello-Cobra hello foo bar           // Hello, foo and bar!
		  Hello-Cobra hello --reverse foo bar // !rab dna oof ,olleH
		`,
		},
		false,
	}

	cmdHello.Command.RunE = cmdHello.sayHelloTo

	// Define flags for `ext` command.
	cmdHello.Flags().BoolVarP(
		&cmdHello.isReverse, "reverse", "r", false, "Reverses the output.",
	)

	// Add child command "world" to "hello".
	cmdHello.AddCommand(cmdworld.New())

	return cmdHello.Command
}

// ----------------------------------------------------------------------------
//  Private Functions
// ----------------------------------------------------------------------------

// reverseString reverses/flip the input string.
func reverseString(input string) string {
	var msgTmp string

	for _, v := range input {
		msgTmp = string(v) + msgTmp
	}

	return msgTmp
}

// sayHelloTo is the main function of "hello" command.
// It will print the greetings to the given args.
func (c *Command) sayHelloTo(cmd *cobra.Command, args []string) error {
	msgToGreet := "Hi!"

	if len(args) > 0 {
		names := strings.Join(args, " and ")
		msgToGreet = "Hello, " + names + "!"
	}

	if c.isReverse {
		msgToGreet = reverseString(msgToGreet)
	}

	fmt.Println(msgToGreet)

	return nil
}
