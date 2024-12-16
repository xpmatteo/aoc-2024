package day6

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

const sample = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const sampleSolved = `....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XXXXXXX.
.XXXXXXX#.
#XXXXXXX..
......#X..`

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         mapping.Map
		expected      mapping.Map
		expectedCount int
	}{
		{
			name: "move up",
			input: mapping.Map{
				"...",
				"...",
				".^.",
				"...",
			},
			expected: mapping.Map{
				".X.",
				".X.",
				".X.",
				"...",
			},
			expectedCount: 3,
		},
		{
			name: "move right",
			input: mapping.Map{
				"....",
				">...",
				"....",
			},
			expected: mapping.Map{
				"....",
				"XXXX",
				"....",
			},
			expectedCount: 4,
		},
		{
			name: "move down",
			input: mapping.Map{
				"....",
				".v..",
				"....",
			},
			expected: mapping.Map{
				"....",
				".X..",
				".X..",
			},
			expectedCount: 2,
		},
		{
			name: "up then right",
			input: mapping.Map{
				".#..",
				"....",
				".^..",
			},
			expected: mapping.Map{
				".#..",
				".XXX",
				".X..",
			},
			expectedCount: 4,
		},
		{
			name: "u turn",
			input: mapping.Map{
				".#....",
				".....#",
				"......",
				".^....",
			},
			expected: mapping.Map{
				".#....",
				".XXXX#",
				".X..X.",
				".X..X.",
			},
			expectedCount: 8,
		},
		{
			name: "round",
			input: mapping.Map{
				".#........",
				".........#",
				"..........",
				"....#.....",
				"........#.",
				".^........",
			},
			expected: mapping.Map{
				".#...X....",
				".XXXXXXXX#",
				".X...X..X.",
				".X..#XXXX.",
				".X......#.",
				".X........",
			},
			expectedCount: 19,
		},
		{
			name:          "sample",
			input:         mapping.ParseMap(sample),
			expected:      mapping.ParseMap(sampleSolved),
			expectedCount: 41,
		},
		{
			name:          "real",
			input:         mapping.ParseMap(day1.ReadFile("day6.txt")),
			expectedCount: 5153,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			markPredictedRoute(test.input)
			assert.Equal(t, test.expectedCount, countVisited(test.input))
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected.String(), test.input.String())
			}
		})
	}
}

func Test_detectLoop(t *testing.T) {
	tests := []struct {
		name     string
		input    mapping.Map
		expected bool
	}{
		{
			name: "no loop",
			input: mapping.Map{
				"...",
				"...",
				".^.",
				"...",
			},
			expected: false,
		},
		{
			name: "loop",
			input: mapping.Map{
				".#........",
				".........#",
				"..........",
				"#.........",
				"........#.",
				".^........",
			},
			expected: true,
		},
		{
			name: "tight loop",
			input: mapping.Map{
				//1234
				".#...", // 0
				"#.#..", // 1
				"#..<.", // 2
				".#...", // 3
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loopDetected := detectLoop(test.input)
			assert.Equal(t, test.expected, loopDetected)
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name          string
		input         mapping.Map
		expectedCount int
	}{
		{
			name: "no loops possible",
			input: mapping.Map{
				"...",
				"...",
				".^.",
				"...",
			},
			expectedCount: 0,
		},
		{
			name: "one loop possible",
			input: mapping.Map{
				".#........",
				".........#",
				"..........",
				"....#.....",
				"........#.",
				".^........",
			},
			expectedCount: 1,
		},
		{
			name:          "sample",
			input:         mapping.ParseMap(sample),
			expectedCount: 6,
		},
		{
			name:          "real",
			input:         mapping.ParseMap(day1.ReadFile("day6.txt")),
			expectedCount: 1711,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedCount, countPossibleLoops(test.input))
		})
	}
}
