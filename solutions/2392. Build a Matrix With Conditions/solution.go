import "errors"

type Graph struct {
	Verts    int
	AdjList  map[int][]int
	Roots    []int
	Indegree map[int]int
}

func NewGraph(k int, conditions [][]int) (*Graph, error) {
	g := Graph{
		Verts:    k,
		AdjList:  make(map[int][]int),
		Roots:    make([]int, 0),
		Indegree: make(map[int]int),
	}
	for i := 1; i < k+1; i += 1 {
		g.Indegree[i] = 0
	}
	for _, cond := range conditions {
		u, v := cond[0], cond[1]
		g.AdjList[u] = append(g.AdjList[u], v)
		g.Indegree[v] += 1
	}
	for key, val := range g.Indegree {
		if val == 0 {
			g.Roots = append(g.Roots, key)
		}
	}
	if len(g.Roots) == 0 {
		return nil, errors.New("")
	}
	return &g, nil
}

func (g *Graph) TopologicalSort() (map[int]int, error) {
	order, idx := make(map[int]int), 0
	for len(g.Roots) != 0 {
		top := g.Roots[0]
		g.Roots = g.Roots[1:]
		order[top] = idx
		idx += 1
		for _, adj := range g.AdjList[top] {
			g.Indegree[adj] -= 1
			if g.Indegree[adj] == 0 {
				g.Roots = append(g.Roots, adj)
			}
		}
	}
	if len(order) != g.Verts {
		return nil, errors.New("")
	}
	return order, nil
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowGraph, err := NewGraph(k, rowConditions)
	if err != nil {
		return make([][]int, 0)
	}
	colGraph, err := NewGraph(k, colConditions)
	if err != nil {
		return make([][]int, 0)
	}

	rowTopo, err := rowGraph.TopologicalSort()
	if err != nil {
		return make([][]int, 0)
	}
	colTopo, err := colGraph.TopologicalSort()
	if err != nil {
		return make([][]int, 0)
	}
	rowGraph, colGraph = nil, nil

	ret := make([][]int, k)
	for i := 0; i < k; i += 1 {
		ret[i] = make([]int, k)
	}

	for i := 1; i < k+1; i += 1 {
		ret[rowTopo[i]][colTopo[i]] = i
	}

	return ret
}
