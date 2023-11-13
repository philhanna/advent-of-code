package day8

import (
	"bufio"
	"log"
	"os"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Data []Vector

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
	output := make([]Vector, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := make(Vector, len(line))
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

func (data Data) GetCol(col int) Vector {
	vector := Vector{}
	for _, rowVector := range data {
		entry := rowVector[col]
		vector = append(vector, entry)
	}
	return vector
}

func (data Data) GetReversedCol(col int) Vector {
	vector := data.GetCol(col)
	return vector.Reverse()
}

func (data Data) GetRow(row int) Vector {
	vector := Vector{}
	for col := 0; col < len(data[row]); col++ {
		vector = append(vector, data[row][col])
	}
	return vector
}

func (data Data) GetReversedRow(row int) Vector {
	vector := data.GetRow(row)
	return vector.Reverse()
}
