/*
Package cmd flag is a sub-command that detects flags and it's value from the arguments.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flagCmd represents the flag command
var flagCmd = &cobra.Command{
	Use:   "flag",
	Short: "Displays the value of the flag.",
	Long: `About:
	This command detects the flag in the argument and outputs it's value.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd)
	},
}

// Alpha stores the value of the flag
var Alpha string

func init() {
	// Add flag options of this command

	// --source/-s Stores it's value to "Alpha" global variable.
	flagCmd.Flags().StringVarP(&Alpha, "alpha", "a", "", "Displays the given option value string.")

	// Add this command to parent(rootCmd)
	rootCmd.AddCommand(flagCmd)
}
