type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func (q *Queue[T]) Push(el T) {
	q.data = append(q.data, el)
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) Pop() T {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	diam1 := findDiameter(edges1)
	diam2 := findDiameter(edges2)

	ans := getRadius(diam1) + 1 + getRadius(diam2)
	ans = max(ans, diam1)
	ans = max(ans, diam2)

	return ans
}

func findDiameter(edges [][]int) int {
	n := len(edges) + 1
	conn := connections(n, edges)
	farthest, _ := bfs(n, conn, 0)
	_, diameter := bfs(n, conn, farthest)
	return diameter
}

func getRadius(d int) int {
	return d - d/2
}

func connections(n int, edges [][]int) [][]int {
	conn := make([][]int, n)
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		conn[node1] = append(conn[node1], node2)
		conn[node2] = append(conn[node2], node1)
	}

	return conn
}

type queueElement struct {
	node   int
	height int
}

func bfs(n int, conn [][]int, startNode int) (farthestNode int, maxHeight int) {
	q := NewQueue[queueElement]()
	q.Push(queueElement{
		node:   startNode,
		height: 0,
	})

	visited := make(map[int]struct{}, n)

	for !q.IsEmpty() {
		c := q.Pop()

		if _, ok := visited[c.node]; ok {
			continue
		}
		visited[c.node] = struct{}{}

		if c.height > maxHeight {
			maxHeight = c.height
			farthestNode = c.node
		}

		for _, neighbor := range conn[c.node] {
			if _, ok := visited[neighbor]; !ok {
				q.Push(queueElement{
					node:   neighbor,
					height: c.height + 1,
				})
			}
		}
	}

	return farthestNode, maxHeight
}
