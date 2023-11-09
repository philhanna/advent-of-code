package main

import (
	"bufio"
	"os"
	"strings"
)

// HandleInput creates a root node, then reads the input and applies it
// to that node.
func HandleInput(filename string) (*DirNode, error) {

	// Create the root node
	root := NewDirNode(nil, "/")

	// Get a pointer to the current working directory
	cwd := root

	// Load the input and apply it one line at a time
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "$ "):
			line = line[2:]
			switch {
			case strings.HasPrefix(line, "cd "):
				dirName := line[3:]
				if dirName == "/" {
					cwd = root
				} else {
					cwd = NewDirNode(cwd, dirName)
				}
			}
		}
	}

	// Done. Everything OK.
	return root, nil
}
