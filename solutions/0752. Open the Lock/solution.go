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

var next = map[byte]string{
	'0': "1", '1': "2", '2': "3", '3': "4", '4': "5",
	'5': "6", '6': "7", '7': "8", '8': "9", '9': "0",
}

var prev = map[byte]string{
	'0': "9", '1': "0", '2': "1", '3': "2", '4': "3",
	'5': "4", '6': "5", '7': "6", '8': "7", '9': "8",
}

type combination struct {
	val   string
	depth int
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]struct{})
	for _, deadend := range deadends {
		visited[deadend] = struct{}{}
	}

	if _, ok := visited["0000"]; ok {
		return -1
	}

	q := NewQueue[combination]()
	q.Push(combination{"0000", 0})
	visited["0000"] = struct{}{}

	for !q.IsEmpty() {
		curr := q.Pop()
		val := curr.val
		depth := curr.depth

		if val == target {
			return depth
		}

		for i := 0; i < 4; i++ {
			nextVal := val[:i] + next[val[i]] + val[i+1:]
			if _, ok := visited[nextVal]; !ok {
				visited[nextVal] = struct{}{}
				q.Push(combination{nextVal, depth + 1})
			}

			prevVal := val[:i] + prev[val[i]] + val[i+1:]
			if _, ok := visited[prevVal]; !ok {
				visited[prevVal] = struct{}{}
				q.Push(combination{prevVal, depth + 1})
			}
		}
	}

	return -1
}
