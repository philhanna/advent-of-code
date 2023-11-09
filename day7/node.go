package main

import (
	"fmt"
	// "strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// INode is the interface that DirNode and FileNode must implement.
type INode interface {
	Name() string     // Returns the node name
	Parent() *DirNode // Returns the parent directory node
	Size() int        // Returns the file or directory size
}

// Node is the abstract base type that DirNode and FileNode extend.
type Node struct {
	INode
}

// DirNode is a Node that represents a directory.
type DirNode struct {
	*Node
	name     string
	parent   *DirNode
	children []*Node
}

// FileNode is a Node that represents a file.
type FileNode struct {
	*Node
	name   string
	parent *DirNode
	size   int
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

func NewDirNode(parent *DirNode, name string) *DirNode {
	// Create the abstract base struct (Node)
	node := new(Node)

	// Create a DirNode object with that base struct
	dirNode := &DirNode{node, name, parent, nil}

	// Give it an empty list of children
	dirNode.children = make([]*Node, 0)

	// Tell the parent node to add this as a child
	if parent != nil {
		parent.children = append(parent.children, node)
	}

	// Tell the base struct that this directory is its contents
	node.INode = dirNode

	return dirNode
}

func NewFileNode(parent *DirNode, name string, size int) *FileNode {
	// Create the abstract base struct
	node := new(Node)

	// Create the FileNode object with that base struct
	fileNode := &FileNode{node, name, nil, size}

	// Tell the parent node to add this as a child
	if parent != nil {
		parent.children = append(parent.children, node)
	}

	// Tell the base struct that this file is its contents
	node.INode = fileNode

	return fileNode
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Implementation of INode interface methods
// ---------------------------------------------------------------------

func (pDir *DirNode) Name() string {
	return pDir.name
}

func (pFile *FileNode) Name() string {
	return pFile.name
}

func (pDir *DirNode) Size() int {
	total := 0
	for _, child := range pDir.children {
		part := child.Size()
		total += part
	}
	return total
}

func (pFile *FileNode) Size() int {
	return pFile.size
}

// ---------------------------------------------------------------------
// Mainline
// ---------------------------------------------------------------------

func main() {
	root := NewDirNode(nil, "/")
	a := NewDirNode(root, "a")
	e := NewDirNode(a, "e")

	file := NewFileNode(a, "bfile.txt", 123)

	fmt.Printf("root base=%v,name=%s,parent=%v\n", root.Node, root.Name(), root.parent)
	fmt.Printf("a    base=%v,name=%s,parent=%v\n", a.Node, a.Name(), a.parent)
	fmt.Printf("e    base=%v,name=%s,parent=%v\n", e.Node, e.Name(), e.parent)
	fmt.Println("file =", file)
}

/*
func (p *Node) String() string {
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

*/
