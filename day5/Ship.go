package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Ship is a collection of the stacks we are working with
type Ship struct {
	Stacks []Stack
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
	ps := new(Ship)
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
		ps.Stacks = append(ps.Stacks, stack)
	}

	return ps, nil
}

func MakeStack(line string) (Stack, error) {
	var err error
	var crate string

	stack := Stack{}

	// The last character of the string is the stack number
	c, line := Peel(line)
	stack.IDNumber, err = strconv.Atoi(c)
	if err != nil {
		return stack, err
	}

	// Get the crate names from the input line, reversing the order so
	// that they are a stack
	for {
		if line == "" {
			break
		}
		crate, line = Peel(line)
		stack.Crates = append(stack.Crates, crate)
	}
	return stack, nil
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
