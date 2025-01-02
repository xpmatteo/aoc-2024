package day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"math"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		input    mapping.Map
		expected int
	}{
		{
			name: "simple",
			input: mapping.Map{
				"#######",
				"#S...E#",
				"#######",
			},
			expected: 4,
		},
		{
			name: "elbow",
			input: mapping.Map{
				"#######",
				"#S....#",
				"#####.#",
				"#####E#",
				"#######",
			},
			expected: 1006,
		},
		{
			name: "dead end",
			input: mapping.Map{
				"#######",
				"#S....#",
				"##.##.#",
				"##.##E#",
				"#######",
			},
			expected: 1006,
		},
		{
			name: "all directions",
			input: mapping.Map{
				"########",
				"#S.....#",
				"######.#",
				"#...E#.#",
				"#.####.#",
				"#......#",
				"########",
			},
			expected: 4019,
		},
		{
			name: "loops",
			input: mapping.Map{
				"############",
				"#S........E#",
				"#.###.####.#",
				"#..........#",
				"############",
			},
			expected: 9,
		},
		{
			name: "first sample",
			input: mapping.ParseMap(`
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`),
			expected: 7036,
		},
		{
			name: "second sample",
			input: mapping.ParseMap(`
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`),
			expected: 11048,
		},
		{
			name:     "real part 1",
			input:    mapping.ParseMap(day1.ReadFile("day16.txt")),
			expected: 99448,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			maze := NewMaze(test.input)
			assert.Equal(t, test.expected, maze.LowestScore())
		})
	}
}

func Test_scoreComparison(t *testing.T) {
	tests := []struct {
		name                     string
		neighborScore, hereScore Score
		neighborDir              mapping.Direction
		expectedScore            Score
	}{
		{
			name:          "both maxint",
			neighborScore: Score{value: math.MaxInt},
			hereScore:     Score{value: math.MaxInt},
			expectedScore: Score{value: math.MaxInt},
		},
		{
			name:          "moving east-west",
			neighborScore: Score{10, mapping.DirectionEast},
			neighborDir:   mapping.DirectionWest,
			hereScore:     Score{value: math.MaxInt},
			expectedScore: Score{11, mapping.DirectionEast},
		},
		{
			name:          "moving east-west then elbow south",
			neighborScore: Score{10, mapping.DirectionEast},
			neighborDir:   mapping.DirectionNorth,
			hereScore:     Score{value: math.MaxInt},
			expectedScore: Score{1011, mapping.DirectionSouth},
		},
		{
			name:          "moving east-west then elbow north",
			neighborScore: Score{10, mapping.DirectionEast},
			neighborDir:   mapping.DirectionSouth,
			hereScore:     Score{value: 100000, dir: mapping.DirectionEast},
			expectedScore: Score{1011, mapping.DirectionNorth},
		},
		{
			name:          "elbow north, not convenient",
			neighborScore: Score{10, mapping.DirectionEast},
			neighborDir:   mapping.DirectionSouth,
			hereScore:     Score{value: 1010, dir: mapping.DirectionEast},
			expectedScore: Score{value: 1010, dir: mapping.DirectionEast},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updatedScore := test.hereScore.ImproveScore(test.neighborDir, test.neighborScore)
			assert.Equal(t, test.expectedScore, updatedScore)
		})
	}
}
