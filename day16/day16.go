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

type Maze struct {
	theMap     mapping.Map
	start, end mapping.Coord
	scores     [][]int
}

func NewMaze(input mapping.Map) *Maze {
	scores := matrix.New[int](input.Rows(), input.Cols())
	for r := 0; r < input.Rows(); r++ {
		for c := 0; c < input.Cols(); c++ {
			scores[r][c] = math.MaxInt
		}
	}
	return &Maze{
		theMap: input,
		start:  findObject(input, objectStart),
		end:    findObject(input, objectEnd),
		scores: scores,
	}
}

func (m *Maze) setScore(c mapping.Coord, score int) {
	m.scores[c.Row][c.Col] = score
}

func (m *Maze) getScore(c mapping.Coord) int {
	return m.scores[c.Row][c.Col]
}

func (m *Maze) LowestScore() int {
	m.setScore(m.start, 0)
	more := true
	for more {
		more = false
		m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
			if value == objectNone || value == objectEnd {
				updatedScore := m.bestScore(c.OrthoNeighbors()) + 1
				if m.getScore(c) > updatedScore {
					m.setScore(c, updatedScore)
					more = true
				}
			}
		})
	}
	return m.getScore(m.end)
}

func (m *Maze) bestScore(neighbors []mapping.Coord) int {
	lowest := math.MaxInt
	for _, neighbor := range neighbors {
		sc := m.getScore(neighbor)
		if sc < lowest {
			lowest = sc
		}
	}
	return lowest
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
