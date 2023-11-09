package main

import "strings"

type DirNode struct {
	parent   *DirNode
	name     string
	children []*Node
}

func NewDirNode(parent *DirNode, name string) *DirNode {
	p := new(DirNode)
	p.parent = parent
	p.name = name
	p.children = make([]*Node, 0)
	return p
}

func (pNode *DirNode) Name() string {
	return pNode.name
}

func (pNode *DirNode) IsDir() bool {
	return true
}

func (pNode *DirNode) Size() int {
	total := 0
	for _, child := range pNode.Children() {
		total += (*child).Size()
	}
	return total
}

func (pNode *DirNode) Children() []*Node {
	return pNode.children
}

func (p *DirNode) String() string {
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
