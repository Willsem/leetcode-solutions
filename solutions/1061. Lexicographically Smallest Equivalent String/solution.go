import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	graph := make(map[byte]map[byte]struct{})
	for i := range s1 {
		if _, ok := graph[s1[i]]; !ok {
			graph[s1[i]] = make(map[byte]struct{})
		}
		if _, ok := graph[s2[i]]; !ok {
			graph[s2[i]] = make(map[byte]struct{})
		}

		graph[s1[i]][s2[i]] = struct{}{}
		graph[s2[i]][s1[i]] = struct{}{}
	}

	var dfs func(char byte, visited map[byte]struct{}) byte
	dfs = func(char byte, visited map[byte]struct{}) byte {
		if visited == nil {
			visited = make(map[byte]struct{})
		}
		if _, ok := visited[char]; ok {
			return 'z' + 1
		}

		visited[char] = struct{}{}

		minEq := byte('z') + 1
		for next := range graph[char] {
			minEq = minB(minEq, next)
			minEq = minB(minEq, dfs(next, visited))
		}
		return minEq
	}

	mapToSmallest := make(map[byte]byte)
	for i := byte('a'); i <= 'z'; i++ {
		mapToSmallest[i] = minB(i, dfs(i, nil))
	}

	result := &strings.Builder{}
	for i := range baseStr {
		result.WriteByte(mapToSmallest[baseStr[i]])
	}
	return result.String()
}

func minB(a, b byte) byte {
	return byte(min(int(a), int(b)))
}
