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

type element struct {
	node   int
	length int
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = []int{i + 1}
	}
	graph[n-1] = []int{}

	bfs := func() int {
		q := NewQueue[element]()
		q.Push(element{
			node:   0,
			length: 0,
		})
		visited := make(map[int]struct{})
		visited[0] = struct{}{}

		for !q.IsEmpty() {
			curr := q.Pop()

			if curr.node == n-1 {
				return curr.length
			}

			for _, next := range graph[curr.node] {
				if _, ok := visited[next]; !ok {
					q.Push(element{
						node:   next,
						length: curr.length + 1,
					})
					visited[next] = struct{}{}
				}
			}
		}

		return -1
	}

	result := make([]int, 0, len(queries))
	for _, q := range queries {
		from, to := q[0], q[1]

		graph[from] = append(graph[from], to)
		result = append(result, bfs())
	}

	return result
}
