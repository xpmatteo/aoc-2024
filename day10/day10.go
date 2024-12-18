package day10

import "github.com/xpmatteo/aoc-2024/maps"

func scoreAllTrails(input maps.Map) int {
	trailHead := maps.Coord{0, 3}
	return scoreTrailHead(trailHead, input)
}

func scoreTrailHead(trailHead maps.Coord, m maps.Map) int {
	score := 0
	front := NewFront(trailHead)
	for len(front) > 0 {
		front = front.Advance(m)
		score += front.Score(m)
	}
	return score
}
