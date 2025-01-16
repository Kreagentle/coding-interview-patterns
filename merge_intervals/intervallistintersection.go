package merge_intervals

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var itera, iterb int
	var intersection [][]int
	for itera < len(firstList) && iterb < len(secondList) {
		lateststartingtime := max(firstList[itera][0], secondList[iterb][0])
		earliestendingtime := min(firstList[itera][1], secondList[iterb][1])
		if lateststartingtime <= earliestendingtime {
			intersection = append(intersection, []int{lateststartingtime, earliestendingtime})
		}
		if firstList[itera][1] < secondList[iterb][1] {
			itera += 1
		} else {
			iterb += 1
		}
	}
	return intersection
}
