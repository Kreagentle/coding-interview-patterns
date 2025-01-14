package merge_intervals

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var result [][]int
	i := 0
	// add not overlapping intervals before new interval
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i += 1
	}
	// generate and merge overlapping intervals
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i += 1
	}
	result = append(result, newInterval)

	// return the rest intervals
	return append(result, intervals[i:]...)
}
