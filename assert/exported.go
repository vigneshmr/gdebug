package assert

/*

- This file contains assertion parallels that are found in testify/assert
  to enable packages to be swapped.
- Methods that are overridden here will just passthough call into testify/assert

*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	*assert.Assertions
	tHook assert.TestingT
}

// New makes a new Assertions object for the specified TestingT.
func New(t *testing.T) *Assertions {
	return &Assertions{
		Assertions: assert.New(t),
		tHook:      t,
	}
}
