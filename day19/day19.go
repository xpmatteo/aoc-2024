package day19

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
	for _, t := range towels {
		if p == t {
			return true
		}
	}
	return false
}
