package day10

import "github.com/xpmatteo/aoc-2024/maps"

type Front []maps.Coord

func (f Front) Advance(m maps.Map) Front {
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
	set := make(map[maps.Coord]struct{})
	for _, coord := range f {
		set[coord] = struct{}{}
	}
	return len(set)
}

func (f Front) ScorePart2() int {
	return len(f)
}

func (f Front) Ongoing(m maps.Map) bool {
	for _, coord := range f {
		// map is not empty
		return m.At(coord) != trailEnd
	}
	// map is empty
	return false
}

func NewFront(head maps.Coord) Front {
	return Front{head}
}
