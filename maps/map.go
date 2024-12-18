package maps

import "strings"

type Map []string

func ParseMap(s string) Map {
	trimmed := strings.Trim(s, "\n")
	return strings.Split(trimmed, "\n")
}

func (m Map) Set(r int, c int, value int32) {
	old := m[r]
	m[r] = old[:c] + string(value) + old[c+1:]
}

func (m Map) SetSafe(coord Coord, value int32) {
	if m.IsValid(coord) {
		m.Set(coord.Row, coord.Col, value)
	}
}

func (m Map) SetCoord(c Coord, val int32) {
	m.Set(c.Row, c.Col, val)
}

func (m Map) Get(r int, c int) int32 {
	return int32(m[r][c])
}

func (m Map) Cols() int {
	return len(m[0])
}

func (m Map) Rows() int {
	return len(m)
}

func (m Map) ForEach(f func(r int, c int, value int32)) {
	for r, row := range m {
		for c, value := range row {
			f(r, c, value)
		}
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

func (m Map) IsValid(coord Coord) bool {
	r := coord.Row
	c := coord.Col
	return r >= 0 && c >= 0 && r < m.Rows() && c < m.Cols()
}

func (m Map) At(coord Coord) int32 {
	return int32(m[coord.Row][coord.Col])
}
