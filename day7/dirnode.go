package day7

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type DirNode struct {
	name     string
	children []*Node
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------
func NewDirNode(name string) *DirNode {
	p := new(DirNode)
	p.name = name
	p.children = make([]*Node, 0)
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------
func (pNode *DirNode) Name() string {
	return pNode.name
}

func (pNode *DirNode) IsDir() bool {
	return true
}

func (pNode *DirNode) Size() int {
	total := 0
	for _, child := range pNode.Children() {
		total += (*child).Size()
	}
	return total
}

func (pNode *DirNode) Children() []*Node {
	return pNode.children
}
