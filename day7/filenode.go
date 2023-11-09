package main

import "strings"

type FileNode struct {
	parent *DirNode
	name string
	size int
}

func NewFileNode(parent *DirNode, name string, size int) *FileNode {
	p := new(FileNode)
	p.parent = parent
	p.name = name
	p.size = size
	return p
}

func (pNode *FileNode) Name() string {
	return pNode.name
}

func (pNode *FileNode) IsDir() bool {
	return false
}

func (pNode *FileNode) Size() int {
	return pNode.size
}

func (pNode *FileNode) Children() []*Node {
	return nil
}

func (p *FileNode) String() string {
	sb := strings.Builder{}
	if p.parent != nil {
		if p.parent.parent != nil {
			sb.WriteString(p.parent.Name())
		}
		sb.WriteString("/")
	}
	sb.WriteString(p.Name())
	return sb.String()
}
