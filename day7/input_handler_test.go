package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const SAMPLE_INPUT = "sampleinput.txt"

func TestHandleInput(t *testing.T) {
	context := NewContext()
	err := context.HandleInput(SAMPLE_INPUT)
	assert.Nil(t, err)
	assert.NotNil(t, context)
	assert.NotNil(t, context.root)
	assert.NotNil(t, context.cwd)
}

func TestInputContext_HandleLS(t *testing.T) {
	context := NewContext()
	err := context.HandleLS("$ ls")
	assert.Nil(t, err)
}

func TestInputContext_HandleCD(t *testing.T) {
	context := NewContext()
	err := context.HandleCD("$ cd a")
	assert.Nil(t, err)
	assert.Equal(t, "/a", context.cwd.FullPath())
}

func TestInputContext_HandleCDRoot(t *testing.T) {
	context := NewContext()
	err := context.HandleCD("$ cd a")
	assert.Nil(t, err)
	err = context.HandleCD("$ cd /")
	assert.Nil(t, err)
	assert.Equal(t, "/", context.cwd.FullPath())
}

func TestInputContext_HandleCDParent(t *testing.T) {
	var err error
	context := NewContext()
	err = context.HandleCD("$ cd a")
	assert.Nil(t, err)
	err = context.HandleCD("$ cd e")
	assert.Nil(t, err)
	err = context.HandleCD("$ cd ..")
	assert.Nil(t, err)
	assert.Equal(t, "/a", context.cwd.FullPath())
}
