/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeQueries(root *TreeNode, queries []int) []int {
	maxHeightAfterRemove := make(map[int]int)
	maxHeight := 0

	var dfs func(node *TreeNode, height int, isLeftFirst bool)
	dfs = func(node *TreeNode, height int, isLeftFirst bool) {
		if node == nil {
			return
		}

		maxHeightAfterRemove[node.Val] = max(maxHeightAfterRemove[node.Val], maxHeight)
		maxHeight = max(maxHeight, height)

		first, second := node.Right, node.Left
		if isLeftFirst {
			first, second = node.Left, node.Right
		}

		dfs(first, height+1, isLeftFirst)
		dfs(second, height+1, isLeftFirst)
	}

	dfs(root, 0, false)
	maxHeight = 0
	dfs(root, 0, true)

	result := make([]int, 0, len(queries))
	for _, query := range queries {
		result = append(result, maxHeightAfterRemove[query])
	}
	return result
}
