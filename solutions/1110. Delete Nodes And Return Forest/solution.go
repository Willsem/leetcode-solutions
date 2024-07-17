/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	toDeleteMap := make(map[int]struct{})
	for _, node := range toDelete {
		toDeleteMap[node] = struct{}{}
	}

	result := make([]*TreeNode, 0)
	var dfs func(node *TreeNode, parent *TreeNode, isLeft bool)
	dfs = func(node *TreeNode, parent *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		left := node.Left
		right := node.Right

		if _, ok := toDeleteMap[node.Val]; ok {
			node = nil
			if parent != nil {
				if isLeft {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}
		} else if parent == nil {
			result = append(result, node)
		}

		dfs(left, node, true)
		dfs(right, node, false)
	}

	dfs(root, nil, false)
	return result
}
