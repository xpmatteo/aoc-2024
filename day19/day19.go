package day19

import (
	"regexp"
	"strings"
)

func solvePart1(towels []string, patterns []string) int {
	re, err := regexp.Compile(toRegexp(towels))
	if err != nil {
		panic(err)
	}
	countOk := 0
	for _, p := range patterns {
		if re.MatchString(p) {
			countOk++
		}
	}
	return countOk
}

func toRegexp(ss []string) string {
	return "^(" + strings.Join(ss, "|") + ")+$"
}

func parseTowels(input string) []string {
	firstLine := strings.Split(input, "\n")[0]
	return strings.Split(firstLine, ", ")
}

func parsePatterns(input string) []string {
	trimmedInput := strings.TrimRight(input, "\n")
	return strings.Split(trimmedInput, "\n")[2:]
}
