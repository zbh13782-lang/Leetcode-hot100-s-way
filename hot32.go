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
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists1(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists1(list1, list2.Next)
		return list2
	}
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
