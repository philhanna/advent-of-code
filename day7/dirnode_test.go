package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirNode_String(t *testing.T) {
	root := NewRootNode()
	assert.Equal(t, "/", root.String())
	a := root.AddDirectory("a")
	assert.Equal(t, "/a", a.String())
	e := a.AddDirectory("e")
	assert.Equal(t, "/a/e", e.String())
}
