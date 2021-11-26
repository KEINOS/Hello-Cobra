package cmdroot_test

import (
	"testing"

	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello"
	"github.com/KEINOS/Hello-Cobra/cmd/cmdhello/cmdworld"
	"github.com/KEINOS/Hello-Cobra/cmd/cmdroot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	obj1 := cmdroot.New("v0.0.0-" + t.Name())
	obj2 := cmdroot.New("v0.0.0-" + t.Name())

	assert.NotSame(t, obj1, obj2, "it should not reference the same object")
}

func TestNew_empty_version(t *testing.T) {
	mother := cmdroot.New("")

	expect := "(unknown)"
	actual := mother.Version

	assert.Equal(t, expect, actual)
}

func TestNew_has_child(t *testing.T) {
	mother := cmdroot.New("v0.0.0-" + t.Name())

	require.True(t, mother.HasSubCommands(),
		"new root object should contain sub-command/s")

	children := mother.Commands()
	expectChild := cmdhello.New()

	require.IsType(t, expectChild, children[0],
		"command 'root' should contain 'hello' as a child")

	grandChildren := children[0].Commands()
	expectGrandChild := cmdworld.New()

	assert.IsType(t, expectGrandChild, grandChildren[0],
		"command 'root' should contain 'world' as a grandchild")
}
