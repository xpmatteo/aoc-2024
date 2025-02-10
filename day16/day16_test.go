package day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"math"
	"strings"
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
			neighborScore: Score{value: 10, dir: mapping.DirectionEast},
			neighborDir:   mapping.DirectionWest,
			hereScore:     Score{value: math.MaxInt},
			expectedScore: Score{value: 11, dir: mapping.DirectionEast},
		},
		{
			name:          "moving east-west then elbow south",
			neighborScore: Score{value: 10, dir: mapping.DirectionEast},
			neighborDir:   mapping.DirectionNorth,
			hereScore:     Score{value: math.MaxInt},
			expectedScore: Score{value: 1011, dir: mapping.DirectionSouth},
		},
		{
			name:          "moving east-west then elbow north",
			neighborScore: Score{value: 10, dir: mapping.DirectionEast},
			neighborDir:   mapping.DirectionSouth,
			hereScore:     Score{value: 100000, dir: mapping.DirectionEast},
			expectedScore: Score{value: 1011, dir: mapping.DirectionNorth},
		},
		{
			name:          "elbow north, not convenient",
			neighborScore: Score{value: 10, dir: mapping.DirectionEast},
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

func Test_part2(t *testing.T) {
	tests := []struct {
		name          string
		input         mapping.Map
		expectedMap   string
		expectedCount int
		skip          bool
	}{
		{
			name: "simple",
			input: mapping.Map{
				"#######",
				"#S...E#",
				"#######",
			},
			expectedCount: 5,
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
			expectedCount: 7,
			expectedMap: mapping.Map{
				"#######",
				"#OOOOO#",
				"##.##O#",
				"##.##O#",
				"#######",
			}.String(),
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
			expectedCount: 10,
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
			expectedMap: `
###############
#.......#....O#
#.#.###.#.###O#
#.....#.#...#O#
#.###.#####.#O#
#.#.#.......#O#
#.#.#####.###O#
#..OOOOOOOOO#O#
###O#O#####O#O#
#OOO#O....#O#O#
#O#O#O###.#O#O#
#OOOOO#...#O#O#
#O###.#.#.#O#O#
#O..#.....#OOO#
###############`,
			expectedCount: 45,
		},
		{
			skip: true,
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
			expectedCount: 64,
			expectedMap: `
#################
#...#...#...#..O#
#.#.#.#.#.#.#.#O#
#.#.#.#...#...#O#
#.#.#.#.###.#.#O#
#OOO#.#.#.....#O#
#O#O#.#.#.#####O#
#O#O..#.#.#OOOOO#
#O#O#####.#O###O#
#O#O#..OOOOO#OOO#
#O#O###O#####O###
#O#O#OOO#..OOO#.#
#O#O#O#####O###.#
#O#O#OOOOOOO..#.#
#O#O#O#########.#
#O#OOO..........#
#################`,
		},
		//		{
		//			name:          "real part 1",
		//			input:         mapping.ParseMap(day1.ReadFile("day16.txt")),
		//			expectedCount: 99448,
		//		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.skip {
				t.Skip()
			}
			maze := NewMaze(test.input)
			scores := maze.computeScoresFrom(maze.start, 0, mapping.DirectionEast)
			assert.Equal(t, test.expectedCount, maze.CountBestTilesToSit(scores))
			if len(test.expectedMap) > 0 {
				assert.Equal(t, strings.TrimLeft(test.expectedMap, "\n"), maze.ShowBestPath())
			}
		})
	}
}
