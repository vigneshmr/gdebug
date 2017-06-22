package require

/*

- This file contains assertion parallels that are found in testify/require
  to enable packages to be swapped.
- Methods that are overridden here will just passthough call into testify/require

*/

import (
	"github.com/stretchr/testify/require"
)

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	*require.Assertions
	tHook require.TestingT
}

// New makes a new Assertions object for the specified TestingT.
func New(t require.TestingT) *Assertions {
	return &Assertions{
		Assertions: require.New(t),
		tHook:      t,
	}
}

type TestingT struct {
	require.TestingT
}

func NoError(t require.TestingT, err error, msgAndArgs ...interface{}) {
	require.NoError(t, err, msgAndArgs)
}
