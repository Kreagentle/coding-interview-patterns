package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	result := head
	var prev *ListNode
	leftCounter := left
	for leftCounter > 1 {
		prev = head
		head = head.Next
		leftCounter -= 1
	}
	reversedhead, nodeafterresulthead := reverseLinkedList(head, right-left)
	if prev != nil {
		prev.Next = reversedhead
	}
	head.Next = nodeafterresulthead
	if left == 1 {
		return reversedhead
	}
	return result
}

func reverseLinkedList(head *ListNode, k int) (*ListNode, *ListNode) {
	var prev *ListNode
	for k >= 0 {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
		k -= 1
	}
	return prev, head
}
