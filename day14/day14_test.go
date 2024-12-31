package day14

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

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
				size: point{11, 11},
				robots: []*Robot{{
					position: point{0, 4},
					speed:    point{3, -3},
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
				size: point{11, 11},
				robots: []*Robot{
					{
						position: point{0, 4},
						speed:    point{3, -3},
					},
					{
						position: point{6, 3},
						speed:    point{-1, -2},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, parseLobby(point{11, 11}, test.input))
		})
	}
}

/*
	----- tests ----
	X 2 robots moving
	- 2 robots in the same tile
	X wrap around
*/

func Test_movement(t *testing.T) {
	tests := []struct {
		name     string
		input    Lobby
		seconds  int
		expected mapping.Map
	}{
		{
			name:    "one robot",
			input:   parseLobby(point{11, 4}, "p=0,0 v=1,2"),
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
			input:   parseLobby(point{11, 4}, "p=0,0 v=1,2"),
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
			input: parseLobby(point{11, 4},
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lobby := test.input
			lobby.Simulate(test.seconds)
			assert.Equal(t, test.expected, lobby.Map())
		})
	}
}
