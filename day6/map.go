package day6

import "strings"

const (
	Visited        = 'X'
	DirectionUp    = '^'
	DirectionRight = '>'
	DirectionDown  = 'v'
	DirectionLeft  = '<'
	Obstacle       = '#'
)

type Map []string

func ParseMap(s string) Map {
	trimmed := strings.Trim(s, "\n")
	return strings.Split(trimmed, "\n")
}

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

func (m Map) IsFacingObstacle(r int, c int, dir int32) bool {
	return dir == DirectionUp && r > 0 && m[r-1][c] == Obstacle ||
		dir == DirectionLeft && c > 0 && m[r][c-1] == Obstacle ||
		dir == DirectionRight && c+1 < m.Cols() && m[r][c+1] == Obstacle ||
		dir == DirectionDown && r+1 < m.Rows() && m[r+1][c] == Obstacle
}

func (m Map) Turn90DegreesRight(dir int32) int32 {
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

func (m Map) String() string {
	return strings.Join(m, "\n")
}

func (m Map) Clone() Map {
	clone := make(Map, len(m))
	_ = copy(clone, m)
	return clone
}
