package hot

import "slices"

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

func generateParenthesis(n int) (ans []string) {
	var dfs func(int, int, []byte)

	dfs = func(left, right int, cur []byte) {
		if left == 0 && right == 0 {
			ans = append(ans, string(slices.Clone(cur)))
			return
		}

		if left > 0 {
			cur = append(cur, '(')
			dfs(left-1, right, cur)
			cur = cur[:len(cur)-1]
		}

		if right > left {
			cur = append(cur, ')')
			dfs(left, right-1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(n, n, []byte{})
	return
}
