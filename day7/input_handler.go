package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

var (
	re_size_and_name = regexp.MustCompile(`^(\d+) (\S+)`)
)

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
	case strings.HasPrefix(line, "$ cd"):
		context.HandleCD(line)
	case re_size_and_name.MatchString(line):
		context.HandleFileSizeAndName(line)
	}
	return nil
}

// HandleLS handles the ls command
func (context *InputContext) HandleLS(line string) ([]INode, error) {
	return context.cwd.children, nil
}

// HandleCD handles the cd command
func (context *InputContext) HandleCD(line string) error {

	dirname := line[5:]

	switch dirname {

	case "/": // Change to root directory
		context.cwd = context.root
		return nil

	case "..": // Change to parent of current working directory
		if context.cwd.parent == nil {
			return fmt.Errorf("cannot get parent of root")
		}
		context.cwd = context.cwd.parent
		return nil

	default: // Change to specified directory
		child := context.cwd.LookupChild(dirname)
		if child == nil {
			context.cwd = NewDirNode(context.cwd, dirname)
		} else {

			// Yes, you *can* cast an interface to a struct that
			// implements that interface, e.g., make an INode into a
			// *DirNode, once we know for sure it is a directory.
			context.cwd = child.(*DirNode)
		}
		return nil
	}
}

func (context *InputContext) HandleFileSizeAndName(line string) error {
	groups := re_size_and_name.FindStringSubmatch(line)
	if groups == nil {
		return fmt.Errorf("%q is not a file size and name line", line)
	}
	fileSize, _ := strconv.Atoi(groups[1])
	fileName := groups[2]
	child := context.cwd.LookupChild(fileName)
	if child != nil {
		// The file by this name is already a child of the cwd
		return nil
	}
	// Add a new file by this name to the cwd
	NewFileNode(context.cwd, fileName, fileSize)
	return nil
}
