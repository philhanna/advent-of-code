package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	// Define a data structure for an Elf.
	type Elf struct {
		number int
		cals   int
	}

	// Make an array to keep track of elf numbers and their total
	// calories.  Each elf entry is an Elf object (elf number and total
	// calories), and elves is an array of elf entries.
	elves := make([]Elf, 0)

	// Read the input data into a byte array, with a blank line
	// separating the entries between the calorie records for each elf.
	data, _ := os.ReadFile("input.dat")
	calorieData := bytes.Split(data, []byte{'\n', '\n'})

	// Calculate the total calories for each elf
	for i, calorieEntries := range calorieData {

		// New elf
		elf := Elf{number: i + 1, cals: 0}

		// Calculate all this elf's calories
		for _, entry := range bytes.Split(calorieEntries, []byte{'\n'}) {
			cals, _ := strconv.Atoi(string(entry))
			elf.cals += cals
		}

		// Add this elf to the array
		elves = append(elves, elf)
	}

	// Sort the elf array descending by total calories
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].cals > elves[j].cals
	})

	// Part 1
	topElf := elves[0]
	fmt.Printf("Part1: Elf %d has %d calories\n", topElf.number, topElf.cals)

	// Part2
	total := elves[0].cals + elves[1].cals + elves[2].cals
	fmt.Printf("Part2: Top three elves have %d calories\n", total)
}
