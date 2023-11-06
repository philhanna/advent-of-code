package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigLines(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{"dummy file", "testdata/dummy.txt", []string{"[T]     [Z]", "[D] [S] [R]", " 1   2   3 "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := LoadConfigLines(tt.filename)
			assert.Nil(t, err)
			assert.Equal(t, len(tt.want), len(lines))
			assert.Equal(t, tt.want, lines)
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