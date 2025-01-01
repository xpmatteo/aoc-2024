package day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/mapping"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		input    mapping.Map
		expected int
	}{
		{
			name: "simple",
			input: mapping.Map{
				"#######",
				"#S...E#",
				"#######",
			},
			expected: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, lowestScore(test.input))
		})
	}
}
