package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/matrix"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		input    string
		expected int
	}{
		{
			name:     "2x2 empty",
			size:     2,
			input:    "",
			expected: 3,
		},
		{
			name:     "3x3 empty",
			size:     3,
			input:    "",
			expected: 5,
		},
		{
			name: "5x5 with obstacle",
			size: 5,
			// .#...
			// ...#.
			// ...#.
			// ...#.
			// ...#.

			input: "" +
				"1,0\n" +
				"3,1\n" +
				"3,2\n" +
				"3,3\n" +
				"3,4\n" +
				"",
			expected: 11,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, solvePart1(test.size, test.input))
		})
	}
}

type point struct{ x, y int }

func (p point) plus(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

var directions = []point{
	{1, 0}, // right
	{0, 1}, // bottom
	{-1, 0},
	{0, -1},
}

func solvePart1(size int, input string) int {
	seen := matrix.New[bool](size, size)
	blocked := matrix.New[bool](size, size)
	frontier := []point{{0, 0}}
	length := 1
	seen[0][0] = true
	for {
		newFrontier := []point{}
		length++
		for _, p := range frontier {
			for _, dir := range directions {
				newPoint := p.plus(dir)
				switch {
				case isExit(size, newPoint):
					return length
				case outOfRange(size, newPoint):
				case get(seen, newPoint):
				case get(blocked, newPoint):
					continue
				default:
					set(seen, newPoint)
					newFrontier = append(newFrontier, newPoint)
				}
			}
		}
		frontier = newFrontier
	}
}

func isExit(size int, p point) bool {
	return p.x == size-1 && p.y == size-1
}

func set(m [][]bool, p point) {
	m[p.x][p.y] = true
}

func get(seen [][]bool, p point) bool {
	return seen[p.x][p.y]
}

func outOfRange(size int, p point) bool {
	return p.x < 0 || p.y < 0 || p.x >= size || p.y >= size
}
