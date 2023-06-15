/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	sumOfLevel := make([]int, 0)

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(sumOfLevel) == level {
			sumOfLevel = append(sumOfLevel, node.Val)
		} else {
			sumOfLevel[level] += node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	maxIndex := 0
	for i := range sumOfLevel {
		if sumOfLevel[i] > sumOfLevel[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex + 1
}
