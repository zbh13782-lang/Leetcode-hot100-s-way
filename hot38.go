package hot

func isSymmetric(root *TreeNode) bool {
	return ismirror(root.Left, root.Right)
}

func ismirror(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	return L.Val == R.Val && ismirror(L.Left, R.Right) && ismirror(L.Right, R.Left)
}
