package main

import "fmt"

// ---------------------------------------------------------------------
// Mainline
// ---------------------------------------------------------------------

func main() {
	root := NewDirNode(nil, "/")
	a := NewDirNode(root, "a")
	e := NewDirNode(a, "e")

	file := NewFileNode(e, "bfile.txt", 123)
	rootFile := NewFileNode(root, "rootFile.txt", 86)

	fmt.Printf("root: %s\n", root)
	fmt.Printf("rootFile: %s\n", rootFile)
	fmt.Printf("a   : %s\n", a)
	fmt.Printf("e   : %s\n", e)
	fmt.Printf("file: %s\n", file)
}
