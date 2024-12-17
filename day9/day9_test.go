package day9

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"testing"
)

func Test_parseDisk(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected disk
	}{
		{
			name:     "simple",
			input:    "12345",
			expected: []block{0, e, e, 1, 1, 1, e, e, e, e, 2, 2, 2, 2, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, parseDisk(test.input))
		})
	}
}

func Test_parseSolution(t *testing.T) {
	expected := disk{0, 0, 1, 1, e, e}
	assert.Equal(t, expected, parseSolution("0011.."))
}

const sample = "2333133121414131402"

func Test_compact(t *testing.T) {
	tests := []struct {
		name     string
		input    disk
		steps    int
		expected disk
	}{
		{
			name:     "1 step",
			input:    parseDisk("12345"),
			steps:    1,
			expected: []block{0, 2, e, 1, 1, 1, e, e, e, e, 2, 2, 2, 2, e},
		},
		{
			name:     "2 steps",
			input:    parseDisk("12345"),
			steps:    2,
			expected: []block{0, 2, 2, 1, 1, 1, e, e, e, e, 2, 2, 2, e, e},
		},
		{
			name:     "unlimited steps",
			input:    parseDisk("12345"),
			steps:    -1,
			expected: []block{0, 2, 2, 1, 1, 1, 2, 2, 2, e, e, e, e, e, e},
		},
		{
			name:     "sample",
			input:    parseDisk(sample),
			steps:    -1,
			expected: parseSolution("0099811188827773336446555566.............."),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compact(test.input, test.steps)
			assert.Equal(t, test.expected, test.input)
		})
	}
}

func Test_checksum(t *testing.T) {
	assert.Equal(t, 11, checksum(parseSolution("12.3")))
}

func Test_part1(t *testing.T) {
	d := parseDisk(day1.ReadFile("day9.txt"))
	compact(d, -1)
	assert.Equal(t, 6334655979668, checksum(d))
}

func Test_compactWholeFiles(t *testing.T) {
	tests := []struct {
		name     string
		input    disk2
		steps    int
		expected disk2
	}{
		{
			name:     "no room",
			input:    parseDisk2("123"),
			steps:    -1,
			expected: []blockSpan{file(0, 1), empty(2), file(1, 3)},
		},
		{
			name:     "fits exactly",
			input:    parseDisk2("122"),
			steps:    -1,
			expected: []blockSpan{file(0, 1), file(1, 2), empty(2)},
		},
		{
			name:     "fits with extra room",
			input:    parseDisk2("132"),
			steps:    -1,
			expected: []blockSpan{file(0, 1), file(1, 2), empty(1), empty(2)},
		},
		//{
		//	name:     "compact empties",
		//	input:    parseDisk2("1821"),
		//	steps:    -1,
		//	expected: []blockSpan{file(0, 1), file(1, 2), empty(7)},
		//},
		//{
		//	name:  "sample",
		//	input: parseDisk2(sample),
		//	steps: -1,
		//	// 00992 111 777.44.333....5555.6666.....8888..
		//	expected: []blockSpan{
		//		file(0, 2),
		//		file(9, 2),
		//		file(2, 1),
		//		file(1, 3),
		//		file(7, 3),
		//		empty(1),
		//		file(4, 2),
		//		empty(1),
		//		file(3, 3),
		//		empty(4),
		//		file(5, 4),
		//		empty(1),
		//		file(6, 4),
		//		empty(5),
		//		file(8, 4),
		//		empty(2),
		//	},
		//},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//assert.Equal(t, test.expected, compact2(test.input))
			assert.Equal(t, test.expected.String(), compact2(test.input).String())
		})
	}
}
