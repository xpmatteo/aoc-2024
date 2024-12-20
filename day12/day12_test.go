package day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/mapping"
	"slices"
	"testing"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func Test_part1(t *testing.T) {
	tests := []struct {
		name              string
		plot              mapping.Map
		expected          Report
		expectedTotalCost int
	}{
		{
			name:              "A",
			plot:              mapping.Map{"A"},
			expected:          Report{{'A', 1, 4}},
			expectedTotalCost: 4,
		},
		{
			name:     "AA",
			plot:     mapping.Map{"AA"},
			expected: Report{{'A', 2, 6}},
		},
		{
			name: "AA + AA",
			plot: mapping.Map{
				"AA",
				"AA",
			},
			expected: Report{{'A', 4, 8}},
		},
		{
			name:     "AB",
			plot:     mapping.Map{"AB"},
			expected: Report{{'A', 1, 4}, {'B', 1, 4}},
		},
		{
			name: "AB + AB",
			plot: mapping.Map{
				"AB",
				"AB"},
			expected: []ReportLine{{'A', 2, 6}, {'B', 2, 6}},
		},
		{
			name:     "the dreaded ABA",
			plot:     mapping.Map{"ABA"},
			expected: Report{{'A', 1, 4}, {'B', 1, 4}, {'A', 1, 4}},
		},
		{
			name: "first given example",
			plot: mapping.ParseMap(`
AAAA
BBCD
BBCC
EEEC`),
			expected: Report{{'A', 4, 10}, {'B', 4, 8}, {'C', 4, 10}, {'D', 1, 4}, {'E', 3, 8}},
		},
		{
			name: "second given example",
			plot: mapping.ParseMap(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`),
			expected: Report{{'O', 21, 36}, {'X', 1, 4}, {'X', 1, 4}, {'X', 1, 4}, {'X', 1, 4}},
		},
		{
			name: "third given example",
			plot: mapping.ParseMap(`
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`),
			expected: Report{
				{'R', 12, 18},
				{'I', 4, 8},
				{'C', 1, 4},
				{'C', 14, 28},
				{'F', 10, 18},
				{'V', 13, 20},
				{'J', 11, 20},
				{'E', 13, 18},
				{'I', 14, 22},
				{'M', 5, 12},
				{'S', 3, 8},
			},
			expectedTotalCost: 1930,
		},
		{
			name:              "real",
			plot:              mapping.ParseMap(day1.ReadFile("day12.txt")),
			expectedTotalCost: 1370258,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := NewRegionSet(test.plot).ReportPart1()
			if test.expectedTotalCost > 0 {
				assert.Equal(t, test.expectedTotalCost, report.TotalCost())
			}
			if test.expected != nil {
				assert.Equal(t, normalize(test.expected), normalize(report))
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name              string
		plot              mapping.Map
		expected          Report
		expectedTotalCost int
	}{
		{
			name:              "A",
			plot:              mapping.Map{"A"},
			expected:          Report{{'A', 1, 4}},
			expectedTotalCost: 4,
		},
		{
			name:     "AA",
			plot:     mapping.Map{"AA"},
			expected: Report{{'A', 2, 4}},
		},
		{
			name: "AA + AA",
			plot: mapping.Map{
				"AA",
				"AA",
			},
			expected: Report{{'A', 4, 4}},
		},
		{
			name: "first given example",
			plot: mapping.ParseMap(`
AAAA
BBCD
BBCC
EEEC`),
			expected:          Report{{'A', 4, 4}, {'B', 4, 4}, {'C', 4, 8}, {'D', 1, 4}, {'E', 3, 4}},
			expectedTotalCost: 80,
		},
		{
			name: "the window",
			plot: mapping.ParseMap(`
OOO
OXO
OOO`),
			expected:          Report{{'O', 8, 8}, {'X', 1, 4}},
			expectedTotalCost: 68,
		},
		{
			name: "second given example",
			plot: mapping.ParseMap(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`),
			expected:          Report{{'O', 21, 20}, {'X', 1, 4}, {'X', 1, 4}, {'X', 1, 4}, {'X', 1, 4}},
			expectedTotalCost: 436,
		},
		{
			name: "the E shape",
			plot: mapping.ParseMap(`
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`),
			expected:          Report{{'E', 17, 12}, {'X', 4, 4}, {'X', 4, 4}},
			expectedTotalCost: 236,
		},
		{
			name: "the A B B A shape",
			plot: mapping.ParseMap(`
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`),
			expected: Report{{'A', 28, 12}, {'B', 4, 4}, {'B', 4, 4}},
		},
		{
			name: "third given example",
			plot: mapping.ParseMap(`
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`),
			expected: Report{
				{'R', 12, 10},
				{'I', 4, 4},
				{'C', 14, 22},
				{'F', 10, 12},
				{'V', 13, 10},
				{'J', 11, 12},
				{'C', 1, 4},
				{'E', 13, 8},
				{'I', 14, 16},
				{'M', 5, 6},
				{'S', 3, 6},
			},
			expectedTotalCost: 1206,
		},
		{
			name:              "real",
			plot:              mapping.ParseMap(day1.ReadFile("day12.txt")),
			expectedTotalCost: 805814,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := NewRegionSet(test.plot).ReportPart2()
			if test.expectedTotalCost > 0 {
				assert.Equal(t, test.expectedTotalCost, report.TotalCost())
			}
			if test.expected != nil {
				assert.Equal(t, normalize(test.expected), normalize(report))
			}
		})
	}
}

func normalize(report Report) []string {
	result := report.Strings()
	slices.Sort(result)
	return result
}
