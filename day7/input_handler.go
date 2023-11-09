package main

import (
	"bufio"
	"os"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type InputContext struct {
	root *DirNode // Root directory
	cwd  *DirNode // Current working directory
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

// HandleInput creates a root node, then reads the input and applies it
// to that node
func HandleInput(filename string) (*InputContext, error) {

	// Create the input context
	context := new(InputContext)
	context.root = NewDirNode(nil, "/")
	context.cwd = context.root

	// Load the input
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// Read the input line by line and apply it to the context
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		err := context.HandleLine(line)
		if err != nil {
			return nil, err
		}
	}

	// Done
	return context, nil
}

// HandleLine handles one line of input in the context of the input
// handler
func (context *InputContext) HandleLine(line string) error {
	return nil
}
