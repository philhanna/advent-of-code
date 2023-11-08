package day7

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type FileNode struct {
	name string
	size int
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

func NewFileNode(name string, size int) *FileNode {
	p := new(FileNode)
	p.name = name
	p.size = size
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------
func (pNode *FileNode) Name() string {
	return pNode.name
}

func (pNode *FileNode) IsDir() bool {
	return false
}

func (pNode *FileNode) Size() int {
	return pNode.size
}

func (pNode *FileNode) Children() []*Node {
	return nil
}