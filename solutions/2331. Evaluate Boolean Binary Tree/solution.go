/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
	False int = iota
	True
	Or
	And
)

func evaluateTree(root *TreeNode) bool {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		switch node.Val {
		case False:
			return false
		case True:
			return true
		case Or:
			return dfs(node.Left) || dfs(node.Right)
		case And:
			return dfs(node.Left) && dfs(node.Right)
		}

		return false
	}

	return dfs(root)
}
