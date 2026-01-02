package hot

//1.dp
//时间N,空间N
func trap1(height []int) int {
	length := len(height)
	leftmax := make([]int, length)
	rightmax := make([]int, length)
	res := 0

	leftmax[0] = height[0]
	for i := 1; i < length; i++ {
		leftmax[i] = max(leftmax[i-1], height[i])
	}

	rightmax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightmax[i] = max(rightmax[i+1], height[i])
	}

	for i := 0; i < length; i++ {
		curwater := min(rightmax[i], leftmax[i]) - height[i]
		res += curwater
	}
	return res
}

//2.对撞指针
//时间N,空间1
func trap(height []int) int {
	leftmax, rightmax := 0, 0
	left := 0
	right := len(height) - 1
	res := 0
	curwater := 0
	for left <= right {
		leftmax = max(leftmax, height[left])
		rightmax = max(rightmax, height[right])
		if height[left] < height[right] {
			curwater = leftmax - height[left]
			left++
		} else {
			curwater = rightmax - height[right]
			right--
		}
		res += curwater
	}
	return res
}

