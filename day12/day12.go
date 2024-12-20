package day12

import (
	"fmt"
	"github.com/xpmatteo/aoc-2024/mapping"
	"github.com/xpmatteo/aoc-2024/matrix"
)

type Plant int32

type RegionSet struct {
	ids  [][]RegionId
	plot mapping.Map
}

type RegionId int

type Report []ReportLine
type ReportLine struct {
	plant     Plant
	area      int
	perimeter int
}

func NewRegionSet(plot mapping.Map) RegionSet {
	rs := RegionSet{
		ids:  initRegionIds(plot),
		plot: plot,
	}
	rs.mergeRegionIds()
	return rs
}

func initRegionIds(plot mapping.Map) [][]RegionId {
	ids := matrix.New[RegionId](plot.Rows(), plot.Cols())
	i := 0
	plot.ForEach(func(r int, c int, value int32) {
		i++
		ids[r][c] = RegionId(i)
	})
	return ids
}

func (rs RegionSet) ReportPart1() Report {
	perimeterF := rs.perimeterPart1
	return rs.ReportFunc(perimeterF)
}

func (rs RegionSet) ReportPart2() Report {
	perimeterF := rs.perimeterPart2
	return rs.ReportFunc(perimeterF)
}

func (rs RegionSet) ReportFunc(perimF func(c mapping.Coord, id RegionId) int) Report {
	areas := make(map[RegionId]int)
	perims := make(map[RegionId]int)
	plants := make(map[RegionId]Plant)
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		id := rs.ids[c.Row][c.Col]
		plants[id] = Plant(value)
		areas[id] = areas[id] + 1
		perims[id] = perims[id] + perimF(c, id)
	})
	var result []ReportLine
	for id, area := range areas {
		result = append(result, ReportLine{
			plant:     plants[id],
			area:      area,
			perimeter: perims[id],
		})
	}
	return result
}

func (rs RegionSet) perimeterPart1(c mapping.Coord, id RegionId) int {
	var perimeter int
	for _, coord := range c.OrthoNeighbors() {
		if !rs.isValid(coord) || rs.ids[coord.Row][coord.Col] != id {
			perimeter++
		}
	}
	return perimeter
}

// I suspect this function is wrong bc it's not symmetric!
func (rs RegionSet) perimeterPart2(c mapping.Coord, id RegionId) int {
	var perimeter int
	{
		alreadyCounted := rs.sameRegion(c.West(), id) && !rs.sameRegion(c.NorthWest(), id)
		if !alreadyCounted && !rs.sameRegion(c.North(), id) {
			perimeter++
		}
	}
	{
		alreadyCounted := rs.sameRegion(c.North(), id) && !rs.sameRegion(c.NorthEast(), id)
		if !alreadyCounted && !rs.sameRegion(c.East(), id) {
			perimeter++
		}
	}
	{
		alreadyCounted := rs.sameRegion(c.West(), id) && !rs.sameRegion(c.SouthWest(), id)
		if !alreadyCounted && !rs.sameRegion(c.South(), id) {
			perimeter++
		}
	}
	{
		alreadyCounted := rs.sameRegion(c.North(), id) && !rs.sameRegion(c.NorthWest(), id)
		if !alreadyCounted && !rs.sameRegion(c.West(), id) {
			perimeter++
		}
	}
	return perimeter
}

func (rs RegionSet) sameRegion(c mapping.Coord, id RegionId) bool {
	return rs.isValid(c) && rs.idOf(c) == id
}

func (rs RegionSet) idOf(c mapping.Coord) RegionId {
	return rs.ids[c.Row][c.Col]
}

func (rs RegionSet) mergeRegionIds() {
	more := true
	for more {
		more = rs.mergeIdsOnce()
	}
}

func (rs RegionSet) mergeIdsOnce() bool {
	more := false
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		id := rs.regionIdOf(c)
		for _, neighbor := range c.OrthoNeighbors() {
			if rs.isValid(neighbor) {
				neighborPlant := Plant(rs.plot.At(neighbor))
				neighborId := rs.regionIdOf(neighbor)
				if Plant(value) == neighborPlant && id != neighborId {
					least := min(id, neighborId)
					rs.setRegionId(c, least)
					rs.setRegionId(neighbor, least)
					more = true
				}
			}
		}
	})
	return more
}

func (rs RegionSet) isValid(neighbor mapping.Coord) bool {
	return rs.plot.IsValid(neighbor)
}

func (rs RegionSet) regionIdOf(c mapping.Coord) RegionId {
	return rs.ids[c.Row][c.Col]
}

func (rs RegionSet) setRegionId(c mapping.Coord, id RegionId) {
	rs.ids[c.Row][c.Col] = id
}

func (r Report) Strings() []string {
	var result []string
	for _, line := range r {
		result = append(result, fmt.Sprintf("%c %d %d, ", line.plant, line.area, line.perimeter))
	}
	return result
}

func (r Report) TotalCost() int {
	var total int
	for _, line := range r {
		total += line.area * line.perimeter
	}
	return total
}
