package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeek(t *testing.T) {
	stack := NewStack("ABC")

	topCrate := stack.Peek()
	assert.Equal(t, 'C', topCrate)
	assert.Equal(t, 3, stack.Length())

	stack = NewStack("")
	assert.Panics(t, func() {
		stack.Peek()
	})
}

func TestPop(t *testing.T) {
	stack := NewStack("ABC")

	topCrate := stack.Pop()
	assert.Equal(t, 'C', topCrate)
	expected := NewStack("AB")
	assert.Equal(t, expected, stack)

	topCrate = stack.Pop()
	assert.Equal(t, 'B', topCrate)
	expected = NewStack("A")
	assert.Equal(t, expected, stack)

	topCrate = stack.Pop()
	assert.Equal(t, 'A', topCrate)
	expected = NewStack("")
	assert.Equal(t, expected, stack)

	assert.Panics(t, func() {
		stack.Pop()
	})

}

func TestStack_Push(t *testing.T) {
	stack := NewStack("ABC")
	stack.Push('D')
	assert.Equal(t, "ABCD", stack.Contents)
}
