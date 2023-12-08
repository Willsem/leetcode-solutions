import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	res := strings.Builder{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		res.WriteString(strconv.Itoa(node.Val))

		if node.Left != nil {
			res.WriteString("(")
			dfs(node.Left)
			res.WriteString(")")
		}

		if node.Right != nil {
			if node.Left == nil {
				res.WriteString("()")
			}

			res.WriteString("(")
			dfs(node.Right)
			res.WriteString(")")
		}
	}

	dfs(root)
	return res.String()
}
