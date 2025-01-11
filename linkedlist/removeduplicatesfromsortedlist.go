package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return result
}
