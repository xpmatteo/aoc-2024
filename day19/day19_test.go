package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
// invariant
// Pattern is a list of attempts
// an attempt is a list of strings to match
//
// r, wr, b, g, bwu, rb, gb, br
//
// brwrr -> rwrr -> wrr -> r -> OK
//          wrr -> r -> OK
//

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		towels   []string
		patterns []string
		expected int
	}{
		{
			name:     "1 towel 1 pattern ok",
			towels:   []string{"wr"},
			patterns: []string{"wr"},
			expected: 1,
		},
		{
			name:     "1 towel 2 pattern ok",
			towels:   []string{"wr"},
			patterns: []string{"wr", "wr"},
			expected: 2,
		},
		{
			name:     "1 towel 1 pattern ok, 1 pattern not ok",
			towels:   []string{"wr"},
			patterns: []string{"wr", "xx"},
			expected: 1,
		},
		{
			name:     "2 towel 2 pattern ok",
			towels:   []string{"wr", "bb"},
			patterns: []string{"bb", "bb"},
			expected: 2,
		},
		{
			name:     "1 towel , 1 pattern composed ok",
			towels:   []string{"wr", "bb"},
			patterns: []string{"wrbb"},
			expected: 1,
		},
		{
			name:     "2 towel, 2 patterns composed ok",
			towels:   []string{"wr", "bb"},
			patterns: []string{"wrbb", "bbwr"},
			expected: 2,
		},
		{
			name:     "2 towel, 3 patterns composed ok, 1 not ok",
			towels:   []string{"wr", "bb"},
			patterns: []string{"wrbb", "bbwr", "bbbb", "www"},
			expected: 3,
		},
		{
			name:     "a false start",
			towels:   []string{"wr", "w", "rxx"},
			patterns: []string{"wrxx"},
			expected: 1,
		},
		{
			name:   "small example part 1",
			towels: []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			patterns: []string{
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			expected: 6,
		},
		//{
		//	name:     "test part 1",
		//	towels:   parseTowels(day1.ReadFile("input.txt")),
		//	patterns: parsePatterns(day1.ReadFile("input.txt")),
		//	expected: 0,
		//},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, solvePart1(test.towels, test.patterns))
		})
	}
}

const sampleInput = `aa, bbb, cccc

brubbru
bopbop
aaaa
`

func Test_parseTowels(t *testing.T) {
	assert.Equal(t, []string{"aa", "bbb", "cccc"}, parseTowels(sampleInput))
}

func Test_parsePatterns(t *testing.T) {
	assert.Equal(t, []string{"brubbru", "bopbop", "aaaa"}, parsePatterns(sampleInput))
}

func TestContinuations(t *testing.T) {
	tests := []struct {
		name     string
		towels   []string
		pattern  string
		expected []string
	}{
		{
			name:     "no continuations",
			towels:   []string{"aa", "bb"},
			pattern:  "abc",
			expected: nil,
		},
		{
			name:     "one continuation",
			towels:   []string{"aa", "bb"},
			pattern:  "aapippo",
			expected: []string{"pippo"},
		},
		{
			name:     "many continuation",
			towels:   []string{"aa", "aap", "a"},
			pattern:  "aapippo",
			expected: []string{"pippo", "ippo", "apippo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, continuations(test.towels, test.pattern))
		})
	}
}
