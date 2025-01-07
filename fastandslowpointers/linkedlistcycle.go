package fastandslowpointers

func detectCycle(head *LinkedListNode) bool {
	slowPointer, fastPointer := head, head
	for fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
