package hot

func permute(nums []int) (ans [][]int) {
	var process func(int)
	var swap func(*int, *int)
	swap = func(i1, i2 *int) {
		*i1, *i2 = *i2, *i1
	}
	process = func(i1 int) {
		if i1 == len(nums) {
			cp := make([]int, i1)
			copy(cp, nums)
			ans = append(ans, cp)
			return
		}
		for i := i1; i < len(nums); i++ {
			swap(&nums[i1], &nums[i])
			process(i + 1)
			swap(&nums[i1], &nums[i])
		}
	}
	process(0)
	return
}
