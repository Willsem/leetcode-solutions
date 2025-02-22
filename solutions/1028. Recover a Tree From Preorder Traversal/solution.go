/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
	levels := make([]*TreeNode, 0)

	n := len(traversal)
	for i := 0; i < n; {
		depth := 0
		for i < n && traversal[i] == '-' {
			depth++
			i++
		}

		value := 0
		for i < n && traversal[i] >= '0' && traversal[i] <= '9' {
			value = value*10 + int(traversal[i]-'0')
			i++
		}

		node := &TreeNode{
			Val: value,
		}

		if depth < len(levels) {
			levels[depth] = node
		} else {
			levels = append(levels, node)
		}

		if depth > 0 {
			parent := levels[depth-1]
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}
	}

	return levels[0]
}
