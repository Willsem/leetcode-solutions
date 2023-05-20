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
