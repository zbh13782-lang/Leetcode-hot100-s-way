package hot

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(i interface{}) {
	h.IntSlice = append(h.IntSlice, i.(int))
}

func (h *hp) Pop() (i interface{}) {
	a = h.IntSlice
	i = a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return i
}

// 大根堆 nlogn n
func maxSlidingWindow1(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)
	res := []int{}
	res = append(res, nums[q.IntSlice[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		res = append(res, nums[q.IntSlice[0]])
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	var q []int

	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	res := []int{}
	res = append(res, nums[q[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		res = append(res, nums[q[0]])
	}
	return res
}
