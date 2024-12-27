package fastandslowpointers

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func detectCycle(head *EduLinkedListNode) bool {
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
