package main

import "fmt"

const (
	FILENAME = "sampleinput.txt"
	LIMIT = 100000
)

func main() {
	context := NewContext()
	err := context.HandleInput(FILENAME)
	if err != nil {
		fmt.Printf("Could not read %s: %s\n", FILENAME, err)
	}
}
