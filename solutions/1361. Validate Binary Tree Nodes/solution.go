func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root, hasRoot := findRoot(n, leftChild, rightChild)
	if !hasRoot {
		return false
	}

	visited := make([]bool, n)

	var dfs func(i int) bool
	dfs = func(node int) bool {
		if visited[node] {
			return false
		}

		visited[node] = true

		left, right := leftChild[node], rightChild[node]
		if left != -1 && !dfs(left) {
			return false
		}
		if right != -1 && !dfs(right) {
			return false
		}

		return true
	}

	if !dfs(root) {
		return false
	}

	for _, visit := range visited {
		if !visit {
			return false
		}
	}

	return true
}

func findRoot(n int, leftChild []int, rightChild []int) (int, bool) {
	isChild := make([]bool, n)

	for _, child := range leftChild {
		if child >= 0 {
			isChild[child] = true
		}
	}

	for _, child := range rightChild {
		if child >= 0 {
			isChild[child] = true
		}
	}

	root := -1
	rootCount := 0
	for i := range isChild {
		if !isChild[i] {
			root = i
			rootCount++
		}
	}

	if rootCount != 1 {
		return -1, false
	}

	return root, true
}
