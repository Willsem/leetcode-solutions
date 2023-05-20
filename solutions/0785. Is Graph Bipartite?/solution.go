type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Clear() {
	q.data = []int{}
}

func (q *Queue) Push(el int) {
	q.data = append(q.data, el)
}

func (q *Queue) Peek() int {
	return q.data[0]
}

func (q *Queue) Pop() int {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	q := NewQueue()

	bfs := func(start int) bool {
		q.Clear()
		q.Push(start)

		for !q.IsEmpty() {
			curr := q.Pop()

			for _, next := range graph[curr] {
				if colors[next] == 0 {
					colors[next] = -colors[curr]
					q.Push(next)
				} else {
					if colors[next] == colors[curr] {
						return false
					}

					continue
				}
			}
		}

		return true
	}

	for i := range graph {
		if colors[i] == 0 {
			colors[i] = 1
			if !bfs(i) {
				return false
			}
		}
	}

	return true
}
