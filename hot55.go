package hot

func subsets(nums []int) (res [][]int) {
	var dfs func(int, []int)
	dfs = func(i int, cur []int) {
		if i == len(nums) {
			cp := make([]int, len(cur))
			copy(cp, cur)
			res = append(res, cp)
			return
		}
		cur = append(cur, nums[i])
		dfs(i+1, cur)
		cur = cur[:len(cur)-1]
		dfs(i+1, cur)
	}
	dfs(0, []int{})
	return	
}
