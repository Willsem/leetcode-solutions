func maxNumEdgesToRemove(n int, edges [][]int) int {
	aliceG := NewGraph(AliceType, n, edges)
	bobG := NewGraph(BobType, n, edges)

	aliceCompopentsCount := aliceG.GetComponentsCount()
	bobCompopentsCount := bobG.GetComponentsCount()

	if aliceCompopentsCount != 1 || bobCompopentsCount != 1 {
		return -1
	}

	commonG := NewGraph(CommonType, n, edges)
	commonCompopentsCount := commonG.GetComponentsCount()

	goodCommodEdges := n - commonCompopentsCount
	goodAliceEdges := commonCompopentsCount - 1
	goodBobEdges := goodAliceEdges // (commonCompopentsCount - 1)

	sumOfGoodEdges := goodCommodEdges + goodAliceEdges + goodBobEdges
	return len(edges) - sumOfGoodEdges
}

type GraphType int

const (
	AliceType GraphType = iota
	BobType
	CommonType
)

type Graph struct {
	n int

	g         map[int][]int
	graphType GraphType

	used []bool
}

func NewGraph(graphType GraphType, n int, edges [][]int) *Graph {
	newG := &Graph{
		n: n,

		g:         make(map[int][]int),
		graphType: graphType,

		used: make([]bool, n+1),
	}

	for _, edge := range edges {
		newG.addEdge(edge)
	}

	return newG
}

func (g *Graph) addEdge(edge []int) {
	edgeType, fromV, toV := edge[0], edge[1], edge[2]

	switch g.graphType {
	case AliceType:
		if edgeType == 1 || edgeType == 3 {
			g.g[fromV] = append(g.g[fromV], toV)
			g.g[toV] = append(g.g[toV], fromV)
		}
	case BobType:
		if edgeType == 2 || edgeType == 3 {
			g.g[fromV] = append(g.g[fromV], toV)
			g.g[toV] = append(g.g[toV], fromV)
		}
	case CommonType:
		if edgeType == 3 {
			g.g[fromV] = append(g.g[fromV], toV)
			g.g[toV] = append(g.g[toV], fromV)
		}
	}
}

func (g *Graph) GetComponentsCount() int {
	componentsCount := 0
	for startV := 1; startV <= g.n; startV++ {
		if !g.used[startV] {
			g.dfs(startV)
			componentsCount++
		}
	}

	return componentsCount
}

func (g *Graph) dfs(v int) {
	if g.used[v] {
		return
	}
	g.used[v] = true
	for _, nextV := range g.g[v] {
		g.dfs(nextV)
	}
}
