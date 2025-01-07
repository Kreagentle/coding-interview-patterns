package fastandslowpointers

type LinkedList[T any] struct {
	Data T
	Next *LinkedList[T]
}

func twinSum(head *LinkedList[int]) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *LinkedList[int]
	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}
	var maxsum int
	curr := head
	for prev != nil {
		res := curr.Data + prev.Data
		if res > maxsum {
			maxsum = res
		}
		prev = prev.Next
		curr = curr.Next
	}
	return maxsum
}
