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
	
	var (
		err error
		context *InputContext
	)

	context = NewContext()
	err = context.HandleCD("$ cd a")
	assert.Nil(t, err)
	assert.Equal(t, "/a", context.cwd.FullPath())
}
