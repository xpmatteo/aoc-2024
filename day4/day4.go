package day4

import (
	"regexp"
	"strings"
)

type Matrix [][]rune

func (m *Matrix) numRows() int {
	return len(*m)
}

func (m *Matrix) numCols() int {
	return len((*m)[0])
}

func (m *Matrix) get(r, c int) rune {
	if m.isValidCoord(r, c) {
		return (*m)[r][c]
	}
	return ' '
}

func (m *Matrix) isValidCoord(r int, c int) bool {
	return r >= 0 && c >= 0 && r < m.numRows() && c < m.numCols()
}

func (m *Matrix) forAll(f func(ch rune, r, c int)) {
	for r, row := range *m {
		for c, ch := range row {
			f(ch, r, c)
		}
	}
}

func SearchXmas(input string) int {
	return search(input) + search(rotate(input)) + search(diag1(input)) + search(diag1(flipHor(input)))
}

func SearchCrossMas(input string) int {
	m, _, _ := toMatrix(input)
	result := 0
	m.forAll(func(ch rune, r, c int) {
		if isCrossMas(m, r, c) {
			result++
		}
	})
	return result
}

func isCrossMas(m Matrix, r int, c int) bool {
	nw := m.get(r-1, c-1)
	ne := m.get(r-1, c+1)
	center := m.get(r, c)
	sw := m.get(r+1, c-1)
	se := m.get(r+1, c+1)
	return center == 'A' && isMas(nw, se) && isMas(ne, sw)
}

func isMas(a rune, b rune) bool {
	return a == 'M' && b == 'S' || a == 'S' && b == 'M'
}

func search(input string) int {
	fwd := regexp.MustCompile("XMAS")
	bwd := regexp.MustCompile("SAMX")
	allFwd := fwd.FindAllString(input, -1)
	allBwd := bwd.FindAllString(input, -1)
	return len(allFwd) + len(allBwd)
}

func makeMatrix(rows, cols int) Matrix {
	a := make(Matrix, rows)
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

func toMatrix(input string) (m Matrix, nr int, nc int) {
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
