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

// CountVisible returns the number of visible trees in this input data
func (data Data) CountVisible() int {
	count := 0
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data); col++ {
			if data.Visible(row, col) {
				count++
			}
		}
	}
	return count
}

// Visible returns true if the tree at this row and column is visible
// from any of the four directions.
func (data Data) Visible(row, col int) bool {

	// Make a list of functions to test with
	x := []func(int, int) bool{
		data.VisibleUp,
		data.VisibleDown,
		data.VisibleLeft,
		data.VisibleRight,
	}
	for i := 0; i < len(x); i++ {
		if x[i](row, col) {
			return true
		}
	}
	return false
}

// VisibleUp returns true if the tree at this row and column is visible
// from above
func (data Data) VisibleUp(row, col int) bool {
	thisHeight := data[row][col]
	for r := row - 1; r >= 0; r-- {
		thatHeight := data[r][col]
		if thatHeight >= thisHeight {
			return false
		}
	}
	return true
}

// VisibleDown returns true if the tree at this row and column is visible
// from below
func (data Data) VisibleDown(row, col int) bool {
	nRows := len(data)
	thisHeight := data[row][col]
	for r := row + 1; r < nRows; r++ {
		thatHeight := data[r][col]
		if thatHeight >= thisHeight {
			return false
		}
	}
	return true
}

// VisibleLeft returns true if the tree at this row and column is visible
// from the left
func (data Data) VisibleLeft(row, col int) bool {
	thisHeight := data[row][col]
	for c := col - 1; c >= 0; c-- {
		thatHeight := data[row][c]
		if thatHeight >= thisHeight {
			return false
		}
	}
	return true
}

// VisibleRight returns true if the tree at this row and column is visible
// from the right
func (data Data) VisibleRight(row, col int) bool {
	nCols := len(data)
	thisHeight := data[row][col]
	for c := col + 1; c < nCols; c++ {
		thatHeight := data[row][c]
		if thatHeight >= thisHeight {
			return false
		}
	}
	return true
}
