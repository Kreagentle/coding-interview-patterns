package linkedlist

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	slow, fast, list, prev := head, head, head, new(ListNode)
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	middle := reverseList(slow)
	prev.Next = nil
	var pointer int
	for list != nil && middle != nil {
		if pointer%2 == 0 {
			temp := list.Next
			list.Next = middle
			list = temp
		} else {
			temp := middle.Next
			middle.Next = list
			middle = temp
		}
		pointer += 1
	}
}

func reverseList(curr *ListNode) *ListNode {
	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
