package day10

import mymap "github.com/xpmatteo/aoc-2024/maps"

type Front map[mymap.Coord]struct{}

func (f Front) Advance(m mymap.Map) Front {
	newFront := make(Front)
	for coord := range f {
		current := m.At(coord)
		if current < trailStart || current > trailEnd {
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

func (f Front) Ongoing(m mymap.Map) bool {
	for coord := range f {
		// map is not empty
		return m.At(coord) != trailEnd
	}
	// map is empty
	return false
}

func NewFront(head mymap.Coord) Front {
	return Front{head: struct{}{}}
}
