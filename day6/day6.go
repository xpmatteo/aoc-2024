package day6

import (
	"github.com/xpmatteo/aoc-2024/mapping"
	"slices"
)

const (
	Visited        = 'X'
	DirectionUp    = '^'
	DirectionRight = '>'
	DirectionDown  = 'v'
	DirectionLeft  = '<'
	Obstacle       = '#'
)

func countVisited(m mapping.Map) int {
	count := 0
	m.ForEach(func(r int, c int, value int32) {
		if value == Visited {
			count++
		}
	})
	return count
}

func findInitialPosition(m mapping.Map) (r int, c int, dir int32) {
	values := []int32{DirectionUp, DirectionRight, DirectionDown, DirectionLeft}
	m.ForEach(func(rr int, cc int, vv int32) {
		if slices.Index(values, vv) >= 0 {
			r = rr
			c = cc
			dir = vv
		}
	})
	return
}

func markPredictedRoute(m mapping.Map) {
	r, c, curDir := findInitialPosition(m)
	for r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols() {
		if isFacingObstacle(m, r, c, curDir) {
			curDir = Turn90DegreesRight(curDir)
		}
		m.Set(r, c, Visited)
		switch curDir {
		case DirectionUp:
			r--
		case DirectionRight:
			c++
		case DirectionDown:
			r++
		case DirectionLeft:
			c--
		default:
			panic("heading in unknown direction: " + string(curDir))
		}
	}
}

func detectLoop(m mapping.Map) (loopDetected bool) {
	r, c, curDir := findInitialPosition(m)
	log := NewVisitLog()
	for r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols() {
		if log.DejaVu(r, c, curDir) {
			return true
		}
		log.Log(r, c, curDir)
		if isFacingObstacle(m, r, c, curDir) {
			curDir = Turn90DegreesRight(curDir)
		} else {
			switch curDir {
			case DirectionUp:
				r--
			case DirectionRight:
				c++
			case DirectionDown:
				r++
			case DirectionLeft:
				c--
			default:
				panic("heading in unknown direction: " + string(curDir))
			}
		}
	}
	return false
}

func countPossibleLoops(input mapping.Map) int {
	guardRow, guardCol, _ := findInitialPosition(input)
	count := 0
	input.ForEach(func(r, c int, value int32) {
		if r == guardRow && c == guardCol {
			return
		}
		whatWasThere := input.Get(r, c)
		input.Set(r, c, Obstacle)
		loopDetected := detectLoop(input)
		if loopDetected {
			count++
		}
		input.Set(r, c, whatWasThere)
	})
	return count
}

func isFacingObstacle(m mapping.Map, r int, c int, dir int32) bool {
	return dir == DirectionUp && r > 0 && m[r-1][c] == Obstacle ||
		dir == DirectionLeft && c > 0 && m[r][c-1] == Obstacle ||
		dir == DirectionRight && c+1 < m.Cols() && m[r][c+1] == Obstacle ||
		dir == DirectionDown && r+1 < m.Rows() && m[r+1][c] == Obstacle
}

func Turn90DegreesRight(dir int32) int32 {
	switch dir {
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	default:
		panic("don't know where to turn! " + string(dir))
	}
}
