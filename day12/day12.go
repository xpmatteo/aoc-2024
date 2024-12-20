package day12

import "github.com/xpmatteo/aoc-2024/mapping"

type Plant int32

type RegionSet struct {
	regions map[regionId][]mapping.Coord
	plot    mapping.Map
}

type regionId int

type ReportLine struct {
	plant     Plant
	area      int
	perimeter int
}

func NewRegionSet(plot mapping.Map) RegionSet {
	return RegionSet{
		regions: make(map[regionId][]mapping.Coord),
		plot:    plot,
	}
}

func (rs RegionSet) Report() []ReportLine {
	areas := make(map[Plant]int)
	perims := make(map[Plant]int)
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		plant := Plant(value)
		areas[plant] = areas[plant] + 1
		perims[plant] = perims[plant] + rs.perimeter(c)
	})
	var result []ReportLine
	for plant, area := range areas {
		result = append(result, ReportLine{
			plant:     plant,
			area:      area,
			perimeter: perims[plant],
		})
	}
	return result
}

func (rs RegionSet) perimeter(c mapping.Coord) int {
	var perimeter int
	plant := rs.plot.At(c)
	for _, coord := range c.OrthoNeighbors() {
		if rs.plot.At(coord) != plant {
			perimeter++
		}
	}
	return perimeter
}
