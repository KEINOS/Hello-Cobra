/*
Package cmd flag is a sub-command that detects flags and it's value from the arguments.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

/*
-------------------------------------------------------------------------------
 Define global variables.

 These variables may be defined as local. But mostly option/flag arguments are
 referenced through out the application, so we set here as global to let the
 other packages/functions to reference them.
-------------------------------------------------------------------------------
*/

// Alpha stores the given value of the "--alpha" "-a" flag option. Default: empty
var Alpha string

// Beta stores the bool state. Default: false, with the '--beta' flag: true.
var Beta bool

/*
-------------------------------------------------------------------------------
 Define local variables
-------------------------------------------------------------------------------
*/

// flagCmd represents the flag command
var flagCmd = &cobra.Command{
	Use:   "flag",
	Short: "Displays the internal value according to the flags provided.",
	Long: `About:
	This command detects the flag in the argument, and outputs it's internal value.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		echoFlagStatuses(cmd, args)
	},
}

/*
-------------------------------------------------------------------------------
 Define local functions
-------------------------------------------------------------------------------
*/

// echoFlagStatuses is the actual body function of flagCmd.
func echoFlagStatuses(cmd *cobra.Command, args []string) {
	echoValueAlpha()
	echoValueBeta()
}

func echoValueAlpha() {
	if Alpha == "" {
		Echo("The value of Alpha: '' (empty. Option not provided)")
	} else {
		Echo("The value of Alpha: " + Alpha)
	}
}

func echoValueBeta() {
	if Beta {
		Echo("The value of Beta: true")
	} else {
		Echo("The value of Beta: false (Option not provided)")
	}
}

// Initialization when `cmd.Execute()` was called.
func init() {
	// Add flag options of this command:
	// * The flag --alpha/-a will store it's value into "Alpha" global variable.
	flagCmd.Flags().StringVarP(&Alpha, "alpha", "a", "", "Displays the given string of option-value. A value is required.")
	// * The flag --beta/-b will store it's bool state into "Beta" global variable.
	flagCmd.Flags().BoolVarP(&Beta, "beta", "b", false, "Sets the 'beta' boolean flag to 'true'. (Default: false)")

	// Add this command to the parent(rootCmd)
	rootCmd.AddCommand(flagCmd)
}
