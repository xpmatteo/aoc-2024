package day10

import "github.com/xpmatteo/aoc-2024/maps"

const trailStart = '0'
const trailEnd = '9'

func scoreAllTrailsPart1(input maps.Map) int {
	score := 0
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := maps.Coord{r, c}
			front := scoreTrailHead(trailHead, input)
			score += front.ScorePart1()
		}
	})
	return score
}

func scoreAllTrailsPart2(input maps.Map) int {
	score := 0
	input.ForEach(func(r, c int, value int32) {
		if value == trailStart {
			trailHead := maps.Coord{r, c}
			front := scoreTrailHead(trailHead, input)
			score += front.ScorePart2()
		}
	})
	return score
}

func scoreTrailHead(trailHead maps.Coord, m maps.Map) Front {
	front := NewFront(trailHead)
	for front.Ongoing(m) {
		front = front.Advance(m)
	}
	return front
}
