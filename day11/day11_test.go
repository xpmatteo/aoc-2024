package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseStones(t *testing.T) {
	tests := []struct {
		stones   string
		expected StoneList
	}{
		{
			stones:   "1",
			expected: make(StoneList).Add(1, 1),
		},
		{
			stones:   "1 2 3",
			expected: make(StoneList).Add(1, 1).Add(2, 1).Add(3, 1),
		},
		{
			stones:   "1 1 1",
			expected: make(StoneList).Add(1, 3),
		},
	}
	for _, test := range tests {
		t.Run(test.stones, func(t *testing.T) {
			assert.Equal(t, test.expected, parseStones(test.stones))
		})
	}
}

func Test_StoneList_Size(t *testing.T) {
	tests := []struct {
		stones   string
		expected int
	}{
		{
			stones:   "1",
			expected: 1,
		},
		{
			stones:   "1 1 2",
			expected: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.stones, func(t *testing.T) {
			assert.Equal(t, test.expected, parseStones(test.stones).Size())
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         StoneList
		steps         int
		expected      StoneList
		expectedCount int
	}{
		{
			name:          "first example",
			input:         parseStones("0 1 10 99 999"),
			steps:         1,
			expected:      parseStones("1 2024 1 0 9 9 2021976"),
			expectedCount: 7,
		},
		{
			name:          "iterate 2",
			input:         parseStones("0"),
			steps:         2,
			expected:      parseStones("2024"),
			expectedCount: 1,
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
			name:          "real 25",
			input:         parseStones("0 89741 316108 7641 756 9 7832357 91"),
			steps:         25,
			expectedCount: 193899,
		},
		{
			name:     "repeated numbers",
			input:    parseStones("0 0 0 0"),
			steps:    1,
			expected: parseStones("1 1 1 1"),
		},
		{
			name:          "real 75",
			input:         parseStones("0 89741 316108 7641 756 9 7832357 91"),
			steps:         75,
			expectedCount: 229682160383225,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evolved := blink1(test.input, test.steps)
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected, evolved)
			}
			if test.expectedCount > 0 {
				assert.Equal(t, test.expectedCount, evolved.Size())
			}
		})
	}
}

func Test_StoneList_Add(t *testing.T) {
	sl := make(StoneList)
	sl.Add(1, 3)
	sl.Add(1, 4)

	assert.Equal(t, 0, sl[0])
	assert.Equal(t, 7, sl[1])
}
