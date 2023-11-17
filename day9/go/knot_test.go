package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnot_Move(t *testing.T) {
	tests := []struct {
		name      string
		row       int
		col       int
		direction string
		wantRow   int
		wantCol   int
	}{
		{"U", 1, 5, "U", 0, 5},
		{"D", 1, 5, "D", 2, 5},
		{"L", 1, 5, "L", 1, 4},
		{"R", 1, 5, "R", 1, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := NewPoint(tt.row, tt.col)
			head := NewKnot("Head", point)
			head.Move(tt.direction)
			assert.Equal(t, tt.wantRow, head.point.row)
			assert.Equal(t, tt.wantCol, head.point.col)
		})
	}
}

func TestKnot_Follow(t *testing.T) {
	tests := []struct {
		name      string
		hrow      int
		hcol      int
		trow      int
		tcol      int
		direction string
		wantRow   int
		wantCol   int
	}{
		{"tail starts one to the left, head moves right", 1, 2, 1, 1, "R", 1, 2},
		{"Same starting point", 1, 3, 1, 3, "U", 1, 3},
		{"Same starting point", 1, 3, 1, 3, "D", 1, 3},
		{"Same starting point", 1, 3, 1, 3, "L", 1, 3},
		{"Same starting point", 1, 3, 1, 3, "R", 1, 3},
		{"tail starts one to the left, head moves up", 2, 6, 2, 5, "U", 2, 5},
		{"tail starts one to the right, head moves up", 2, 6, 2, 7, "U", 2, 7},
		{"tail starts one up, head moves up", 2, 6, 1, 6, "U", 1, 6},
		{"tail southeast one, head moves up", 1, 6, 2, 5, "U", 1, 6},
		{"tail southeast, head moves up", 3, 3, 4, 2, "U", 3, 3},
		{"tail southeast, head moves right", 3, 3, 4, 2, "R", 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewKnot("H", NewPoint(tt.hrow, tt.hcol))
			tail := NewKnot("T", NewPoint(tt.trow, tt.tcol))
			head.Move(tt.direction)
			tail.Follow(head, tt.direction)
			assert.Equal(t, tt.wantRow, tail.point.row)
			assert.Equal(t, tt.wantCol, tail.point.col)
		})
	}
}
