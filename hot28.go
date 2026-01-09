package hot

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	plen := 0
	for p != nil {
		plen++
		p = p.Next
	}
	if plen <= n {
		return head.Next
	}
	precnt := plen - n
	p = head
	pre := &ListNode{}
	for i := 0; i < precnt; i++ {
		pre = p
		p = p.Next
	}
	pre.Next = p.Next
	return head
}
