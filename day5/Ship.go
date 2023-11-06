package main

import (
	"bufio"
	"os"
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

	return ps, nil
}

// LoadConfigLines reads from the configuration file until the first blank line,
// returning the lines as a slice of strings.  We will transpose this list with
// another function.
func LoadConfigLines(filename string) ([]string, error) {
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

	// Make all the lines the same length
	maxlen := 0
	for _, line := range lines {
		if len(line) > maxlen {
			maxlen = len(line)
		}
	}
	sLines := make([]string, 0)
	for _, line := range lines {
		for len(line) < maxlen {
			line += " "
		}
		sLines = append(sLines, line)
	}

	// Get a long list of all the characters in the lines, left to right
	// and top to bottom.
	chars := make([]byte, 0)
	for _, line := range sLines {
		for i := 0; i < len(line); i++ {
			c := line[i]
			chars = append(chars, c)
		}
	}

	// Now construct the matrix of transposed lines
	matrix := make([][]byte, maxlen)
	for i := 0; i < maxlen; i++ {
		matrix[i] = make([]byte, len(lines))
	}
	p := 0
	for i := 0; i < maxlen; i++ {
		for j := 0; j < len(lines); j++ {
			c := chars[p]
			p++
			matrix[i][j] = c
		}
	}

	// Convert the 2D matrix of bytes into an array of strings
	out := make([]string, maxlen)
	for i := 0; i < len(matrix[i]); i++ {
		s := ""
		for j := 0; j < len(lines); j++ {
			s += string(matrix[i][j])
		}
		out = append(out, s)
	}
	return out
}
