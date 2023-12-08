import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	res := ""

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		res += strconv.Itoa(node.Val)

		if node.Left != nil {
			res += "("
			dfs(node.Left)
			res += ")"
		}

		if node.Right != nil {
			if node.Left == nil {
				res += "()"
			}

			res += "("
			dfs(node.Right)
			res += ")"
		}
	}

	dfs(root)
	return res
}
