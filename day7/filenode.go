package main

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
// Methods
// ---------------------------------------------------------------------

// Name returns the file name
func (pNode *FileNode) Name() string {
	return pNode.name
}

// Size returns the file size
func (pNode *FileNode) Size() int {
	return pNode.size
}

// Children returns nil (only directories have children)
func (pNode *FileNode) Children() []*AbstractNode {
	return nil
}
