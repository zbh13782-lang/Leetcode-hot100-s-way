package hot

func search(nums []int, target int) int {
	n := len(nums)
	if target == nums[n-1] {
		return n - 1
	}
	var FindMin func() int
	FindMin = func() int {
		left, right := -1, n
		for right-left > 1 {
			middle := left + (right-left)/2
			if nums[middle] < nums[n-1] {
				right = middle
			} else {
				left = middle
			}
		}
		return right
	}

	minindex := FindMin()
	var bresc func(left, right int) int
	bresc = func(left, right int) int {
		for right-left > 1 {
			middle := left + (right-left)/2
			if nums[middle] >= target {
				right = middle
			} else {
				left = middle
			}
		}
		return right
	}
	if target > nums[n-1] {
		res := bresc(-1, minindex)
		if res >= n || nums[res] != target {
			return -1
		}
		return res
	} else {
		res := bresc(minindex-1, n)
		if res >= n || nums[res] != target {
			return -1
		}
		return res
	}

}
