package hot

func findMin(nums []int) int {
	n := len(nums)
	left, right := -1, n
	var isgreen66 func(num int) bool
	isgreen66 = func(num int) bool {
		return num <= nums[n-1]
	}
	for right-left > 1 {
		middle := left + (right-left)/2
		if isgreen66(nums[middle]) {
			right = middle
		} else {
			left = middle
		}
	}
	return nums[right]
}
