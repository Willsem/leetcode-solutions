import (
	"container/heap"
	"math"
)

type Edge struct {
	to     int
	weight float64
}

type Node struct {
	node   int
	weight float64
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight > pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]Edge, n)
	for i, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], Edge{edge[1], math.Log(succProb[i])})
		graph[edge[1]] = append(graph[edge[1]], Edge{edge[0], math.Log(succProb[i])})
	}

	dist := make([]float64, n)
	for i := range dist {
		dist[i] = math.Inf(-1)
	}
	dist[start] = 0

	pq := PriorityQueue{}
	heap.Push(&pq, &Node{start, 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Node)
		if curr.node == end {
			return math.Exp(dist[end])
		}
		if curr.weight < dist[curr.node] {
			continue
		}
		for _, edge := range graph[curr.node] {
			if nextDist := dist[curr.node] + edge.weight; nextDist > dist[edge.to] {
				dist[edge.to] = nextDist
				heap.Push(&pq, &Node{edge.to, nextDist})
			}
		}
	}

	return 0
}
