package day12

import "github.com/xpmatteo/aoc-2024/mapping"

type Plant int32

type RegionSet struct {
	regions map[regionId][]mapping.Coord
	plot    mapping.Map
}

type regionId int

func (rs RegionSet) Add(id regionId, letter int32, coord mapping.Coord) RegionSet {
	return rs
}

func (rs RegionSet) Area() int {
	area := 0
	rs.plot.ForEach(func(r int, c int, value int32) {
		area++
	})
	return area
}

func (rs RegionSet) Perimeter() int {
	var perimeter int
	rs.plot.ForEachCoord(func(c mapping.Coord, value int32) {
		for _, coord := range c.OrthoNeighbors() {
			if rs.plot.At(coord) != 'A' {
				perimeter++
			}
		}
	})
	return perimeter
}

func NewRegionSet(plot mapping.Map) RegionSet {
	return RegionSet{
		regions: make(map[regionId][]mapping.Coord),
		plot:    plot,
	}
}
