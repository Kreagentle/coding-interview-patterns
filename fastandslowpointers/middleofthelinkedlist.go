package fastandslowpointers

func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
