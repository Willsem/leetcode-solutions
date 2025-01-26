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

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	inDegree := make([]int, n)
	for _, f := range favorite {
		inDegree[f]++
	}

	q := NewQueue[int]()
	for i := range n {
		if inDegree[i] == 0 {
			q.Push(i)
		}
	}

	longestChain := make([]int, n)
	for !q.IsEmpty() {
		i := q.Pop()
		fav := favorite[i]
		longestChain[fav] = max(longestChain[fav], longestChain[i]+1)
		inDegree[fav]--
		if inDegree[fav] == 0 {
			q.Push(fav)
		}
	}

	maxCycleSize := 0
	mutualPairs := 0
	for i := range n {
		if inDegree[i] == 0 {
			continue
		}
		cycleSize := 0
		cur := i
		for inDegree[cur] > 0 {
			inDegree[cur] = 0
			cur = favorite[cur]
			cycleSize++
		}
		if cycleSize == 2 {
			mutualPairs += 2 + longestChain[i] + longestChain[favorite[i]]
		} else {
			maxCycleSize = max(maxCycleSize, cycleSize)
		}
	}

	return max(maxCycleSize, mutualPairs)
}
