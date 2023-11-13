package main

import (
	"fmt"

	"github/philhanna/advent-of-code/day8"
)

const filename = "testdata/input.dat"
func main() {
	data := day8.LoadData(filename)
	most := 0
	mostRow := 0
	mostCol := 0
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data); col++ {
			ss := data.ScenicScore(row, col)
			if ss > most {
				most = ss
				mostRow = row
				mostCol = col
			}
		}
	}
	fmt.Printf("Most scenic spot is (%d,%d). Score=%d\n", mostRow, mostCol, most)
}
