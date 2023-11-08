package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Move struct {
	Count           int
	FromStackNumber int
	ToStackNumber   int
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

var re = regexp.MustCompile(`^\s*move (\d+) from (\d+) to (\d+)`)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// LoadMoves creates a list of moves from the input data file.
// The moves are located after the first empty line.
func LoadMoves(filename string) ([]Move, error) {
	moves := make([]Move, 0)
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	state := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case 0:
			// Skip stack configuration section and look for first blank line
			if strings.TrimSpace(line) == "" {
				state = 1
			}
		case 1:
			fallthrough
		default:
			move, err := ParseMove(line)
			if err != nil {
				return nil, err
			}
			moves = append(moves, move)
		}
	}
	return moves, nil
}

// ParseMove creates a Move structure from an input line in the format:
//
//	move nn from n to n
func ParseMove(line string) (Move, error) {
	groups := re.FindStringSubmatch(line)
	if groups == nil {
		err := fmt.Errorf("Invalid format for move input: %s", line)
		return Move{}, err
	}
	count, _ := strconv.Atoi(groups[1])
	fromStack, _ := strconv.Atoi(groups[2])
	toStack, _ := strconv.Atoi(groups[3])
	return Move{
		Count:           count,
		FromStackNumber: fromStack,
		ToStackNumber:   toStack,
	}, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Equal returns true if the specified Move object is equal to this one.
func (this *Move) Equal(that Move) bool {
	return this.Count == that.Count &&
		this.FromStackNumber == that.FromStackNumber &&
		this.ToStackNumber == that.ToStackNumber
}

// String returns a string representation of the move
func (p *Move) String() string {
	parts := []string{
		fmt.Sprintf("%d", p.Count),
		fmt.Sprintf("%d", p.FromStackNumber),
		fmt.Sprintf("%d", p.ToStackNumber),
	}
	return strings.Join(parts, ",")
}
