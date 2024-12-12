package day4

import (
	"regexp"
	"strings"
)

func SearchXmas(input string) int {
	return search(input) + search(rotate(input)) + search(diag1(input)) + search(diag1(flipHor(input)))
}

func search(input string) int {
	fwd := regexp.MustCompile("XMAS")
	bwd := regexp.MustCompile("SAMX")
	allFwd := fwd.FindAllString(input, -1)
	allBwd := bwd.FindAllString(input, -1)
	return len(allFwd) + len(allBwd)
}

func makeMatrix(rows, cols int) [][]rune {
	a := make([][]rune, rows)
	for i := range a {
		a[i] = make([]rune, cols)
	}
	return a
}

func rotate(input string) string {
	lines := strings.Split(input, "\n")
	out := makeMatrix(len(lines[0]), len(lines))
	for row, line := range lines {
		for col, r := range []rune(line) {
			out[col][row] = r
		}
	}
	result := ""
	for _, runes := range out {
		for _, r := range runes {
			result += string(r)
		}
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}

func diag1(input string) string {
	m, nr, nc := toMatrix(input)
	result := ""
	maxDiag := nr + nc - 1
	for diag := range maxDiag {
		col, row := diag, 0
		for row < nr {
			if col >= 0 && row >= 0 && row < nr && col < nc {
				result += string(m[row][col])
			}
			col--
			row++
		}
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}

func flipHor(input string) string {
	m, _, nc := toMatrix(input)
	result := ""
	for _, row := range m {
		for i, _ := range row {
			result += string(row[nc-i-1])
		}
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}

func toMatrix(input string) (m [][]rune, nr int, nc int) {
	lines := strings.Split(input, "\n")
	nr = len(lines)
	nc = len(lines[0])
	m = makeMatrix(nr, nc)
	for row, line := range lines {
		for col, r := range []rune(line) {
			m[row][col] = r
		}
	}
	return m, len(m), len(m[0])
}
