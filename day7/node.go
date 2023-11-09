package main

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Node is an interface that represents an entry in the file system,
// either a directory or a file
type Node interface {
	Name() string              // Returns the file or directory name
	Size() int                 // Returns the direct or indirect size of the node
	Children() []*AbstractNode // Returns the children of the node
}

// AbstractNode is the base class for DirNode and FileNode
type AbstractNode struct {
	Node
}
