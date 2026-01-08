package hot

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var first *ListNode = nil
	second := head
	third := head.Next

	for second != nil {
		third = second.Next
		second.Next = first
		first = second
		second = third

	}
	return first
}
