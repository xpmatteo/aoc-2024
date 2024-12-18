package day8

import (
	"github.com/xpmatteo/aoc-2024/mapping"
)

const Antinode = int32('#')

func plotAntinodes(input maps.Map) maps.Map {
	result := input.Clone()
	for _, frequency := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890" {
		coords := findCoordinates(input, frequency)
		forAllPairs(coords, func(c0, c1 maps.Coord) {
			delta := c1.Minus(c0)
			result.SetSafe(c0.Minus(delta), Antinode)
			result.SetSafe(c1.Plus(delta), Antinode)
		})
	}
	return result
}

func plotAntinodesPart2(input maps.Map) maps.Map {
	result := input.Clone()
	for _, frequency := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890" {
		coords := findCoordinates(input, frequency)
		forAllPairs(coords, func(c0, c1 maps.Coord) {
			delta := c1.Minus(c0)
			applyDelta(result, c0, delta)
		})
	}
	return result
}

func applyDelta(m maps.Map, startPoint maps.Coord, delta maps.Coord) {
	for p := startPoint; m.IsValid(p); p = p.Minus(delta) {
		m.SetCoord(p, Antinode)
	}
	for p := startPoint; m.IsValid(p); p = p.Plus(delta) {
		m.SetCoord(p, Antinode)
	}
}

func forAllPairs(c []maps.Coord, f func(c0 maps.Coord, c1 maps.Coord)) {
	for i := 0; i < len(c)-1; i++ {
		for j := i + 1; j < len(c); j++ {
			f(c[i], c[j])
		}
	}
}

func findCoordinates(m maps.Map, target int32) []maps.Coord {
	var result []maps.Coord
	m.ForEach(func(r int, c int, value int32) {
		if value == target {
			result = append(result, maps.Coord{r, c})
		}
	})
	return result
}

func countAntiNodes(m maps.Map) int {
	var result int
	m.ForEach(func(r int, c int, value int32) {
		if value == Antinode {
			result++
		}
	})
	return result
}
