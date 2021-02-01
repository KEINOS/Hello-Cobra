/*
Package main main.go is the entry point to runs/execute the package "cmd".
And it exits the app with the staus code according to the command's result.
*/
package main

import (
	"os"

	"github.com/KEINOS/Hello-Cobra/cmd"
)

// osExit is a copy of `os.Exit` to ease the "exit status" test.
// See: https://stackoverflow.com/a/40801733/8367711
var osExit = os.Exit

func main() {
	// `cmd.Execute()` initializes the app such as adding sub-commands to the root command.
	// Then executes the designated command.
	// If `cmd.Execute()` returns an error then it will print the error message to STDERR and exits with 1.
	osExit(cmd.EchoStdErrIfError(cmd.Execute()))
}
