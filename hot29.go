package hot

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	node0 := dummy
	node1 := dummy.Next
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		node0 = node1
		node1 = node3

	}
	return dummy.Next
}
