package hot

func maxDepth(root *TreeNode) int {
	var dfs func(node *TreeNode, dep int)
	res := 0
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		dep++
		res = max(res, dep)
		dfs(node.Left, dep)
		dfs(node.Right, dep)
	}
	return res
}
