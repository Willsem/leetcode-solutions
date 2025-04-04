/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}

		lnode, ld := dfs(node.Left)
		rnode, rd := dfs(node.Right)

		if ld > rd {
			return lnode, ld + 1
		}

		if ld < rd {
			return rnode, rd + 1
		}

		return node, ld + 1
	}

	node, _ := dfs(root)
	return node
}
