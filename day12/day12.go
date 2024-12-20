package day12

import (
	"fmt"
	"github.com/xpmatteo/aoc-2024/mapping"
	"github.com/xpmatteo/aoc-2024/matrix"
)

type Plant int32

type RegionSet struct {
	regions map[RegionId][]mapping.Coord
	plot    mapping.Map
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
		regions: make(map[RegionId][]mapping.Coord),
		plot:    plot,
	}
}

func (rs RegionSet) Report() Report {
	ids := rs.initRegionIds()
	areas := make(map[RegionId]int)
	perims := make(map[RegionId]int)
	plants := make(map[RegionId]Plant)
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		id := ids[c.Row][c.Col]
		plant := Plant(value)
		plants[id] = plant

		areas[id] = areas[id] + 1

		var perimeter int
		for _, coord := range c.OrthoNeighbors() {
			if !rs.plot.IsValid(coord) || ids[coord.Row][coord.Col] != id {
				perimeter++
			}
		}
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

func (rs RegionSet) initRegionIds() [][]RegionId {
	ids := matrix.New[RegionId](rs.plot.Rows(), rs.plot.Cols())
	i := 0
	rs.plot.ForEach(func(r int, c int, value int32) {
		i++
		ids[r][c] = RegionId(i)
	})
	return ids
}

func (r Report) Strings() []string {
	var result []string
	for _, line := range r {
		result = append(result, fmt.Sprintf("%c %d %d, ", line.plant, line.area, line.perimeter))
	}
	return result
}
