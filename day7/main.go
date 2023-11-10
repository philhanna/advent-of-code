package main

import (
	"fmt"
	"strings"
)

const (
	FILENAME = "sampleinput.txt"
	LIMIT    = 100000
)

var (
	totalSize = 0
)

func main() {
	context := NewContext()
	err := context.HandleInput(FILENAME)
	if err != nil {
		fmt.Printf("Could not read %s: %s\n", FILENAME, err)
	}
	Tree(context.root, 0)
}

func Tree(dir *DirNode, level int) {
	indent := func(level int) string {
		return strings.Repeat(" ", 4*level)
	}
	fmt.Printf("%s%s\n", indent(level), dir.Name())
	for _, child := range dir.children {
		switch v := child.(type) {
		case *DirNode:
			Tree(v, level + 1)
		case *FileNode:
			fmt.Printf("%s%s\n", indent(level+1), v.Name())
		default:
			fmt.Printf("BUG: Unknown file type %s for %s\n", v, child)
		}
	}
}

func Main2() {
	context := NewContext()
	err := context.HandleInput(FILENAME)
	if err != nil {
		fmt.Printf("Could not read %s: %s\n", FILENAME, err)
	}

	HandleDirectory(context.root)
	fmt.Printf("Total size is %d\n", totalSize)
}

func HandleDirectory(dir *DirNode) {
	if dir.Size() <= LIMIT {
		Accept(dir)
	}
	for _, child := range dir.children {
		switch v := child.(type) {
		case *DirNode:
			HandleDirectory(v)
		case *FileNode:
		default:
			fmt.Printf("Node %s, unknown type %T\n", child, v)
		}
	}
}

func Accept(dir *DirNode) {
	fmt.Printf("Accepted directory %s of size %d\n", dir.FullPath(), dir.Size())
	totalSize += dir.Size()
}
