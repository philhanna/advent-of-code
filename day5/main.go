package main

import (
	"fmt"
	"log"
)

const filename = "testdata/moves.txt"

func main() {
	pShip, err := LoadShip(filename)
	if err != nil {
		log.Fatalf("Could not load ship: %s\n", err)
	}
	moves, err := LoadMoves(filename)
	if err != nil {
		log.Fatalf("Could not load moves: %s\n", err)
	}
	for i, move := range moves {
		err = pShip.MakeMove(move)
		if err != nil {
			log.Fatalf("Could not make move %d: %s\n", i+1, err)
		}
	}
	stackTops := pShip.StackTops()
	fmt.Printf("%s\n", pShip)
	fmt.Printf("Stack tops: %s\n", stackTops)
}