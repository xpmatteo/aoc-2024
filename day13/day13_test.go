package day13

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"testing"
)

const sample = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func Test_part1(t *testing.T) {
	tests := []struct {
		name           string
		input          MachineList
		expectedTokens int
	}{
		{
			name: "simplest machine",
			input: MachineList{Machine{
				buttonA: Button{tokens: 3, advance: Point{100, 200}},
				buttonB: Button{tokens: 1, advance: Point{10, 20}},
				prize:   Point{1000, 2000},
			}},
			expectedTokens: 30,
		},
		{
			name: "no win",
			input: MachineList{Machine{
				buttonA: Button{tokens: 3, advance: Point{100, 200}},
				buttonB: Button{tokens: 1, advance: Point{10, 20}},
				prize:   Point{1, 2},
			}},
			expectedTokens: 0,
		},
		{
			name: "first machine",
			input: parseMachineList(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`, 0),
			expectedTokens: 280,
		},
		{
			name:           "sample",
			input:          parseMachineList(sample, 0),
			expectedTokens: 480,
		},
		{
			name:           "real",
			input:          parseMachineList(day1.ReadFile("day13.txt"), 0),
			expectedTokens: 30413,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.input.tokensNeeded()
			assert.Equal(t, test.expectedTokens, actual)
		})
	}
}

func Test_part2(t *testing.T) {
	const addition = 10_000_000_000_000
	tests := []struct {
		name           string
		input          MachineList
		expectedTokens int
	}{
		{
			name:           "sample",
			input:          parseMachineList(sample, addition),
			expectedTokens: 480,
		},
		{
			name:           "real",
			input:          parseMachineList(day1.ReadFile("day13.txt"), addition),
			expectedTokens: 30413,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ml := test.input
			actual := ml.tokensNeeded()
			assert.Equal(t, test.expectedTokens, actual)
		})
	}
}

// x*94+y*22=8400
func Test_parseMachineList(t *testing.T) {
	ml := parseMachineList(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`, 10000)

	assert.Equal(t, ml, MachineList{
		Machine{
			buttonA: Button{tokens: 3, advance: Point{94, 34}},
			buttonB: Button{tokens: 1, advance: Point{22, 67}},
			prize:   Point{18400, 15400},
		},
	}, ml.String())
}
