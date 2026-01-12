package hot

func pathSum(root *TreeNode, targetSum int) (ans int) {
	var dfs func(node *TreeNode, s int)
	submap := map[int]int{}
	submap[0] = 1
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s += node.Val

		ans = ans + submap[s-targetSum]
		submap[s]++
		dfs(node.Left, s)
		dfs(node.Right, s)
		submap[s]--
	}
	dfs(root, 0)
	return ans
}
