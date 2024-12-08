package day3

import (
	"github.com/stretchr/testify/assert"
	"github.com/xpmatteo/aoc-2024/day1"
	"testing"
)

const sample = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const sample2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func Test_filterMul(t *testing.T) {
	expected := []mul{
		{2, 4}, {5, 5}, {11, 8}, {8, 5},
	}
	assert.Equal(t, expected, filterMul(sample))
}

func Test_day3_part1_sample(t *testing.T) {
	assert.Equal(t, 161, sumMul(filterMul(sample)))
}

func Test_day3_part1(t *testing.T) {
	assert.Equal(t, 159833790, sumMul(filterMul(day1.ReadFile("day3.txt"))))
}

func Test_day3_part2_sample(t *testing.T) {
	assert.Equal(t, 48, sumMul(filterMul(skipExcluded(sample2))))
}

func Test_day3_part2(t *testing.T) {
	assert.Equal(t, 89349241, sumMul(filterMul(skipExcluded(day1.ReadFile("day3.txt")))))
}

func Test_skipExcluded(t *testing.T) {
	assert.Equal(t, "abcd", skipExcluded("abcd"))

	expected := "AC" //                     01234567890123
	assert.Equal(t, expected, skipExcluded("Adon't()Bdo()C"))

	expected2 := "pippo--pluto--paperino"
	assert.Equal(t, expected2, skipExcluded("pippo-don't()-skip-do()-pluto-don't()-skip-do()-paperino"))
}
