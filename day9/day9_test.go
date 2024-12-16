package day9

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"math"
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

func parseSolution(s string) disk {
	var result disk
	for _, val := range s {
		value := block(val - '0')
		if val == '.' {
			value = emptyBlock
		}
		result = append(result, value)
	}
	return result
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
		//{
		//	name:     "real",
		//	input:    parseDisk(day1.ReadFile("day9.txt")),
		//	steps:    -1,
		//	expected: parseSolution("0099811188827773336446555566.............."),
		//},
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
	assert.Equal(t, 1, checksum(d))
}

func checksum(solution disk) int {
	var result int
	for i, val := range solution {
		if val != emptyBlock {
			result += i * int(val)
		}
	}
	return result
}

func compact(d disk, steps int) {
	if steps == -1 {
		steps = math.MaxInt
	}
	left, right := 0, len(d)-1
	for i := 0; i < steps && left < right; {
		if d[left] != e {
			left++
		} else if d[right] == e {
			right--
		} else {
			d[left], d[right] = d[right], d[left]
			i++
		}
	}
}
