func numTilePossibilities(tiles string) int {
	result := make(map[string]struct{})
	visited := make([]bool, len(tiles))

	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(path) > 0 {
			result[string(path)] = struct{}{}
		}

		for i := range tiles {
			if !visited[i] {
				visited[i] = true
				path = append(path, tiles[i])

				dfs(path)

				visited[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs([]byte{})
	return len(result)
}
