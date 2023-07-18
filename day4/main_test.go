package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		name  string
		pair  string
		want1 int
		want2 int
	}{
		{"same", "3-3", 3, 3},
		{"diffOf2", "3-5", 3, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n1, n2 := Expand(tt.pair)
			assert.Equal(t, tt.want1, n1)
			assert.Equal(t, tt.want2, n2)
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		range1 string
		range2 string
		want   bool
	}{
		{"contains", "2-6", "3-4", true},
		{"starts with same", "3-6", "3-4", true},
		{"not contains", "3-4", "2-6", false},
		{"exact", "10-50", "10-50", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := Contains(tt.range1, tt.range2)
			want := tt.want
			assert.Equal(t, have, want)
		})
	}
}
