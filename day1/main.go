package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Elf struct {
	number int // Elf number (1, 2, ...)
	total int // Total calories
}

func (elf *Elf) String() string {
	return fmt.Sprintf("\tElf %d has total of %d calories", elf.number, elf.total)	
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

var elves = make([]Elf, 0)

const data =
`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

// ---------------------------------------------------------------------
// Mainline
// ---------------------------------------------------------------------


func main() {

	elves := loadElves()
	dumpElves("Original:", elves)
	sortElves(elves)
	dumpElves("Sorted:", elves)
	dumpTop("Top Elf:", elves)
	dumpTopThree("Sum of top three:", elves)
}
func dumpTopThree(label string, elves []Elf) {
	fmt.Printf("Sum of top 3:\n")
	fmt.Printf("\t%d calories\n", sumOfTopThree(elves))

}
// Dump the list of elves
func dumpElves(label string, elves []Elf) {
	fmt.Println(label)
	for _, elf := range elves {
		fmt.Printf("\tElf %d has total of %d calories\n", elf.number, elf.total)
	}
}

// Dumps the top elf
func dumpTop(label string, elves []Elf) {
	fmt.Println(label)
	fmt.Println(elves[0].String())
}

// Load the calorie data into an array of Elf structures
func loadElves() []Elf {

	var number int // Current elf number (starts at zero)

	// Create an array of elf calorie totals
	elves := []Elf{}
	elves = append(elves, Elf{number, 0})

	// Read the data into the array, converting from strings to integers
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			number++ // Start the next elf
			elves = append(elves, Elf{number, 0})
		} else {
			calories, _ := strconv.Atoi(text)
			elves[number].total += calories
		}
	}

	return elves
}

// Sort the elf array in place in descending order of total calories
func sortElves(elves []Elf) {
	sortFunction := func(i, j int) bool {
		return elves[i].total >= elves[j].total
	}
	sort.Slice(elves, sortFunction)
}

// Returns the sum of the calories of the top three elves
func sumOfTopThree(elves []Elf) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += elves[i].total
	}
	return sum
}