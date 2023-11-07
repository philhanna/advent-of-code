package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Move struct {
	Count     int
	FromStack int
	ToStack   int
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// ParseMove creates a Move structure from an input line in the format:
//
//	move nn from n to n
func ParseMove(line string) (Move, error) {
	re := regexp.MustCompile(`^\s*move (\d+) from (\d+) to (\d+)`)
	groups := re.FindStringSubmatch(line)
	if groups == nil {
		err := fmt.Errorf("Invalid format for move input: %s", line)
		return Move{}, err
	}
	count, _ := strconv.Atoi(groups[1])
	fromStack, _ := strconv.Atoi(groups[2])
	toStack, _ := strconv.Atoi(groups[3])
	return Move{
		Count:     count,
		FromStack: fromStack,
		ToStack:   toStack,
	}, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Equal returns true if the specified Move object is equal to this one.
func (this *Move) Equal(that Move) bool {
	return this.Count == that.Count &&
		this.FromStack == that.FromStack &&
		this.ToStack == that.ToStack
}
