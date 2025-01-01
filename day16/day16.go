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
	scores     [][]Score
}

type Score struct {
	dir   mapping.Direction
	value int
}

func NewMaze(input mapping.Map) *Maze {
	scores := matrix.New[Score](input.Rows(), input.Cols())
	for r := 0; r < input.Rows(); r++ {
		for c := 0; c < input.Cols(); c++ {
			scores[r][c].value = math.MaxInt
		}
	}
	return &Maze{
		theMap: input,
		start:  input.FindObject(objectStart),
		end:    input.FindObject(objectEnd),
		scores: scores,
	}
}

func (m *Maze) setScore(c mapping.Coord, score int, dir mapping.Direction) {
	m.scores[c.Row][c.Col].value = score
	m.scores[c.Row][c.Col].dir = dir
}

func (m *Maze) getScore(c mapping.Coord) Score {
	return m.scores[c.Row][c.Col]
}

func (m *Maze) LowestScore() int {
	m.setScore(m.start, 0, mapping.DirectionNone)
	more := true
	for more {
		more = false
		m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
			if value == objectNone || value == objectEnd {
				updatedScore := m.updatedScore(c, c.OrthoNeighbors1())
				if updatedScore.value < m.getScore(c).value {
					m.setScore(c, updatedScore.value, "")
					more = true
				}
			}
		})
	}
	return m.getScore(m.end).value
}

func (m *Maze) updatedScore(c mapping.Coord, neighbors []mapping.CoordDir) Score {
	lowest := math.MaxInt
	for _, neighbor := range neighbors {
		sc := m.getScore(neighbor.C)
		if sc.value < lowest {
			lowest = sc.value
		}
	}
	return Score{value: lowest + 1}
}
