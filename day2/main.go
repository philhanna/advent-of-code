package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Load input data into memory
	input, err := os.ReadFile("input.dat")
	if err != nil {
		log.Fatal(err)
	}
	s := string(input)                 // Convert byte array to string
	s = strings.ReplaceAll(s, " ", "") // Remove spaces
	strategy := strings.Split(s, "\n") // Split lines into array

	totalScore := 0
	for _, move := range strategy {
		totalScore +=

			// Give yourself 1 if you chose rock, 2 if paper, 3 if scissors
			func(m byte) int {
				return 1 + strings.Index("XYZ", string(m))
			}(move[1]) +

				// Add the part that is determined by the outcome (W/L/D)
				func(move string) int {
					switch {
					case strings.Contains("AZ,BX,CY", move):
						return 0 // I lost
					default:
						return 3 // Draw
					case strings.Contains("AY,BZ,CX", move):
						return 6 // I won
					}
				}(move)
	}
	fmt.Printf("Total score is %d\n", totalScore)
}
