package day11

import (
	"github.com/xpmatteo/aoc-2024/day1"
	"strconv"
	"strings"
)

type Stone int
type StoneList map[Stone]int

func (sl StoneList) Add(stone Stone, count int) StoneList {
	n, ok := sl[stone]
	if ok {
		sl[stone] = n + count
	} else {
		sl[stone] = count
	}
	return sl
}

func (sl StoneList) Size() int {
	output := 0
	for _, count := range sl {
		output += count
	}
	return output
}

func blink(stones StoneList, steps int) StoneList {
	for range steps {
		stones = blinkOnce(stones)
	}
	return stones
}

func blinkOnce(sl StoneList) StoneList {
	var output = make(StoneList)
	for stone, count := range sl {
		blinkOneStoneOnce(output, stone, count)
	}
	return output
}

func blinkOneStoneOnce(sl StoneList, sto Stone, count int) {
	if sto == 0 {
		sl.Add(1, count)
	} else if sto.HasEvenDigits() {
		sl.Add(sto.LeftHalf(), count)
		sl.Add(sto.RightHalf(), count)
	} else {
		sl.Add(sto*2024, count)
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
	return stoneFrom(str[:len(str)/2])
}

func (s Stone) RightHalf() Stone {
	str := s.String()
	return stoneFrom(str[len(str)/2:])
}

func parseStones(input string) StoneList {
	result := make(StoneList)
	split := strings.Split(input, " ")
	for _, str := range split {
		result.Add(stoneFrom(str), 1)
	}
	return result
}

func stoneFrom(str string) Stone {
	return Stone(day1.Atoi(str))
}
