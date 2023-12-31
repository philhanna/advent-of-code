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
	children []INode
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDirNode creates a new directory node with the specified parent
// directory and name.
func NewDirNode(parent *DirNode, name string) *DirNode {
	node := new(Node)
	dirNode := &DirNode{node, name, parent, nil}
	dirNode.children = make([]INode, 0)
	if parent != nil {
		parent.children = append(parent.children, dirNode)
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

// StringLS returns the string that represents this directory when
// listed with the ls command
func (pDir *DirNode) StringLS() string {
	return fmt.Sprintf("dir %s", pDir.name)
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// ChildrenAsStrings returns the output of "ls" on the specified
// directory
func (p *DirNode) ChildrenAsStrings() []string {
	n := len(p.children)
	list := make([]string, n)
	for i, child := range p.children {
		list[i] = child.StringLS()
	}
	return list
}

// DeleteChild removes the child of the specified name, if it exists.
// If the child is a directory, removes all its children first.
func (p *DirNode) DeleteChild(name string) {
	childNode := p.LookupChild(name)
	if childNode == nil {
	} else {
		switch v := childNode.(type) {
		case *DirNode:
			for _, child := range v.children {
				v.DeleteChild(child.Name())
			}
		case *FileNode:
		default:
		}
		newList := make([]INode, 0)
		for _, child := range p.children {
			if child != childNode {
				newList = append(newList, child)
			}
		}
		p.children = newList
	}
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

// LookupChild looks through the children of this directory node to find
// a child with the specified name.  If found, returns that child,
// otherwise returns nil
func (p *DirNode) LookupChild(childName string) INode {
	for _, child := range p.children {
		if childName == child.Name() {
			return child
		}
	}
	return nil
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

// Tree applies a function to this directory then recursively applies
// the function to each of its children
func (p *DirNode) Tree(level int, apply func(INode, int)) {
	apply(p, level)
	level++
	for _, child := range p.children {
		switch v := child.(type) {
		case *DirNode:
			v.Tree(level, apply)
		case *FileNode:
			apply(v, level)
		default:
			fmt.Printf("BUG: Unknown file type %s for %s\n", v, child)
		}
	}

}
