package day16

import (
	"github.com/xpmatteo/aoc-2024/mapping"
	"github.com/xpmatteo/aoc-2024/matrix"
	"math"
	"slices"
	"testing"
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

func (sc Scores) setScore(c mapping.Coord, score Score) {
	sc[c.Row][c.Col] = score
}

func (sc Scores) getScore(c mapping.Coord) Score {
	return sc[c.Row][c.Col]
}

func (m *Maze) LowestScore() int {
	scores := m.computeScoresFrom(m.start, Score{0, mapping.DirectionEast})
	return scores.getScore(m.end).value
}

type coordAndScore struct {
	c mapping.Coord
	s Score
}

var memoized = make(map[coordAndScore]Score)

func (m *Maze) endScore(c mapping.Coord, s Score) Score {
	//key := coordAndScore{c, s}
	//score, ok := memoized[key]
	//if ok {
	//	return score
	//}
	//score = m.computeScoresFrom(c, s).getScore(m.end)
	//memoized[key] = score
	//return score

	return m.computeScoresFrom(c, s).getScore(m.end)
}

func (m *Maze) computeScoresFrom(start mapping.Coord, startScore Score) Scores {
	scores := NewScores(m.theMap.Rows(), m.theMap.Cols())
	scores.setScore(start, startScore)
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
					scores.setScore(c, scoreHere)
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

func (m *Maze) isOnBestPath(c mapping.Coord) bool {
	return m.onBestPath[c.Row][c.Col]
}

func (m *Maze) setOnBestPath(tile mapping.Coord) {
	m.onBestPath[tile.Row][tile.Col] = true
}

func (m *Maze) propagateBestTileToSit(sc Scores, t *testing.T) []mapping.Coord {
	toExplore := []mapping.Coord{m.end}
	bestPlaces := []mapping.Coord{}
	endScore := sc.getScore(m.end)
	it := 0
	for len(toExplore) > 0 {
		t.Logf("It %3d, to explore %2d, best path %3d", it, len(toExplore), len(bestPlaces))
		c := toExplore[0]
		toExplore = toExplore[1:]
		bestPlaces = append(bestPlaces, c)
		for _, neighbor := range c.OrthoNeighbors() {
			if m.theMap.At(neighbor) == objectWall || slices.Contains(toExplore, neighbor) || slices.Contains(bestPlaces, neighbor) {
				continue
			}
			neighborScore := sc.getScore(neighbor)
			endScoreFromNeighbor := m.endScore(neighbor, neighborScore)
			if endScoreFromNeighbor == endScore {
				toExplore = append(toExplore, neighbor)
			}
		}
		it++
	}
	return bestPlaces
}

func (m *Maze) ShowBestPath(bestPlaces []mapping.Coord) string {
	clone := m.theMap.Clone()
	for _, c := range bestPlaces {
		clone.SetCoord(c, 'O')
	}
	return clone.String()
}
