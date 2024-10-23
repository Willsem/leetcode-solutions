/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	depthSum := make(map[int]int)

	var calcSum func(node *TreeNode, depth int)
	calcSum = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		depthSum[depth] += node.Val

		calcSum(node.Left, depth+1)
		calcSum(node.Right, depth+1)
	}
	calcSum(root, 0)

	var replaceValues func(node *TreeNode, brothersSum int, depth int)
	replaceValues = func(node *TreeNode, brothersSum int, depth int) {
		if node == nil {
			return
		}

		node.Val = depthSum[depth] - brothersSum

		childBrothersSum := 0
		if node.Left != nil {
			childBrothersSum += node.Left.Val
		}
		if node.Right != nil {
			childBrothersSum += node.Right.Val
		}

		replaceValues(node.Left, childBrothersSum, depth+1)
		replaceValues(node.Right, childBrothersSum, depth+1)
	}

	replaceValues(root, root.Val, 0)
	return root
}
