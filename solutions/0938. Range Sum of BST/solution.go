/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	switch {
	case root.Val < low:
		return rangeSumBST(root.Right, low, high)

	case root.Val <= high:
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)

	default:
		return rangeSumBST(root.Left, low, high)
	}
}
