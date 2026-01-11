package hot

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		curlevel := []int{}
		size:=len(queue)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			curlevel = append(curlevel, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, curlevel)
	}
	return
}
