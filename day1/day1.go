package day1

import (
	"log"
	"os"
	"strconv"
)

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ReadFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("cant read file", err)
	}
	return string(bytes)
}
