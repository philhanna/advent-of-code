package main

import (
	"fmt"
)

func main() {
	root := NewDirNode(nil, "/")
	fmt.Printf("DEBUG: root=%s\n", root)
}