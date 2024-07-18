/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const MAX_DISTANCE = 10

func countPairs(root *TreeNode, distance int) int {
	count := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return make([]int, MAX_DISTANCE)
		}

		if node.Left == nil && node.Right == nil {
			res := make([]int, MAX_DISTANCE)
			res[0] = 1
			return res
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		for i := 0; i < distance; i++ {
			for j := 0; j < distance-i-1; j++ {
				count += left[i] * right[j]
			}
		}

		res := make([]int, MAX_DISTANCE)
		for i := 0; i < MAX_DISTANCE-1; i++ {
			res[i+1] = left[i] + right[i]
		}
		return res
	}

	dfs(root)
	return count
}
