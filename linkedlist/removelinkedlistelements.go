package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	result := head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	if result.Val == val {
		return result.Next
	}
	return result
}
