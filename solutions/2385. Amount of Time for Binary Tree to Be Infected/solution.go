/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)

	var dfs func(node *TreeNode, parent *int)
	dfs = func(node *TreeNode, parent *int) {
		if node == nil {
			return
		}

		graph[node.Val] = make([]int, 0, 3)

		if parent != nil {
			graph[node.Val] = append(graph[node.Val], *parent)
		}

		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			dfs(node.Left, &node.Val)
		}

		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			dfs(node.Right, &node.Val)
		}
	}

	dfs(root, nil)

	maxDist := 0
	visited := make(map[int]struct{})

	var calcDist func(node int, dist int)
	calcDist = func(node int, dist int) {
		if _, ok := visited[node]; ok {
			return
		}

		visited[node] = struct{}{}

		if dist > maxDist {
			maxDist = dist
		}

		for _, n := range graph[node] {
			calcDist(n, dist+1)
		}
	}

	calcDist(start, 0)

	return maxDist
}
