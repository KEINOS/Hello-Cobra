package cmdroot

import (
	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello"
	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
//  Commnad Struct
// ----------------------------------------------------------------------------

// Command is the struct to hold cobra.Command and it's flag options.
type Command struct {
	*cobra.Command
}

// ----------------------------------------------------------------------------
//  Public Functions
// ----------------------------------------------------------------------------

// New returns a new "root" command object.
//
// The argument appVersion will be used when "--version" option was called.
func New(appVersion string) *cobra.Command {
	if appVersion == "" {
		appVersion = "(unknown)"
	}

	cmdRoot := &Command{
		&cobra.Command{
			// The first word in "Use:" will be used as a command name.
			Use: "hello-cobra",
			// One-line description of the app. This will not appear on command help
			// but in the generated document.
			Short: "A sample of Cobra usage.",
			// Detailed description of the app.
			Long: `About:
  A simple CLI app to see how Cobra works.`,
			Example: `
  Hello-Cobra hello
  Hello-Cobra hello ext --reverse foo bar

  Hello-Cobra hello --help
  Hello-Cobra help hello
  Hello-Cobra help hello ext`,
			Version: appVersion,
		},
	}

	cmdRoot.PersistentFlags().Bool("verbose", false, "displays debug info if any")

	// Add child commands to the "root" command.
	cmdRoot.AddCommand(
		cmdhello.New(), // Add "hello" command (with "world" grand child command)
	)

	return cmdRoot.Command
}
