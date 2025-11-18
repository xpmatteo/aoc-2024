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
			name:     "1 towel composed of 2 patterns",
			towels:   []string{"wr", "bb"},
			patterns: []string{"wrbb"},
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, solvePart1(test.towels, test.patterns))
		})
	}
}
