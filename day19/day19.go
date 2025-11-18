package day19

import (
	"log"
	"slices"
	"strings"
)

func solvePart1(towels []string, patterns []string) int {
	countOk := 0
	for _, p := range patterns {
		//if isPossibleIterative(towels, []string{p}) {
		if isPossible(towels, p) {
			countOk++
		}
	}
	return countOk
}

func isPossibleIterative(towels []string, patterns []string) bool {
	iteration := 0
	for {
		log.Printf("Iteration %d Patterns: %d", iteration, len(patterns))
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
		iteration++
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
	log.Printf("Pattern: %s", p)
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
