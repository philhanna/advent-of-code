package main

import (
	"fmt"
	"os"
)

const LIMIT = 4

func main() {
	FILENAME := "input.txt"
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	p, prev := findMarker(string(data))
	fmt.Printf("p=%d, prev=%s\n", p, prev)
}

func findMarker(s string) (int, string) {
	for p := range s {
		if p >= LIMIT {
			prev := s[p-LIMIT : p]
			if !dups(prev) {
				return p, prev
			}
		}
	}
	return 0, ""
}

func dups(s string) bool {
	for i, ci := range s {
		for j, cj := range s {
			if i != j && ci == cj {
				return true
			}
		}
	}
	return false
}
