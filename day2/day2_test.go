package day2

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"testing"
)

const sample = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_parse(t *testing.T) {
	twoLines := `
7 6 4 2 1
1 2 7 8 9`
	expected := []Report{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
	}
	assert.Equal(t, expected, parse(twoLines))
}

// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
func Test_safety(t *testing.T) {
	tests := []struct {
		input        string
		expectedSafe bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			report := parseReport(test.input)
			assert.Equal(t, test.expectedSafe, report.isSafe())
		})
	}
}

func Test_day2_part1_sample(t *testing.T) {
	assert.Equal(t, 2, countSafeReports(parse(sample)))
}

func Test_day2_part1(t *testing.T) {
	assert.Equal(t, 371, countSafeReports(parse(day1.ReadFile("day2.txt"))))
}

func Test_remove(t *testing.T) {
	assert.Equal(t, Report{2, 3, 4}, remove(Report{1, 2, 3, 4}, 0))
	assert.Equal(t, Report{1, 3, 4}, remove(Report{1, 2, 3, 4}, 1))
	assert.Equal(t, Report{1, 2, 4}, remove(Report{1, 2, 3, 4}, 2))
	assert.Equal(t, Report{1, 2, 3}, remove(Report{1, 2, 3, 4}, 3))
}

func Test_Dampen(t *testing.T) {
	report := Report{1, 3, 2, 4, 5}
	assert.False(t, report.isSafe())
	assert.Equal(t, Report{1, 2, 4, 5}, report.Dampen())
}

func Test_day2_part2_sample(t *testing.T) {
	assert.Equal(t, 4, countSafeReports(dampen(parse(sample))))
}

func Test_day2_part2(t *testing.T) {
	assert.Equal(t, 426, countSafeReports(dampen(parse(day1.ReadFile("day2.txt")))))
}
