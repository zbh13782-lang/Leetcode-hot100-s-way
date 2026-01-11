package hot

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
}
