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

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	q := NewQueue()
	q.Push(0)

	seen := make(map[int]struct{}, 0)

	for !q.IsEmpty() {
		start := q.Pop()
		if start == len(s) {
			return true
		}

		for end := start + 1; end <= len(s); end++ {
			if _, ok := seen[end]; ok {
				continue
			}

			if _, ok := words[s[start:end]]; ok {
				q.Push(end)
				seen[end] = struct{}{}
			}
		}
	}

	return false
}
