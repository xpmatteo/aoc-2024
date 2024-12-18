package day10

import (
	"github.com/xpmatteo/aoc-2024/mapping"
)

const trailStart = '0'
const trailEnd = '9'

func scoreAllTrails(input mapping.Map) int {
	score := 0
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := mapping.Coord{r, c}
			front := exploreTrails(trailHead, input)
			score += front.ScorePart1()
		}
	})
	return score
}

func rateAllTrails(input mapping.Map) int {
	score := 0
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := mapping.Coord{r, c}
			front := exploreTrails(trailHead, input)
			score += front.Rating()
		}
	})
	return score
}

func exploreTrails(trailHead mapping.Coord, m mapping.Map) Front {
	front := NewFront(trailHead)
	for front.Ongoing(m) {
		front = front.Advance(m)
	}
	return front
}
