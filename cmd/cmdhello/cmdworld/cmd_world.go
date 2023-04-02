/*
Package cmdworld defines the "world" command which is the child of "hello" command.
*/
package cmdworld

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
//  Commnad Struct
// ----------------------------------------------------------------------------

// Command is the struct to hold cobra.Command and it's flag options.
type Command struct {
	*cobra.Command
	isReverse bool // flag value for "--reverse" option
}

// ----------------------------------------------------------------------------
//  Public Functions
// ----------------------------------------------------------------------------

// New returns the newly created object pointer of the "world" command.
func New() *cobra.Command {
	command := new(cobra.Command)

	command.Use = "world"
	command.Short = "Says hello to the world."
	command.Long = heredoc.Doc(`
		About:
			'world' is a sub command of 'hello' which displays "Hello, world!".
	`)
	command.Example = heredoc.Doc(`
		Hello-Cobra hello world           // Hello, world!
		Hello-Cobra hello world --reverse // !dlrow ,olleH
	`)

	// Instantiate new object
	cmdWorld := &Command{
		command,
		false,
	}

	// Add method to RunE
	cmdWorld.RunE = cmdWorld.sayHelloWorld

	// Define flags for `world` command.
	cmdWorld.Flags().BoolVarP(
		&cmdWorld.isReverse, "reverse", "r", cmdWorld.isReverse, "Reverses the output.",
	)

	return cmdWorld.Command
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

// sayHelloWorld is the main function of "world" command.
// It prints "Hello, world!" or if "--reverse" flag was set then prints
// "!dlrow ,olleH".
func (c *Command) sayHelloWorld(_ *cobra.Command, _ []string) error {
	msgToGreet := "Hello, world!"

	if c.isReverse {
		msgToGreet = reverseString(msgToGreet)
	}

	//nolint:forbidigo // fmt.Println is OK for CLI
	fmt.Println(msgToGreet)

	return nil
}
