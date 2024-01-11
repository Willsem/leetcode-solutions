import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	ans := 0

	var maxDiff func(node *TreeNode, min, max int)
	maxDiff = func(node *TreeNode, min, max int) {
		if node == nil {
			return
		}

		ans = int(math.Max(
			float64(ans),
			math.Max(
				math.Abs(float64(min-node.Val)),
				math.Abs(float64(max-node.Val)),
			),
		))

		min = int(math.Min(float64(node.Val), float64(min)))
		max = int(math.Max(float64(node.Val), float64(max)))

		maxDiff(node.Left, min, max)
		maxDiff(node.Right, min, max)
	}

	maxDiff(root, root.Val, root.Val)
	return ans
}
