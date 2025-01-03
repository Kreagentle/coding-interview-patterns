package fastandslowpointers

func SplitCircularLinkedList(head *LinkedListNode) []*LinkedListNode {
  	slow, fast := head, head
  	for fast.next != head && fast.next.next != head {
  	  slow = slow.next
  	  fast = fast.next.next
  	}
  	result1, result2 := head, slow.next
  	slow.next = result1
  	fast = result2
  	for fast.next != head {
  	  fast = fast.next
  	}
  	fast.next = result2
    return []*LinkedListNode{result1,result2}
}

type LinkedListNode struct {
	data int
	next *LinkedListNode
}
