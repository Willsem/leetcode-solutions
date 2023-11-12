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

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	adjList := make(map[int][]int)
	for route := range routes {
		for _, stop := range routes[route] {
			adjList[stop] = append(adjList[stop], route)
		}
	}

	q := NewQueue[int]()
	visited := make(map[int]struct{})
	for _, route := range adjList[source] {
		q.Push(route)
		visited[route] = struct{}{}
	}

	busCount := 1
	for !q.IsEmpty() {
		size := q.Size()

		for i := 0; i < size; i++ {
			route := q.Pop()

			for _, stop := range routes[route] {
				if stop == target {
					return busCount
				}

				for _, nextRoute := range adjList[stop] {
					if _, ok := visited[nextRoute]; !ok {
						visited[nextRoute] = struct{}{}
						q.Push(nextRoute)
					}
				}
			}
		}
		busCount++
	}

	return -1
}
