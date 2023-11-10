package main

import "fmt"

const (
	FILENAME = "input.txt"
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