package day1

import (
	"fmt"
	"time"
)

func Stopwatch() func(string) {
	start := time.Now()
	lap := start
	format := "%20s: %14v %14v\n"
	fmt.Printf(format, "", "elapsed", "lap")
	return func(s string) {
		fmt.Printf(format, s, time.Since(start), time.Since(lap))
		lap = time.Now()
	}
}
