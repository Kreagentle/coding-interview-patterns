package greedy

import (
	"sort"
)

func findContentChildren(greedFactors []int, cookieSizes []int) int {
	sort.Ints(greedFactors)
	sort.Ints(cookieSizes)
	greeds, cookies, result := 0, 0, 0
	for greeds != len(greedFactors) && cookies != len(cookieSizes) {
		if greedFactors[greeds] <= cookieSizes[cookies] {
			greeds += 1
			cookies += 1
			result += 1
		} else {
			cookies += 1
		}
	}
	return result
}
