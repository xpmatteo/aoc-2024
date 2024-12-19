package day11

import (
	"github.com/xpmatteo/aoc-2024/day1"
	"strconv"
	"strings"
)
import "github.com/samber/lo"

type Stone int

func blink(stones []Stone, steps int) []Stone {
	for range steps {
		stones = blinkOnce(stones)
	}
	return stones
}

func blinkOnce(input []Stone) []Stone {
	var output []Stone
	for _, sto := range input {
		output = append(output, blinkStone(sto)...)
	}
	return output
}

func blinkStone(sto Stone) []Stone {
	if sto == 0 {
		return []Stone{1}
	} else if sto.HasEvenDigits() {
		return []Stone{sto.LeftHalf(), sto.RightHalf()}
	} else {
		return []Stone{sto * 2024}
	}
}

func (s Stone) HasEvenDigits() bool {
	return len(s.String())%2 == 0
}

func (s Stone) String() string {
	return strconv.Itoa(int(s))
}

func (s Stone) LeftHalf() Stone {
	str := s.String()
	return Stone(day1.Atoi(str[:len(str)/2]))
}

func (s Stone) RightHalf() Stone {
	str := s.String()
	return Stone(day1.Atoi(str[len(str)/2:]))
}

func parseStones(input string) []Stone {
	split := strings.Split(input, " ")
	return lo.Map(split, func(item string, index int) Stone {
		return Stone(day1.Atoi(item))
	})
}
