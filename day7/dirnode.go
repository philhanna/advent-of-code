package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// DirNode is a Node that represents a directory.
type DirNode struct {
	*Node
	name     string
	parent   *DirNode
	children []*Node
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDirNode creates a new directory node with the specified parent
// directory and name.
func NewDirNode(parent *DirNode, name string) *DirNode {
	node := new(Node)
	dirNode := &DirNode{node, name, parent, nil}
	dirNode.children = make([]*Node, 0)
	if parent != nil {
		parent.children = append(parent.children, node)
	}
	node.INode = dirNode
	return dirNode
}

// ---------------------------------------------------------------------
// Implementation of INode interface methods
// ---------------------------------------------------------------------

// Name returns the directory name
func (pDir *DirNode) Name() string {
	return pDir.name
}

// Size returns the size of this directory, which is the recursive sum
// of the sizes of all its children.
func (pDir *DirNode) Size() int {
	total := 0
	for _, child := range pDir.children {
		part := child.Size()
		total += part
	}
	return total
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// HasChild returns true if this directory has the specified node as a
// direct child.
func (p *DirNode) HasChild(childName string) bool {
	for _, child := range p.children {
		if childName == child.Name() {
			return true
		}
	}
	return false
}

// String returns a string representation of this directory node
func (p *DirNode) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("name=%s", p.Name()))
	parts = append(parts, fmt.Sprintf("path=%s", p.FullPath()))
	if p.parent == nil {
		parts = append(parts, "parent=nil")
	} else {
		parts = append(parts, fmt.Sprintf("parent=%s", p.parent.Name()))
	}
	kids := make([]string, 0)
	for _, child := range p.children {
		kids = append(kids, child.Name())
	}
	kidString := "children=[" + strings.Join(kids, ",") + "]"
	parts = append(parts, kidString)

	result := strings.Join(parts, ",")
	return result
}

// FullPath returns the full path for this directory node
func (p *DirNode) FullPath() string {
	sb := strings.Builder{}
	if p.parent == nil {
		sb.WriteString("/")
	} else {
		if p.parent.parent != nil {
			sb.WriteString(p.parent.FullPath())
		}
		sb.WriteString("/")
		sb.WriteString(p.Name())
	}
	return sb.String()
}
