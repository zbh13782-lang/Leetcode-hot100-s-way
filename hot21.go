package hot

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	plen, qlen := 0, 0
	for p != nil {
		plen++
		p = p.Next
	}
	for q != nil {
		qlen++
		q = q.Next
	}
	var cnt int
	p, q = headA, headB
	if plen > qlen {
		cnt = plen - qlen
		for i := 0; i < cnt; i++ {
			p = p.Next
		}
	} else if qlen > plen {
		cnt = qlen - plen
		for i := 0; i < cnt; i++ {
			q=q.Next
		}
	}
	for p != q {
		p = p.Next
		q = q.Next
		if p == nil || q == nil {
			return nil
		}
	}
	return p
}
