package main

import (
	"fmt"
)

const (
	FILENAME = "input.txt"
	TOTAL_DISK_SPACE = 70000000
	REQUIRED_SPACE = 30000000
)

func main() {
	context := NewContext()
	err := context.HandleInput(FILENAME)
	if err != nil {
		fmt.Printf("Could not read %s: %s\n", FILENAME, err)
	}

	var (
		smallestDir *DirNode = nil
		smallestSize = TOTAL_DISK_SPACE + 1
		outermostDirSize = context.root.Size()
		unusedSize = TOTAL_DISK_SPACE - outermostDirSize
	)

	willWork := func(dir *DirNode) bool {
		return unusedSize + dir.Size() >= REQUIRED_SPACE
	}

	context.root.Tree(0, func(node INode, level int) {
		switch v := node.(type) {
		case *DirNode:
			if willWork(v) {
				if v.Size() <= smallestSize {
					smallestDir = v
					smallestSize = v.Size()
				}
			}	
		case *FileNode:
		default:
			fmt.Printf("BUG: Unknown type %s\n", v)
		}
	})

	fmt.Printf("Smallest directory=%s\n", smallestDir.FullPath())
	fmt.Printf("Smallest size=%d\n", smallestSize)
}
