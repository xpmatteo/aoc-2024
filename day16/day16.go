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

type Scores [][]Score

type Maze struct {
	theMap     mapping.Map
	start, end mapping.Coord
	onBestPath [][]bool
}

type Score struct {
	value int
	dir   mapping.Direction
}

func NewMaze(input mapping.Map) *Maze {
	bestPaths := matrix.New[bool](input.Rows(), input.Cols())
	return &Maze{
		theMap:     input,
		start:      input.FindObject(objectStart),
		end:        input.FindObject(objectEnd),
		onBestPath: bestPaths,
	}
}

func NewScores(rows, cols int) Scores {
	sc := matrix.New[Score](rows, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			sc[r][c].value = math.MaxInt
		}
	}
	return sc
}

func (sc Scores) setScore(c mapping.Coord, score int, dir mapping.Direction) {
	sc[c.Row][c.Col].value = score
	sc[c.Row][c.Col].dir = dir
}

func (sc Scores) getScore(c mapping.Coord) Score {
	return sc[c.Row][c.Col]
}

func (m *Maze) LowestScore() int {
	scores := m.computeScoresFrom(m.start, 0, mapping.DirectionEast)
	return scores.getScore(m.end).value
}

func (m *Maze) computeScoresFrom(start mapping.Coord, startScore int, startDir mapping.Direction) Scores {
	scores := NewScores(m.theMap.Rows(), m.theMap.Cols())
	scores.setScore(start, startScore, startDir)
	more := true
	for more {
		more = false
		m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
			if value == objectNone || value == objectEnd {
				scoreHere := scores[c.Row][c.Col]
				scoreHere = scoreHere.ImproveScore(mapping.DirectionNorth, scores.getScore(c.North()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionEast, scores.getScore(c.East()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionSouth, scores.getScore(c.South()))
				scoreHere = scoreHere.ImproveScore(mapping.DirectionWest, scores.getScore(c.West()))
				if scoreHere != scores.getScore(c) {
					scores.setScore(c, scoreHere.value, scoreHere.dir)
					more = true
				}
			}
		})
	}
	return scores
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

func (m *Maze) CountBestTilesToSit(sc Scores) int {
	m.setOnBestPath(m.end)
	more := true
	for more {
		more = m.propagateBestTileToSit(sc)
	}
	total := 0
	m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
		if m.isOnBestPath(c) {
			total++
		}
	})
	return total
}

func (m *Maze) isOnBestPath(c mapping.Coord) bool {
	return m.onBestPath[c.Row][c.Col]
}

func (m *Maze) setOnBestPath(tile mapping.Coord) {
	m.onBestPath[tile.Row][tile.Col] = true
}

func (m *Maze) propagateBestTileToSit(sc Scores) bool {
	more := false
	m.theMap.ForEachCoord(func(c mapping.Coord, value int32) {
		if !m.isOnBestPath(c) {
			return
		}
		scoreHere := sc.getScore(c).value
		for _, neighbor := range c.OrthoNeighbors() {
			neighborScore := sc.getScore(neighbor).value
			//if (neighborScore == scoreHere.value-1 || neighborScore == scoreHere.value-1001) && !m.isOnBestPath(neighbor) {
			if (neighborScore < scoreHere || neighborScore == scoreHere+999) && !m.isOnBestPath(neighbor) {
				m.setOnBestPath(neighbor)
				more = true
			}
		}
	})
	return more
}

func (m *Maze) ShowBestPath() string {
	clone := m.theMap.Clone()
	clone.ForEachCoord(func(c mapping.Coord, object int32) {
		if m.isOnBestPath(c) {
			clone.SetCoord(c, 'O')
		}
	})
	return clone.String()
}
