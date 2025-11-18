package day19

import (
	"slices"
	"strings"
)

func solvePart1(towels []string, patterns []string) int {
	countOk := 0
	for _, p := range patterns {
		if isPossibleIterative(towels, []string{p}) {
			countOk++
		}
	}
	return countOk
}

func isPossibleIterative(towels []string, patterns []string) bool {
	for {
		if len(patterns) == 0 {
			return false
		}
		var nextGeneration []string
		for _, p := range patterns {
			if slices.Contains(towels, p) {
				return true
			}
			nextGeneration = append(nextGeneration, continuations(towels, p)...)
		}
		patterns = nextGeneration
	}
}

func continuations(towels []string, pattern string) []string {
	var result []string
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			result = append(result, pattern[len(towel):])
		}
	}
	return result
}

func isPossible(towels []string, p string) bool {
	if len(p) == 0 {
		return true
	}
	for _, t := range towels {
		if strings.HasPrefix(p, t) && isPossible(towels, p[len(t):]) {
			return true
		}
	}
	return false
}

func parseTowels(input string) []string {
	firstLine := strings.Split(input, "\n")[0]
	return strings.Split(firstLine, ", ")
}

func parsePatterns(input string) []string {
	trimmedInput := strings.TrimRight(input, "\n")
	return strings.Split(trimmedInput, "\n")[2:]
}
