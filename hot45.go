package hot

func flatten2(root *TreeNode) {
	var dfs func(node *TreeNode)
	nodes := []*TreeNode{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
}

func flatten(root *TreeNode) {
	var dfs func(node *TreeNode)
	var head *TreeNode
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		dfs(node.Left)
		node.Right = head
		node.Left = nil
		head = node
	}
	dfs(root)
}
