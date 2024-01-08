/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= low && node.Val <= high {
			res += node.Val

			dfs(node.Left)
			dfs(node.Right)
		} else if node.Val < low {
			dfs(node.Right)
		} else {
			dfs(node.Left)
		}
	}

	dfs(root)

	return res
}
