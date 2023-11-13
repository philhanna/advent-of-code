package main

import (
	"fmt"

	"github/philhanna/advent-of-code/day8"
)

const filename = "testdata/sample.dat"
func main() {
	data := day8.LoadData(filename)
	count := data.CountVisible()
	fmt.Printf("Number of visible trees = %d\n", count)
}
