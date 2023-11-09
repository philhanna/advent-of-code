package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirNode_String(t *testing.T) {
	root := NewDirNode(nil, "/")
	assert.Equal(t, "/", root.String())
	a := NewDirNode(root, "a")
	assert.Equal(t, "/a", a.String())
	e := NewDirNode(a, "e")
	assert.Equal(t, "/a/e", e.String())
}
