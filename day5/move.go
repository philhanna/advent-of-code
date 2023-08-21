package main

import (
	"regexp"
	"strconv"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Move struct {
	Count          int
	FromStackIndex int // The 1-based "from" stack
	ToStackIndex   int // The 1-based "to" stack
}

func NewMove(movestr string) *Move {
	pm := new(Move)
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	m := re.FindStringSubmatch(movestr)
	if m != nil {
		pm.Count, _ = strconv.Atoi(m[1])
		pm.FromStackIndex, _ = strconv.Atoi(m[2])
		pm.ToStackIndex, _ = strconv.Atoi(m[3])
	}
	return pm
}
