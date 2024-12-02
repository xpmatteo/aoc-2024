package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"slices"
	"strings"
	"testing"
)

const sample = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func Test_day1_part1_sample(t *testing.T) {
	assert.Equal(t, 11, distance(parse(sample)))
}

func Test_day1_part1(t *testing.T) {
	assert.Equal(t, 2285373, distance(parse(ReadFile("day1.txt"))))
}

func Test_day1_part2_sample(t *testing.T) {
	assert.Equal(t, 31, similarity(parse(sample)))
}

func Test_day1_part2(t *testing.T) {
	assert.Equal(t, 31, similarity(parse(ReadFile("day1.txt"))))
}

func distance(a []int, b []int) int {
	if len(a) != len(b) {
		log.Fatalln("lengths differ", len(a), len(b))
	}
	slices.Sort(a)
	slices.Sort(b)
	result := 0
	for i := range len(a) {
		result += Abs(a[i] - b[i])
	}
	return result
}

func parse(s string) ([]int, []int) {
	a := []int{}
	b := []int{}
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		var n1 int
		var n2 int
		parseLine(line, &n1, &n2)
		a = append(a, n1)
		b = append(b, n2)
	}
	return a, b
}

func parseLine(line string, n1 *int, n2 *int) {
	_, err := fmt.Sscanf(line, "%d %d", n1, n2)
	if err != nil {
		log.Fatalln("nooo", err)
	}
}

func Test_parseLine(t *testing.T) {
	var n1 int
	var n2 int
	parseLine("2   5", &n1, &n2)

	assert.Equal(t, 2, n1)
	assert.Equal(t, 5, n2)
}

func similarity(a []int, b []int) int {
	occurrencesB := make(map[int]int)
	for _, i := range b {
		n, ok := occurrencesB[i]
		if !ok {
			occurrencesB[i] = 0
		}
		occurrencesB[i] = n + 1
	}
	var result int
	for _, i := range a {
		result += i * occurrencesB[i]
	}
	return result
}
