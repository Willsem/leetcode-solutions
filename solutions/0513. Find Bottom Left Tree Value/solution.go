/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	maxLevel := -1
	res := 0

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level > maxLevel {
			maxLevel = level
			res = node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return res
}
