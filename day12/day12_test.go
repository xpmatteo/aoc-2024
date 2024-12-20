package day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/mapping"
	"slices"
	"testing"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		plot     mapping.Map
		expected Report
	}{
		{
			name:     "A",
			plot:     mapping.Map{"A"},
			expected: Report{{'A', 1, 4}},
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
			// A region of R plants with price 12 * 18 = 216.
			//A region of I plants with price  4 *  8 = 32.
			//A region of C plants with price 14 * 28 = 392.
			//A region of F plants with price 10 * 18 = 180.
			//A region of V plants with price 13 * 20 = 260.
			//A region of J plants with price 11 * 20 = 220.
			//A region of C plants with price  1 *  4 = 4.
			//A region of E plants with price 13 * 18 = 234.
			//A region of I plants with price 14 * 22 = 308.
			//A region of M plants with price  5 * 12 = 60.
			//A region of S plants with price  3 *  8 = 24.
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
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewRegionSet(test.plot)
			assert.Equal(t, normalize(test.expected), normalize(r.Report()))
		})
	}
}

func normalize(report Report) []string {
	result := report.Strings()
	slices.Sort(result)
	return result
}
