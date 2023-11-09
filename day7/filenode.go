package day7

type FileNode struct {
	name string
	size int
}

func NewFileNode(name string, size int) *FileNode {
	p := new(FileNode)
	p.name = name
	p.size = size
	return p
}

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