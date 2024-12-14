package day7

import (
	"github.com/xpmatteo/aoc-2024/day1"
	"strconv"
	"strings"
)

type Operator string

func (o Operator) Evaluate(a int, b int) int {
	switch o {
	case OpPlus:
		return a + b
	case OpTimes:
		return a * b
	case OpConcat:
		return day1.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	default:
		panic("bad operator: " + o)
	}
}

const OpPlus = Operator("+")
const OpTimes = Operator("*")
const OpConcat = Operator("||")

var Operators2 = []Operator{OpPlus, OpTimes}
var Operators3 = []Operator{OpPlus, OpTimes, OpConcat}

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

func (e Equation) IsSolvable(availableOperators []Operator) bool {
	numOperators := len(e.operands) - 1
	numOpCombinations := power(len(availableOperators), numOperators)
	for i := range numOpCombinations {
		ops := genCombination(numOperators, i, availableOperators)
		if e.IsSolvedBy(ops) {
			return true
		}
	}
	return false
}

func genCombination(length int, index int, availableOperators []Operator) []Operator {
	result := make([]Operator, length)
	for i := range length {
		op := index % len(availableOperators)
		result[i] = availableOperators[op]
		index /= len(availableOperators)
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

func sumOfSolvableEquations(eqns []Equation, availableOperators []Operator) int {
	result := 0
	for _, eqn := range eqns {
		if eqn.IsSolvable(availableOperators) {
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
