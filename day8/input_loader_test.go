package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		row      int
		col      int
		expected byte
	}{		
		{"topLeft", "testdata/sample.dat", 0, 0, '3'},
		{"bottomRight", "testdata/sample.dat", 4, 4, '0'},
		{"interior", "testdata/sample.dat", 3, 2, '5'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := LoadInput(tt.filename)
			assert.True(t, len(data) >= tt.row)
			rowData := data[tt.row]
			assert.True(t, len(rowData) >= tt.col)
			actual := rowData[tt.col]
			assert.Equal(t, tt.expected, actual)
		})
	}
}
