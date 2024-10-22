import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	sums := make(map[int]int)

	var calcSums func(node *TreeNode, depth int)
	calcSums = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		sums[depth] += node.Val
		calcSums(node.Left, depth+1)
		calcSums(node.Right, depth+1)
	}
	calcSums(root, 0)

	sumVals := make([]int, 0, len(sums))
	for _, v := range sums {
		sumVals = append(sumVals, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sumVals)))
	if k-1 >= len(sumVals) {
		return -1
	}

	return int64(sumVals[k-1])
}
