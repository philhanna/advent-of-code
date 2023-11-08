package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Ship is a collection of the stacks we are working with
type Ship struct {
	Stacks map[int]*Stack
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewShip creates a new ship with empty stacks
func NewShip() *Ship {
	p := new(Ship)
	p.Stacks = make(map[int]*Stack)
	return p
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// LoadShip creates a ship from information in a file.
// The initial configuration looks like this:
//
// [G]                 [D] [R]
// [W]         [V]     [C] [T] [M]
// [L]         [P] [Z] [Q] [F] [V]
// [J]         [S] [D] [J] [M] [T] [V]
// [B]     [M] [H] [L] [Z] [J] [B] [S]
// [R] [C] [T] [C] [T] [R] [D] [R] [D]
// [T] [W] [Z] [T] [P] [B] [B] [H] [P]
// [D] [S] [R] [D] [G] [F] [S] [L] [Q]
//  1   2   3   4   5   6   7   8   9
//
// move 1 from 3 to 5
// move 5 from 5 to 4
// move 6 from 7 to 3
// ...
// move 1 from 3 to 4

func LoadShip(filename string) (*Ship, error) {

	// Create a new Ship structure
	ps := NewShip()

	// Load the starting configuration from the specified file
	lines, err := LoadStackLines(filename)
	if err != nil {
		return nil, err
	}
	lines = TransposeLines(lines)

	// Extract the stacks that occur every 4 lines:
	//
	// 0: "[[[[[[[[ ",
	// 1: "GWLJBRTD1",  <--
	// 2: "]]]]]]]] ",
	// 3: "         ",
	// 4: "     [[[ ",
	// 5: "     CWS2",  <--
	// 6: "     ]]] ",
	// 7: "         ",
	// 8: "    [[[[ ",
	// 9: "    MTZR3",  <--
	// 10: "    ]]]] ",
	// etc...
	for p := 1; p < len(lines); p += 4 {
		line := lines[p]
		stack, err := MakeStack(line)
		if err != nil {
			return nil, err
		}
		ps.Stacks[stack.IDNumber] = stack
	}

	// Done
	return ps, nil
}

// LoadStackLines reads from the input data file until the first blank
// line, returning the lines as a slice of strings.  We will transpose
// this list with another function.
func LoadStackLines(filename string) ([]string, error) {
	lines := make([]string, 0)
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// MakeStack creates a stack structure consisting of the ID number and
// the array of crates.
func MakeStack(line string) (*Stack, error) {
	var err error
	var crate string

	pStack := new(Stack)

	// The last character of the string is the stack number
	c, line := Peel(line)
	pStack.IDNumber, err = strconv.Atoi(c)
	if err != nil {
		return nil, err
	}

	// Get the crate names from the input line, reversing the order so
	// that they are a stack
	for {
		crate, line = Peel(line)
		if strings.TrimSpace(crate) == "" {
			break
		}
		pStack.Crates = append(pStack.Crates, crate)
	}

	// All OK
	return pStack, nil
}

// Peel returns the last character of a string and the original string
// without that last character
func Peel(s string) (string, string) {
	i := len(s) - 1
	if i < 0 {
		return "", ""
	}
	c := s[i:]
	s = s[:i]
	return c, s
}

// TransposeLines creates a transposition of the matrix of line characters.
func TransposeLines(lines []string) []string {
	var nRows, nCols int

	nRows = len(lines)

	// Make sure all the lines are the same length
	maxlen := 0
	for _, line := range lines {
		if len(line) > maxlen {
			maxlen = len(line)
		}
	}
	nCols = maxlen
	sLines := make([]string, 0)
	for _, line := range lines {
		for len(line) < nCols {
			line += " "
		}
		sLines = append(sLines, line)
	}

	// Create the transposed list of strings
	out := make([]string, nCols)
	for i := 0; i < nCols; i++ {
		out[i] = ""
		for j := 0; j < nRows; j++ {
			ch := sLines[j][i : i+1]
			out[i] += ch
		}
	}
	return out
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// MakeMove applies a move to a from and to stack on this ship
func (ps *Ship) MakeMove(move Move) error {
	var err error

	pFromStack, ok := ps.Stacks[move.FromStackNumber]
	if !ok {
		return fmt.Errorf("Invalid fromStack %d", move.FromStackNumber)
	}
	pToStack, ok := ps.Stacks[move.ToStackNumber]
	if !ok {
		return fmt.Errorf("Invalid toStack %d", move.ToStackNumber)
	}

	for count := 0; count < move.Count; count++ {
		if len(pFromStack.Crates) < 1 {
			return fmt.Errorf("From stack %d is empty", move.FromStackNumber)
		}
		fromCratesCount := len(pFromStack.Crates)
		crate := pFromStack.Crates[fromCratesCount-1]
		pFromStack.Crates = pFromStack.Crates[:fromCratesCount-1]
		pToStack.Crates = append(pToStack.Crates, crate)
	}

	return err
}

// SortedStackNumbers returns a slice of the stack numbers sorted in ascending order
func (ps *Ship) SortedStackNumbers() []int {
	keys := make([]int, 0)
	for _, pStack := range ps.Stacks {
		keys = append(keys, pStack.IDNumber)
	}
	sort.Ints(keys)
	return keys
}

// StackTops returns a string containing the top crate on each stack
func (ps *Ship) StackTops() string {
	sb := strings.Builder{}
	keys := ps.SortedStackNumbers()
	for _, key := range keys {
		pStack := ps.Stacks[key]
		n := len(pStack.Crates)
		crate := pStack.Crates[n-1]
		sb.WriteString(crate)
	}
	return sb.String()
}

// String returns a string representation of the stacks
func (ps *Ship) String() string {
	parts := make([]string, 0)
	keys := ps.SortedStackNumbers()
	for _, key := range keys {
		pStack := ps.Stacks[key]
		parts = append(parts, fmt.Sprintf("%s", pStack))
	}
	return strings.Join(parts, "\n")
}
