package main

import (
	"os"
	"strconv"
	"strings"
)

func Expand(pair string) (int, int) {
	p := strings.Index(pair, "-")
	n1, _ := strconv.Atoi(pair[:p])
	n2, _ := strconv.Atoi(pair[p+1:])
	return n1, n2
}

func Contains(range1, range2 string) bool {
	n1a, n1b := Expand(range1)
	n2a, n2b := Expand(range2)
	return (n1a <= n2a && n1b >= n2b)
}

func main() {
	blob, _ := os.ReadFile("input")
	data := string(blob)
	for _, line := range strings.Split(data, "\n") {
		_ = line
	}
}
