package day13

import "math"

type Point struct{ x, y int }

func (p Point) plus(point Point) Point {
	return Point{
		x: p.x + point.x,
		y: p.y + point.y,
	}
}

func (p Point) times(n int) Point {
	return Point{
		x: p.x * n,
		y: p.y * n,
	}
}

type Button struct {
	tokens  int
	advance Point
}

type Machine struct {
	buttonA, buttonB Button
	prizePoint       Point
}

const maxPushes = 100
const noWin = math.MaxInt

func (m *Machine) tokensNeededToWin() int {
	var minTokens = math.MaxInt
	for pushA := range maxPushes {
		for pushB := range maxPushes {
			tokensSpent := m.attemptToWIn(pushA, pushB)
			if tokensSpent != noWin && tokensSpent < minTokens {
				minTokens = tokensSpent
			}
		}
	}
	return minTokens
}

func (m *Machine) attemptToWIn(a int, b int) int {
	pa := m.buttonA.advance.times(a)
	pb := m.buttonB.advance.times(b)
	if pa.plus(pb) == m.prizePoint {
		return a*m.buttonA.tokens + b*m.buttonB.tokens
	}
	return noWin
}

type MachineList []Machine

func (l MachineList) tokensNeeded() int {
	var tokensTotal int
	for _, machine := range l {
		tokensTotal += machine.tokensNeededToWin()
	}
	return tokensTotal
}

func parseMachineList(s string) MachineList {
	return nil
}
