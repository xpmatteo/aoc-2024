package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

type mul struct {
	x, y int
}

func filterMul(s string) []mul {
	var result []mul
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	findString := re.FindAllStringSubmatch(s, -1)
	for _, match := range findString {
		result = append(result, mul{atoi(match[1]), atoi(match[2])})
	}
	print(findString)
	return result
}

func sumMul(muls []mul) int {
	result := 0
	for _, m := range muls {
		result += m.x * m.y
	}
	return result
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprint("not an integer", err))
	}
	return i
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
