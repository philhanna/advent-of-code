package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirNode_HasChild(t *testing.T) {
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

	tests := []struct {
		name   string
		parent *DirNode
		childName string
		want   bool
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
			assert.Equal(t, tt.want, tt.parent.HasChild(tt.childName))
		})
	}
}
