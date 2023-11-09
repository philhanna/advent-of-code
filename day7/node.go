package main

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// INode is the interface that DirNode and FileNode must implement.
type INode interface {
	Name() string     // Returns the node name
	Parent() *DirNode // Returns the parent directory node
	Size() int        // Returns the file or directory size
	StringLS() string // Returns the string output by the ls command for this node
}

// Node is the abstract base type that DirNode and FileNode extend.
type Node struct {
	INode
}
