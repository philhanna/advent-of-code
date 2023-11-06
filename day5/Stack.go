package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Stack struct {
	IDNumber int
	Crates   []string
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of the stack
func (ps *Stack) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d:", ps.IDNumber))
	for _, crate := range ps.Crates {
		sb.WriteString(fmt.Sprintf(" [%s]", crate))
	}
	return sb.String()
}