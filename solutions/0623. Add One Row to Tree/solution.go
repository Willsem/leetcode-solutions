/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	var dfs func(node *TreeNode, currDepth int)
	dfs = func(node *TreeNode, currDepth int) {
		if node == nil {
			return
		}

		if currDepth == depth-1 {
			node.Left = &TreeNode{
				Val:  val,
				Left: node.Left,
			}
			node.Right = &TreeNode{
				Val:   val,
				Right: node.Right,
			}
			return
		}

		dfs(node.Left, currDepth+1)
		dfs(node.Right, currDepth+1)
	}

	dfs(root, 1)
	return root
}
