package day19

import (
	"strings"
)

func solvePart1(towels []string, patterns []string) int {
	var memo = make(map[string]bool)
	countOk := 0
	for _, p := range patterns {
		if isPossible(memo, towels, p) {
			countOk++
		}
	}
	return countOk
}

func isPossible(memo map[string]bool, towels []string, p string) bool {
	if len(p) == 0 {
		return true
	}
	memoizedResult, ok := memo[p]
	if ok {
		return memoizedResult
	}
	for _, t := range towels {
		if strings.HasPrefix(p, t) && isPossible(memo, towels, p[len(t):]) {
			memo[p] = true
			return true
		}
	}
	memo[p] = false
	return false
}

func solvePart2(towels []string, patterns []string) int {
	var memo = make(map[string]bool)
	countOk := 0
	for _, p := range patterns {
		if isPossible(memo, towels, p) {
			countOk++
		}
	}
	return countOk
}

func parseTowels(input string) []string {
	firstLine := strings.Split(input, "\n")[0]
	return strings.Split(firstLine, ", ")
}

func parsePatterns(input string) []string {
	trimmedInput := strings.TrimRight(input, "\n")
	return strings.Split(trimmedInput, "\n")[2:]
}
