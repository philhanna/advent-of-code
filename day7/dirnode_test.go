package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetTestData() (*DirNode, *DirNode, *DirNode, *DirNode) {
	root := NewDirNode(nil, "/")
	a := NewDirNode(root, "a")
	e := NewDirNode(a, "e")
	NewFileNode(e, "i", 584)
	NewFileNode(a, "f", 29116)
	NewFileNode(a, "g", 2557)
	NewFileNode(a, "h.lst", 62596)
	NewFileNode(root, "b.txt", 14848514)
	NewFileNode(root, "c.dat", 8504156)
	d := NewDirNode(root, "d")
	NewFileNode(d, "j", 4060174)
	NewFileNode(d, "d.log", 8033020)
	NewFileNode(d, "d.ext", 5626152)
	NewFileNode(d, "k", 7214296)
	return root, a, e, d
}

func TestDirNode_LookupChild(t *testing.T) {

	root, a, e, _ := GetTestData()

	tests := []struct {
		name      string
		parent    *DirNode
		childName string
		want      bool
	}{
		{"/ has child d", root, "d", true},
		{"/a has child e", a, "e", true},
		{"/a has child file f", a, "f", true},
		{"/a doesn't have bogus child", a, "bogus", false},
		{"/a/e has child file i", e, "i", true},
		{"/ doesn't have indirect child f", root, "f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child := tt.parent.LookupChild(tt.childName)
			if tt.want {
				assert.NotNil(t, child)
				assert.Equal(t, tt.childName, child.Name())
			} else {
				assert.Nil(t, child)
			}
		})
	}
}

func TestDirNode_ChildrenAsStrings(t *testing.T) {

	root, a, e, _ := GetTestData()

	tests := []struct {
		name     string
		dirNode  *DirNode
		expected []string
	}{
		{"Children of /", root, []string{
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
		},
		},
		{"Children of /a", a, []string{
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
		},
		},
		{"Children of /a/e", e, []string{
			"584 i",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.dirNode.ChildrenAsStrings()
			expected := tt.expected
			assert.ElementsMatch(t, expected, actual)
		})
	}
}

func TestDirNode_Size(t *testing.T) {

	root, a, e, d := GetTestData()

	tests := []struct {
		name     string
		dirNode  *DirNode
		expected int
	}{
		{"/", root, 48381165},
		{"/a", a, 94853},
		{"/a/e", e, 584},
		{"/d", d, 24933642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.dirNode.Size())
		})
	}
}
