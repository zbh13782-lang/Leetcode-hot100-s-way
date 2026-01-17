package hot

import (
	"fmt"
	"math"
	"slices"
)

func solveNQueens(n int) (ans [][]string) {
	cur := make([]string, n)
	s := make([]byte, n)
	for i, _ := range s {
		s[i] = '.'
	}
	dp := make([]int, n)
	var check func(int, int) bool
	check = func(i, k int) bool {
		for j := 0; j < k; j++ {
			if dp[j] == i || math.Abs(float64(i-dp[j])) == math.Abs(float64(k-j)) {
				return false
			}
		}
		return true
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			fmt.Println(len(cur))
			ans = append(ans, slices.Clone(cur))
			return
		}
		for j := 0; j < n; j++ {
			if check(j, i) {
				dp[i] = j
				s[j] = 'Q'
				cur[i] = string(s)
				s[j] = '.'
				dfs(i + 1)

				dp[i] = 0

			}
		}
	}
	dfs(0)
	return
}
