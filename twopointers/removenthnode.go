package twopointers

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {
	dummy := &EduLinkedListNode{next: head}
	left, right := dummy, head
	temp := n
	for temp > 0 {
		right = right.next
		temp -= 1
	}
	for right != nil {
		left = left.next
		right = right.next
	}
	left.next = left.next.next
	return dummy.next
}
