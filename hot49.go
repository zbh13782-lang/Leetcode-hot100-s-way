package hot

import "math"

func maxPathSum(root *TreeNode) int {
	maxpath := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) (listlen int) {
		//base情况
		if node == nil {
			return 0
		}
		//左右链长(一定大于等于0!!!)
		left := dfs(node.Left)
		right := dfs(node.Right)
		//带上这个节点的链长
		listlen = node.Val + max(left, right)
		listlen = max(0, listlen)

		//最大路径和
		//本路和
		thispathsum := left + right + node.Val
		maxpath = max(maxpath, thispathsum)
		return
	}
	dfs(root)
	return maxpath
}
