package main

import (
	"bufio"
	"fmt"
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
// Constructor
// ---------------------------------------------------------------------

// NewContext creates a new input handler context with an initialized
// root and current working directory
func NewContext() *InputContext {
	context := new(InputContext)
	context.root = NewDirNode(nil, "/")
	context.cwd = context.root
	return context
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// HandleInput reads input from the specified file and applies it to the
// context
func (context *InputContext) HandleInput(filename string) error {

	// Load the input
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	// Read the input line by line and apply it to the context
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		err := context.HandleLine(line)
		if err != nil {
			return err
		}
	}

	// Done. Indicate that there was no error.
	return nil
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

// HandleCD handles the cd command
func (context *InputContext) HandleCD(line string) error {
	
	dirname := line[5:]

	switch dirname {

	case "/":	// Change to root directory
		context.cwd = context.root
		return nil

	case "..":	// Change to parent of current working directory
		if context.cwd.parent == nil {
			return fmt.Errorf("cannot get parent of root")
		}
		context.cwd = context.cwd.parent
		return nil

	default:	// Change to specified directory
		if !context.cwd.HasChild(dirname) {
			context.cwd = NewDirNode(context.cwd, dirname)
		}
		return nil
	}
}