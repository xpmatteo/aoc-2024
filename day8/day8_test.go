package day8

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
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
		input         maps.Map
		expected      maps.Map
		expectedCount int
	}{
		{
			name: "1 antenna",
			input: maps.Map{
				"..........",
				"........a.",
				"..........",
			},
			expected: maps.Map{
				"..........",
				"........a.",
				"..........",
			},
			expectedCount: 0,
		},
		{
			name: "2 antennas",
			input: maps.Map{
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
			expected: maps.Map{
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
			input: maps.Map{
				".a......a.",
			},
			expected: maps.Map{
				".a......a.",
			},
			expectedCount: 0,
		},
		{
			name: "3 nodes",
			input: maps.Map{
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
			expected: maps.Map{
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
			input: maps.Map{
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
			expected: maps.Map{
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
			input:         maps.ParseMap(sample),
			expected:      maps.ParseMap(sampleSolved),
			expectedCount: 14,
		},
		{
			name:          "real",
			input:         maps.ParseMap(day1.ReadFile("day8.txt")),
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
		input         maps.Map
		expected      maps.Map
		expectedCount int
	}{
		{
			name: "simple",
			input: maps.Map{
				".....a...a........",
			},
			expected: maps.Map{
				".#...#...#...#...#",
			},
			expectedCount: 5,
		},
		{
			name:          "sample2",
			input:         maps.ParseMap(sample2),
			expected:      maps.ParseMap(sample2Solved),
			expectedCount: 9,
		},
		{
			name:          "real",
			input:         maps.ParseMap(day1.ReadFile("day8.txt")),
			expectedCount: 1246,
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
