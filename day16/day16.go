package day16

import (
	"github.com/xpmatteo/aoc-2024/mapping"
	"github.com/xpmatteo/aoc-2024/matrix"
	"math"
)

const (
	objectStart = 'S'
	objectEnd   = 'E'
	objectNone  = '.'
	objectWall  = '#'
)

func lowestScore(input mapping.Map) int {
	scores := matrix.New[int](input.Rows(), input.Cols())
	for r := 0; r < input.Rows(); r++ {
		for c := 0; c < input.Cols(); c++ {
			scores[r][c] = math.MaxInt
		}
	}
	start := findObject(input, objectStart)
	end := findObject(input, objectEnd)
	setScore(scores, start, 0)
	more := true
	for more {
		more = false
		input.ForEachCoord(func(c mapping.Coord, value int32) {
			if value == objectNone || value == objectEnd {
				updatedScore := bestScore(scores, c.OrthoNeighbors()) + 1
				if getScore(scores, c) > updatedScore {
					setScore(scores, c, updatedScore)
					more = true
				}
			}
		})
	}
	return getScore(scores, end)
}

func bestScore(scores [][]int, neighbors []mapping.Coord) int {
	lowest := math.MaxInt
	for _, neighbor := range neighbors {
		sc := getScore(scores, neighbor)
		if sc < lowest {
			lowest = sc
		}
	}
	return lowest
}

func getScore(scores [][]int, c mapping.Coord) int {
	return scores[c.Row][c.Col]
}

func setScore(scores [][]int, c mapping.Coord, score int) {
	scores[c.Row][c.Col] = score
}

func findObject(input mapping.Map, lookingFor int32) mapping.Coord {
	var result mapping.Coord
	input.ForEachCoord(func(c mapping.Coord, value int32) {
		if value == lookingFor {
			result = c
		}
	})
	return result
}
