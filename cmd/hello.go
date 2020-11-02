/*
Package cmd hello is a sub-command that echoes simply "Hello, world!".
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Says hello to the world.",
	Long: `About:
  It will return "Hello, world!".
  `,
	// One of the best practice in Cobra is to use RunE instead of
	// Run and return error if an error occures.
	Run: func(cmd *cobra.Command, args []string) {
		sayHello()
	},
}

func init() {
	// Add/register helloCmd as a child command to the root command (rootCmd)
	rootCmd.AddCommand(helloCmd)
}

// sayHello is the main function of "hello"(helloCmd)
func sayHello() {
	Echo("Hello, world!")
}
