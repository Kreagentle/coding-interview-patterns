package heap

import "container/heap"

func largestInteger(num int) int {
	if num == 0 {
		return 0
	}

	var oddHeap []int
	var evenHeap []int
	var controller []bool

	for num != 0 {
		number := num % 10
		if number%2 == 1 {
			oddHeap = append(oddHeap, number)
			controller = append([]bool{true}, controller...)
		} else {
			evenHeap = append(evenHeap, number)
			controller = append([]bool{false}, controller...)
		}
		num /= 10
	}

	maxHeapOdd := &MaxHeap{}
	*maxHeapOdd = oddHeap

	maxHeapEven := &MaxHeap{}
	*maxHeapEven = evenHeap

	heap.Init(maxHeapOdd)
	heap.Init(maxHeapEven)

	var result int
	var counter int
	for maxHeapOdd.Len() > 0 || maxHeapEven.Len() > 0 {
		var temp int
		if maxHeapOdd.Len() > 0 && maxHeapEven.Len() > 0 {
			if controller[counter] {
				temp = heap.Pop(maxHeapOdd).(int)
			} else {
				temp = heap.Pop(maxHeapEven).(int)
			}
		} else {
			if maxHeapOdd.Len() > 0 {
				temp = heap.Pop(maxHeapOdd).(int)
			}
			if maxHeapEven.Len() > 0 {
				temp = heap.Pop(maxHeapEven).(int)
			}
		}
		result = result*10 + temp
		counter++
	}

	return result
}

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(*h)-1]
	*h = old[:len(*h)-1]
	return x
}
