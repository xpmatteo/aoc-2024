package day10

import "github.com/xpmatteo/aoc-2024/maps"

type Front []maps.Coord

func (f Front) Advance(m maps.Map) Front {
	var newFront Front
	for _, coord := range f {
		current := m.At(coord)
		if current == '9' {
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

func (f Front) Score(m maps.Map) int {
	score := 0
	for _, coord := range f {
		if m.At(coord) == '9' {
			score++
		}
	}
	return score
}

func NewFront(head maps.Coord) Front {
	return Front{head}
}
