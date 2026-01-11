package hot

func mergeTwoLists3(list1, list2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {

	m := len(lists)

	for step := 1; step < m; step = step * 2 {
		for i := 0; i < m-step; i = i + step*2 {
			lists[i] = mergeTwoLists3(lists[i], lists[i+step])
		}
	}
	return lists[0]
}
