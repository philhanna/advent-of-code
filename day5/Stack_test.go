package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_String(t *testing.T) {
	type fields struct {
		IDNumber int
		Crates   []string
	}
	tests := []struct {
		name     string
		fields   fields
		expected string
	}{
		{"Simple", fields{4, []string{"A", "B", "C"}}, "4: [A] [B] [C]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &Stack{
				IDNumber: tt.fields.IDNumber,
				Crates:   tt.fields.Crates,
			}
			expected := tt.expected
			actual := ps.String()
			assert.Equal(t, expected, actual)
		})
	}
}
