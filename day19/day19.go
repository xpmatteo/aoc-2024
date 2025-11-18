package day19

import "strings"

func solvePart1(towels []string, patterns []string) int {
	countOk := 0
	for _, p := range patterns {
		if isPossible(towels, p) {
			countOk++
		}
	}
	return countOk
}

func isPossible(towels []string, p string) bool {
	if len(p) == 0 {
		return true
	}
	for _, t := range towels {
		if strings.HasPrefix(p, t) {
			return isPossible(towels, p[len(t):])
		}
	}
	return false
}
