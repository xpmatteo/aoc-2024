package day14

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

const sample = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func Test_parseLobby(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Lobby
	}{
		{
			name:  "one robot",
			input: "p=0,4 v=3,-3",
			expected: Lobby{
				size: Point{11, 11},
				robots: []*Robot{{
					position: Point{0, 4},
					speed:    Point{3, -3},
				}},
			},
		},
		{
			name: "two robots",
			input: join(
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-2",
			),
			expected: Lobby{
				size: Point{11, 11},
				robots: []*Robot{
					{
						position: Point{0, 4},
						speed:    Point{3, -3},
					},
					{
						position: Point{6, 3},
						speed:    Point{-1, -2},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ParseLobby(Point{11, 11}, test.input))
		})
	}
}

/*
	----- tests ----
	X 2 robots moving
	X 2 robots in the same tile
	X wrap around
	X sample simulation
	- find safety factor
*/

func Test_simulation(t *testing.T) {
	tests := []struct {
		name                 string
		input                Lobby
		seconds              int
		expected             mapping.Map
		expectedSafetyFactor int
	}{
		{
			name:    "one robot",
			input:   ParseLobby(Point{11, 4}, "p=0,0 v=1,2"),
			seconds: 1,
			expected: mapping.Map{
				"...........",
				"...........",
				".1.........",
				"...........",
			},
		},
		{
			name:    "wrap around",
			input:   ParseLobby(Point{11, 4}, "p=0,0 v=1,2"),
			seconds: 2,
			expected: mapping.Map{
				"..1........",
				"...........",
				"...........",
				"...........",
			},
		},
		{
			name: "2 robots",
			input: ParseLobby(Point{11, 4},
				"p=0,0 v=1,2\n"+
					"p=0,0 v=-1,-1\n"),
			seconds: 1,
			expected: mapping.Map{
				"...........",
				"...........",
				".1.........",
				"..........1",
			},
		},
		{
			name: "2 robots in the same tile",
			input: ParseLobby(Point{11, 4},
				"p=0,0 v=1,2\n"+
					"p=0,0 v=1,2\n"),
			seconds: 1,
			expected: mapping.Map{
				"...........",
				"...........",
				".2.........",
				"...........",
			},
		},
		{
			name:    "sample 100",
			input:   ParseLobby(Point{11, 7}, sample),
			seconds: 100,
			expected: mapping.Map{
				"......2..1.",
				"...........",
				"1..........",
				".11........",
				".....1.....",
				"...12......",
				".1....1....",
			},
			expectedSafetyFactor: 12,
		},
		{
			name:                 "real",
			input:                ParseLobby(Point{101, 103}, day1.ReadFile("day14.txt")),
			seconds:              100,
			expectedSafetyFactor: 221616000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lobby := test.input
			lobby.Simulate(test.seconds)
			if test.expectedSafetyFactor > 0 {
				assert.Equal(t, test.expectedSafetyFactor, lobby.SafetyFactor())
			}
			if test.expected != nil {
				assert.Equal(t, test.expected.String(), lobby.Map().String())
			}
		})
	}
}

func Test_easterEgg(t *testing.T) {
	lobby := ParseLobby(Point{101, 103}, day1.ReadFile("day14.txt"))
	start := 36155
	lobby.Simulate(start)
	for seconds := start + 1; seconds < 100000; seconds++ {
		lobby.Simulate(1)
		if lobby.Matrix()[50][0] > 0 && lobby.Matrix()[49][1] > 0 && lobby.Matrix()[51][1] > 0 {
			println("seconds:", seconds)
			println(lobby.Map().String())
			break
		}
	}
}

func Test_proportion(t *testing.T) {
	tests := []struct {
		name               string
		input              Lobby
		expectedProportion float64
		expectedMap        mapping.Map
	}{
		{
			name: "simple",
			input: ParseLobby(Point{11, 4},
				"p=0,0 v=0,0\n"+
					"p=0,2 v=0,0\n"+
					"p=1,2 v=0,0\n"+
					"p=2,2 v=0,0\n"),
			expectedMap: mapping.Map{
				"1..........",
				"...........",
				"111........",
				"...........",
			},
			expectedProportion: 1 / 3.0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lobby := test.input
			if test.expectedMap != nil {
				assert.Equal(t, test.expectedMap.String(), lobby.Map().String())
			}
			assert.Equal(t, test.expectedProportion, lobby.UpperVsLowerProportion())
		})
	}
}
