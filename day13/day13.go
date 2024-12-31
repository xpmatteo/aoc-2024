package day13

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xpmatteo/aoc-2024/day1"
	"math"
	"regexp"
	"strings"
)

type Point struct{ X, Y int }

func (p Point) Plus(point Point) Point {
	return Point{
		X: p.X + point.X,
		Y: p.Y + point.Y,
	}
}

func (p Point) times(n int) Point {
	return Point{
		X: p.X * n,
		Y: p.Y * n,
	}
}

func (p Point) increaseBy(point Point) {
	p.X += point.X
	p.Y += point.Y
}

type Button struct {
	tokens  int
	advance Point
}

type Machine struct {
	buttonA, buttonB Button
	prize            Point
}

const maxPushes = 100
const noWin = math.MaxInt

func (m *Machine) tokensNeededToWin() int {
	/*
	   pushA * a.X

	*/

	var minTokens = math.MaxInt
	for pushA := range maxPushes {
		for pushB := range maxPushes {
			tokensSpent := m.attemptToWIn(pushA, pushB)
			if tokensSpent < minTokens {
				minTokens = tokensSpent
			}
		}
	}
	return minTokens
}

func (m *Machine) attemptToWIn(a int, b int) int {
	pa := m.buttonA.advance.times(a)
	pb := m.buttonB.advance.times(b)
	if pa.Plus(pb) == m.prize {
		return a*m.buttonA.tokens + b*m.buttonB.tokens
	}
	return noWin
}

func (m *Machine) String() string {
	format := `
Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d`

	return fmt.Sprintf(format,
		m.buttonA.advance.X, m.buttonA.advance.X,
		m.buttonB.advance.X, m.buttonB.advance.X,
		m.prize.X, &m.prize.Y,
	)
}

type MachineList []Machine

func (l MachineList) tokensNeeded() int {
	var tokensTotal int
	for _, machine := range l {
		tokens := machine.tokensNeededToWin()
		if tokens != noWin {
			tokensTotal += tokens
		}
	}
	return tokensTotal
}

func (l MachineList) String() string {
	ss := lo.Map(l, func(m Machine, index int) string { return m.String() })
	return strings.Join(ss, "\n\n")
}

func parseMachineList(s string, prizeOffset int) MachineList {
	split := strings.Split(s, "\n\n")
	return lo.Map(split, func(item string, index int) Machine {
		return parseOneMachine(item, prizeOffset)
	})
}

func parseOneMachine(s string, offset int) Machine {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	if len(matches) != 6 {
		panic(fmt.Sprintf("Unexpected # of numbers: %v", matches))
	}
	numbers := lo.Map(matches, func(s string, index int) int {
		return day1.Atoi(s)
	})

	return Machine{
		buttonA: Button{
			tokens:  3,
			advance: Point{numbers[0], numbers[1]},
		},
		buttonB: Button{
			tokens:  1,
			advance: Point{numbers[2], numbers[3]},
		},
		prize: Point{numbers[4] + offset, numbers[5] + offset},
	}
}
