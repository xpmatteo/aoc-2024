package day2

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"log"
	"strconv"
	"strings"
	"testing"
)

const sample = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

type Report []int

func (r Report) isSafe() bool {
	sign := sgn(r[0] - r[1])
	for i := range len(r) - 1 {
		d := r[i] - r[i+1]
		absD := abs(d)
		if absD < 1 || absD > 3 {
			return false
		}
		if sgn(d) != sign {
			return false
		}
	}
	return true
}

func sgn(n int) interface{} {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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

func countSafeReports(reports []Report) int {
	result := 0
	for _, report := range reports {
		if report.isSafe() {
			result++
		}
	}
	return result
}

func parse(s string) []Report {
	var result []Report
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		report := parseReport(line)
		result = append(result, report)
	}
	return result
}

func parseReport(line string) Report {
	report := Report{}
	for _, token := range strings.Split(line, " ") {
		if token == "" {
			continue
		}
		n, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalln("Cannot parse", token)
		}
		report = append(report, n)
	}
	return report
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
