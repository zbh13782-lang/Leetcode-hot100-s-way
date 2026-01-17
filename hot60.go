package hot

import "slices"

// 示例 1：

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：

// 输入：s = "a"
// 输出：[["a"]]

func partition(s string) (ans [][]string) {
	
	
	var check func([]byte) bool
	check = func(b []byte) bool {
		t := slices.Clone(b)
		slices.Reverse(b)
		if string(t) == string(b) {
			return true
		}
		return false
	}

	cur := []string{}
	var dfs func(s string)
	dfs = func(s string) {
		n := len(s)
		if n == 0 {
			ans = append(ans, slices.Clone(cur))
			return
		}
		for i := 1; i <= n; i++ {
			substr := s[:i]

			if check([]byte(substr)) {
				cur = append(cur, substr)
				dfs(s[i:])
				cur = cur[:len(cur)-1]
			}
		}

	}
	dfs(s)
	return
}