package day8

import (
	"bufio"
	"log"
	"os"
)

// LoadInput reads the problem input and converts it into a
// two-dimensional slice of byte slices.
func LoadInput(filename string) [][]byte {

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
