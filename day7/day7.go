package day7

import (
	"github.com/xpmatteo/aoc-2024/day1"
	"strings"
)

type Operator string

func (o Operator) Evaluate(a int, b int) int {
	switch o {
	case OpPlus:
		return a + b
	case OpTimes:
		return a * b
	default:
		panic("bad operand: " + o)
	}
}

const OpPlus = Operator("0")
const OpTimes = Operator("1")

var Operators = []Operator{OpPlus, OpTimes}

type Equation struct {
	result   int
	operands []int
}

func (e Equation) IsSolvedBy(ops []Operator) bool {
	if len(e.operands)-1 != len(ops) {
		panic("wrong number of operands")
	}
	return e.Evaluate(ops) == e.result
}

func (e Equation) IsSolvable() bool {
	numOperators := len(e.operands) - 1
	numOpCombinations := power(len(Operators), numOperators)
	for i := range numOpCombinations {
		ops := genCombination(numOperators, i)
		if e.IsSolvedBy(ops) {
			return true
		}
	}
	return false
}

func genCombination(length int, index int) []Operator {
	result := make([]Operator, length)
	for i := range length {
		op := index % len(Operators)
		result[i] = Operators[op]
		index /= len(Operators)
	}
	return result
}

func (e Equation) Evaluate(ops []Operator) int {
	result := ops[0].Evaluate(e.operands[0], e.operands[1])
	for i := 1; i < len(ops); i++ {
		result = ops[i].Evaluate(result, e.operands[i+1])
	}
	return result
}

func power(n int, exp int) int {
	result := 1
	for range exp {
		result *= n
	}
	return result
}

func sumOfSolvableEquations(eqns []Equation) int {
	result := 0
	for _, eqn := range eqns {
		if eqn.IsSolvable() {
			result += eqn.result
		}
	}
	return result
}

func parseEquation(s string) Equation {
	equation := Equation{}
	firstSplit := strings.Split(s, ":")
	equation.result = day1.Atoi(firstSplit[0])
	secondSplit := strings.Split(firstSplit[1], " ")
	for _, s2 := range secondSplit {
		if len(s2) > 0 {
			equation.operands = append(equation.operands, day1.Atoi(s2))
		}
	}
	return equation
}
