/*
Package cmd cmd_root.go defines `root` command, the mother command of all the commands.

Define common things here. Also note that each child(sub commands) of root command
should be added/registered themself in their `init()` via `rootCmd.AddCommand()`.
See other command's `init()` func.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KEINOS/Hello-Cobra/conf"
	"github.com/spf13/cobra"
)

// ===============================================================================
//  Constants
// ===============================================================================

const (
	// SUCCESS is an alias of exit status code to ease read.
	SUCCESS int = 0
	// FAILURE is an alias of exit status code to ease read.
	FAILURE int = 1
)

// ===============================================================================
//  Application Settings
// ===============================================================================

// TypeConfUser defines the data structure to store values from a conf file. Viper
// will read these values from the config file or env variables. `mapstructure`
// defines the key name in the conf file.
//   Ex) `{"name_to_greet": "foo"}`  will be `NameToGreet = "foo"```
type TypeConfUser struct {
	NameToGreet string `mapstructure:"name_to_greet"`
}

var (
	// ConfApp is the basic app settings.
	ConfApp = conf.TypeConfigApp{
		PathDirConf:        ".",
		NameFileConf:       "config",
		NameTypeConf:       "json",
		PathFileConf:       "",    // User defined file path
		IsUsingDefaultConf: false, // Set to true if conf file not fond
	}

	// ConfUser holds the values read from the config file. The values here are
	// the default.
	ConfUser = TypeConfUser{
		NameToGreet: "", // Conf for `hello` and `hello ext` command.
	}
)

// osExit is a copy of `os.Exit` to ease the "exit status" test.
// See: https://stackoverflow.com/a/40801733/8367711
var osExit = os.Exit

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
  A simple CLI app to see how Cobra works to create commands.`,
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

// init runs on app initialization.
// Regardless of whether a command was specified or not.
func init() {}

// loadConfig sets the object in the arg with the results exits with an error if user defined conf file didn't exist.
// Otherwise searches the default file and if not found then use the default value.
func loadConfig(configApp *conf.TypeConfigApp, configUser interface{}) {
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

// ============================================================================
//  Exported functions of cmd package. (https://tour.golang.org/basics/3)
//
//  Define functions that were used outside this package such as `main` pkg.
//  These will work as `cmd` method.
//    Ex) cmd.EchoStdErrIfError(err)
// ============================================================================

// EchoStdErrIfError is an STDERR wrappter and returns 0(zero) or 1.
// It does nothing if the error is nil and returns 0.
func EchoStdErrIfError(err error) int {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return FAILURE
	}

	return SUCCESS
}

// Execute is the main function of `cmd` package.
// It adds all the child commands to the root's command tree and sets their flag
// settings. Then runs/executes the `rootCmd` to find appropriate matches for child
// commands with corresponding flags and args.
//
// Usually `cmd.Execute` will be called by the `main.main()` and it only needs to
// happen once to the rootCmd.
// Returns `error` when it fails to execute.
func Execute() error {
	// Read conf file values to ConfUser with ConfApp settings
	return rootCmd.Execute()
}
