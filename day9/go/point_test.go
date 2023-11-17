package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_touches(t *testing.T) {
	tests := []struct {
		name string
		row1 int
		col1 int
		row2 int
		col2 int
		want bool
	}{
		{"Same point", 3, 4, 3, 4, true},
		{"One to the right", 2, 3, 2, 4, true},
		{"One to the left", 2, 3, 2, 2, true},
		{"Up one row", 2, 3, 1, 3, true},
		{"Down one row", 2, 3, 3, 3, true},
		{"Diagonally adjacent", 1, 5, 2, 6, true},
		{"Two cols apart", 1, 1, 1, 3, false},
		{"Too far apart", 3, 5, 1, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewPoint(tt.row1, tt.col1)
			tail := NewPoint(tt.row2, tt.col2)
			expected := tt.want
			actual := head.Touches(tail)
			assert.Equal(t, expected, actual)
		})
	}
}
