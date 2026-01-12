package hot

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	m := map[*TreeNode]*TreeNode{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			m[node.Left] = node
		}
		if node.Right != nil {
			m[node.Right] = node
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	list1, list2 := p, q
	plen, qlen := 0, 0
	for list1 != root {
		plen++
		list1 = m[list1]
	}
	for list2 != root {
		qlen++
		list2 = m[list2]
	}
	list1, list2 = p, q
	if plen > qlen {
		distance := plen - qlen
		for i := 0; i < distance; i++ {
			list1 = m[list1]
		}
	} else {
		distance := qlen - plen
		for i := 0; i < distance; i++ {
			list2 = m[list2]
		}
	}

	for list1 != list2 {
		list1 = m[list1]
		list2 = m[list2]
	}
	return list1
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
