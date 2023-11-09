package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirNode_String(t *testing.T) {
	root := NewRootNode()
	
	rootString := root.String()
	assert.Equal(t, "/", rootString)

	a := root.AddDirectory("a")
	aString := a.String()
	assert.Equal(t, "/a", aString)

	e := a.AddDirectory("e")
	eString := e.String()
	assert.Equal(t, "/a/e", eString)
}
