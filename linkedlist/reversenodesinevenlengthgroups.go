package linkedlist

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var prev, saver *ListNode
	counter := 0
	curr := head
	for curr != nil {
		if counter%2 != 0 {
			temp := curr
			tempcounter := counter + 1
			for tempcounter > 0 && curr != nil {
				curr = curr.Next
				tempcounter -= 1
			}
			headreversed, lastreversed := reversegroup(temp, counter+1)
			prev.Next = headreversed
			saver = lastreversed
			lastreversed.Next = curr
		} else {
			if curr.Next == nil {
				return head
			}
			temp := curr
			tempcounter := counter
			for tempcounter > 0 && curr != nil {
				curr = curr.Next
				tempcounter -= 1
			}
			if curr == nil {
				headreversed, lastreversed := reversegroup(temp, counter-tempcounter)
				saver.Next = headreversed
				lastreversed.Next = curr
				return head
			}
			prev = curr
			curr = curr.Next
		}
		counter += 1
	}
	return head
}

func reversegroup(current *ListNode, k int) (*ListNode, *ListNode) {
	var prev *ListNode
	last := current
	for k > 0 {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
		k -= 1
	}
	return prev, last
}
