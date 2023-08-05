/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	memory := make(map[int]map[int][]*TreeNode)

	var solve func(start, end int) []*TreeNode
	solve = func(start, end int) []*TreeNode {
		result := make([]*TreeNode, 0)

		if start > end {
			return []*TreeNode{nil}
		}

		if _, ok := memory[start]; ok {
			if v, ok := memory[start][end]; ok {
				return v
			}
		}

		for i := start; i <= end; i++ {
			left := solve(start, i-1)
			right := solve(i+1, end)

			for _, l := range left {
				for _, r := range right {
					result = append(result, &TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					})
				}
			}
		}

		if _, ok := memory[start]; !ok {
			memory[start] = make(map[int][]*TreeNode, 1)
		}

		memory[start][end] = result

		return result
	}

	return solve(1, n)
}
