package hot

//没太理解
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := make(map[int]int, len(inorder))
	for i, v := range inorder {
		index[v] = i
	}
	var dfs func(int, int, int) *TreeNode
	dfs = func(preL, preR, inL int) *TreeNode {
		if preL == preR {
			return nil
		}

		leftsize := index[preorder[preL]] - inL
		lefttree := dfs(preL+1, preL+leftsize+1, inL)
		rightree := dfs(preL+leftsize+1, preR, inL+leftsize+1)
		return &TreeNode{
			preorder[preL], lefttree, rightree,
		}
	}
	return dfs(0, len(preorder), 0)
}
