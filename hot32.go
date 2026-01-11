package hot

func middlenode(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	fast := list.Next
	slow := list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoLists1(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := middlenode(head)
	l2 := m.Next
	m.Next = nil
	l1 := sortList(head)
	l2 = sortList(l2)
	return mergeTwoLists1(l1, l2)
}
