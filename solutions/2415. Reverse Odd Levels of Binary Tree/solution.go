/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	levels := make(map[int][]*TreeNode)

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth%2 == 1 {
			levels[depth] = append(levels[depth], node)
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	for _, level := range levels {
		n := len(level)
		for i := 0; i < n/2; i++ {
			level[i].Val, level[n-1-i].Val = level[n-1-i].Val, level[i].Val
		}
	}

	return root
}
