package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	blob, _ := os.ReadFile("../input")
	data := string(blob)
	count := 0

	expand := func(pair string) (int, int) {
		p := strings.Index(pair, "-")
		n1, _ := strconv.Atoi(pair[:p])
		n2, _ := strconv.Atoi(pair[p+1:])
		return n1, n2
	}

	overlaps := func(range1, range2 string) bool {
		n1a, n1b := expand(range1)
		n2a, n2b := expand(range2)
		common := make(map[int]int)
		for i := n1a; i <= n1b; i++ {
			common[i] = 1
		}
		for i := n2a; i <= n2b; i++ {
			common[i] += 1
		}
		for _, v := range common {
			if v > 1 {
				return true
			}
		}
		return false
	}

	for _, line := range strings.Split(data, "\n") {
		p := strings.Index(line, ",")
		if overlaps(line[:p], line[p+1:]) {
			count++
		}
	}
	fmt.Printf("Count=%d\n", count)
}
