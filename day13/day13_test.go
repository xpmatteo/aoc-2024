package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name           string
		input          MachineList
		expectedTokens int
	}{
		{
			name: "simplest machine",
			input: MachineList{Machine{
				buttonA:    Button{tokens: 3, advance: Point{100, 200}},
				buttonB:    Button{tokens: 1, advance: Point{10, 20}},
				prizePoint: Point{1000, 2000},
			}},
			expectedTokens: 30,
		},
		{
			name: "first machine",
			input: parseMachineList(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`),
			expectedTokens: 280,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.input.tokensNeeded()
			assert.Equal(t, test.expectedTokens, actual)
		})
	}
}
