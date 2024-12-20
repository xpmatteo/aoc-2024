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
	return RegionSet{
		ids:  initRegionIds(plot),
		plot: plot,
	}
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

func (rs RegionSet) Report() Report {
	rs.mergeRegionIds()
	areas := make(map[RegionId]int)
	perims := make(map[RegionId]int)
	plants := make(map[RegionId]Plant)
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		id := rs.ids[c.Row][c.Col]
		plant := Plant(value)
		plants[id] = plant
		areas[id] = areas[id] + 1
		perimeter := rs.perimeter(c, id)
		perims[id] = perims[id] + perimeter
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

func (rs RegionSet) perimeter(c mapping.Coord, id RegionId) int {
	var perimeter int
	for _, coord := range c.OrthoNeighbors() {
		if !rs.plot.IsValid(coord) || rs.ids[coord.Row][coord.Col] != id {
			perimeter++
		}
	}
	return perimeter
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
		plant := Plant(value)
		id := rs.regionIdOf(c)
		for _, neighbor := range c.OrthoNeighbors() {
			if rs.plot.IsValid(neighbor) {
				neighborPlant := Plant(rs.plot.At(neighbor))
				neighborId := rs.regionIdOf(neighbor)
				if plant == neighborPlant && id != neighborId {
					l := min(id, neighborId)
					rs.setRegionId(c, l)
					rs.setRegionId(neighbor, l)
					more = true
				}
			}
		}
	})
	return more
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
