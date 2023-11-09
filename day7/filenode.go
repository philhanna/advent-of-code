package main

import "strings"

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// FileNode is a node representing a file
type FileNode struct {
	*AbstractNode
	parent *DirNode
	name   string
	size   int
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

func (pNode *FileNode) Name() string {
	return pNode.name
}

func (pNode *FileNode) Size() int {
	return pNode.size
}

func (pNode *FileNode) Children() []*AbstractNode {
	return nil
}

func (p *FileNode) String() string {
	sb := strings.Builder{}
	if p.parent == nil {
		sb.WriteString("/")
	} else {
		if p.parent.parent != nil {
			sb.WriteString(p.parent.String())
		}
		sb.WriteString("/")
		sb.WriteString(p.Name())
	}
	return sb.String()
}
