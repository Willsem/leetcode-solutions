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

type State struct {
	mask, node, dist int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	allVisited := (1 << n) - 1
	q := NewQueue[State]()
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		q.Push(State{
			mask: 1 << i,
			node: i,
			dist: 0,
		})
		visited[(1<<i)*16+i] = true
	}

	for !q.IsEmpty() {
		cur := q.Pop()

		if cur.mask == allVisited {
			return cur.dist
		}

		for _, neighbor := range graph[cur.node] {
			newMask := cur.mask | (1 << neighbor)
			hashValue := newMask*16 + neighbor

			if !visited[hashValue] {
				visited[hashValue] = true
				q.Push(State{
					mask: newMask,
					node: neighbor,
					dist: cur.dist + 1,
				})
			}
		}
	}

	return -1
}
