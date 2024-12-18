package day10

import (
	mymap "github.com/xpmatteo/aoc-2024/maps"
	"maps"
)

const trailStart = '0'
const trailEnd = '9'

func scoreAllTrails(input mymap.Map) int {
	fronts := make(Front)
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := mymap.Coord{r, c}
			front := advanceFront(trailHead, input)
			maps.Copy(fronts, front)
		}
	})
	return len(fronts)
}

func advanceFront(trailHead mymap.Coord, m mymap.Map) Front {
	front := NewFront(trailHead)
	for front.Ongoing(m) {
		front = front.Advance(m)
	}
	return front
}
