package cmdhello_test

import (
	"testing"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello"
	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello/cmdworld"
	"github.com/KEINOS/Hello-Cobra/cmd/cmdroot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

// ----------------------------------------------------------------------------
//  Test Functions
// ----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	t.Parallel()

	obj1 := cmdhello.New()
	obj2 := cmdhello.New()

	assert.NotSame(t, obj1, obj2, "it should not reference the same object")
}

func TestNew_has_child(t *testing.T) {
	t.Parallel()

	mother := cmdhello.New()
	children := mother.Commands()
	expectChild := cmdworld.New()

	require.True(t, mother.HasSubCommands(), "command 'hello' should contain a sub-command")
	assert.IsType(t, expectChild, children[0], "command 'hello' should contain 'world'")
}

func Test_sayHelloTo(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		expect string
		args   []string
	}{
		{
			"Hi!", []string{},
		},
		{
			"!iH", []string{"--reverse"},
		},
		{
			"About:", []string{"--help"},
		},
		{
			"Hello, foo and bar!", []string{"foo", "bar"},
		},
		{
			"!dlrow ,olleH", []string{"world", "--reverse"},
		},
	} {
		args := append([]string{"hello"}, test.args...)

		mother := cmdroot.New("v0.0.0-" + t.Name())
		mother.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			err := mother.Execute()

			require.NoError(t, err,
				"Input args: %v\nExpect out: %v\n", test.args, test.expect)
		})

		assert.Contains(t, out, test.expect,
			"Expect: %v\nActual: %v\n", test.expect, out)
	}
}
