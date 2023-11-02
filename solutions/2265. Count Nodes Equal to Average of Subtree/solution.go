/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
	result := 0

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		sumLeft, countLeft := dfs(node.Left)
		sumRight, countRight := dfs(node.Right)

		sum := sumLeft + sumRight + node.Val
		count := countLeft + countRight + 1

		if node.Val == sum/count {
			result++
		}

		return sum, count
	}

	dfs(root)
	return result
}
