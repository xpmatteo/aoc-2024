package day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		plot     mapping.Map
		expected []ReportLine
	}{
		{
			name:     "A",
			plot:     mapping.Map{"A"},
			expected: []ReportLine{{'A', 1, 4}},
		},
		{
			name:     "AA",
			plot:     mapping.Map{"AA"},
			expected: []ReportLine{{'A', 2, 6}},
		},
		{
			name: "AA + AA",
			plot: mapping.Map{
				"AA",
				"AA",
			},
			expected: []ReportLine{{'A', 4, 8}},
		},
		{
			name:     "AB",
			plot:     mapping.Map{"AB"},
			expected: []ReportLine{{'A', 1, 4}, {'B', 1, 4}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewRegionSet(test.plot)
			assert.Equal(t, test.expected, r.Report(), "area")
		})
	}
}
