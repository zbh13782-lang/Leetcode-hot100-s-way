package hot

func searchRange(nums []int, target int) []int {
	left, right := -1, len(nums)

	for right-left > 1 {
		middle := left + (right-left)/2
		if isGreen(nums[middle], target) {
			right = middle
		} else {
			left = middle
		}
	}

	if nums[right] != target {
		return []int{0, 0}
	}
	resleft := left
	right = len(nums)

	for right-left > 1 {
		middle := left + (right-left)/2
		if isGreen64(nums[middle], target) {
			right = middle
		} else {
			left = middle
		}
	}

	resright := left
	return []int{resleft, resright}
}

func isGreen(num, target int) bool {
	return num >= target
}

func isGreen64(num, target int) bool {
	return num > target
}
