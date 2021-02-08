package cmd

import (
	"bytes"
	"testing"

	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type TDataProvider []struct {
	command  string // command args of the app
	expect   string // expected output
	msgError string // message to display when the test fails
}

// ----------------------------------------------------------------------------
//  Test functions
// ----------------------------------------------------------------------------

func Test_helloExtCmd_NoFlags(t *testing.T) {
	// test cases for default behavior
	var cases TDataProvider = TDataProvider{
		// no options
		{
			command: "hello ext", expect: "Hello, world!\n",
			msgError: "'Hello, world!' should return by default.",
		},
		// simple arg
		{
			command: "hello ext foo", expect: "Hello, foo!\n",
			msgError: "The arg value should be used as a name.",
		},
		{
			command: "hello ext foo bar", expect: "Hello, foo bar!\n",
			msgError: "The arg value should be used as a name.",
		},
	}

	runTestCasesForHelloCmdExt(t, cases)
}

func Test_helloExtCmd_WhoFlag(t *testing.T) {
	// test cases for who flag
	var cases TDataProvider = TDataProvider{
		{
			command: "hello ext -w foo", expect: "Hello, foo!\n",
			msgError: "The value of option -w should be used as a name.",
		},
		{
			command: "hello ext --who foo", expect: "Hello, foo!\n",
			msgError: "The value of option --who should be used as a name.",
		},
		{
			command: "hello ext --who foo bar", expect: "Hello, foo and bar!\n",
			msgError: "The value of option --who and arg should be used as a name.",
		},
		{
			command: "hello ext -w foo bar buz", expect: "Hello, foo and bar buz!\n",
			msgError: "The value of option -w and args should be used as a name.",
		},
	}

	runTestCasesForHelloCmdExt(t, cases)
}

func Test_helloExtCmd_ReverseFlag(t *testing.T) {
	// test cases for reverse flag
	var cases TDataProvider = TDataProvider{
		{
			command: "hello ext -r", expect: "!dlrow ,olleH\n",
			msgError: "The result should be reversed by '-r' opt.",
		},
		{
			command: "hello ext --reverse", expect: "!dlrow ,olleH\n",
			msgError: "The result should be reversed by '--reverse' opt.",
		},
		{
			command: "hello ext -r foo", expect: "!oof ,olleH\n",
			msgError: "The result should include arg and reversed by '-r' opt.",
		},
		{
			command: "hello ext foo -r", expect: "!oof ,olleH\n",
			msgError: "The result should include arg and reversed by '-r' opt.",
		},
		{
			command: "hello ext -r foo bar", expect: "!rab oof ,olleH\n",
			msgError: "The result should include arg and reversed by '-r' opt.",
		},
		{
			command: "hello ext foo -r bar", expect: "!rab oof ,olleH\n",
			msgError: "The result should include arg and reversed by '-r' opt.",
		},
		{
			command: "hello ext foo bar -r", expect: "!rab oof ,olleH\n",
			msgError: "The result should include arg and reversed by '-r' opt.",
		},
	}

	runTestCasesForHelloCmdExt(t, cases)
}

// ----------------------------------------------------------------------------
//  Helper Functions (Functions for testing)
// ----------------------------------------------------------------------------

func convertShellArgsToSlice(t *testing.T, str string) []string {
	t.Helper() // With this call go test prints correct lines of code for failed tests.

	cmdArgs, err := shellwords.Parse(str)
	if err != nil {
		t.Fatalf("args parse error: %+v\n", err)
	}

	if 0 == len(cmdArgs) {
		t.Fatalf("args parse error. Command contains fatal strings: %+v\n", str)
	}

	// `hello cmd` dependent check
	if "hello" != cmdArgs[0] {
		t.Fatal("format error. The command must start with 'hello'.")
	}

	if "ext" != cmdArgs[1] {
		t.Fatal("format error. The command must start with 'hello cmd'.")
	}

	return cmdArgs[2:] // trim the first two args "hello" and "cmd".
}

func runTestCasesForHelloCmdExt(t *testing.T, cases TDataProvider) {
	t.Helper() // With this call go test prints correct lines of code for failed tests.

	var (
		helloExtCmd *cobra.Command
		buffTmp     = new(bytes.Buffer)
		argsTmp     []string

		expect    string
		actual    string
		msgErrTmp string
	)

	// Loop test cases
	for _, c := range cases {
		helloExtCmd = createHelloExtCmd()
		argsTmp = convertShellArgsToSlice(t, c.command)

		// init
		buffTmp.Reset()
		helloExtCmd.SetOut(buffTmp)
		helloExtCmd.SetArgs(argsTmp)

		// Run `hello ext`
		if err := helloExtCmd.Execute(); err != nil {
			assert.FailNowf(t, "Failed to execute 'helloExtCmd.Execute()'.", "Error msg: %v", err)
		}

		expect = c.expect
		actual = buffTmp.String() // resotre buffer
		msgErrTmp = c.msgError
		assert.Equal(t, expect, actual, msgErrTmp)
	}
}
