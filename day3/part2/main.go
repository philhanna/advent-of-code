package main

import (
	"fmt"
	"os"
	"strings"
)

// Convert a string to an array of distinct members
func toSet(s string) []string {
	contents := map[string]bool{}
	for i := 0; i < len(s); i++ {
		value := string(s[i])
		contents[value] = true
	}
	set := []string{}
	for value := range contents {
		if contents[value] {
			set = append(set, value)
		}
	}
	return set
}

func intersection(set1, set2 []string) []string {
	result := make([]string, 0)
	for _, s := range set1 {
		for _, t := range set2 {
			if s == t {
				result = append(result, s)
			}
		}
	}
	return result
}

func priority(s string) int {
	switch {
	case s >= "a" && s <= "z":
		return 1 + int(s[0] - 'a')
	case s >= "A" && s <= "Z":
		return 27 + int(s[0] - 'A')
	}
	return 0
}

func main() {
	inputBytes, _ := os.ReadFile("../testdata/input")
	data := strings.Split(string(inputBytes), "\n")
	sumOfPriorities := 0
	for i := 0;; i += 3 {
		if i+2 > len(data) {
			break
		}
		// Get the intersection of the contents of the rucksacks of each elf in this group
		items1 := toSet(data[i])
		items2 := toSet(data[i+1])
		result := intersection(items1, items2)
		items3 := toSet(data[i+2])
		result = intersection(result, items3)
		p := priority(result[0])
		sumOfPriorities += p
	}
	fmt.Printf("Sum of priorities = %d\n", sumOfPriorities)
}
