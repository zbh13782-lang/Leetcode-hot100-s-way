package hot

import "fmt"

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p0 := head

	listlength := 0
	for p0 != nil {
		listlength++
		p0 = p0.Next
	}
	p0 = dummy
	pre := dummy
	cur := head

	cnt := listlength / k
	fmt.Println(cnt)
	for j := 0; j < cnt; j++ {
		for i := 0; i < k; i++ {
			nex := cur.Next
			cur.Next = pre
			pre = cur
			cur = nex
		}

		temp := p0.Next

		p0.Next.Next = cur
		p0.Next = pre

		p0 = temp

	}
	return dummy.Next
}
