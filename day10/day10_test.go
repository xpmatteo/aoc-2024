package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         mapping.Map
		expectedScore int
	}{
		{
			name:          "one complete trail",
			input:         mapping.Map{"...0123456789..."},
			expectedScore: 1,
		},
		{
			name:          "one incomplete trail",
			input:         mapping.Map{"...012345678..."},
			expectedScore: 0,
		},
		{
			name:          "forked trail",
			input:         mapping.Map{"9876543210123456789..."},
			expectedScore: 2,
		},
		{
			name: "twisting trail",
			input: mapping.Map{
				"0123456789...",
				" 0123456789...",
			},
			expectedScore: 3,
		},
		{
			name: "sample 4",
			input: mapping.ParseMap(`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`),
			expectedScore: 4,
		},
		{
			name: "sample two trailheads 1+2",
			input: mapping.ParseMap(`
10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`),
			expectedScore: 3,
		},
		{
			name: "larger sample",
			input: mapping.ParseMap(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`),
			expectedScore: 36,
		},
		{
			name:          "real",
			input:         mapping.ParseMap(day1.ReadFile("day10.txt")),
			expectedScore: 535,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, scoreAllTrails(test.input))
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name          string
		input         mapping.Map
		expectedScore int
	}{
		{
			name: "twisting trail",
			input: mapping.Map{
				"0123456789...",
				" 0123456789...",
			},
			expectedScore: 11,
		},
		{
			name: "rating 3",
			input: mapping.ParseMap(`
.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`),
			expectedScore: 3,
		},
		{
			name: "sample two trailheads 1+2",
			input: mapping.ParseMap(`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`),
			expectedScore: 13,
		},
		{
			name: "227",
			input: mapping.ParseMap(`
012345
123456
234567
345678
4.6789
56789.`),
			expectedScore: 227,
		},
		{
			name: "larger sample",
			input: mapping.ParseMap(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`),
			expectedScore: 81,
		},
		{
			name:          "real",
			input:         mapping.ParseMap(day1.ReadFile("day10.txt")),
			expectedScore: 1186,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, rateAllTrails(test.input))
		})
	}
}
