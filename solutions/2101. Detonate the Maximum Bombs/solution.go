func distance(A []int, B []int) int64 {
	return int64(A[0]-B[0])*int64(A[0]-B[0]) + int64(A[1]-B[1])*int64(A[1]-B[1])
}

func dfs(curr int, adjList [][]int, visited []bool, count *int) {
	if visited[curr] {
		return
	}
	visited[curr] = true
	(*count)++
	for _, next := range adjList[curr] {
		dfs(next, adjList, visited, count)
	}
}

func maximumDetonation(bombs [][]int) int {
	var n int = len(bombs)
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var distSquare int64 = distance(bombs[i], bombs[j])
			if distSquare <= int64(bombs[i][2])*int64(bombs[i][2]) {
				adjList[i] = append(adjList[i], j)
			}
			if distSquare <= int64(bombs[j][2])*int64(bombs[j][2]) {
				adjList[j] = append(adjList[j], i)
			}
		}
	}
	var ans int = 0
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		var count int = 0
		dfs(i, adjList, visited, &count)
		if count > ans {
			ans = count
		}
	}
	return ans
}
