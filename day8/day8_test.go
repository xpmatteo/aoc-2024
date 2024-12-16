package day8

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/day6"
	"testing"
)

const sample = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

const sampleSolved = `......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....#.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.`

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         day6.Map
		expected      day6.Map
		expectedCount int
	}{
		{
			name: "1 antenna",
			input: day6.Map{
				"..........",
				"........a.",
				"..........",
			},
			expected: day6.Map{
				"..........",
				"........a.",
				"..........",
			},
			expectedCount: 0,
		},
		{
			name: "2 antennas",
			input: day6.Map{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"..........",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			expected: day6.Map{
				"..........",
				"...#......",
				"..........",
				"....a.....",
				"..........",
				".....a....",
				"..........",
				"......#...",
				"..........",
				"..........",
			},
			expectedCount: 2,
		},
		{
			name: "antinode falling outside",
			input: day6.Map{
				".a......a.",
			},
			expected: day6.Map{
				".a......a.",
			},
			expectedCount: 0,
		},
		{
			name: "3 nodes",
			input: day6.Map{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"........a.",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			expected: day6.Map{
				"..........",
				"...#......",
				"#.........",
				"....a.....",
				"........a.",
				".....a....",
				"..#.......",
				"......#...",
				"..........",
				"..........",
			},
			expectedCount: 4,
		},
		{
			name: "different frequencies",
			input: day6.Map{
				"..........",
				"..........",
				"..........",
				"....b.....",
				"........b.",
				".....b....",
				"..........",
				"......A...",
				"..........",
				"..........",
			},
			expected: day6.Map{
				"..........",
				"...#......",
				"#.........",
				"....b.....",
				"........b.",
				".....b....",
				"..#.......",
				"......#...",
				"..........",
				"..........",
			},
			expectedCount: 4,
		},
		{
			name:          "sample",
			input:         day6.ParseMap(sample),
			expected:      day6.ParseMap(sampleSolved),
			expectedCount: 14,
		},
		{
			name:          "real",
			input:         day6.ParseMap(day1.ReadFile("day8.txt")),
			expectedCount: 390,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := plotAntinodes(test.input)
			assert.Equal(t, test.expectedCount, countAntiNodes(m))
			if test.expected != nil {
				assert.Equal(t, test.expected.String(), m.String())
			}
		})
	}
}

const sample2 = `
T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`

const sample2Solved = `
#....#....
...#......
.#....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........`

func Test_part2(t *testing.T) {
	tests := []struct {
		name          string
		input         day6.Map
		expected      day6.Map
		expectedCount int
	}{
		{
			name: "simple",
			input: day6.Map{
				".....a...a........",
			},
			expected: day6.Map{
				".#...#...#...#...#",
			},
			expectedCount: 5,
		},
		{
			name:          "sample2",
			input:         day6.ParseMap(sample2),
			expected:      day6.ParseMap(sample2Solved),
			expectedCount: 9,
		},
		{
			name:          "real",
			input:         day6.ParseMap(day1.ReadFile("day8.txt")),
			expectedCount: 9,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := plotAntinodesPart2(test.input)
			assert.Equal(t, test.expectedCount, countAntiNodes(m))
			if test.expected != nil {
				assert.Equal(t, test.expected.String(), m.String())
			}
		})
	}
}
