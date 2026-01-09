package hot

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	plen, qlen := 0, 0
	l1tail, l2tail := l1, l2
	for p != nil {
		plen++
		if p.Next == nil {
			l1tail = p
		}
		p = p.Next

	}
	for q != nil {
		qlen++
		if q.Next == nil {
			l2tail = q
		}
		q = q.Next

	}
	distance := 0
	if plen > qlen {
		distance = plen - qlen
		for i := 0; i < distance; i++ {
			l2tail.Next = &ListNode{Val: 0}
			l2tail = l2tail.Next
		}
	} else {
		distance = qlen - plen
		for i := 0; i < distance; i++ {

			l1tail.Next = &ListNode{Val: 0}
			l1tail = l1tail.Next
		}
	}
	catch := 0
	p, q = l1, l2
	for p != nil && q != nil {

		s := catch + p.Val + q.Val
		if s >= 10 {
			catch = 1
		}else{
			catch = 0
		}
		p.Val = s % 10
		p, q = p.Next, q.Next
	}
	if catch == 1 {
		l1tail.Next = &ListNode{Val: 1}
	}
	return l1
}
