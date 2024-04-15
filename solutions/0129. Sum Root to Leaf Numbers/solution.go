/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, path int) int
	dfs = func(node *TreeNode, path int) int {
		if node == nil {
			return 0
		}

		path = path*10 + node.Val

		left := dfs(node.Left, path)
		right := dfs(node.Right, path)
		if left == 0 && right == 0 {
			return path
		}

		return left + right
	}

	return dfs(root, 0)
}
