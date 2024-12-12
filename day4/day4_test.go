package day4

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"regexp"
	"testing"
)

func Test_toMatrix(t *testing.T) {
	m, nr, nc := toMatrix("abc\ndef")
	assert.Equal(t, 2, len(m))
	assert.Equal(t, 2, nr)
	assert.Equal(t, 3, len(m[0]))
	assert.Equal(t, 3, nc)
	assert.Equal(t, 'a', m[0][0])
	assert.Equal(t, 'f', m[1][2])
}

func Test_rotate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"ab", "a\nb"},
		{"abc", "a\nb\nc"},
		{"a\nb", "ab"},
		{"a\nb\nc", "abc"},
		{"AB\nCD\nDF", "ACD\nBDF"},
		{"ABC\nDEF\nGHI", "ADG\nBEH\nCFI"},
		{".X.\n.M.\n.A.\n.S.", "....\nXMAS\n...."},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(t, stripWhitespace(test.expected), stripWhitespace(rotate(test.input)))
		})
	}
}

func Test_diags1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"abc", "a\nb\nc"},
		// ab
		// cd
		{"ab\ncd", "a\nbc\nd"},
		// abc
		// def
		// ghi
		{"abc\ndef\nghi", "a\nbd\nceg\nfh\ni"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(t, test.expected, diag1(test.input))
		})
	}
}

func stripWhitespace(s string) string {
	re := regexp.MustCompile("[ 	]")
	return re.ReplaceAllString(s, "")
}

func Test_flip(t *testing.T) {
	input := "abc\n123"
	expected := "cba\n321"
	assert.Equal(t, expected, flipHor(input))
}

const sample1 = `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

const sample2 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_xmasSearch(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"aaa", 0},
		{"xxXMASxx", 1},
		{"xxXMASxxXMASxx", 2},
		{"..SAMX..", 1},
		{"..SAMXMAS..", 2},
		{"XMAS\nXMAS\nXMAS", 3},
		{".X.\n.M.\n.A.\n.S.", 1},
		{".S.\n.A.\n.M.\n.X.", 1},
		// ...X
		// ..M.
		// .A..
		// S...
		{"...X\n..M.\n.A..\nS...", 1},
		{"X...\n.M..\n..A.\n...S", 1},
		{sample1, 4},
		{sample2, 18},
		{day1.ReadFile("day4.txt"), 2642},
	}
	for _, test := range tests {
		name := test.input[:min(20, len(test.input))]
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, SearchXmas(test.input))
		})
	}
}

func Test_crossMasSearch(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"aaa", 0},
		{"S.M\n.A.\nS.S", 0},
		{"M.M\n.A.\nS.S", 1},
		{"S.S\n.A.\nM.M", 1},
		{"S.M\n.A.\nS.M", 1},
		{"M.S\n.A.\nM.S", 1},
		{".M.S\n..A.\n.M.S", 1},
		{sample2, 9},
		{day1.ReadFile("day4.txt"), 1974},
	}
	for _, test := range tests {
		name := test.input[:min(20, len(test.input))]
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, SearchCrossMas(test.input))
		})
	}
}
