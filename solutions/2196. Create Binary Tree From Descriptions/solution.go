/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	hasParent := make(map[int]bool)
	nodes := make(map[int]*TreeNode)

	for _, d := range descriptions {
		parent, child, isLeft := d[0], d[1], d[2]

		hasParent[child] = true
		if _, ok := hasParent[parent]; !ok {
			hasParent[parent] = false
		}

		if _, ok := nodes[child]; !ok {
			nodes[child] = &TreeNode{
				Val: child,
			}
		}

		if _, ok := nodes[parent]; !ok {
			nodes[parent] = &TreeNode{
				Val: parent,
			}
		}

		parentNode := nodes[parent]
		if isLeft == 1 {
			parentNode.Left = nodes[child]
		} else {
			parentNode.Right = nodes[child]
		}
	}

	for nodeNum, nodeHasParent := range hasParent {
		if !nodeHasParent {
			return nodes[nodeNum]
		}
	}

	return nil
}
