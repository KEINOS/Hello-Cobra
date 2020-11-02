/*
Package main
*/
package main

import (
	"os"

	"github.com/KEINOS/Hello-Cobra/cmd"
)

// https://stackoverflow.com/a/40801733/8367711
var osExit = os.Exit

func main() {
	// `cmd.Execute()` initializes Cobra such as adding sub-commands to the root
	// command and executes the designated command.
	// If `cmd.Execute()` returns an error then it will print the error message
	// to STDERR and exits with 1.
	osExit(cmd.EchoIfError(cmd.Execute()))
}
