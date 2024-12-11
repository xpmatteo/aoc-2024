package day4

import (
	"regexp"
	"strings"
)

func SearchXmas(input string) int {
	return search(input) + search(rotate(input))
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
	for diag := range nr + nc - 1 {
		for col := 0; col < nc; col++ {
			for row := diag; row >= 0; row-- {
				result += string(m[row][col])
			}
		}
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}

func toMatrix(input string) (matrix [][]rune, nr int, nc int) {
	lines := strings.Split(input, "\n")
	in := makeMatrix(len(lines), len(lines[0]))
	for row, line := range lines {
		for col, r := range []rune(line) {
			in[row][col] = r
		}
	}
	return in, len(in), len(in[0])
}
