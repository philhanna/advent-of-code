package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("testdata/input")
	sumOfPriorities := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		n := len(line) / 2
		if n > 0 {
			compartment1 := line[:n]
			compartment2 := line[n:]
			p := strings.IndexAny(compartment1, compartment2)
			if p > -1 {
				c := line[p]
				var priority int
				switch {
				case c >= 'a' && c <= 'z':
					priority = int(c - 'a') + 1
				case c >= 'A' && c <= 'Z':
					priority = int(c - 'A') + 27
				}
				sumOfPriorities += priority
			}
		}
	}
	fmt.Printf("Sum of priorities=%d\n", sumOfPriorities)
}
