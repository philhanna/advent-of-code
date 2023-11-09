package main

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Node is an interface that represents an entry in the file system,
// either a directory or a file
type Node interface {
	Name() string      // Returns the file or directory name
	IsDir() bool       // Returns true if the node is a directory
	Size() int         // Returns the direct or indirect size of the node
	Children() []*Node // Returns the children of the node
}
