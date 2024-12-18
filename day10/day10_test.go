package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/maps"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name          string
		input         maps.Map
		expectedScore int
	}{
		{
			name:          "one complete trail",
			input:         maps.Map{"...0123456789..."},
			expectedScore: 1,
		},
		{
			name:          "one incomplete trail",
			input:         maps.Map{"...012345678..."},
			expectedScore: 0,
		},
		{
			name:          "forked trail",
			input:         maps.Map{"9876543210123456789..."},
			expectedScore: 2,
		},
		{
			name: "sample 4",
			input: maps.ParseMap(`
..90..9
...1.98
...2..7
6543456
765.987
876....
987....`),
			expectedScore: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, scoreAllTrails(test.input))
		})
	}
}
