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

func (ps *Stack) String() string {
	parts := make([]string, 0)
	for _, crate := range ps.Crates {
		parts = append(parts, fmt.Sprintf("[%s]", crate))
	}
	return fmt.Sprintf("%d: %s", ps.IDNumber, strings.Join(parts, " "))
}