package day8

import (
	"bufio"
	"log"
	"os"
)

func LoadInput(filename string) [][]byte {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not read %s: %s\n", filename, err)
	}
	defer fp.Close()

	output := make([][]byte, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			byteValue := line[i] - '0' // Convert char to int
			newRow[i] = byteValue
		}
		output = append(output, newRow)
	}
	return output
}
