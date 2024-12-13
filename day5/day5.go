package day5

import (
	"strconv"
	"strings"
)

type Page int

type Rule struct {
	ante, post Page
}

func parseUpdates(updates string) []Update {
	var result []Update
	for _, r := range strings.Split(updates, "\n") {
		if len(r) == 0 {
			continue
		}
		var update Update
		for _, s := range strings.Split(r, ",") {
			update = append(update, toPage(s))
		}
		result = append(result, update)
	}
	return result
}

func parseRules(rules string) []Rule {
	var result []Rule
	for _, r := range strings.Split(rules, "\n") {
		if len(r) == 0 {
			continue
		}
		split := strings.Split(r, "|")
		rule := Rule{toPage(split[0]), toPage(split[1])}
		result = append(result, rule)
	}
	return result
}

func toPage(s string) Page {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return Page(result)
}

func sumValidUpdates(rules []Rule, updates []Update) int {
	result := 0
	for _, update := range updates {
		if update.ObeysAll(rules) {
			result += int(update.MiddleValue())
		}
	}
	return result
}

func fixAndSumInvalidUpdates(rules []Rule, updates []Update) int {
	result := 0
	for _, update := range updates {
		if !update.ObeysAll(rules) {
			update.Sort(rules)
			result += int(update.MiddleValue())
		}
	}
	return result
}
