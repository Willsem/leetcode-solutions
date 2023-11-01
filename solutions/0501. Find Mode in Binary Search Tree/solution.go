/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	counts := make(map[int]int)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		counts[node.Val]++

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	maxCount := 0
	res := make([]int, 0)
	for value, count := range counts {
		if count > maxCount {
			maxCount = count
			res = []int{value}
		} else if count == maxCount {
			res = append(res, value)
		}
	}

	return res
}
