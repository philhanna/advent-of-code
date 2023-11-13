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

func TestData_GetCol(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name      string
		colNumber int
		want      Vector
	}{
		{"left", 0, []byte{3, 2, 6, 3, 3}},
		{"interior", 1, []byte{0, 5, 5, 3, 5}},
		{"right", 4, []byte{3, 2, 2, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := data.GetCol(tt.colNumber)
			want := tt.want
			assert.Equal(t, want, have)
		})
	}
}

func TestData_GetReversedCol(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name      string
		colNumber int
		want      Vector
	}{
		{"left", 0, []byte{3, 3, 6, 2, 3}},
		{"interior", 1, []byte{5, 3, 5, 5, 0}},
		{"right", 4, []byte{0, 9, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := data.GetReversedCol(tt.colNumber)
			want := tt.want
			assert.Equal(t, want, have)
		})
	}
}

func TestData_GetRow(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name      string
		rowNumber int
		want      Vector
	}{
		{"top", 0, []byte{3, 0, 3, 7, 3}},
		{"interior", 2, []byte{6, 5, 3, 3, 2}},
		{"bottom", 4, []byte{3, 5, 3, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := data.GetRow(tt.rowNumber)
			want := tt.want
			assert.Equal(t, want, have)
		})
	}
}

func TestData_GetReversedRow(t *testing.T) {
	data := LoadData("testdata/sample.dat")

	tests := []struct {
		name      string
		rowNumber int
		want      Vector
	}{
		{"top", 0, []byte{3, 7, 3, 0, 3}},
		{"interior", 2, []byte{2, 3, 3, 5, 6}},
		{"bottom", 4, []byte{0, 9, 3, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := data.GetReversedRow(tt.rowNumber)
			want := tt.want
			assert.Equal(t, want, have)
		})
	}
}

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
