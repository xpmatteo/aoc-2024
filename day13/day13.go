package day13

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xpmatteo/aoc-2024/day1"
	"math"
	"regexp"
	"strings"
)

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
	prize            Point
}

const maxPushes = 100
const noWin = math.MaxInt

func (m *Machine) tokensNeededToWin() int {
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
	if pa.plus(pb) == m.prize {
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
		m.buttonA.advance.x, m.buttonA.advance.x,
		m.buttonB.advance.x, m.buttonB.advance.x,
		m.prize.x, &m.prize.y,
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

func parseMachineList(s string) MachineList {
	split := strings.Split(s, "\n\n")
	return lo.Map(split, func(item string, index int) Machine {
		return parseOneMachine(item)
	})
}

func parseOneMachine(s string) Machine {
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
		prize: Point{numbers[4], numbers[5]},
	}
}
