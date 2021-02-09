package cmd

import (
	"fmt"

	"github.com/KEINOS/Hello-Cobra/conf"
	"github.com/spf13/cobra"
)

// ===============================================================================
//  Root command
// ===============================================================================

// rootCmd is the mother command of all other commands.
// Each child/sub command must register themself in their `init()` to `rootCmd`.
//   - See how: hello.go
var rootCmd = createRootCmd()

// createRootCmd creates the `root` command.
func createRootCmd() *cobra.Command {
	cmdTmp := &cobra.Command{
		// The first word in "Use:" will be used as a command name.
		Use: "Hello-Cobra",
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
	}

	// OnInitialize appends the passed function to initializers to be run when
	// each command's `Execute` method was called after `init`.
	cobra.OnInitialize(func() {
		// Load user conf file if exists.
		loadConfig(&ConfApp, &ConfUser)
	})

	// Additinal global flags of the app.
	// These flags will be available to every command under root command.
	cmdTmp.PersistentFlags().StringVarP(
		&ConfApp.PathFileConf,
		"config",
		"c",
		"",
		"File path of config.",
	)

	// Additional flags of root command.
	// Flags() is the drop-in replacement for Go's flag package to implement
	// POSIX/GNU-style --flags. For the detailed usage see:
	//   https://github.com/spf13/pflag
	cmdTmp.Flags().BoolP(
		"toggle",
		"t",
		false,
		"A flag for the main command but does nothing.",
	)

	// Return the actual root command
	return cmdTmp
}

// init runs on app initialization. Regardless of whether a command was specified
// or not.
//
// NOTE: that each child(sub commands) of root command should be added/registered
// themself in their `init()` via `rootCmd.AddCommand()`.
// See other command's `init()` func.
func init() {}

// loadConfig sets the object in the arg with the results exits with an error if user defined conf file didn't exist.
// Otherwise searches the default file and if not found then use the default value.
func loadConfig(configApp *conf.TConfigApp, configUser interface{}) {
	// Overwrite "configUser" with conf file value if file found.
	if err := conf.LoadConfig(*configApp, &configUser); err != nil {
		// Exits if user defined conf file fails to read
		if "" != configApp.PathFileConf {
			msg := fmt.Errorf("Failed to read configuration file.\n  Error msg: %v", err)
			osExit(EchoStdErrIfError(msg))
		}
		// Conf file not found. Using default. Set flag to true.
		configApp.IsUsingDefaultConf = true
	}
}
