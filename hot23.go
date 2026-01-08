package hot

func isPalindrome(head *ListNode) bool {
	p := head
	cnt := 0

	for p != nil {
		cnt++
		p = p.Next
	}
	if cnt%2 == 0 {
		p = head
		cnt /= 2
		for i := 0; i < cnt; i++ {
			p = p.Next
		}
	} else {
		p = head
		cnt /= 2
		p = p.Next
		for i := 0; i < cnt; i++ {
			p = p.Next
		}
	}

	var first, second *ListNode = nil, p
	for second != nil {
		third := second.Next
		second.Next = first
		first = second
		second = third
	}
	p, q := head, first
	for p != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}
	return true
}
