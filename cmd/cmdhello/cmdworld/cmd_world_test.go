package cmdworld_test

import (
	"testing"

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

	obj1 := cmdworld.New()
	obj2 := cmdworld.New()

	assert.NotSame(t, obj1, obj2, "it should not reference the same object")
}

func Test_sayHelloWorld(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		expect string
		args   []string
	}{
		{"Hello, world!", []string{}},
		{"!dlrow ,olleH", []string{"-r"}},
		{"!dlrow ,olleH", []string{"--reverse"}},
	} {
		args := append([]string{"hello", "world"}, test.args...)

		grandMother := cmdroot.New("v0.0.0-" + t.Name())
		grandMother.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			err := grandMother.Execute()

			require.NoError(t, err)
		})

		assert.Contains(t, out, test.expect)
	}
}
