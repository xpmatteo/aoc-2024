package day15

import (
	"github.com/stretchr/testify/assert"
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

func Test_parseMoves(t *testing.T) {
	assert.Equal(t, []Move{moveUp, moveDown, moveLeft, moveRight}, parseMoves("^v<>"))
}

/*
tests
X move right
X move left
- move down
X move right blocked by wall
- move right blocked by box and wall
- move right pushing box
*/
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
		//{
		//	name:        "sample",
		//	inputMap:    mapping.ParseMap(smallSampleMap),
		//	inputMoves:  parseMoves(smallSampleMoves),
		//	expectedMap: mapping.ParseMap(smallSampleSolution),
		//},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solvedMap := predictRobot(test.inputMap, test.inputMoves)
			assert.Equal(t, test.expectedMap.String(), solvedMap.String())
		})
	}
}
