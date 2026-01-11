package hot

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	res := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1

		}
		Llen := dfs(node.Left) + 1
		Rlen := dfs(node.Right) + 1
		res = max(res, Llen+Rlen)
		return max(Llen, Rlen)
	}
	dfs(root)
	return res
}
