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

	ctx := new(InputContext)

	// Create the root node
	ctx.root = NewDirNode(nil, "/")

	// Get a pointer to the current working directory
	ctx.cwd = ctx.root

	// Load the input and apply it one line at a time
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		err := ctx.HandleLine(line)
		if err != nil {
			return nil, err
		}
	}

	// Done. Everything OK.
	return ctx, nil
}

// HandleLine handles one line of input in the context of the input
// handler
func (ctx *InputContext) HandleLine(line string) error {
	return nil
}
