package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name     string
		row      int
		col      int
		expected byte
	}{
		{"topLeft", 0, 0, 3},
		{"bottomRight", 4, 4, 0},
		{"interior", 3, 2, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, len(data) >= tt.row)
			rowData := data[tt.row]
			assert.True(t, len(rowData) >= tt.col)
			actual := rowData[tt.col]
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestData_EdgeTreesAllVisible(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	nRows := len(data)
	nCols := nRows

	for col := 0; col < nCols; col++ {
		assert.True(t, data.VisibleUp(0, col))
		assert.True(t, data.VisibleDown(nRows-1, col))
	}

	for row := 0; row < nRows; row++ {
		assert.True(t, data.VisibleLeft(row, 0))
		assert.True(t, data.VisibleRight(row, nCols-1))
	}
}

func TestData_VisibleUp(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name string // test name
		row  int    // tree row
		col  int    // tree column
		want bool
	}{
		{"1,1", 1, 1, true},
		{"2,1", 2, 1, false},
		{"2,2", 2, 2, false},
		{"4,1", 4, 1, false},
		{"4,3", 4, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, data.VisibleUp(tt.row, tt.col))
		})
	}
}

func TestData_VisibleDown(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name string // test name
		row  int    // tree row
		col  int    // tree column
		want bool
	}{
		{"1, 2", 1, 2, false},
		{"2, 0", 2, 0, true},
		{"3, 1", 3, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, data.VisibleDown(tt.row, tt.col))
		})
	}
}

func TestData_VisibleLeft(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name string // test name
		row  int    // tree row
		col  int    // tree column
		want bool
	}{
		{"0, 1", 0, 1, false},
		{"0, 2", 0, 2, false},
		{"0, 3", 0, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, data.VisibleLeft(tt.row, tt.col))
		})
	}
}
