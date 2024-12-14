package day6

import "slices"

func countVisited(m Map) int {
	count := 0
	m.forEach(func(r int, c int, value int32) {
		if value == visited {
			count++
		}
	})
	return count
}

func findInitialPosition(m Map) (r int, c int, value int32) {
	values := []int32{guardianUp, guardianRight, guardianDown, guardianLeft}
	m.forEach(func(rr int, cc int, vv int32) {
		if slices.Index(values, vv) >= 0 {
			r = rr
			c = cc
			value = vv
		}
	})
	return
}

func markPredictedRoute(m Map) {
	r, c, curDir := findInitialPosition(m)
	for r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols() {
		if m.facingObstacle(r, c, curDir) {
			curDir = m.TurnGuardian(curDir)
		} else {
			m.Set(r, c, visited)
			switch curDir {
			case guardianUp:
				r--
			case guardianRight:
				c++
			case guardianDown:
				r++
			case guardianLeft:
				c--
			}
		}
	}
}
