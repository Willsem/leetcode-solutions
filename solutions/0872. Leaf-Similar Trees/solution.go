/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := dfs(root1)
	leaf2 := dfs(root2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	return append(dfs(node.Left), dfs(node.Right)...)
}
