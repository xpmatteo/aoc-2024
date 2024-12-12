package day5

import "slices"

type Update []Page

func (u Update) ObeysAll(rules []Rule) bool {
	for _, rule := range rules {
		if !u.Obeys(rule) {
			return false
		}
	}
	return true
}

func (u Update) Obeys(rule Rule) bool {
	indexAnte := slices.Index(u, rule.ante)
	indexPost := slices.Index(u, rule.post)
	return indexAnte == -1 || indexPost == -1 || indexAnte < indexPost
}

func (u Update) MiddleValue() Page {
	middleIndex := len(u) / 2
	return u[middleIndex]
}
