package hot

func searchInsert(nums []int, target int) (ans int) {
	left, right := -1, len(nums)

	for right-left > 1 {
		middle := left + (right-left)/2
		if isGreen(nums[middle], target) {
			right = middle
		} else {
			left = middle
		}
	}
	return right
}

