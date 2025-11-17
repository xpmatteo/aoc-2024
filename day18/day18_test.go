package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xpmatteo/aoc-2024/matrix"
	"os"
	"strconv"
	"strings"
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
			expected: 2,
		},
		{
			name:     "3x3 empty",
			size:     3,
			input:    "",
			expected: 4,
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
			expected: 10,
		},
		{
			name:     "small example from problem statement",
			size:     7,
			input:    "5,4\n4,2\n4,5\n3,0\n2,1\n6,3\n2,4\n1,5\n0,6\n3,3\n2,6\n5,1",
			expected: 22,
		},
		{
			name:     "part1",
			size:     71,
			input:    readFile("input.txt"),
			expected: 22,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, solvePart1(test.size, test.input))
		})
	}
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func TestParseInput(t *testing.T) {
	input := "" +
		"1,0\n" +
		"0,2\n" +
		"2,2"
	expected := matrix.New[bool](3, 3)
	expected[1][0] = true
	expected[0][2] = true

	mat, err := parseInput(3, 2, input)
	require.NoError(t, err)
	assert.Equal(t, expected, mat)
}

func parseInput(size int, maxBytes int, input string) ([][]bool, error) {
	lines := strings.Split(input, "\n")
	result := matrix.New[bool](size, size)
	for i, line := range lines {
		if i == maxBytes {
			return result, nil
		}
		if len(line) == 0 {
			continue
		}
		numbers := strings.Split(line, ",")
		if len(numbers) != 2 {
			return nil, fmt.Errorf("bad line: <%s>", line)
		}
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, fmt.Errorf("bad line 1: <%s>", line)
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, fmt.Errorf("bad line 1: <%s>", line)
		}
		result[x][y] = true
	}
	return result, nil
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
	blocked, err := parseInput(size, 1024, input)
	if err != nil {
		panic(err)
	}
	frontier := []point{{0, 0}}
	length := 0
	seen[0][0] = true
	for {
		if len(frontier) == 0 {
			panic(fmt.Errorf("empty frontier at length %d", length))
		}
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
