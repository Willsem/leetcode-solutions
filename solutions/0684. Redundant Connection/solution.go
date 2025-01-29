type DisjointSet struct {
	parent []int
	size   []int
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	ds := NewDisjointSet(n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		if ds.Find(u) == ds.Find(v) {
			return edge
		}
		ds.Union(u, v)
	}
	return nil
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range ds.parent {
		ds.parent[i] = i
		ds.size[i] = 1
	}
	return ds
}

func (ds *DisjointSet) Find(v int) int {
	if ds.parent[v] != v {
		ds.parent[v] = ds.Find(ds.parent[v])
	}
	return ds.parent[v]
}

func (ds *DisjointSet) Union(a, b int) {
	aRoot, bRoot := ds.Find(a), ds.Find(b)
	if aRoot != bRoot {
		if ds.size[aRoot] < ds.size[bRoot] {
			ds.parent[aRoot] = bRoot
			ds.size[bRoot] += ds.size[aRoot]
		} else {
			ds.parent[bRoot] = aRoot
			ds.size[aRoot] += ds.size[bRoot]
		}
	}
}
