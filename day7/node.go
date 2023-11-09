package main

import "fmt"

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type INode interface {
	Name() string     // Returns the node name
	Parent() *DirNode // Returns the parent directory node
	Size() int        // Returns the file or directory size
}

type Node struct {
	INode
}

type DirNode struct {
	*Node
	name     string
	parent   *DirNode
	children []*Node
}

type FileNode struct {
	*Node
	name string
}

// ---------------------------------------------------------------------
// Functions and methods
// ---------------------------------------------------------------------

func (pDir *DirNode) Name() string {
	return pDir.name
}

func (pFile *FileNode) Name() string {
	return pFile.name
}

func NewDirNode(theName string) *DirNode {
	a := new(Node)
	r := &DirNode{a, theName, nil, nil}
	a.INode = r
	return r
}

func NewFileNode(theName string) *FileNode {
	a := new(Node)
	r := &FileNode{a, theName}
	a.INode = r
	return r
}

func main() {
	var dA INode = NewDirNode("/")
	var dB INode = NewFileNode("bfile.txt")

	fmt.Printf("dA=%v\n", dA)
	fmt.Printf("dB=%v\n", dB)
}
