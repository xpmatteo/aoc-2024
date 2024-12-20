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
			report := NewRegionSet(test.plot).Report()
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
