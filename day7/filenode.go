package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// FileNode is a Node that represents a file.
type FileNode struct {
	*Node
	name   string
	parent *DirNode
	size   int
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewFileNode creates a new file node with the specified parent
// directory, name, and size.
func NewFileNode(parent *DirNode, name string, size int) *FileNode {
	node := new(Node)
	fileNode := &FileNode{node, name, parent, size}
	if parent != nil {
		parent.children = append(parent.children, node)
	}
	node.INode = fileNode
	return fileNode
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of this file node
func (p *FileNode) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("name=%s", p.Name()))
	parts = append(parts, fmt.Sprintf("parent=%s", p.parent.Name()))
	parts = append(parts, fmt.Sprintf("path=%s", p.FullPath()))
	result := strings.Join(parts, ",")
	return result
}

// FullPath returns the full path to this file
func (p *FileNode) FullPath() string {
	sb := strings.Builder{}
	if p.parent.parent != nil {
		sb.WriteString(p.parent.FullPath())
	}
	sb.WriteString("/")
	sb.WriteString(p.name)
	return sb.String()
}

// ---------------------------------------------------------------------
// Implementation of INode interface methods
// ---------------------------------------------------------------------

// Name returns the simple name of this file
func (pFile *FileNode) Name() string {
	return pFile.name
}

// Size returns the size of this file
func (pFile *FileNode) Size() int {
	return pFile.size
}
