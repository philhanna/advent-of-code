package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigLines(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{"dummy file", "testdata/dummy.txt", []string{"[T]     [Z]", "[D] [S] [R]", " 1   2   3 "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := LoadConfigLines(tt.filename)
			assert.Nil(t, err)
			assert.Equal(t, len(tt.want), len(lines))
			assert.Equal(t, tt.want, lines)
		})
	}
}

func TestTransposeLines(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"4x3", []string{"abc", "def", "ghi", "jkl",}, []string{"adgj", "behk", "cfil"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := TransposeLines(tt.lines)
			assert.Equal(t, tt.want, have)
		})
	}
}
