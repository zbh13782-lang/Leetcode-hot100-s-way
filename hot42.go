package hot

func isValidBST(root *TreeNode) bool {
	var inorder []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		inorder = append(inorder, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	for i := 0; i < len(inorder)-1; i++ {
		if inorder[i+1] < inorder[i] {
			return false
		}
	}
	return true
}
