package day10

import (
	"github.com/xpmatteo/aoc-2024/mapping"
)

type Front []mapping.Coord

func (f Front) Advance(m mapping.Map) Front {
	var newFront Front
	for _, coord := range f {
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
				newFront = append(newFront, neighbor)
			}
		}
	}
	return newFront
}

func (f Front) ScorePart1() int {
	set := make(map[mapping.Coord]struct{})
	for _, coord := range f {
		set[coord] = struct{}{}
	}
	return len(set)
}

func (f Front) Rating() int {
	return len(f)
}

func (f Front) Ongoing(m mapping.Map) bool {
	for _, coord := range f {
		// map is not empty
		return m.At(coord) != trailEnd
	}
	// map is empty
	return false
}

func NewFront(head mapping.Coord) Front {
	return Front{head}
}
