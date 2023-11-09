package main

import (
	"bufio"
	"os"
	"strings"
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

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// NewContext creates a new input handler context with an initialized
// root and current working directory
func NewContext() *InputContext {
	context := new(InputContext)
	context.root = NewDirNode(nil, "/")
	context.cwd = context.root
	return context
}

// HandleInput creates a root node, then reads the input and applies it
// to that node
func HandleInput(filename string) (*InputContext, error) {

	// Create the input context
	context := NewContext()

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
	switch {
	case strings.HasPrefix(line, "$ ls"):
		context.HandleLS(line)
	}
	return nil
}

// HandleLS handles the ls command
func (context *InputContext) HandleLS(line string) error {
	// Nothing to do. Return no error
	return nil
}