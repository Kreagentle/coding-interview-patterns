package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	var prev int
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= result[prev][1] {
			result[prev][1] = max(intervals[i][1], result[prev][1])
		} else {
			result = append(result, intervals[i])
			prev += 1
		}
	}
	return result
}
