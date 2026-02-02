package hot

func largestRectangleArea(heights []int) int {
	stack := []int{}
	n := len(heights)
	curArea := 0
	maxArea := 0
	for i, v := range heights {
		for len(stack) > 0 && v < heights[stack[len(stack)-1]] {
			hg := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			curArea = (i - left - 1) * hg
			maxArea = max(maxArea, curArea)
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		left := -1
		right := n
		hg := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		curArea = (right - left - 1) * hg
		maxArea = max(maxArea, curArea)
	}
	return maxArea
}
