package hot

//右->左
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}

//左->右
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
