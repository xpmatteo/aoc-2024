package day6

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"strings"
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
			input: []string{
				"...",
				"...",
				".^.",
				"...",
			},
			expected: []string{
				".X.",
				".X.",
				".X.",
				"...",
			},
			expectedCount: 3,
		},
		{
			name: "move right",
			input: []string{
				"....",
				">...",
				"....",
			},
			expected: []string{
				"....",
				"XXXX",
				"....",
			},
			expectedCount: 4,
		},
		{
			name: "move down",
			input: []string{
				"....",
				".v..",
				"....",
			},
			expected: []string{
				"....",
				".X..",
				".X..",
			},
			expectedCount: 2,
		},
		{
			name: "up then right",
			input: []string{
				".#..",
				"....",
				".^..",
			},
			expected: []string{
				".#..",
				".XXX",
				".X..",
			},
			expectedCount: 4,
		},
		{
			name: "u turn",
			input: []string{
				".#....",
				".....#",
				"......",
				".^....",
			},
			expected: []string{
				".#....",
				".XXXX#",
				".X..X.",
				".X..X.",
			},
			expectedCount: 8,
		},
		{
			name: "round",
			input: []string{
				".#........",
				".........#",
				"..........",
				"....#.....",
				"........#.",
				".^........",
			},
			expected: []string{
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
			input:         parse(sample),
			expected:      parse(sampleSolved),
			expectedCount: 41,
		},
		{
			name:          "real",
			input:         parse(day1.ReadFile("day6.txt")),
			expectedCount: 5153,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			markPredictedRoute(test.input)
			assert.Equal(t, test.expectedCount, countVisited(test.input))
			if len(test.expected) > 0 {
				assert.Equal(t, join(test.expected), join(test.input))
			}
		})
	}
}

func parse(s string) Map {
	return strings.Split(s, "\n")
}

func join(m Map) string {
	return strings.Join(m, "\n")
}

func Test_set(t *testing.T) {
	m := Map([]string{"aaa"})
	m.Set(0, 1, 'X')
	assert.Equal(t, Map([]string{"aXa"}), m)
}
