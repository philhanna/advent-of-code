package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("testdata/input")
	sum := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		n := len(line) / 2
		if n > 0 {
			comp1 := line[:n]
			comp2 := line[n:]
			p := strings.IndexAny(comp1, comp2)
			if p > -1 {
				c := line[p]
				var priority int
				switch {
				case c >= 'a' && c <= 'z':
					priority = int(c - 'a') + 1
				case c >= 'A' && c <= 'Z':
					priority = int(c - 'A') + 27
				}
				sum += priority
			}
		}
	}
	fmt.Printf("Sum of priorities=%d\n", sum)
}
