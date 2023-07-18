package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	blob, _ := os.ReadFile("input")
	data := string(blob)
	count := 0

	expand := func(pair string) (int, int) {
		p := strings.Index(pair, "-")
		n1, _ := strconv.Atoi(pair[:p])
		n2, _ := strconv.Atoi(pair[p+1:])
		return n1, n2			
	}

	contains := func(range1, range2 string) bool {
		n1a, n1b := expand(range1)
		n2a, n2b := expand(range2)
		return (n1a <= n2a && n1b >= n2b) || (n2a <= n1a && n2b >= n1b)
	}

	for _, line := range strings.Split(data, "\n") {
		p := strings.Index(line, ",")
		if contains(line[:p], line[p+1:]) {
			count++
		}
	}
	fmt.Printf("Count=%d\n", count)
}
