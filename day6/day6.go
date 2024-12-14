package day6

import "slices"

func countVisited(m Map) int {
	count := 0
	m.forEach(func(r int, c int, value int32) {
		if value == Visited {
			count++
		}
	})
	return count
}

func findInitialPosition(m Map) (r int, c int, dir int32) {
	values := []int32{DirectionUp, DirectionRight, DirectionDown, DirectionLeft}
	m.forEach(func(rr int, cc int, vv int32) {
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
		if m.facingObstacle(r, c, curDir) {
			curDir = m.Turn90Degrees(curDir)
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
