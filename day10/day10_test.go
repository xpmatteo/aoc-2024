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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedScore, scoreAllTrails(test.input))
		})
	}
}
