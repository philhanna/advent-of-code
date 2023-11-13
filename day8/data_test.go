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
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		data Data
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.VisibleUp(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("Data.VisibleUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
