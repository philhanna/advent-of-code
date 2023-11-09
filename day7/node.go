package main

import "strings"

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Node is an interface that represents an entry in the file system,
// either a directory or a file
type Node interface {
	Name() string              // Returns the file or directory name
	Parent() *DirNode           // Returns the parent node
	Size() int                 // Returns the direct or indirect size of the node
	Children() []*AbstractNode // Returns the children of the node
}

// AbstractNode is the base class for DirNode and FileNode
type AbstractNode struct {
	Node
}

// String returns a string representation of this node and all its
// parents, the full path
func (p *AbstractNode) String() string {
	sb := strings.Builder{}
	if p.Parent() == nil {
		sb.WriteString("/")
	} else {
		if p.Parent().Parent() != nil {
			sb.WriteString(p.Parent().String())
		}
		sb.WriteString("/")
		sb.WriteString(p.Name())
	}
	return sb.String()
}
