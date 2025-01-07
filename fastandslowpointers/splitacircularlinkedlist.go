package fastandslowpointers

func SplitCircularLinkedList(head *LinkedListNode) []*LinkedListNode {
	slow, fast := head, head
	for fast.next != head && fast.next.next != head {
		slow = slow.next
		fast = fast.next.next
	}
	result1, result2 := head, slow.next
	if fast.next == head {
		fast.next = result2
	} else if fast.next.next == head {
		fast.next.next = result2
	}
	slow.next = result1
	return []*LinkedListNode{result1, result2}
}

type LinkedListNode struct {
	data int
	next *LinkedListNode
}
