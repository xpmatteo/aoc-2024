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
	value int
	dir   mapping.Direction
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
	m.setScore(m.start, 0, mapping.DirectionEast)
	more := true
	for more {
		more = false
		m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
			if value == objectNone || value == objectEnd {
				scoreHere := m.getScore(c)
				scoreHere = scoreHere.ImproveScore(mapping.DirectionNorth, m.getScore(c.North()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionEast, m.getScore(c.East()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionSouth, m.getScore(c.South()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionWest, m.getScore(c.West()))
				if scoreHere != m.getScore(c) {
					m.setScore(c, scoreHere.value, scoreHere.dir)
					more = true
				}
			}
		})
	}
	return m.getScore(m.end).value
}

func (here Score) ImproveScore(neighborIs mapping.Direction, neighborScore Score) Score {
	if neighborScore.dir == mapping.DirectionNone {
		return here
	}
	if neighborScore.dir == neighborIs.Opposite() {
		if neighborScore.value+1 < here.value {
			return Score{
				value: neighborScore.value + 1,
				dir:   neighborScore.dir,
			}
		}
		return here
	}
	if neighborScore.value+1000+1 < here.value {
		return Score{
			value: neighborScore.value + 1000 + 1,
			dir:   neighborIs.Opposite(),
		}
	}
	return here
}
