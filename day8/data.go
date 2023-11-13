package day8

import (
	"bufio"
	"log"
	"os"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Data [][]byte

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------
// LoadData reads the problem input and converts it into a
// two-dimensional slice of byte slices.
func LoadData(filename string) Data {

	// Open the input file or die trying
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not read %s: %s\n", filename, err)
	}
	defer fp.Close()

	// Read each line and convert it to a slice of bytes
	output := make([][]byte, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			byteValue := line[i] - '0' // Convert char to int
			newRow[i] = byteValue
		}

		// Store byte slice into output slice
		output = append(output, newRow)
	}

	// Done
	return output
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// VisibleUp returns true if the tree at this row and column is visible
// from above
func (data Data) VisibleUp(row, col int) bool {
	return false
}

// VisibleDown returns true if the tree at this row and column is visible
// from below
func (data Data) VisibleDown(row, col int) bool {
	return false
}

// VisibleLeft returns true if the tree at this row and column is visible
// from the left
func (data Data) VisibleLeft(row, col int) bool {
	return false
}

// VisibleRight returns true if the tree at this row and column is visible
// from the right
func (data Data) VisibleRight(row, col int) bool {
	return false
}
