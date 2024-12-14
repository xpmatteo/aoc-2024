package day7

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"strings"
	"testing"
)

const sample = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_equation(t *testing.T) {
	tests := []struct {
		name             string
		input            Equation
		shouldBeSolvable bool
	}{
		{
			name:             "simple solvable with +",
			input:            parseEquation("5: 2 3"),
			shouldBeSolvable: true,
		},
		{
			name:             "simple solvable with *",
			input:            parseEquation("6: 2 3"),
			shouldBeSolvable: true,
		},
		{
			name:             "simple UNsolvable",
			input:            parseEquation("6: 2 100"),
			shouldBeSolvable: false,
		},
		{
			name:             "three operands, solvable",
			input:            parseEquation("111: 1 10 100"),
			shouldBeSolvable: true,
		},
		{
			name:             "three operands, solvable",
			input:            parseEquation("1100: 1 10 100"),
			shouldBeSolvable: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.shouldBeSolvable, test.input.IsSolvable(Operators2))
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		eqns     []Equation
		expected int
	}{
		{
			name:     "one solvable eq",
			eqns:     []Equation{parseEquation("12: 2 2 3")},
			expected: 12,
		},
		{
			name:     "sample",
			eqns:     parseManyEquations(sample),
			expected: 3749,
		},
		{
			name:     "real",
			eqns:     parseManyEquations(day1.ReadFile("day7.txt")),
			expected: 6231007345478,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, sumOfSolvableEquations(test.eqns, Operators2))
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name     string
		eqns     []Equation
		expected int
	}{
		{
			name:     "one solvable eq",
			eqns:     []Equation{parseEquation("223: 2 2 3")},
			expected: 223,
		},
		{
			name:     "sample",
			eqns:     parseManyEquations(sample),
			expected: 11387,
		},
		{
			name:     "real",
			eqns:     parseManyEquations(day1.ReadFile("day7.txt")),
			expected: 333027885676693,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, sumOfSolvableEquations(test.eqns, Operators3))
		})
	}
}

func parseManyEquations(input string) []Equation {
	var result []Equation
	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			continue
		}
		result = append(result, parseEquation(s))
	}
	return result
}

func Test_parseEquation(t *testing.T) {
	e := parseEquation("465821964765: 83 413 37 153 365")
	expected := Equation{
		result:   465821964765,
		operands: []int{83, 413, 37, 153, 365},
	}
	assert.Equal(t, expected, e)
}

func Test_power(t *testing.T) {
	assert.Equal(t, 8, power(2, 3))
}

func Test_genCombination(t *testing.T) {
	tests := []struct {
		name     string
		len      int
		index    int
		expected []Operator
	}{
		{
			name:     "simple 0",
			len:      1,
			index:    0,
			expected: []Operator{OpPlus},
		},
		{
			name:     "simple 1",
			len:      1,
			index:    1,
			expected: []Operator{OpTimes},
		},
		{
			name:     "00",
			len:      2,
			index:    0,
			expected: []Operator{OpPlus, OpPlus},
		},
		{
			name:     "01",
			len:      2,
			index:    1,
			expected: []Operator{OpTimes, OpPlus},
		},
		{
			name:     "10",
			len:      2,
			index:    2,
			expected: []Operator{OpPlus, OpTimes},
		},
		{
			name:     "11",
			len:      2,
			index:    3,
			expected: []Operator{OpTimes, OpTimes},
		},
		{
			name:     "0001",
			len:      4,
			index:    1,
			expected: []Operator{OpTimes, OpPlus, OpPlus, OpPlus},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, genCombination(test.len, test.index, Operators2))
		})
	}
}

func Test_evaluate(t *testing.T) {
	tests := []struct {
		name     string
		e        Equation
		ops      []Operator
		expected int
	}{
		{
			name:     "simple",
			e:        parseEquation("0: 2 3"),
			ops:      []Operator{OpPlus},
			expected: 5,
		},
		{
			name:     "3 operants",
			e:        parseEquation("0: 1 10 100"),
			ops:      []Operator{OpPlus, OpPlus},
			expected: 111,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.e.Evaluate(test.ops))
		})
	}
}
