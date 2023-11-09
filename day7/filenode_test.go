package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileNode_Size(t *testing.T) {
	tests := []struct {
		name     string
		filesize int
	}{
		{"86 bytes", 86},
		{"0 bytes", 0},
		{"-1 bytes", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := NewContext()
			file := NewFileNode(context.cwd, "theFile", tt.filesize)
			assert.Equal(t, tt.filesize, file.Size())
		})
	}
}
