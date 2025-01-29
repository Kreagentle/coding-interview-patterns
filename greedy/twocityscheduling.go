package greedy

import (
	"sort"
)

type Sorting [][]int

func (a Sorting) Len() int           { return len(a) }
func (a Sorting) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sorting) Less(i, j int) bool { return a[i][1]-a[i][0] > a[j][1]-a[j][0] }

func twoCityScheduling(costs [][]int) int {
	totalCost := 0
	sort.Sort(Sorting(costs))

	for i := 0; i < len(costs)/2; i++ {
		totalCost += costs[i][0]
	}
	for i := len(costs) / 2; i < len(costs); i++ {
		totalCost += costs[i][1]
	}
	return totalCost
}
