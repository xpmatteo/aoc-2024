package day6

const (
	visited       = 'X'
	guardianUp    = '^'
	guardianRight = '>'
	guardianDown  = 'v'
	guardianLeft  = '<'
	obstacle      = '#'
)

type Map []string

func (m Map) Set(r int, c int, value int32) {
	old := m[r]
	m[r] = old[:c] + string(value) + old[c+1:]
}

func (m Map) forEach(f func(r int, c int, value int32)) {
	for r, row := range m {
		for c, value := range row {
			f(r, c, value)
		}
	}
}

func (m Map) Cols() int {
	return len(m[0])
}

func (m Map) Rows() int {
	return len(m)
}

func (m Map) facingObstacle(r int, c int, dir int32) bool {
	return dir == guardianUp && r > 0 && m[r-1][c] == obstacle ||
		dir == guardianLeft && c > 0 && m[r][c-1] == obstacle ||
		dir == guardianRight && c+1 < m.Cols() && m[r][c+1] == obstacle ||
		dir == guardianDown && r+1 < m.Rows() && m[r+1][c] == obstacle

}

func (m Map) TurnGuardian(dir int32) int32 {
	switch dir {
	case guardianUp:
		return guardianRight
	case guardianRight:
		return guardianDown
	case guardianDown:
		return guardianLeft
	case guardianLeft:
		return guardianUp
	default:
		panic("don't know where to turn! " + string(dir))
	}
}
