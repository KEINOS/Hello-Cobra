/*
Package cmdworld defines the "world" command which is the child of "hello" command.
*/
package cmdworld

import (
	"fmt"

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
	// Instantiate new object
	cmdWorld := &Command{
		&cobra.Command{
			Use:   "world",
			Short: "Says hello to the world.",
			Long: `About:
		  'world' is a sub command of 'hello' which displays "Hello, world!".`,
			Example: `
		  Hello-Cobra hello world           // Hello, world!
		  Hello-Cobra hello world --reverse // !dlrow ,olleH
			`,
		},
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
func (c *Command) sayHelloWorld(cmd *cobra.Command, args []string) error {
	msgToGreet := "Hello, world!"

	if c.isReverse {
		msgToGreet = reverseString(msgToGreet)
	}

	fmt.Println(msgToGreet)

	return nil
}
