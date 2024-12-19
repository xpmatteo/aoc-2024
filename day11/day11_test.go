package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         []Stone
		steps         int
		expected      []Stone
		expectedCount int
	}{
		{
			name:     "first example",
			input:    parseStones("0 1 10 99 999"),
			steps:    1,
			expected: parseStones("1 2024 1 0 9 9 2021976"),
		},
		{
			name:     "iterate 2",
			input:    parseStones("0"),
			steps:    2,
			expected: parseStones("2024"),
		},
		{
			name:     "sample 6 blinks",
			input:    parseStones("125 17"),
			steps:    6,
			expected: parseStones("2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2"),
		},
		{
			name:          "sample 25 blinks",
			input:         parseStones("125 17"),
			steps:         25,
			expectedCount: 55312,
		},
		{
			name:          "real",
			input:         parseStones("0 89741 316108 7641 756 9 7832357 91"),
			steps:         25,
			expectedCount: 193899,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evolved := blink(test.input, test.steps)
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected, evolved)
			}
			if test.expectedCount > 0 {
				assert.Equal(t, test.expectedCount, len(evolved))
			}
		})
	}
}
