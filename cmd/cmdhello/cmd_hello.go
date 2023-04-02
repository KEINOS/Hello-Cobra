/*
Package cmdhello defines the "hello" command.
*/
package cmdhello

import (
	"fmt"
	"strings"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello/cmdworld"
	"github.com/MakeNowJust/heredoc"
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

// New returns the newly created object pointer of the "hello" command.
func New() *cobra.Command {
	command := new(cobra.Command)

	command.Use = "hello [name [name] ...]"
	command.Short = "Greets to the given arg."
	command.Long = heredoc.Doc(`
		About:
			'hello' prints a greeting message to the given name in the arguments.
	`)
	command.Example = heredoc.Doc(`
		Hello-Cobra hello foo bar           // Hello, foo and bar!
		Hello-Cobra hello --reverse foo bar // !rab dna oof ,olleH
	`)

	// Instantiate new object
	cmdHello := &Command{
		command,
		false,
	}

	// Assign the method as RunE function
	cmdHello.Command.RunE = cmdHello.sayHelloTo

	// Define flags for `hello` command.
	cmdHello.Flags().BoolVarP(
		&cmdHello.isReverse, "reverse", "r", false, "Reverses the output.",
	)

	// Add child command "world" to "hello".
	cmdHello.AddCommand(cmdworld.New())

	return cmdHello.Command
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// sayHelloTo prints the greetings to the given args. It is the main function of
// "hello" command.
func (c *Command) sayHelloTo(_ *cobra.Command, args []string) error {
	msgToGreet := "Hi!"

	if len(args) > 0 {
		names := strings.Join(args, " and ")
		msgToGreet = "Hello, " + names + "!"
	}

	if c.isReverse {
		msgToGreet = reverseString(msgToGreet)
	}

	//nolint:forbidigo // fmt.Println is OK for CLI
	fmt.Println(msgToGreet)

	return nil
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
