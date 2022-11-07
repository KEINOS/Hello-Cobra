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

// New returns the newly created object pointer of the "root" command.
//
// The argument appVersion will be used when "--version" option was called. If
// empty it will auto-detect.
func New(appVersion string) *cobra.Command {
	if appVersion == "" {
		appVersion = "(unknown)"
	}

	cmdRoot := &Command{
		//nolint:exhaustruct // We don't need to set all the fields.
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
  $ hello-cobra --version

  $ hello-cobra hello
  $ hello-cobra hello --reverse
  $ hello-cobra hello foo bar
  $ hello-cobra hello foo bar --reverse

  $ hello-cobra hello world
  $ hello-cobra hello world --reverse

  $ hello-cobra completion --help
  $ hello-cobra completion bash > hello-cobra_completion.sh`,
			Version: appVersion,
		},
	}

	// Define persistent flags for the app.
	cmdRoot.PersistentFlags().Bool("verbose", false, "displays debug info if any")

	// Add child commands to the "root" command.
	cmdRoot.AddCommand(
		cmdhello.New(), // Add "hello" command (with grand child command "world")
	)

	return cmdRoot.Command
}
