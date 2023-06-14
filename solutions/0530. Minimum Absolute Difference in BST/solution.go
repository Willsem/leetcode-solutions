/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	min                = -1
	prevNode *TreeNode = nil
)

func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	dfs(node.Left)

	if prevNode != nil {
		diff := node.Val - prevNode.Val
		if min == -1 || diff < min {
			min = diff
		}
	}
	prevNode = node

	dfs(node.Right)
}

func getMinimumDifference(root *TreeNode) int {
	min = -1
	prevNode = nil
	dfs(root)
	return min
}
