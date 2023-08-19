import "sort"

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// Save initial indexes before sorting
	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	// Sort by weight for using it in Kruskal algo
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	// Reusable DSU struct
	dsu := newDSU(n)

	// Find the weight of MST using all edges
	mstWeight := findMstWeight(n, edges, nil, &dsu)
	critical, pseudoCritical := []int{}, []int{}

	// Iterate through all edges
	for _, edge := range edges {
		// Force to use this edge
		if findMstWeight(n, edges, edge, &dsu) > mstWeight {
			// This edge is not part of any mst
			continue
		}

		// Make this edge weight 0, so the algo skips it
		edgeWeight := edge[2]
		edge[2] = 0

		if findMstWeight(n, edges, nil, &dsu) == mstWeight {
			// A part of some msts, but not all, because found
			// another path without this edge but the same weight
			pseudoCritical = append(pseudoCritical, edge[3])
		} else {
			// It's greater than mstWeight or equal to -1:
			// A part of all possible msts, otherwise the algo
			// would have found another path with the same weight
			critical = append(critical, edge[3])
		}

		// Revert the weight change
		edge[2] = edgeWeight
	}

	return [][]int{critical, pseudoCritical}
}

// ////////////////
// Kruskal algo //
// ////////////////
func findMstWeight(n int, edges [][]int, forcedEdge []int, dsu *DSU) int {
	// Reuse the same dsu instead of creating a new one (for memory optimisation)
	dsu.Reset()
	mstWeight := 0

	if forcedEdge != nil {
		dsu.Union(forcedEdge[0], forcedEdge[1])
		mstWeight += forcedEdge[2]
	}

	for ei := 0; ei < len(edges) && dsu.isles > 1; ei++ {
		u, v, weight := edges[ei][0], edges[ei][1], edges[ei][2]
		if weight == 0 {
			continue
		} // Skip with weight 0
		if dsu.Connected(u, v) {
			continue
		} // Sets were already connected, avoid cycles
		dsu.Union(u, v)
		mstWeight += weight
	}
	return mstWeight
}

type DSU struct {
	parent, rank []int
	n, isles     int
}

func newDSU(n int) DSU {
	dsu := DSU{make([]int, n), make([]int, n), n, n}
	dsu.Reset()
	return dsu
}
func (dsu *DSU) Reset() {
	dsu.isles = dsu.n
	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.rank[i] = 0
	}
}
func (dsu *DSU) Find(i int) int {
	if i != dsu.parent[i] {
		dsu.parent[i] = dsu.Find(dsu.parent[i])
	}
	return dsu.parent[i]
}
func (dsu *DSU) Connected(i, j int) bool {
	return dsu.Find(i) == dsu.Find(j)
}
func (dsu *DSU) Union(i, j int) {
	ip, jp := dsu.Find(i), dsu.Find(j)
	if ip == jp {
		return
	}
	if dsu.rank[ip] < dsu.rank[jp] {
		dsu.parent[ip] = jp
	} else if dsu.rank[jp] < dsu.rank[ip] {
		dsu.parent[jp] = ip
	} else {
		dsu.parent[jp] = ip
		dsu.rank[ip]++
	}
	dsu.isles--
}
