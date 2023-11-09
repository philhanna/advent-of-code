package main

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// DirNode is a node representing a directory
type DirNode struct {
	*AbstractNode
	parent   *DirNode
	name     string
	children []*AbstractNode
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// NewRootNode creates a root node with a name of "/", no parent, and an
// empty list of children
func NewRootNode() *DirNode {

	// Create the new directory node
	root := new(DirNode)
	root.parent = nil
	root.name = "/"
	root.children = make([]*AbstractNode, 0)

	// Make an instance of the base class and make this directory its
	// base node
	nodeStruct := new(AbstractNode)
	nodeStruct.Node = root

	return root
}

// AddDirectory creates a new directory node as a child of the specified
// parent node.  It works the same as NewRootNode except that it saves a
// pointer to its parent directory and adds itself to the parent
// directory's children.
func (pParent *DirNode) AddDirectory(name string) *DirNode {

	// Create the subdirectory node
	subdir := new(DirNode)
	subdir.parent = pParent
	subdir.name = name
	subdir.children = make([]*AbstractNode, 0)

	// Make an instance of the base class and make this directory its
	// base node
	nodeStruct := new(AbstractNode)
	nodeStruct.Node = subdir
	
	// Append this directory to the parent's list of children
	pParent.children = append(pParent.children, nodeStruct)

	return subdir
}

// AddFile creates a new file node as a child of the specified parent
// directory.
func (pParent *DirNode) AddFile(name string, size int) *FileNode {
	
	// Create a new File node
	file := new(FileNode)
	file.parent = pParent
	file.name = name
	file.size = size

	// Make an instance of the base class and make this file its
	// base node
	nodeStruct := new(AbstractNode)
	nodeStruct.Node = file

	// Append this file to the parent's list of children
	pParent.children = append(pParent.children, nodeStruct)
	
	return file
}

// Name returns the directory name
func (pNode *DirNode) Name() string {
	return pNode.name
}

// Parent returns the parent directory
func (pNode *DirNode) Parent() *DirNode {
	return pNode.parent
}

// Size returns the directory size, which is the sum of the sizes of its
// children (recursively, for subdirectories)
func (pNode *DirNode) Size() int {
	total := 0
	for _, child := range pNode.Children() {
		total += (*child).Size()
	}
	return total
}

// Children returns the list of this directory's child nodes
func (pNode *DirNode) Children() []*AbstractNode {
	return pNode.children
}
