type nodeType int

const (
	noType nodeType = iota
	safe
	terminal
)

func eventualSafeNodes(graph [][]int) []int {
	typeOfNode := make([]nodeType, len(graph))

	for i := range graph {
		if len(graph[i]) == 0 {
			typeOfNode[i] = terminal
		}
	}

	visited := make([]bool, len(graph))
	var dfs func(curr int)
	dfs = func(curr int) {
		visited[curr] = true
		isSafe := true
		for _, i := range graph[curr] {
			if !visited[i] {
				dfs(i)
			}
			if typeOfNode[i] == noType {
				isSafe = false
			}
		}
		if typeOfNode[curr] != terminal && isSafe {
			typeOfNode[curr] = safe
		}
	}
	for i := range graph {
		if !visited[i] {
			dfs(i)
		}
	}

	result := make([]int, 0)
	for i, node := range typeOfNode {
		switch node {
		case safe:
			fallthrough
		case terminal:
			result = append(result, i)
		}
	}

	return result
}
