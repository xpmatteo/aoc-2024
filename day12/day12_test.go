package day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func Test_part1(t *testing.T) {
	tests := []struct {
		name                            string
		plot                            mapping.Map
		expectedArea, expectedPerimeter int
	}{
		{
			name:              "1x1",
			plot:              mapping.Map{"A"},
			expectedArea:      1,
			expectedPerimeter: 4,
		},
		{
			name:              "2x1",
			plot:              mapping.Map{"AA"},
			expectedArea:      2,
			expectedPerimeter: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewRegionSet(test.plot)
			assert.Equal(t, test.expectedArea, r.Area())
			assert.Equal(t, test.expectedPerimeter, r.Perimeter())
		})
	}
}
