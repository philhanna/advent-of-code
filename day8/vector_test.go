package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector_Reverse(t *testing.T) {
	tests := []struct {
		name   string
		vector Vector
		want   Vector
	}{
		{"empty", Vector{}, Vector{}},
		{"odd", Vector{1, 2, 3}, Vector{3, 2, 1}},
		{"even", Vector{1, 2, 3, 4}, Vector{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := tt.vector.Reverse()
			assert.Equal(t, want, have)
		})
	}
}
