package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Define a data structure for an Elf.
type Elf struct {
	number int
	cals   int
}

func main() {
	// Make a slice to keep track of elf numbers and their total calories.
	elves := make([]Elf, 0)

	// Read the input data into a slice, with a blank line separating the entries
	// between the calorie records for each elf.
	inputData, _ := ioutil.ReadFile("input.dat")
	inputLines := strings.Split(string(inputData), "\n\n")

	// Calculate the total calories for each elf.
	for i, entries := range inputLines {
		// New elf.
		elf := Elf{i + 1, 0}
		// Calculate all this elf's calories.
		entryLines := strings.Split(entries, "\n")
		for _, entry := range entryLines {
			cals, _ := strconv.Atoi(entry)
			elf.cals += cals
		}
		// Add this elf to the slice.
		elves = append(elves, elf)
	}

	// Sort the elf slice descending by total calories.
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].cals > elves[j].cals
	})

	// Part 1
	topElf := elves[0]
	fmt.Printf("Part1: Elf %d has %d calories\n", topElf.number, topElf.cals)

	// Part 2
	total := 0
	for _, elf := range elves[:3] {
		total += elf.cals
	}
	fmt.Printf("Part2: Top three elves have %d calories\n", total)
}
