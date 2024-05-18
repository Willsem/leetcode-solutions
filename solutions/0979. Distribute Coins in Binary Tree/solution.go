/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	moves := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftCoins := dfs(node.Left)
		rightCoins := dfs(node.Right)
		moves += abs(leftCoins) + abs(rightCoins)
		return node.Val - 1 + leftCoins + rightCoins
	}

	dfs(root)
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
