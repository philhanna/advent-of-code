package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadStackLines(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{
			"dummy file", "testdata/dummy.txt", []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := LoadStackLines(tt.filename)
			assert.Nil(t, err)
			assert.Equal(t, len(tt.want), len(lines))
			assert.Equal(t, tt.want, lines)
		})
	}
}

func TestShip_MakeMove(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		move     Move
	}{
		{"single move", "testdata/moves.txt", Move{1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps, err := LoadShip(tt.filename)
			assert.Nil(t, err)
			fmt.Printf("\nBEFORE:\n")
			fmt.Printf("%s\n", ps)
			ps.MakeMove(tt.move)
			fmt.Printf("\nAFTER:\n")
			fmt.Printf("%s\n", ps)
		})
	}
}

func TestPeel(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		expectedC    string
		expectedLine string
	}{
		{"abc", "abc", "c", "ab"},
		{"a", "a", "a", ""},
		{"empty", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, line := Peel(tt.line)
			assert.Equal(t, tt.expectedC, c)
			assert.Equal(t, tt.expectedLine, line)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			"dummy file", "testdata/dummy.txt",
			"1: [Z] [N]\n" +
				"2: [M] [C] [D]\n" +
				"3: [P]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps, err := LoadShip(tt.filename)
			assert.Nil(t, err)
			s := ps.String()
			assert.Equal(t, tt.want, s)
			fmt.Printf("%s\n", ps)
		})
	}
}

func TestTransposeLines(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{"4x3", []string{"abc", "def", "ghi", "jkl"}, []string{"adgj", "behk", "cfil"}},
		{"full",
			[]string{
				"[G]                 [D] [R]        ",
				"[W]         [V]     [C] [T] [M]    ",
				"[L]         [P] [Z] [Q] [F] [V]    ",
				"[J]         [S] [D] [J] [M] [T] [V]",
				"[B]     [M] [H] [L] [Z] [J] [B] [S]",
				"[R] [C] [T] [C] [T] [R] [D] [R] [D]",
				"[T] [W] [Z] [T] [P] [B] [B] [H] [P]",
				"[D] [S] [R] [D] [G] [F] [S] [L] [Q]",
				" 1   2   3   4   5   6   7   8   9 ",
			},
			[]string{
				"[[[[[[[[ ",
				"GWLJBRTD1",
				"]]]]]]]] ",
				"         ",
				"     [[[ ",
				"     CWS2",
				"     ]]] ",
				"         ",
				"    [[[[ ",
				"    MTZR3",
				"    ]]]] ",
				"         ",
				" [[[[[[[ ",
				" VPSHCTD4",
				" ]]]]]]] ",
				"         ",
				"  [[[[[[ ",
				"  ZDLTPG5",
				"  ]]]]]] ",
				"         ",
				"[[[[[[[[ ",
				"DCQJZRBF6",
				"]]]]]]]] ",
				"         ",
				"[[[[[[[[ ",
				"RTFMJDBS7",
				"]]]]]]]] ",
				"         ",
				" [[[[[[[ ",
				" MVTBRHL8",
				" ]]]]]]] ",
				"         ",
				"   [[[[[ ",
				"   VSDPQ9",
				"   ]]]]] ",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.lines
			expected := tt.want
			actual := TransposeLines(input)
			assert.Equal(t, expected, actual)
		})
	}
}
