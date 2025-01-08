package linkedlist

func reverse(head *LinkedListNode) *LinkedListNode {
	var prev *LinkedListNode
	curr := head
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	return prev
}

type LinkedListNode struct {
	data int
	next *LinkedListNode
}
