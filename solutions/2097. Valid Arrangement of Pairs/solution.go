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

func validArrangement(pairs [][]int) [][]int {
	pathsMatrix := make(map[int]*Queue[int])
	in, out := make(map[int]int), make(map[int]int)

	for _, pair := range pairs {
		start, end := pair[0], pair[1]
		if pathsMatrix[start] == nil {
			pathsMatrix[start] = NewQueue[int]()
		}

		pathsMatrix[start].Push(end)
		out[start]++
		in[end]++
	}

	result := make([]int, 0)

	var dfs func(node int)
	dfs = func(node int) {
		for pathsMatrix[node] != nil && !pathsMatrix[node].IsEmpty() {
			dfs(pathsMatrix[node].Pop())
		}
		result = append(result, node)
	}

	startNode := -1
	for node := range out {
		if out[node] == in[node]+1 {
			startNode = node
			break
		}
	}
	if startNode == -1 {
		startNode = pairs[0][0]
	}

	dfs(startNode)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	pairsResult := make([][]int, 0)
	for i := 1; i < len(result); i++ {
		pairsResult = append(pairsResult, []int{result[i-1], result[i]})
	}
	return pairsResult
}
