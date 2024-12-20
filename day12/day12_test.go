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
