package day6

import "slices"

func countVisited(m Map) int {
	count := 0
	m.ForEach(func(r int, c int, value int32) {
		if value == Visited {
			count++
		}
	})
	return count
}

func findInitialPosition(m Map) (r int, c int, dir int32) {
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

func markPredictedRoute(m Map) {
	r, c, curDir := findInitialPosition(m)
	for r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols() {
		if m.IsFacingObstacle(r, c, curDir) {
			curDir = m.Turn90DegreesRight(curDir)
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

func detectLoop(m Map) (loopDetected bool) {
	r, c, curDir := findInitialPosition(m)
	log := NewVisitLog()
	for r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols() {
		if log.DejaVu(r, c, curDir) {
			return true
		}
		log.Log(r, c, curDir)
		if m.IsFacingObstacle(r, c, curDir) {
			curDir = m.Turn90DegreesRight(curDir)
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

func countPossibleLoops(input Map) int {
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
