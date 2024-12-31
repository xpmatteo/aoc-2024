package day15

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

const smallSampleMap = `
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`

const smallSampleMoves = `<^^>>>vv<v>>v<<`

const smallSampleSolution = `
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`

const largeSampleMap = `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########`

const largeSampleMoves = `
<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const largeSampleSolution = `
##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`

func Test_parseMoves(t *testing.T) {
	assert.Equal(t, []Move{moveUp, moveDown, moveLeft, moveRight}, parseMoves("^v<>"))
	assert.Equal(t, []Move{moveUp, moveDown}, parseMoves("^\nv"))
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name             string
		inputMap         mapping.Map
		inputMoves       []Move
		expectedMap      mapping.Map
		expectedGpsTotal int
	}{
		{
			name: "move right",
			inputMap: mapping.Map{
				"#.@.#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#..@#",
			},
		},
		{
			name: "move left",
			inputMap: mapping.Map{
				"#.@.#",
			},
			inputMoves: parseMoves("<"),
			expectedMap: mapping.Map{
				"#@..#",
			},
		},
		{
			name: "move right blocked by a wall",
			inputMap: mapping.Map{
				"#.@#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#.@#",
			},
		},
		{
			name: "move right pushing a box",
			inputMap: mapping.Map{
				"#.@O..#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#..@O.#",
			},
		},
		{
			name: "move right with blocked box",
			inputMap: mapping.Map{
				"#.@O#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#.@O#",
			},
		},
		{
			name: "move right with multiple blocked boxes",
			inputMap: mapping.Map{
				"#.@OOOO#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#.@OOOO#",
			},
		},
		{
			name: "move right pushing multiple boxes",
			inputMap: mapping.Map{
				"#.@OOO.#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#..@OOO#",
			},
		},
		{
			name: "move left pushing multiple boxes",
			inputMap: mapping.Map{
				"#..OOO@#",
			},
			inputMoves: parseMoves("<"),
			expectedMap: mapping.Map{
				"#.OOO@.#",
			},
		},
		{
			name: "multiple moves",
			inputMap: mapping.Map{
				"#...@#",
			},
			inputMoves: parseMoves("<<"),
			expectedMap: mapping.Map{
				"#.@..#",
			},
		},
		{
			name: "multiple moves with boxes",
			inputMap: mapping.Map{
				"#...OOO@#",
			},
			inputMoves: parseMoves("<<<<<>>>>"),
			expectedMap: mapping.Map{
				"#OOO...@#",
			},
		},
		{
			name:             "small sample",
			inputMap:         mapping.ParseMap(smallSampleMap),
			inputMoves:       parseMoves(smallSampleMoves),
			expectedMap:      mapping.ParseMap(smallSampleSolution),
			expectedGpsTotal: 2028,
		},
		{
			name:             "large sample",
			inputMap:         mapping.ParseMap(largeSampleMap),
			inputMoves:       parseMoves(largeSampleMoves),
			expectedMap:      mapping.ParseMap(largeSampleSolution),
			expectedGpsTotal: 10092,
		},
		{
			name:             "real",
			inputMap:         mapping.ParseMap(day1.ReadFile("day15-map.txt")),
			inputMoves:       parseMoves(day1.ReadFile("day15-moves.txt")),
			expectedGpsTotal: 1371036,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solvedMap := predictRobot(test.inputMap, test.inputMoves)
			if test.expectedGpsTotal > 0 {
				assert.Equal(t, test.expectedGpsTotal, gpsTotal(solvedMap))
			}
			if test.expectedMap != nil {
				assert.Equal(t, test.expectedMap.String(), solvedMap.String())
			}
		})
	}
}

func Test_enlarge(t *testing.T) {
	tests := []struct {
		name        string
		inputMap    mapping.Map
		expectedMap mapping.Map
	}{
		{
			name: "example",
			inputMap: mapping.ParseMap(`
#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######`),
			expectedMap: mapping.ParseMap(`
##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedMap.String(), enlarge(test.inputMap).String())
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name             string
		inputMap         mapping.Map
		inputMoves       []Move
		expectedMap      mapping.Map
		expectedGpsTotal int
	}{
		{
			name: "move right pushing a box",
			inputMap: mapping.Map{
				"#.@[]..#",
			},
			inputMoves: parseMoves(">"),
			expectedMap: mapping.Map{
				"#..@[].#",
			},
		},
		{
			name: "move up pushing a box from the left side",
			inputMap: mapping.Map{
				"########",
				"#......#",
				"#..[]..#",
				"#..@...#",
			},
			inputMoves: parseMoves("^"),
			expectedMap: mapping.Map{
				"########",
				"#..[]..#",
				"#..@...#",
				"#......#",
			},
		},
		{
			name: "move up pushing a box from the right side",
			inputMap: mapping.Map{
				"########",
				"#......#",
				"#..[]..#",
				"#...@..#",
			},
			inputMoves: parseMoves("^"),
			expectedMap: mapping.Map{
				"########",
				"#..[]..#",
				"#...@..#",
				"#......#",
			},
		},
		{
			name: "move up pushing a box from the left side partially blocked",
			inputMap: mapping.Map{
				"########",
				"#...#..#",
				"#..[]..#",
				"#..@...#",
			},
			inputMoves: parseMoves("^"),
			expectedMap: mapping.Map{
				"########",
				"#...#..#",
				"#..[]..#",
				"#..@...#",
			},
		},
		{
			name: "move up with chain of blocked boxes",
			inputMap: mapping.Map{
				"########",
				"#...#..#",
				"#...[].#",
				"#..[]..#",
				"#..@...#",
			},
			inputMoves: parseMoves("^"),
			expectedMap: mapping.Map{
				"########",
				"#...#..#",
				"#...[].#",
				"#..[]..#",
				"#..@...#",
			},
		},
		{
			name: "example",
			inputMap: mapping.ParseMap(`
##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`),
			inputMoves: parseMoves("<vv<<^^<<^^"),
			expectedMap: mapping.ParseMap(`
##############
##...[].##..##
##...@.[]...##
##....[]....##
##..........##
##..........##
##############`),
		},
		{
			name:       "large sample",
			inputMap:   enlarge(mapping.ParseMap(largeSampleMap)),
			inputMoves: parseMoves(largeSampleMoves),
			expectedMap: mapping.ParseMap(`
####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################`),
			expectedGpsTotal: 9021,
		},
		{
			name:             "real",
			inputMap:         enlarge(mapping.ParseMap(day1.ReadFile("day15-map.txt"))),
			inputMoves:       parseMoves(day1.ReadFile("day15-moves.txt")),
			expectedGpsTotal: 1392847,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solvedMap := predictRobot(test.inputMap, test.inputMoves)
			if test.expectedGpsTotal > 0 {
				assert.Equal(t, test.expectedGpsTotal, gpsTotal(solvedMap))
			}
			if test.expectedMap != nil {
				assert.Equal(t, test.expectedMap.String(), solvedMap.String())
			}
		})
	}
}
