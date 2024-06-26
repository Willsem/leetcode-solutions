import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	vineHead := &TreeNode{Val: 0}
	vineHead.Right = root
	current := vineHead
	for current.Right != nil {
		if current.Right.Left != nil {
			rightRotate(current, current.Right)
		} else {
			current = current.Right
		}
	}

	nodeCount := 0
	current = vineHead.Right
	for current != nil {
		nodeCount++
		current = current.Right
	}

	m := int(math.Pow(2, math.Floor(math.Log2(float64(nodeCount+1))))) - 1
	makeRotations(vineHead, nodeCount-m)
	for m > 1 {
		m /= 2
		makeRotations(vineHead, m)
	}

	balancedRoot := vineHead.Right
	vineHead = nil
	return balancedRoot
}

func rightRotate(parent, node *TreeNode) {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	parent.Right = tmp
}

func leftRotate(parent, node *TreeNode) {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	parent.Right = tmp
}

func makeRotations(vineHead *TreeNode, count int) {
	current := vineHead
	for i := 0; i < count; i++ {
		tmp := current.Right
		leftRotate(current, tmp)
		current = current.Right
	}
}
