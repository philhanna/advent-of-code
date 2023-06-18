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
		score :=
			func(m byte) int { return strings.Index("XYZ", string(m)) }(move[1]) +
			func(move string) int {
				switch {
				case strings.Contains("AY,BZ,CX", move):
					return 6 // I won
				case strings.Contains("AZ,BX,CY", move):
					return 0 // I lost
				default:
					return 3 // Draw
				}
			}(move)
		printMove(move, score)
		totalScore += score
	}
	fmt.Printf("Total score is %d\n", totalScore)
}

func printMove(move string, score int) {
	him := move[0]
	fmt.Print("Opponent's choice was ")
	switch him {
	case 'A':
		fmt.Print("rock")
	case 'B':
		fmt.Print("paper")
	case 'C':
		fmt.Print("scissors")
	}
	me := move[1]
	fmt.Print(", my choice was ")
	switch me {
	case 'X':
		fmt.Print("rock")
	case 'Y':
		fmt.Print("paper")
	case 'Z':
		fmt.Print("scissors")
	}
	iWin := strings.Contains("AY,BZ,CX", move)
	iLose := strings.Contains("AZ,BX,CY", move)
	switch {
	case iWin:
		fmt.Printf(", I win")
	case iLose:
		fmt.Printf(", I lose")
	default:
		fmt.Printf(", we draw")
	}
	fmt.Printf(", score for this move is %d\n", score)

}
