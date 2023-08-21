package main

import (
	"regexp"
	"strconv"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Move struct {
	Count     int
	FromStack int // The 1-based "from" stack
	ToStack   int // The 1-based "to" stack
}

func NewMove(movestr string) *Move {
	pm := new(Move)
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	m := re.FindStringSubmatch(movestr)
	if m != nil {
		pm.Count, _ = strconv.Atoi(m[1])
		pm.FromStack, _ = strconv.Atoi(m[2])
		pm.ToStack, _ = strconv.Atoi(m[3])
	}
	return pm
}
