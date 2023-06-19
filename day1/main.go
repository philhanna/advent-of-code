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

	// Make an array to keep track of elf numbers and their total calories
	elves := make([][2]int, 0)

	// Read the data into the array, converting from strings to integers
	data, err := os.ReadFile("input.dat")
	if err != nil {
		log.Fatal(err)
	}

	// Split the data on blank lines
	calorieData := bytes.Split(data, []byte{'\n', '\n'})
	for i, token := range calorieData {

		// New elf
		elf := [2]int{i + 1, 0}

		// Calculate all this elf's calories
		for _, entry := range bytes.Split(token, []byte{'\n'}) {
			cals, _ := strconv.Atoi(string(entry))
			elf[1] += cals
		}

		// Add this elf to the array
		elves = append(elves, elf)
	}

	// Sort the elf array descending by total calories
	sort.Slice(elves, func(i, j int) bool {
		return elves[i][1] > elves[j][1]
	})

	// Part 1
	fmt.Printf("Part1: Elf %d has %d calories\n", elves[0][0], elves[0][1])

	// Part2
	total := elves[0][1] + elves[1][1] + elves[2][1]
	fmt.Printf("Part2: Top three elves have %d calories\n", total)
}
