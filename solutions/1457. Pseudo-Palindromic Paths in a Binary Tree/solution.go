/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	counts := make([]int, 9)
	res := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		counts[node.Val-1]++

		if node.Left == nil && node.Right == nil {
			sum, odd := 0, 0
			for _, count := range counts {
				sum += count
				if count%2 == 1 {
					odd++
				}
			}

			if sum%2 == 0 && odd == 0 {
				res++
			}

			if sum%2 == 1 && odd == 1 {
				res++
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}

		counts[node.Val-1]--
	}

	dfs(root)
	return res
}
