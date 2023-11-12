package day8

import (
	"bufio"
	"log"
	"os"
)

func LoadInput(filename string) []string {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not read %s: %s\n", filename, err)
	}
	defer fp.Close()

	output := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return output
}