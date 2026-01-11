package hot

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	newlist := &Node{}

	p := head
	m := map[*Node]*Node{}

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}
	p = head
	newlist.Next = m[p]
	for p != nil {
		m[p].Next = m[p.Next]
		m[p].Random = m[p.Random]
		p = p.Next
	}

	return newlist.Next
}

func copyRandomList(head *Node) *Node {
	p := head
	for p != nil {
		p.Next = &Node{Val: p.Val, Next: p.Next}
		p = p.Next.Next
	}

	for p = head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	dummy := &Node{}
	dummy.Next = head.Next
	for p = head; p.Next.Next != nil; p = p.Next {
		clone := p.Next
		p.Next = clone.Next
		clone.Next = clone.Next.Next
	}
	p.Next = nil
	return dummy.Next
}
