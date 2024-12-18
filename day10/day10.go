package day10

import "github.com/xpmatteo/aoc-2024/maps"

const trailStart = '0'
const trailEnd = '9'

func scoreAllTrails(input maps.Map) int {
	score := 0
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := maps.Coord{r, c}
			score += scoreTrailHead(trailHead, input)
		}
	})
	return score
}

func scoreTrailHead(trailHead maps.Coord, m maps.Map) int {
	front := NewFront(trailHead)
	for front.Ongoing(m) {
		front = front.Advance(m)
	}
	return front.Score()
}
