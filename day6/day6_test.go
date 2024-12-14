package day6

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
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
		input         Map
		expected      Map
		expectedCount int
	}{
		{
			name: "move up",
			input: Map{
				"...",
				"...",
				".^.",
				"...",
			},
			expected: Map{
				".X.",
				".X.",
				".X.",
				"...",
			},
			expectedCount: 3,
		},
		{
			name: "move right",
			input: Map{
				"....",
				">...",
				"....",
			},
			expected: Map{
				"....",
				"XXXX",
				"....",
			},
			expectedCount: 4,
		},
		{
			name: "move down",
			input: Map{
				"....",
				".v..",
				"....",
			},
			expected: Map{
				"....",
				".X..",
				".X..",
			},
			expectedCount: 2,
		},
		{
			name: "up then right",
			input: Map{
				".#..",
				"....",
				".^..",
			},
			expected: Map{
				".#..",
				".XXX",
				".X..",
			},
			expectedCount: 4,
		},
		{
			name: "u turn",
			input: Map{
				".#....",
				".....#",
				"......",
				".^....",
			},
			expected: Map{
				".#....",
				".XXXX#",
				".X..X.",
				".X..X.",
			},
			expectedCount: 8,
		},
		{
			name: "round",
			input: Map{
				".#........",
				".........#",
				"..........",
				"....#.....",
				"........#.",
				".^........",
			},
			expected: Map{
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
			input:         ParseMap(sample),
			expected:      ParseMap(sampleSolved),
			expectedCount: 41,
		},
		{
			name:          "real",
			input:         ParseMap(day1.ReadFile("day6.txt")),
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
