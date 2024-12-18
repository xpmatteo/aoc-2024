package day10

import "github.com/xpmatteo/aoc-2024/maps"

type Front map[maps.Coord]struct{}

func (f Front) Advance(m maps.Map) Front {
	newFront := make(Front)
	for coord := range f {
		current := m.At(coord)
		if current < '0' || current > '9' {
			panic("Unexpected current value " + string(current))
		}
		if current == trailEnd {
			continue
		}
		neighbors := coord.OrthoNeighbors()
		for _, neighbor := range neighbors {
			if m.At(neighbor) == current+1 {
				newFront[neighbor] = struct{}{}
			}
		}
	}
	return newFront
}

func (f Front) Score() int {
	return len(f)
}

func (f Front) Ongoing(m maps.Map) bool {
	for coord := range f {
		// map is not empty
		return m.At(coord) != trailEnd
	}
	// map is empty
	return false
}

func NewFront(head maps.Coord) Front {
	return Front{head: struct{}{}}
}
