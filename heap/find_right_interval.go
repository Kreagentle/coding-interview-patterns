package heap

import "container/heap"

func findRightInterval(intervals [][]int) []int {
	mHeapStartTime := &MinHeapArr{}
	mHeapEndTime := &MinHeapArr{}

	result := make([]int, len(intervals))
	for i := range result {
		result[i] = -1
	}

	for count, interval := range intervals {
		*mHeapStartTime = append(*mHeapStartTime, []int{interval[0], count})
		*mHeapEndTime = append(*mHeapEndTime, []int{interval[1], count})
	}

	heap.Init(mHeapStartTime)
	heap.Init(mHeapEndTime)

	for mHeapEndTime.Len() > 0 && mHeapStartTime.Len() > 0 {
		if (*mHeapStartTime)[0][0] >= (*mHeapEndTime)[0][0] {
			start := (*mHeapStartTime)[0]
			end := (*mHeapEndTime)[0]
			result[end[1]] = start[1]
			heap.Pop(mHeapEndTime)
		} else {
			heap.Pop(mHeapStartTime)
		}
	}

	return result
}

type MinHeapArr [][]int

func (h MinHeapArr) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeapArr) Len() int           { return len(h) }
func (h MinHeapArr) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapArr) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeapArr) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
