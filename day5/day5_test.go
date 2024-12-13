package day5

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"testing"
)

const sampleRules = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`

const sampleUpdates = `
75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		rules    []Rule
		updates  []Update
		expected int
	}{
		{
			name:     "simple match",
			rules:    []Rule{{2, 3}},
			updates:  []Update{{1, 2, 3}},
			expected: 2,
		},
		{
			name:     "no match bc order",
			rules:    []Rule{{2, 3}},
			updates:  []Update{{1, 3, 2}},
			expected: 0,
		},
		{
			name:     "ok bc only one present",
			rules:    []Rule{{2, 3}},
			updates:  []Update{{1, 5, 3}},
			expected: 5,
		},
		{
			name:     "break one rule",
			rules:    []Rule{{2, 3}, {2, 5}},
			updates:  []Update{{1, 5, 2, 0, 3}},
			expected: 0,
		},
		{
			name:     "ok 2 updates",
			rules:    []Rule{{2, 3}, {2, 5}},
			updates:  []Update{{1, 2, 3}, {4, 5, 6}},
			expected: 7,
		},
		{
			name:     "sample",
			rules:    parseRules(sampleRules),
			updates:  parseUpdates(sampleUpdates),
			expected: 143,
		},
		{
			name:     "real",
			rules:    parseRules(day1.ReadFile("day5-rules.txt")),
			updates:  parseUpdates(day1.ReadFile("day5-updates.txt")),
			expected: 4462,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, sumValidUpdates(test.rules, test.updates))
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name     string
		rules    []Rule
		updates  []Update
		expected int
	}{
		{
			name:     "updates all correct",
			rules:    []Rule{{2, 3}},
			updates:  []Update{{1, 2, 3}},
			expected: 0,
		},
		{
			name:     "fix one update",
			rules:    []Rule{{2, 3}},
			updates:  []Update{{1, 3, 2}},
			expected: 2,
		},
		{
			name:     "sample",
			rules:    parseRules(sampleRules),
			updates:  parseUpdates(sampleUpdates),
			expected: 123,
		},
		{
			name:     "real",
			rules:    parseRules(day1.ReadFile("day5-rules.txt")),
			updates:  parseUpdates(day1.ReadFile("day5-updates.txt")),
			expected: 4462,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, fixAndSumInvalidUpdates(test.rules, test.updates))
		})
	}
}
