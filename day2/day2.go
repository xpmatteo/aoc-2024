package day2

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Report []int

func (r Report) isSafe() bool {
	sign := sgn(r[0] - r[1])
	for i := range len(r) - 1 {
		d := r[i] - r[i+1]
		absD := abs(d)
		if absD < 1 || absD > 3 {
			return false
		}
		if sgn(d) != sign {
			return false
		}
	}
	return true
}

func remove(r Report, i int) Report {
	result := slices.Clone(r)
	return append(result[0:i], result[i+1:]...)
}

func (r Report) Dampen() Report {
	if r.isSafe() {
		return r
	}
	for i := range len(r) {
		dampened := remove(r, i)
		if dampened.isSafe() {
			return dampened
		}
	}
	return r
}

func sgn(n int) interface{} {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func countSafeReports(reports []Report) int {
	result := 0
	for _, report := range reports {
		if report.isSafe() {
			result++
		}
	}
	return result
}

func parse(s string) []Report {
	var result []Report
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		report := parseReport(line)
		result = append(result, report)
	}
	return result
}

func parseReport(line string) Report {
	report := Report{}
	for _, token := range strings.Split(line, " ") {
		if token == "" {
			continue
		}
		n, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalln("Cannot parse", token)
		}
		report = append(report, n)
	}
	return report
}

func dampen(reports []Report) []Report {
	var result []Report
	for _, r := range reports {
		result = append(result, r.Dampen())
	}
	return result
}
