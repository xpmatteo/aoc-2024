package day8

import (
	"github.com/xpmatteo/aoc-2024/day6"
	"github.com/xpmatteo/aoc-2024/mapping"
)

const Antinode = int32('#')

func countAntiNodes(m day6.Map) int {
	var result int
	m.ForEach(func(r int, c int, value int32) {
		if value == Antinode {
			result++
		}
	})
	return result
}

func plotAntinodes(input day6.Map) day6.Map {
	result := input.Clone()
	for _, frequency := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890" {
		coords := findCoordinates(input, frequency)
		forAllPairs(coords, func(c0, c1 mapping.Coord) {
			delta := c1.Minus(c0)
			result.SetSafe(c0.Minus(delta), Antinode)
			result.SetSafe(c1.Plus(delta), Antinode)
		})
	}
	return result
}

func forAllPairs(c []mapping.Coord, f func(c0 mapping.Coord, c1 mapping.Coord)) {
	for i := 0; i < len(c)-1; i++ {
		for j := i + 1; j < len(c); j++ {
			f(c[i], c[j])
		}
	}
}

func findCoordinates(m day6.Map, target int32) []mapping.Coord {
	var result []mapping.Coord
	m.ForEach(func(r int, c int, value int32) {
		if value == target {
			result = append(result, mapping.Coord{r, c})
		}
	})
	return result
}
