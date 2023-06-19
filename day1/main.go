package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	// Make a structure to keep track of elf totals
	type Elf struct {
		number int
		cals   int
	}
	elves := make([]Elf, 0)

	// Read the data into the array, converting from strings to integers
	data, err := os.ReadFile("input.dat")
	if err != nil {
		log.Fatal(err)
	}

	// Split the data on blank lines
	calorieData := bytes.Split(data, []byte{'\n', '\n'})
	for i, token := range calorieData {

		// New elf
		elf := Elf{number: i + 1, cals: 0}

		// Calculate all this elf's calories
		for _, entry := range bytes.Split(token, []byte{'\n'}) {
			cals, err := strconv.Atoi(string(entry))
			if err != nil {
				log.Fatal(err)
			}
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
	fmt.Printf("Part1: Elf %d has %d calories\n", elves[0].number, elves[0].cals)

	// Part2
	total := 0
	for i := 0; i < 3; i++ {
		total += elves[i].cals
	}
	fmt.Printf("Part2: Top three elves have %d calories\n", total)
}
