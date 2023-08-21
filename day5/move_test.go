package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMove(t *testing.T) {
	tests := []struct {
		name     string
		movestr  string
		expected []int
	}{
		{"Single digits", "move 1 from 2 to 3", []int{1, 2, 3}},
		{"Multiple digits", "move 12 from 1 to 16", []int{12, 1, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMove(tt.movestr)
			assert.Equal(t, tt.expected[0], actual.Count)
			assert.Equal(t, tt.expected[1], actual.FromStack)
			assert.Equal(t, tt.expected[2], actual.ToStack)
		})
	}
}
