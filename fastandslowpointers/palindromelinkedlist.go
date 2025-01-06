package fastandslowpointers

func palindrome(head *LinkedListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	var prev *LinkedListNode
	curr := slow
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	for prev != nil {
		if prev.data != head.data {
			return false
		}
		prev = prev.next
		head = head.next
	}
	return true
}
