package hot

func combinationSum(candidates []int, target int) (ans [][]int) {
	var dfs func(i, sum int, cur []int)
	dfs = func(i int, sum int, cur []int) {
		if sum == target {
			cp := make([]int, len(cur))
			copy(cp, cur)
			ans = append(ans, cp)
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			sum += candidates[j]

			dfs(j, sum, cur)
			cur = cur[:len(cur)-1]
			sum = sum - candidates[j]

		}
	}
	dfs(0, 0, []int{})
	return
}
