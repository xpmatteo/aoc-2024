package day1

import (
	"log"
	"os"
)

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
